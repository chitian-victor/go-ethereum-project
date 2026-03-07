package service

import (
	"testing"

	"github.com/chitian-victor/go-ethereum-project/client"
	"github.com/chitian-victor/go-ethereum-project/consts"
)

func TestCommon(t *testing.T) {
	client.Init()
	//ctx := context.Background()

	t.Run("IsContractAddress", func(t *testing.T) {
		isContractAddress, err := IsContractAddress(client.Client, consts.USER_ADDRESS)
		if err != nil {
			t.Logf("IsContractAddress failed, err: %v", err)
			return
		}
		t.Logf("IsContractAddress: res=%v", isContractAddress)
	})
}
