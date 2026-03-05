package utils

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/params"
)

func WeiToEther(wei *big.Int) *big.Float {
	weiPerEther := new(big.Float).SetInt(big.NewInt(params.Ether))
	return new(big.Float).Quo(new(big.Float).SetInt(wei), weiPerEther)
}

func ToJson(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}
