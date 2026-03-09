package service

import (
	"context"
	"math/big"
	"testing"

	"github.com/chitian-victor/go-ethereum-project/client"
	"github.com/chitian-victor/go-ethereum-project/consts"
	"github.com/chitian-victor/go-ethereum-project/utils"
	"github.com/ethereum/go-ethereum/params"
)

func TestConnect(t *testing.T) {
	client.Init()
	ctx := context.Background()

	t.Run("GetPumpkinTokenBalance", func(t *testing.T) {
		ownerBalance, err := GetPumpkinTokenBalance(client.Client, consts.PUMPKIN_TOKEN, consts.USER_ADDRESS)
		if err != nil {
			t.Logf("GetPumpkinTokenBalance failed, err: %v", err)
			return
		}
		balancePK := utils.WeiToEther(ownerBalance)
		t.Logf("GetPumpkinTokenBalance: %v(wei) === %v(eth)", ownerBalance, balancePK)
	})

	t.Run("MintPumpkinToken", func(t *testing.T) {
		ether := big.NewInt(int64(params.Ether))
		mintAmount := big.NewInt(1).Mul(big.NewInt(100), ether)
		err := MintPumpkinToken(ctx, client.Client, consts.PUMPKIN_TOKEN, consts.OWNER_PRIVATE_KEY, mintAmount)
		if err != nil {
			t.Logf("MintPumpkinToken failed, err: %v", err)
			return
		}
	})

}
