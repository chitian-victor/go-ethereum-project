package service

import (
	"context"
	"log"

	"github.com/chitian-victor/go-ethereum-project/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ConnectPumpkinToken(ctx context.Context, cli *ethclient.Client, contractHex string) (*contract.PumpkinToken, error) {
	pumpkinToken, err := contract.NewPumpkinToken(common.HexToAddress(contractHex), cli)
	if err != nil {
		log.Fatalf("Failed to connect to PumpkinToken, err=%v", err)
		return nil, err
	}
	return pumpkinToken, nil
}
