package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/chitian-victor/go-ethereum-project/contract"
	"github.com/chitian-victor/go-ethereum-project/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func deployPumpkinToken(ctx context.Context, cli *ethclient.Client, privateKeyHex, ownerAddressHex string) (*contract.PumpkinToken, string, error) {
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := cli.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := cli.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("获取 ChainID 失败: %v", err)
	}

	auth := bind.NewKeyedTransactor(privateKey, chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	// 最好不要随意设置 gaslimit，否则容易导致合约部署不上
	//auth.GasLimit
	auth.GasPrice = gasPrice

	ownerAddress := common.HexToAddress(ownerAddressHex)
	address, tx, instance, err := contract.DeployPumpkinToken(auth, cli, ownerAddress)
	if err != nil {
		log.Fatal(err)
	}
	//可以用来等待合约部署完成
	address, err = bind.WaitDeployed(ctx, cli, tx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PumpkinToken Address=", address.Hex())
	fmt.Printf("transaction:\n hash:%v\ngas:%v \n", tx.Hash().Hex(), tx.Gas())
	return instance, address.Hex(), nil
}

func CallContractMethod(cli *ethclient.Client, privateKeyHex, contractAddressHex, methodName string, args ...interface{}) error {
	// 1. 加载私钥 (推导 From 地址)
	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("解析私钥失败: %v", err)
	}
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// 3. 定义合约地址和接收者地址
	contractAddress := common.HexToAddress(contractAddressHex)

	// 4. 核心步骤：通过简化的 ABI 字符串解析合约接口
	abiFilePath := "../abi/PumpkinToken.abi"
	contractABI, err := LoadContractABI(abiFilePath)
	if err != nil {
		log.Printf("LoadContractABI failed, err: %v", err)
		return err
	}

	// 注意：这里的方法名不需要带参数类型，因为我们利用 abi 解析的，这里只是选取方法
	data, err := contractABI.Pack(methodName, args...)
	if err != nil {
		log.Fatalf("打包 Data 失败: %v", err)
	}
	fmt.Printf("打包生成的 Raw Data: 0x%x\n", data)

	// 6. 估算 Gas Limit (合约调用绝不是固定的 21000，必须动态估算！)
	msg := ethereum.CallMsg{
		From: fromAddress,
		To:   &contractAddress,
		Data: data,
	}
	gasLimit, err := cli.EstimateGas(context.Background(), msg)
	if err != nil {
		log.Fatalf("估算 Gas 失败 (可能余额不足或交易必然失败): %v", err)
	}
	// 稍微给 GasLimit 加一点余量，防止执行时状态微调导致 Out of Gas
	gasLimit = uint64(float64(gasLimit) * 1.1)

	// 7. 获取 Nonce, ChainID, TipCap, FeeCap 等网络参数
	nonce, _ := cli.PendingNonceAt(context.Background(), fromAddress)
	chainID, _ := cli.NetworkID(context.Background())
	tipCap, _ := cli.SuggestGasTipCap(context.Background())
	feeCap, _ := cli.SuggestGasPrice(context.Background())

	// 8. 构建 EIP-1559 交易
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: tipCap,
		GasFeeCap: feeCap,
		Gas:       gasLimit,
		To:        &contractAddress, // 注意：这里的 To 是合约地址，不是代币接收者！
		Value:     big.NewInt(0),    // 纯调用合约方法，不随交易发送 ETH 主币
		Data:      data,             // 把打包好的二进制代码塞进去
	})

	// 9. 签名并广播
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	if err != nil {
		log.Fatalf("签名失败: %v", err)
	}

	err = cli.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("广播失败: %v", err)
	}
	fmt.Printf("调用合约成功! TxHash: %s\n", signedTx.Hash().Hex())
	return nil
}

func GetPumpkinTokenBalance(cli *ethclient.Client, contractAddress, userAddress string) (*big.Int, error) {
	pumpkinToken, err := contract.NewPumpkinToken(common.HexToAddress(contractAddress), cli)
	if err != nil {
		log.Printf("Failed to connect to PumpkinToken, err=%v", err)
		return nil, err
	}
	balance, err := pumpkinToken.BalanceOf(nil, common.HexToAddress(userAddress))
	if err != nil {
		log.Printf("Failed to get balance of token, err=%v", err)
		return nil, err
	}
	log.Printf("GetPumpkinTokenBalance: %v(wei) === %v(eth)", balance, utils.WeiToEther(balance))
	return balance, nil
}

func MintPumpkinToken(ctx context.Context, cli *ethclient.Client, contractAddress, privateKeyHex string, mintAmountWei *big.Int) error {
	pumpkinToken, err := contract.NewPumpkinToken(common.HexToAddress(contractAddress), cli)
	if err != nil {
		log.Printf("Failed to connect to MintPumpkinToken, err=%v", err)
		return err
	}
	// 签名
	auth, err := GenerateAuth(ctx, cli, privateKeyHex)
	if err != nil {
		log.Printf("Failed to generate auth, err=%v", err)
		return err
	}
	// 发起交易
	transaction, err := pumpkinToken.Mint(auth, mintAmountWei)
	if err != nil {
		log.Printf("Failed to mint token, err=%v", err)
		return err
	}
	log.Printf("transaction=%v", utils.ToJson(transaction))
	// 等待交易完成
	err = WaitTransaction(ctx, cli, transaction, 60)
	if err != nil {
		log.Printf("Failed to mint token, err=%v", err)
		return err
	}
	return nil
}
