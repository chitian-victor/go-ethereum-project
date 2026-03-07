package service

import (
	"context"
	"log"
	"math/big"

	"github.com/chitian-victor/go-ethereum-project/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBlockHeader(cli *ethclient.Client) (*types.Header, error) {
	header, err := cli.HeaderByNumber(context.Background(), nil) // the latest one
	if err != nil {
		log.Printf("GetBlockInfo failed, err: %v", err)
		return nil, err
	}
	return header, nil
}

func GetETHBalance(ctx context.Context, cli *ethclient.Client, accountHex string) (*big.Int, error) {
	balanceWei, err := cli.BalanceAt(ctx, common.HexToAddress(accountHex), nil)
	if err != nil {
		log.Printf("failed to get balance: %v", err)
		return nil, err
	}
	balanceEth := utils.WeiToEther(balanceWei)
	// todo-hs
	log.Printf("balance: %d wei === %.6f eth:", balanceWei, balanceEth)
	return balanceWei, nil
}
