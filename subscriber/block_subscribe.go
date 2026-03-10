package subscriber

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SubscribeBlockHeader(ctx context.Context, cli *ethclient.Client) {
	headers := make(chan *types.Header)
	sub, err := cli.SubscribeNewHead(ctx, headers)
	if err != nil {
		fmt.Println("SubscribeNewHead failed, err:", err)
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex())
			block, err := cli.BlockByHash(ctx, header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("============", block.Hash().Hex(), "==============")
			fmt.Println(block.Number().Uint64())
			fmt.Println(time.Unix(int64(block.Time()), 0))
			fmt.Println(block.Nonce())
			fmt.Println(len(block.Transactions()))
			fmt.Println()
		}
	}
}
