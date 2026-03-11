package subscriber

import (
    "context"
    "fmt"
    "log"

    "github.com/chitian-victor/go-ethereum-project/contract"
    "github.com/chitian-victor/go-ethereum-project/service"
    "github.com/chitian-victor/go-ethereum-project/utils"
    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
)

// 这里仅仅示例了解析 Transfer 事件
func SubscribeEventByWatch(ctx context.Context, cli *ethclient.Client, contractAddressHex string) {
    contractAddress := common.HexToAddress(contractAddressHex)

    tokenContract, err := contract.NewPumpkinToken(contractAddress, cli)
    if err != nil {
        log.Fatalf("实例化合约失败: %v", err)
    }
    transferSink := make(chan *contract.PumpkinTokenTransfer)

    // 5. 设置 Watch 选项
    watchOpts := &bind.WatchOpts{
        Context: ctx,
    }

    // 启动订阅
    // WatchTransfer 的后两个参数是 from 和 to (对应 Solidity 里的 indexed 参数)。
    // 如果传入 nil，代表监听该合约【所有】的 Transfer 事件。
    // 如果想监听特定人的转账，可以传入包含该地址的数组，例如：[]common.Address{myAddress}
    sub, err := tokenContract.WatchTransfer(watchOpts, transferSink, nil, nil)
    if err != nil {
        log.Fatalf("订阅 Transfer 事件失败: %v", err)
    }
    fmt.Println("开始监听 PumpkinToken 的 Transfer 事件...")

    // 7. 使用 goroutine 或主循环处理事件
    for {
        select {
        case <-ctx.Done():
            fmt.Println("上下文取消，退出订阅")
            return
        case err := <-sub.Err():
            // 处理订阅异常断开
            log.Fatalf("订阅发生错误: %v", err)

        case transferEvent := <-transferSink:
            // 直接获取强类型的解析结果，告别手动解析 Topics 和 Data
            fmt.Printf("\n--- 收到新转账事件 (区块: %d) ---\n", transferEvent.Raw.BlockNumber)
            fmt.Printf("交易 Hash: %s\n", transferEvent.Raw.TxHash.Hex())

            // 直接点出属性，非常丝滑
            fmt.Printf("From  : %s\n", transferEvent.From.Hex())
            fmt.Printf("To    : %s\n", transferEvent.To.Hex())
            fmt.Printf("Value : %s\n", transferEvent.Value.String())
        }
    }
}

// 这里仅仅示例了解析 Transfer 事件，更底层的监听方式
func SubscribeEvent(ctx context.Context, cli *ethclient.Client, contractAddressHex string) {
    contractAddress := common.HexToAddress(contractAddressHex)

    // 可以指定要监听的 event
    logTransferSig := []byte("Transfer(address,address,uint256)")
    logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

    query := ethereum.FilterQuery{
        Addresses: []common.Address{contractAddress},
        Topics:    [][]common.Hash{{logTransferSigHash}},
    }

    logs := make(chan types.Log)
    sub, err := cli.SubscribeFilterLogs(ctx, query, logs)
    if err != nil {
        log.Fatal(err)
    }

    contractABI, err := service.LoadContractABI("../abi/PumpkinToken.abi")
    if err != nil {
        log.Printf("LoadContractABI failed, err: %v", err)
        return
    }

    for {
        select {
        case <-ctx.Done():
            fmt.Println("上下文取消，退出订阅")
            return
        case err := <-sub.Err():
            log.Fatal(err)
        case vLog := <-logs:
            // 接收到新的事件日志
            fmt.Printf("\n--- 收到新事件区块高度: %d ---\n", vLog.BlockNumber)
            fmt.Printf("交易 Hash: %s\n", vLog.TxHash.Hex())
            // 其实就是crypto.Keccak256Hash("transfer(address,uint256)")
            fmt.Printf("签名哈希: %s\n", vLog.Topics[0].Hex())

            // 验证 Topics 长度。对于 ERC-20 Transfer，from 和 to 都是 indexed，所以会在 topics 里
            // Topic[0] 是签名哈希, Topic[1] 是 from, Topic[2] 是 to
            if len(vLog.Topics) == 3 {
                // Address 在 Topic 中是被填充到 32 字节的，需要转换回 Address 类型
                from := common.HexToAddress(vLog.Topics[1].Hex())
                to := common.HexToAddress(vLog.Topics[2].Hex())
                fmt.Printf("From: %s\n", from.Hex())
                fmt.Printf("To:   %s\n", to.Hex())
            }

            // 解析 Data (非 indexed 的参数，这里是 value，代币的数量)
            resp, err := contractABI.Unpack("Transfer", vLog.Data)
            if err != nil {
                log.Printf("Unpack failed, err: %v", err)
                return
            }
            fmt.Println(" value:", utils.ToJson(resp))

        }
    }
}
