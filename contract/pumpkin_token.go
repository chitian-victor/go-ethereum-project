// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PumpkinTokenMetaData contains all meta data concerning the PumpkinToken contract.
var PumpkinTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526040518060400160405280600781526020017f50756d706b696e00000000000000000000000000000000000000000000000000815250600690816100489190610587565b506040518060400160405280600281526020017f504b0000000000000000000000000000000000000000000000000000000000008152506007908161008d9190610587565b50348015610099575f5ffd5b506040516119a53803806119a583398181016040528101906100bb91906106b4565b80600680546100c99061039d565b80601f01602080910402602001604051908101604052809291908181526020018280546100f59061039d565b80156101405780601f1061011757610100808354040283529160200191610140565b820191905f5260205f20905b81548152906001019060200180831161012357829003601f168201915b5050505050600780546101529061039d565b80601f016020809104026020016040519081016040528092919081815260200182805461017e9061039d565b80156101c95780601f106101a0576101008083540402835291602001916101c9565b820191905f5260205f20905b8154815290600101906020018083116101ac57829003601f168201915b505050505081600390816101dd9190610587565b5080600490816101ed9190610587565b5050505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610260575f6040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161025791906106ee565b60405180910390fd5b61026f8161027660201b60201c565b5050610707565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806103b457607f821691505b6020821081036103c7576103c6610370565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026104297fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826103ee565b61043386836103ee565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61047761047261046d8461044b565b610454565b61044b565b9050919050565b5f819050919050565b6104908361045d565b6104a461049c8261047e565b8484546103fa565b825550505050565b5f5f905090565b6104bb6104ac565b6104c6818484610487565b505050565b5f5b828110156104ec576104e15f8284016104b3565b6001810190506104cd565b505050565b601f82111561053f578282111561053e5761050b816103cd565b610514836103df565b61051d856103df565b602086101561052a575f90505b808301610539828403826104cb565b505050505b5b505050565b5f82821c905092915050565b5f61055f5f1984600802610544565b1980831691505092915050565b5f6105778383610550565b9150826002028217905092915050565b61059082610339565b67ffffffffffffffff8111156105a9576105a8610343565b5b6105b3825461039d565b6105be8282856104f1565b5f60209050601f8311600181146105ef575f84156105dd578287015190505b6105e7858261056c565b86555061064e565b601f1984166105fd866103cd565b5f5b82811015610624578489015182556001820191506020850194506020810190506105ff565b86831015610641578489015161063d601f891682610550565b8355505b6001600288020188555050505b505050505050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6106838261065a565b9050919050565b61069381610679565b811461069d575f5ffd5b50565b5f815190506106ae8161068a565b92915050565b5f602082840312156106c9576106c8610656565b5b5f6106d6848285016106a0565b91505092915050565b6106e881610679565b82525050565b5f6020820190506107015f8301846106df565b92915050565b611291806107145f395ff3fe608060405234801561000f575f5ffd5b50600436106100e8575f3560e01c8063715018a61161008a578063a0712d6811610064578063a0712d6814610238578063a9059cbb14610254578063dd62ed3e14610284578063f2fde38b146102b4576100e8565b8063715018a6146101f25780638da5cb5b146101fc57806395d89b411461021a576100e8565b806323b872dd116100c657806323b872dd14610158578063313ce5671461018857806342966c68146101a657806370a08231146101c2576100e8565b806306fdde03146100ec578063095ea7b31461010a57806318160ddd1461013a575b5f5ffd5b6100f46102d0565b6040516101019190610edf565b60405180910390f35b610124600480360381019061011f9190610f90565b610360565b6040516101319190610fe8565b60405180910390f35b610142610382565b60405161014f9190611010565b60405180910390f35b610172600480360381019061016d9190611029565b61038b565b60405161017f9190610fe8565b60405180910390f35b6101906103b9565b60405161019d9190611094565b60405180910390f35b6101c060048036038101906101bb91906110ad565b6103c1565b005b6101dc60048036038101906101d791906110d8565b610412565b6040516101e99190611010565b60405180910390f35b6101fa610457565b005b61020461046a565b6040516102119190611112565b60405180910390f35b610222610492565b60405161022f9190610edf565b60405180910390f35b610252600480360381019061024d91906110ad565b610522565b005b61026e60048036038101906102699190610f90565b61057b565b60405161027b9190610fe8565b60405180910390f35b61029e6004803603810190610299919061112b565b61059d565b6040516102ab9190611010565b60405180910390f35b6102ce60048036038101906102c991906110d8565b61061f565b005b6060600380546102df90611196565b80601f016020809104026020016040519081016040528092919081815260200182805461030b90611196565b80156103565780601f1061032d57610100808354040283529160200191610356565b820191905f5260205f20905b81548152906001019060200180831161033957829003601f168201915b5050505050905090565b5f5f61036a6106a3565b90506103778185856106aa565b600191505092915050565b5f600254905090565b5f5f6103956106a3565b90506103a28582856106bc565b6103ad85858561074f565b60019150509392505050565b5f6012905090565b6103cb338261083f565b803373ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca560405160405180910390a350565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b61045f6108be565b6104685f610945565b565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600480546104a190611196565b80601f01602080910402602001604051908101604052809291908181526020018280546104cd90611196565b80156105185780601f106104ef57610100808354040283529160200191610518565b820191905f5260205f20905b8154815290600101906020018083116104fb57829003601f168201915b5050505050905090565b61052a6108be565b6105343382610a08565b803373ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d412139688560405160405180910390a350565b5f5f6105856106a3565b905061059281858561074f565b600191505092915050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b6106276108be565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610697575f6040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161068e9190611112565b60405180910390fd5b6106a081610945565b50565b5f33905090565b6106b78383836001610a87565b505050565b5f6106c7848461059d565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811015610749578181101561073a578281836040517ffb8f41b2000000000000000000000000000000000000000000000000000000008152600401610731939291906111c6565b60405180910390fd5b61074884848484035f610a87565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036107bf575f6040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526004016107b69190611112565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361082f575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016108269190611112565b60405180910390fd5b61083a838383610c56565b505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036108af575f6040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526004016108a69190611112565b60405180910390fd5b6108ba825f83610c56565b5050565b6108c66106a3565b73ffffffffffffffffffffffffffffffffffffffff166108e461046a565b73ffffffffffffffffffffffffffffffffffffffff1614610943576109076106a3565b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815260040161093a9190611112565b60405180910390fd5b565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610a78575f6040517fec442f05000000000000000000000000000000000000000000000000000000008152600401610a6f9190611112565b60405180910390fd5b610a835f8383610c56565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610af7575f6040517fe602df05000000000000000000000000000000000000000000000000000000008152600401610aee9190611112565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610b67575f6040517f94280d62000000000000000000000000000000000000000000000000000000008152600401610b5e9190611112565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508015610c50578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610c479190611010565b60405180910390a35b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610ca6578060025f828254610c9a9190611228565b92505081905550610d74565b5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610d2f578381836040517fe450d38c000000000000000000000000000000000000000000000000000000008152600401610d26939291906111c6565b60405180910390fd5b8181035f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610dbb578060025f8282540392505081905550610e05565b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610e629190611010565b60405180910390a3505050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610eb182610e6f565b610ebb8185610e79565b9350610ecb818560208601610e89565b610ed481610e97565b840191505092915050565b5f6020820190508181035f830152610ef78184610ea7565b905092915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610f2c82610f03565b9050919050565b610f3c81610f22565b8114610f46575f5ffd5b50565b5f81359050610f5781610f33565b92915050565b5f819050919050565b610f6f81610f5d565b8114610f79575f5ffd5b50565b5f81359050610f8a81610f66565b92915050565b5f5f60408385031215610fa657610fa5610eff565b5b5f610fb385828601610f49565b9250506020610fc485828601610f7c565b9150509250929050565b5f8115159050919050565b610fe281610fce565b82525050565b5f602082019050610ffb5f830184610fd9565b92915050565b61100a81610f5d565b82525050565b5f6020820190506110235f830184611001565b92915050565b5f5f5f606084860312156110405761103f610eff565b5b5f61104d86828701610f49565b935050602061105e86828701610f49565b925050604061106f86828701610f7c565b9150509250925092565b5f60ff82169050919050565b61108e81611079565b82525050565b5f6020820190506110a75f830184611085565b92915050565b5f602082840312156110c2576110c1610eff565b5b5f6110cf84828501610f7c565b91505092915050565b5f602082840312156110ed576110ec610eff565b5b5f6110fa84828501610f49565b91505092915050565b61110c81610f22565b82525050565b5f6020820190506111255f830184611103565b92915050565b5f5f6040838503121561114157611140610eff565b5b5f61114e85828601610f49565b925050602061115f85828601610f49565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806111ad57607f821691505b6020821081036111c0576111bf611169565b5b50919050565b5f6060820190506111d95f830186611103565b6111e66020830185611001565b6111f36040830184611001565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61123282610f5d565b915061123d83610f5d565b9250828201905080821115611255576112546111fb565b5b9291505056fea2646970667358221220ddb3e5da5066f37431c49e5c4b81f9073476bdd78702797f4caaad85e55006f964736f6c63430008220033",
}

// PumpkinTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use PumpkinTokenMetaData.ABI instead.
var PumpkinTokenABI = PumpkinTokenMetaData.ABI

// PumpkinTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PumpkinTokenMetaData.Bin instead.
var PumpkinTokenBin = PumpkinTokenMetaData.Bin

// DeployPumpkinToken deploys a new Ethereum contract, binding an instance of PumpkinToken to it.
func DeployPumpkinToken(auth *bind.TransactOpts, backend bind.ContractBackend, owner common.Address) (common.Address, *types.Transaction, *PumpkinToken, error) {
	parsed, err := PumpkinTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PumpkinTokenBin), backend, owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PumpkinToken{PumpkinTokenCaller: PumpkinTokenCaller{contract: contract}, PumpkinTokenTransactor: PumpkinTokenTransactor{contract: contract}, PumpkinTokenFilterer: PumpkinTokenFilterer{contract: contract}}, nil
}

// PumpkinToken is an auto generated Go binding around an Ethereum contract.
type PumpkinToken struct {
	PumpkinTokenCaller     // Read-only binding to the contract
	PumpkinTokenTransactor // Write-only binding to the contract
	PumpkinTokenFilterer   // Log filterer for contract events
}

// PumpkinTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type PumpkinTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PumpkinTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PumpkinTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PumpkinTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PumpkinTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PumpkinTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PumpkinTokenSession struct {
	Contract     *PumpkinToken     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PumpkinTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PumpkinTokenCallerSession struct {
	Contract *PumpkinTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PumpkinTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PumpkinTokenTransactorSession struct {
	Contract     *PumpkinTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PumpkinTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type PumpkinTokenRaw struct {
	Contract *PumpkinToken // Generic contract binding to access the raw methods on
}

// PumpkinTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PumpkinTokenCallerRaw struct {
	Contract *PumpkinTokenCaller // Generic read-only contract binding to access the raw methods on
}

// PumpkinTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PumpkinTokenTransactorRaw struct {
	Contract *PumpkinTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPumpkinToken creates a new instance of PumpkinToken, bound to a specific deployed contract.
func NewPumpkinToken(address common.Address, backend bind.ContractBackend) (*PumpkinToken, error) {
	contract, err := bindPumpkinToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PumpkinToken{PumpkinTokenCaller: PumpkinTokenCaller{contract: contract}, PumpkinTokenTransactor: PumpkinTokenTransactor{contract: contract}, PumpkinTokenFilterer: PumpkinTokenFilterer{contract: contract}}, nil
}

// NewPumpkinTokenCaller creates a new read-only instance of PumpkinToken, bound to a specific deployed contract.
func NewPumpkinTokenCaller(address common.Address, caller bind.ContractCaller) (*PumpkinTokenCaller, error) {
	contract, err := bindPumpkinToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PumpkinTokenCaller{contract: contract}, nil
}

// NewPumpkinTokenTransactor creates a new write-only instance of PumpkinToken, bound to a specific deployed contract.
func NewPumpkinTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*PumpkinTokenTransactor, error) {
	contract, err := bindPumpkinToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PumpkinTokenTransactor{contract: contract}, nil
}

// NewPumpkinTokenFilterer creates a new log filterer instance of PumpkinToken, bound to a specific deployed contract.
func NewPumpkinTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*PumpkinTokenFilterer, error) {
	contract, err := bindPumpkinToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PumpkinTokenFilterer{contract: contract}, nil
}

// bindPumpkinToken binds a generic wrapper to an already deployed contract.
func bindPumpkinToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PumpkinTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PumpkinToken *PumpkinTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PumpkinToken.Contract.PumpkinTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PumpkinToken *PumpkinTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PumpkinToken.Contract.PumpkinTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PumpkinToken *PumpkinTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PumpkinToken.Contract.PumpkinTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PumpkinToken *PumpkinTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PumpkinToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PumpkinToken *PumpkinTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PumpkinToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PumpkinToken *PumpkinTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PumpkinToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PumpkinToken *PumpkinTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PumpkinToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PumpkinToken *PumpkinTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PumpkinToken.Contract.Allowance(&_PumpkinToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PumpkinToken *PumpkinTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PumpkinToken.Contract.Allowance(&_PumpkinToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PumpkinToken *PumpkinTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PumpkinToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PumpkinToken *PumpkinTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PumpkinToken.Contract.BalanceOf(&_PumpkinToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PumpkinToken *PumpkinTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PumpkinToken.Contract.BalanceOf(&_PumpkinToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PumpkinToken *PumpkinTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PumpkinToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PumpkinToken *PumpkinTokenSession) Decimals() (uint8, error) {
	return _PumpkinToken.Contract.Decimals(&_PumpkinToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PumpkinToken *PumpkinTokenCallerSession) Decimals() (uint8, error) {
	return _PumpkinToken.Contract.Decimals(&_PumpkinToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PumpkinToken *PumpkinTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PumpkinToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PumpkinToken *PumpkinTokenSession) Name() (string, error) {
	return _PumpkinToken.Contract.Name(&_PumpkinToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PumpkinToken *PumpkinTokenCallerSession) Name() (string, error) {
	return _PumpkinToken.Contract.Name(&_PumpkinToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PumpkinToken *PumpkinTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PumpkinToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PumpkinToken *PumpkinTokenSession) Owner() (common.Address, error) {
	return _PumpkinToken.Contract.Owner(&_PumpkinToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PumpkinToken *PumpkinTokenCallerSession) Owner() (common.Address, error) {
	return _PumpkinToken.Contract.Owner(&_PumpkinToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PumpkinToken *PumpkinTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PumpkinToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PumpkinToken *PumpkinTokenSession) Symbol() (string, error) {
	return _PumpkinToken.Contract.Symbol(&_PumpkinToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PumpkinToken *PumpkinTokenCallerSession) Symbol() (string, error) {
	return _PumpkinToken.Contract.Symbol(&_PumpkinToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PumpkinToken *PumpkinTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PumpkinToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PumpkinToken *PumpkinTokenSession) TotalSupply() (*big.Int, error) {
	return _PumpkinToken.Contract.TotalSupply(&_PumpkinToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PumpkinToken *PumpkinTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _PumpkinToken.Contract.TotalSupply(&_PumpkinToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_PumpkinToken *PumpkinTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PumpkinToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_PumpkinToken *PumpkinTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PumpkinToken.Contract.Approve(&_PumpkinToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_PumpkinToken *PumpkinTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PumpkinToken.Contract.Approve(&_PumpkinToken.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_PumpkinToken *PumpkinTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _PumpkinToken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_PumpkinToken *PumpkinTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _PumpkinToken.Contract.Burn(&_PumpkinToken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_PumpkinToken *PumpkinTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _PumpkinToken.Contract.Burn(&_PumpkinToken.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_PumpkinToken *PumpkinTokenTransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _PumpkinToken.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_PumpkinToken *PumpkinTokenSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _PumpkinToken.Contract.Mint(&_PumpkinToken.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_PumpkinToken *PumpkinTokenTransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _PumpkinToken.Contract.Mint(&_PumpkinToken.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PumpkinToken *PumpkinTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PumpkinToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PumpkinToken *PumpkinTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _PumpkinToken.Contract.RenounceOwnership(&_PumpkinToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PumpkinToken *PumpkinTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PumpkinToken.Contract.RenounceOwnership(&_PumpkinToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_PumpkinToken *PumpkinTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PumpkinToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_PumpkinToken *PumpkinTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PumpkinToken.Contract.Transfer(&_PumpkinToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_PumpkinToken *PumpkinTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PumpkinToken.Contract.Transfer(&_PumpkinToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_PumpkinToken *PumpkinTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PumpkinToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_PumpkinToken *PumpkinTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PumpkinToken.Contract.TransferFrom(&_PumpkinToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_PumpkinToken *PumpkinTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PumpkinToken.Contract.TransferFrom(&_PumpkinToken.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PumpkinToken *PumpkinTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PumpkinToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PumpkinToken *PumpkinTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PumpkinToken.Contract.TransferOwnership(&_PumpkinToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PumpkinToken *PumpkinTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PumpkinToken.Contract.TransferOwnership(&_PumpkinToken.TransactOpts, newOwner)
}

// PumpkinTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PumpkinToken contract.
type PumpkinTokenApprovalIterator struct {
	Event *PumpkinTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PumpkinTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PumpkinTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PumpkinTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PumpkinTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PumpkinTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PumpkinTokenApproval represents a Approval event raised by the PumpkinToken contract.
type PumpkinTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PumpkinToken *PumpkinTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PumpkinTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PumpkinToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PumpkinTokenApprovalIterator{contract: _PumpkinToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PumpkinToken *PumpkinTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PumpkinTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PumpkinToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PumpkinTokenApproval)
				if err := _PumpkinToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PumpkinToken *PumpkinTokenFilterer) ParseApproval(log types.Log) (*PumpkinTokenApproval, error) {
	event := new(PumpkinTokenApproval)
	if err := _PumpkinToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PumpkinTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the PumpkinToken contract.
type PumpkinTokenBurnIterator struct {
	Event *PumpkinTokenBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PumpkinTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PumpkinTokenBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PumpkinTokenBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PumpkinTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PumpkinTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PumpkinTokenBurn represents a Burn event raised by the PumpkinToken contract.
type PumpkinTokenBurn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed user, uint256 indexed amount)
func (_PumpkinToken *PumpkinTokenFilterer) FilterBurn(opts *bind.FilterOpts, user []common.Address, amount []*big.Int) (*PumpkinTokenBurnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _PumpkinToken.contract.FilterLogs(opts, "Burn", userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &PumpkinTokenBurnIterator{contract: _PumpkinToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed user, uint256 indexed amount)
func (_PumpkinToken *PumpkinTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *PumpkinTokenBurn, user []common.Address, amount []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _PumpkinToken.contract.WatchLogs(opts, "Burn", userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PumpkinTokenBurn)
				if err := _PumpkinToken.contract.UnpackLog(event, "Burn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed user, uint256 indexed amount)
func (_PumpkinToken *PumpkinTokenFilterer) ParseBurn(log types.Log) (*PumpkinTokenBurn, error) {
	event := new(PumpkinTokenBurn)
	if err := _PumpkinToken.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PumpkinTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the PumpkinToken contract.
type PumpkinTokenMintIterator struct {
	Event *PumpkinTokenMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PumpkinTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PumpkinTokenMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PumpkinTokenMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PumpkinTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PumpkinTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PumpkinTokenMint represents a Mint event raised by the PumpkinToken contract.
type PumpkinTokenMint struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed user, uint256 indexed amount)
func (_PumpkinToken *PumpkinTokenFilterer) FilterMint(opts *bind.FilterOpts, user []common.Address, amount []*big.Int) (*PumpkinTokenMintIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _PumpkinToken.contract.FilterLogs(opts, "Mint", userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &PumpkinTokenMintIterator{contract: _PumpkinToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed user, uint256 indexed amount)
func (_PumpkinToken *PumpkinTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *PumpkinTokenMint, user []common.Address, amount []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _PumpkinToken.contract.WatchLogs(opts, "Mint", userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PumpkinTokenMint)
				if err := _PumpkinToken.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed user, uint256 indexed amount)
func (_PumpkinToken *PumpkinTokenFilterer) ParseMint(log types.Log) (*PumpkinTokenMint, error) {
	event := new(PumpkinTokenMint)
	if err := _PumpkinToken.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PumpkinTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PumpkinToken contract.
type PumpkinTokenOwnershipTransferredIterator struct {
	Event *PumpkinTokenOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PumpkinTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PumpkinTokenOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PumpkinTokenOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PumpkinTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PumpkinTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PumpkinTokenOwnershipTransferred represents a OwnershipTransferred event raised by the PumpkinToken contract.
type PumpkinTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PumpkinToken *PumpkinTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PumpkinTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PumpkinToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PumpkinTokenOwnershipTransferredIterator{contract: _PumpkinToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PumpkinToken *PumpkinTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PumpkinTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PumpkinToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PumpkinTokenOwnershipTransferred)
				if err := _PumpkinToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PumpkinToken *PumpkinTokenFilterer) ParseOwnershipTransferred(log types.Log) (*PumpkinTokenOwnershipTransferred, error) {
	event := new(PumpkinTokenOwnershipTransferred)
	if err := _PumpkinToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PumpkinTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PumpkinToken contract.
type PumpkinTokenTransferIterator struct {
	Event *PumpkinTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PumpkinTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PumpkinTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PumpkinTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PumpkinTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PumpkinTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PumpkinTokenTransfer represents a Transfer event raised by the PumpkinToken contract.
type PumpkinTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PumpkinToken *PumpkinTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PumpkinTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PumpkinToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PumpkinTokenTransferIterator{contract: _PumpkinToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PumpkinToken *PumpkinTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PumpkinTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PumpkinToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PumpkinTokenTransfer)
				if err := _PumpkinToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PumpkinToken *PumpkinTokenFilterer) ParseTransfer(log types.Log) (*PumpkinTokenTransfer, error) {
	event := new(PumpkinTokenTransfer)
	if err := _PumpkinToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
