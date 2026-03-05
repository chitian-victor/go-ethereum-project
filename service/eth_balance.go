package service

import (
	"context"
	"log"
	"math/big"

	"github.com/chitian-victor/go-ethereum-project/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetETHBalance(ctx context.Context, cli *ethclient.Client, accountHex string) (*big.Int, error) {
	balanceWei, err := cli.BalanceAt(ctx, common.HexToAddress(accountHex), nil)
	if err != nil {
		log.Fatalf("failed to get balance: %v", err)
		return nil, err
	}
	balanceEth := utils.WeiToEther(balanceWei)
	// todo-hs
	log.Printf("balance: %d wei === %.6f eth:", balanceWei, balanceEth)
	return balanceWei, nil
}
