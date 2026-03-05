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

	t.Run("GetETHBalance", func(t *testing.T) {
		bw, err := GetETHBalance(ctx, client.Client, consts.USER_ADDRESS)
		if err != nil {
			t.Logf("GetETHBalance failed, err: %v", err)
			return
		}
		t.Logf("GetETHBalance(wei): %v", bw)
	})
}
