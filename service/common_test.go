package service

import (
	"context"
	"log"
	"testing"

	"github.com/chitian-victor/go-ethereum-project/client"
	"github.com/chitian-victor/go-ethereum-project/consts"
	"github.com/chitian-victor/go-ethereum-project/utils"
)

func TestCommon(t *testing.T) {
	client.Init()
	ctx := context.Background()

	t.Run("GetBlockInfo", func(t *testing.T) {
		info, err := GetBlockInfo(client.Client, nil)
		if err != nil {
			t.Logf("GetBlockInfo failed, err: %v", err)
			return
		}
		for _, tx := range info.Transactions() {
			t.Log(tx.Hash().Hex())        // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
			t.Log(tx.Value().String())    // 10000000000000000
			t.Log(tx.Gas())               // 105000
			t.Log(tx.GasPrice().Uint64()) // 102000000000
			t.Log(tx.Nonce())             // 110644
			t.Log(tx.Data())              // []
			t.Log(tx.To().Hex())          // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

			// 获取 sender
			sender, err := GetSender(client.Client, info.Header(), tx)
			if err != nil || sender == nil {
				return
			}
			t.Logf("msg.Sender=%v", sender.Hex())
			// 获取交易收据
			receipt, err := client.Client.TransactionReceipt(ctx, tx.Hash())
			if err != nil {
				log.Printf("GetBlockInfo failed, err: %v", err)
			}
			t.Log(receipt.Status) // 1
		}
		t.Logf("GetBlockInfo: res=%v", utils.ToJson(info))
	})

	t.Run("GetBlockHeader", func(t *testing.T) {
		header, err := GetBlockHeader(client.Client, nil)
		if err != nil {
			t.Logf("GetBlockHeader failed, err: %v", err)
			return
		}
		t.Logf("GetBlockHeader: block num=%v", header.Number.Int64())
		t.Logf("GetBlockHeader: res=%v", utils.ToJson(header))
	})

	t.Run("IsContractAddress", func(t *testing.T) {
		isContractAddress, err := IsContractAddress(client.Client, consts.USER_ADDRESS)
		if err != nil {
			t.Logf("IsContractAddress failed, err: %v", err)
			return
		}
		t.Logf("IsContractAddress: res=%v", isContractAddress)
	})
}
