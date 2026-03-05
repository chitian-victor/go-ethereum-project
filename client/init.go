package client

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

var Client *ethclient.Client

func Init() {
	//client, err := ethclient.Dial("https://cloudflare-eth.com")
	blockClient, err := ethclient.Dial("http://localhost:8545") //ganache
	if err != nil {
		panic(err)
	}
	Client = blockClient
}
