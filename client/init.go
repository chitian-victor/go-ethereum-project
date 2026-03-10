package client

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	Client    *ethclient.Client
	SubClient *ethclient.Client
)

func Init() {
	// 因为都是 foundry 本地节点，接口调用默认明文传输不加密，所以都是 http 和 ws
	// 另外foundry的 http 和 ws 都是一样的端口，其他服务器不一定的

	blockClient, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		panic(err)
	}
	subBlockClient, err := ethclient.Dial("ws://localhost:8545")
	if err != nil {
		panic(err)
	}

	Client = blockClient
	SubClient = subBlockClient
}
