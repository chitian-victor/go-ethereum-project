package service

import (
	"context"
	"testing"

	"github.com/chitian-victor/go-ethereum-project/client"
	"github.com/chitian-victor/go-ethereum-project/consts"
)

func TestETHBalance(t *testing.T) {
	client.Init()
	ctx := context.Background()

	t.Run("TransferETH", func(t *testing.T) {
		err := TransferETH(client.Client, consts.OWNER_PRIVATE_KEY, consts.USER_ADDRESS, 3)
		if err != nil {
			t.Logf("TransferETH failed, err: %v", err)
			return
		}
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
