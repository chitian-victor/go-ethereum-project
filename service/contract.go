package service

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/chitian-victor/go-ethereum-project/contract"
	"github.com/chitian-victor/go-ethereum-project/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

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
	return balance, nil
}

func MintPumpkinToken(ctx context.Context, cli *ethclient.Client, contractAddress, privateKeyHex string, mintAmountWei *big.Int) error {
	pumpkinToken, err := contract.NewPumpkinToken(common.HexToAddress(contractAddress), cli)
	if err != nil {
		log.Printf("Failed to connect to MintPumpkinToken, err=%v", err)
		return err
	}
	// 1. 加载你的私钥 (切记：千万不要把私钥硬编码在生产环境的代码里！)
	// 私钥需要去掉前缀 ox
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		log.Printf("Failed to parse private key, err=%v", err)
		return err
	}
	// 2. 获取当前链的 ChainID
	chainID, err := cli.ChainID(ctx)
	if err != nil {
		log.Printf("Failed to get chain id, err=%v", err)
		return err
	}
	// 3. 创建一个交易签名者 (TransactOpts)
	auth := bind.NewKeyedTransactor(privateKey, chainID)
	// 4. (可选) 手动设置 Gas 费用和限制。如果不设置，go-ethereum 会自动帮你估算
	// gasPrice, _ := client.SuggestGasPrice(context.Background())
	// auth.GasPrice = gasPrice
	// auth.GasLimit = uint64(300000)

	// 如果需要随交易发送 ETH，可以设置 Value (单位是 Wei)
	// auth.Value = big.NewInt(1000000000000000000) // 1 ETH

	// 5. 发起交易
	transaction, err := pumpkinToken.Mint(auth, mintAmountWei)
	if err != nil {
		log.Printf("Failed to mint token, err=%v", err)
		return err
	}
	// 6. 等待交易上链确认（关键修复点2）
	// 设置超时时间（比如1分钟），避免无限等待
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	receipt, err := bind.WaitMined(ctxWithTimeout, cli, transaction.Hash())
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	// 7. 检查交易是否成功执行
	if receipt.Status != types.ReceiptStatusSuccessful {
		err = fmt.Errorf("mint transaction failed, status: %d, tx hash: %s", receipt.Status, transaction.Hash().Hex())
		log.Printf("Error: %v", err)
		return err
	}

	log.Printf("transaction=%v", utils.ToJson(transaction))
	return nil
}
