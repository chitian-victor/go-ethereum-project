package subscriber

import (
	"context"
	"testing"

	"github.com/chitian-victor/go-ethereum-project/client"
	"github.com/chitian-victor/go-ethereum-project/consts"
)

func TestSubscribeEvent(t *testing.T) {
	client.Init()
	ctx := context.Background()
	SubscribeEvent(ctx, client.SubClient, consts.PUMPKIN_TOKEN)
}
