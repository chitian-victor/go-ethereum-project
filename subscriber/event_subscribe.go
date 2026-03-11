package subscriber

import (
    "context"
    "fmt"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
)

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

    for {
        select {
        case err := <-sub.Err():
            log.Fatal(err)
        case vLog := <-logs:
            // 接收到新的事件日志
            fmt.Printf("\n--- 收到新事件区块高度: %d ---\n", vLog.BlockNumber)
            fmt.Printf("交易 Hash: %s\n", vLog.TxHash.Hex())

            // 验证 Topics 长度。对于 ERC-20 Transfer，from 和 to 都是 indexed，所以会在 topics 里
            // Topic[0] 是签名哈希, Topic[1] 是 from, Topic[2] 是 to
            if len(vLog.Topics) == 3 {
                // Address 在 Topic 中是被填充到 32 字节的，需要转换回 Address 类型
                from := common.HexToAddress(vLog.Topics[1].Hex())
                to := common.HexToAddress(vLog.Topics[2].Hex())
                fmt.Printf("From: %s\n", from.Hex())
                fmt.Printf("To:   %s\n", to.Hex())
            }

            // 解析 Data (非 indexed 的参数，这里是 value)
            // USDT 的精度是 6，为了简单起见，这里直接打印大整数
            transferValue := new(big.Int)
            transferValue.SetBytes(vLog.Data)
            fmt.Printf("Value: %s\n", transferValue.String())
        }
    }
}
