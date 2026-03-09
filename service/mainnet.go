package service

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "log"
    "math/big"
    "strings"

    "github.com/chitian-victor/go-ethereum-project/utils"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/params"
)

func TransferETH(cli *ethclient.Client, privateKeyHex, toAddressHex string, amount int64) error {
    // 1. 加载发送方的私钥
    privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
    if err != nil {
        log.Fatalf("解析私钥失败: %v", err)
    }

    // 2. 从私钥推导发送方地址 (From)
    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("无法将公钥转换为 ECDSA 格式")
    }
    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

    // 3. 获取发送方当前的 Pending Nonce
    nonce, err := cli.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatalf("获取 Nonce 失败: %v", err)
    }

    // 4. 设置转账金额 (注意：以太坊底层单位是 Wei, 1 ETH = 10^18 Wei)
    // 这里演示转账 0.01 ETH
    value := big.NewInt(1).Mul(big.NewInt(params.Ether), big.NewInt(amount))

    // 普通 ETH 转账的 Gas 限制固定为 21000
    gasLimit := uint64(21000)

    // 获取 EIP-1559 的推荐 Gas 参数
    tipCap, err := cli.SuggestGasTipCap(context.Background()) // 小费 (MaxPriorityFeePerGas)
    if err != nil {
        log.Fatalf("获取小费上限失败: %v", err)
    }
    feeCap, err := cli.SuggestGasPrice(context.Background()) // 总费用上限 (MaxFeePerGas)
    if err != nil {
        log.Fatalf("获取 Gas 价格上限失败: %v", err)
    }

    // 设置接收方地址 (To)
    toAddress := common.HexToAddress(toAddressHex)

    // 获取当前网络的 ChainID (用于签名防重放攻击)
    chainID, err := cli.NetworkID(context.Background())
    if err != nil {
        log.Fatalf("获取 ChainID 失败: %v", err)
    }

    // 5. 构建 EIP-1559 交易对象
    tx := types.NewTx(&types.DynamicFeeTx{
        ChainID:   chainID,
        Nonce:     nonce,
        GasTipCap: tipCap,
        GasFeeCap: feeCap,
        Gas:       gasLimit,
        To:        &toAddress,
        Value:     value,
        Data:      nil, // 普通转账不需要附加数据
    })

    // 使用最新的签名器进行私钥签名
    signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
    if err != nil {
        log.Fatalf("交易签名失败: %v", err)
    }

    // 6. 广播交易到网络
    err = cli.SendTransaction(context.Background(), signedTx)
    if err != nil {
        log.Fatalf("广播交易失败: %v", err)
    }

    fmt.Printf("交易哈希 (TxHash): %s\n", signedTx.Hash().Hex())
    return nil
}

func GetETHBalance(ctx context.Context, cli *ethclient.Client, accountHex string) (*big.Int, error) {
    balanceWei, err := cli.BalanceAt(ctx, common.HexToAddress(accountHex), nil)
    if err != nil {
        log.Printf("failed to get balance: %v", err)
        return nil, err
    }
    balanceEth := utils.WeiToEther(balanceWei)
    log.Printf("balance: %d wei === %.6f eth:", balanceWei, balanceEth)
    return balanceWei, nil
}
