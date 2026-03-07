package service

import (
	"context"
	"testing"

	"github.com/chitian-victor/go-ethereum-project/client"
	"github.com/chitian-victor/go-ethereum-project/consts"
	"github.com/chitian-victor/go-ethereum-project/utils"
)

func TestETHBalance(t *testing.T) {
	client.Init()
	ctx := context.Background()

	t.Run("GetBlockHeader", func(t *testing.T) {
		header, err := GetBlockHeader(client.Client)
		if err != nil {
			t.Logf("GetBlockHeader failed, err: %v", err)
			return
		}
		t.Logf("GetBlockHeader: block num=%v", header.Number.Int64())
		t.Logf("GetBlockHeader: res=%v", utils.ToJson(header))
	})

	t.Run("GenerateNewAddress", func(t *testing.T) {
		privateKey, publicKey, address, err := GenerateNewAddress()
		if err != nil {
			t.Logf("GenerateNewAddress failed, err: %v", err)
			return
		}
		t.Logf("GenerateNewAddress:\n privateKey=%v\n publicKey=%v\n address=%v", privateKey, publicKey, address)
	})

	t.Run("GetETHBalance", func(t *testing.T) {
		bw, err := GetETHBalance(ctx, client.Client, consts.USER_ADDRESS)
		if err != nil {
			t.Logf("GetETHBalance failed, err: %v", err)
			return
		}
		t.Logf("GetETHBalance(wei): %v", bw)
	})
}
