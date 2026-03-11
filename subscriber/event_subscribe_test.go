package subscriber

import (
	"context"
	"testing"
	"time"

	"github.com/chitian-victor/go-ethereum-project/client"
	"github.com/chitian-victor/go-ethereum-project/consts"
)

func TestSubscribeEvent(t *testing.T) {
	client.Init()
	ctx := context.Background()
	t.Run("SubscribeEvent", func(t *testing.T) {
		ctx, _ := context.WithTimeout(ctx, 5*time.Minute)
		SubscribeEvent(ctx, client.SubClient, consts.PUMPKIN_TOKEN)
	})

	t.Run("SubscribeEventByWatch", func(t *testing.T) {
		ctx, _ := context.WithTimeout(ctx, 5*time.Minute)
		SubscribeEventByWatch(ctx, client.SubClient, consts.PUMPKIN_TOKEN)
	})
}
