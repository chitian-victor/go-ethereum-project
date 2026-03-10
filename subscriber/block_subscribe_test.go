package subscriber

import (
	"context"
	"testing"

	"github.com/chitian-victor/go-ethereum-project/client"
)

func TestSubscribeBlockHeader(t *testing.T) {
	client.Init()
	ctx := context.Background()
	SubscribeBlockHeader(ctx, client.SubClient)
}
