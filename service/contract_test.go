package service

import (
	"context"
	"testing"

	"github.com/chitian-victor/go-ethereum-project/client"
	"github.com/chitian-victor/go-ethereum-project/consts"
	"github.com/chitian-victor/go-ethereum-project/utils"
)

func TestConnect(t *testing.T) {
	client.Init()
	ctx := context.Background()

	t.Run("ConnectPumpkinToken", func(t *testing.T) {
		pumpkinToken, err := ConnectPumpkinToken(ctx, client.Client, consts.PUMPKIN_TOKEN)
		if err != nil {
			t.Logf("ConnectPumpkinToken failed, err: %v", err)
			return
		}
		t.Logf("ConnectPumpkinToken: %v", utils.ToJson(pumpkinToken))
	})
}
