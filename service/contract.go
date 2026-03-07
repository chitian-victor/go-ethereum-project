package service

import (
	"context"
	"log"
	"math/big"

	"github.com/chitian-victor/go-ethereum-project/contract"
	"github.com/chitian-victor/go-ethereum-project/utils"
	"github.com/ethereum/go-ethereum/common"
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
