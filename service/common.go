package service

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBlockInfo(cli *ethclient.Client, num *big.Int) (*types.Block, error) {
	blockInfo, err := cli.BlockByNumber(context.Background(), num) // num=nil, the latest one
	if err != nil {
		log.Printf("BlockByNumber failed, err: %v", err)
		return nil, err
	}
	return blockInfo, nil
}

func GetBlockHeader(cli *ethclient.Client, num *big.Int) (*types.Header, error) {
	header, err := cli.HeaderByNumber(context.Background(), num) // num=nil, the latest one
	if err != nil {
		log.Printf("HeaderByNumber failed, err: %v", err)
		return nil, err
	}
	return header, nil
}

func GetSender(client *ethclient.Client, header *types.Header, tx *types.Transaction) (*common.Address, error) {
	// 步骤1：获取交易所在区块的签名者验证器（ChainID + 区块头）
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Printf("get ChainID failed, err: %v", err)
		return nil, fmt.Errorf("获取ChainID失败: %v", err)
	}
	signer := types.LatestSignerForChainID(chainID)
	// 步骤2：验证并解析签名者（sender）
	sender, err := types.Sender(signer, tx)
	if err != nil {
		log.Printf("get sender failed, err: %v", err)
		return nil, err
	}
	return &sender, nil
}

func IsContractAddress(client *ethclient.Client, addressHex string) (bool, error) {
	address := common.HexToAddress(addressHex)
	bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is the latest block
	if err != nil {
		log.Printf("failed to get code, err: %v", err)
		return false, err
	}
	return len(bytecode) > 0, nil
}

func GenerateNewAddress() (privateKeyHex string, publicKeyHex, addressHex string, err error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Printf("Failed to generate private key, err=%v", err)
		return
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex = hexutil.Encode(privateKeyBytes)
	publicKeyHex = hexutil.Encode(crypto.FromECDSAPub(&privateKey.PublicKey))
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	addressHex = address.Hex()
	return
}

func WaitTransaction(ctx context.Context, cli *ethclient.Client, tx *types.Transaction, seconds int64) error {
	// 1. 等待交易上链确认
	// 设置超时时间（比如1分钟），避免无限等待
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Duration(seconds)*time.Second)
	defer cancel()

	receipt, err := bind.WaitMined(ctxWithTimeout, cli, tx.Hash())
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	// 2. 检查交易是否成功执行
	if receipt.Status != types.ReceiptStatusSuccessful {
		err = fmt.Errorf("mint transaction failed, status: %d, tx hash: %s", receipt.Status, tx.Hash().Hex())
		log.Printf("Error: %v", err)
		return err
	}
	return nil
}

func GenerateAuth(ctx context.Context, cli *ethclient.Client, privateKeyHex string) (*bind.TransactOpts, error) {
	// 1. 加载你的私钥 (切记：千万不要把私钥硬编码在生产环境的代码里！)
	// 私钥需要去掉前缀 ox
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		log.Printf("Failed to parse private key, err=%v", err)
		return nil, err
	}
	// 2. 获取当前链的 ChainID
	chainID, err := cli.ChainID(ctx)
	if err != nil {
		log.Printf("Failed to get chain id, err=%v", err)
		return nil, err
	}
	// 3. 创建一个交易签名者 (TransactOpts)
	auth := bind.NewKeyedTransactor(privateKey, chainID)
	// 4. (可选) 手动设置 Gas 费用和限制。如果不设置，go-ethereum 会自动帮你估算
	// gasPrice, _ := client.SuggestGasPrice(context.Background())
	// auth.GasPrice = gasPrice
	// auth.GasLimit = uint64(300000)

	return auth, nil
}
