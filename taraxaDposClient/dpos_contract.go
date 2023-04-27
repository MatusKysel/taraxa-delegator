// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package taraxaDposClient

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

// DposInterfaceDelegationData is an auto generated low-level Go binding around an user-defined struct.
type DposInterfaceDelegationData struct {
	Account    common.Address
	Delegation DposInterfaceDelegatorInfo
}

// DposInterfaceDelegatorInfo is an auto generated low-level Go binding around an user-defined struct.
type DposInterfaceDelegatorInfo struct {
	Stake   *big.Int
	Rewards *big.Int
}

// DposInterfaceUndelegationData is an auto generated low-level Go binding around an user-defined struct.
type DposInterfaceUndelegationData struct {
	Stake           *big.Int
	Block           uint64
	Validator       common.Address
	ValidatorExists bool
}

// DposInterfaceValidatorBasicInfo is an auto generated low-level Go binding around an user-defined struct.
type DposInterfaceValidatorBasicInfo struct {
	TotalStake           *big.Int
	CommissionReward     *big.Int
	Commission           uint16
	LastCommissionChange uint64
	Owner                common.Address
	Description          string
	Endpoint             string
}

// DposInterfaceValidatorData is an auto generated low-level Go binding around an user-defined struct.
type DposInterfaceValidatorData struct {
	Account common.Address
	Info    DposInterfaceValidatorBasicInfo
}

// TaraxaDposClientMetaData contains all meta data concerning the TaraxaDposClient contract.
var TaraxaDposClientMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CommissionRewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"commission\",\"type\":\"uint16\"}],\"name\":\"CommissionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UndelegateCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UndelegateConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorInfoSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"cancelUndelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"batch\",\"type\":\"uint32\"}],\"name\":\"claimAllRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"end\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"claimCommissionRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"confirmUndelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"batch\",\"type\":\"uint32\"}],\"name\":\"getDelegations\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"internalType\":\"structDposInterface.DelegatorInfo\",\"name\":\"delegation\",\"type\":\"tuple\"}],\"internalType\":\"structDposInterface.DelegationData[]\",\"name\":\"delegations\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"end\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalEligibleVotesCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"batch\",\"type\":\"uint32\"}],\"name\":\"getUndelegations\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"block\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"validator_exists\",\"type\":\"bool\"}],\"internalType\":\"structDposInterface.UndelegationData[]\",\"name\":\"undelegations\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"end\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"total_stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commission_reward\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"commission\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"last_commission_change\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"internalType\":\"structDposInterface.ValidatorBasicInfo\",\"name\":\"validator_info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getValidatorEligibleVotesCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"batch\",\"type\":\"uint32\"}],\"name\":\"getValidators\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"total_stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commission_reward\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"commission\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"last_commission_change\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"internalType\":\"structDposInterface.ValidatorBasicInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structDposInterface.ValidatorData[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"end\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"batch\",\"type\":\"uint32\"}],\"name\":\"getValidatorsFor\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"total_stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commission_reward\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"commission\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"last_commission_change\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"internalType\":\"structDposInterface.ValidatorBasicInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structDposInterface.ValidatorData[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"end\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidatorEligible\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validator_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"reDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vrf_key\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"commission\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"commission\",\"type\":\"uint16\"}],\"name\":\"setCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"name\":\"setValidatorInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TaraxaDposClientABI is the input ABI used to generate the binding from.
// Deprecated: Use TaraxaDposClientMetaData.ABI instead.
var TaraxaDposClientABI = TaraxaDposClientMetaData.ABI

// TaraxaDposClient is an auto generated Go binding around an Ethereum contract.
type TaraxaDposClient struct {
	TaraxaDposClientCaller     // Read-only binding to the contract
	TaraxaDposClientTransactor // Write-only binding to the contract
	TaraxaDposClientFilterer   // Log filterer for contract events
}

// TaraxaDposClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaraxaDposClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaraxaDposClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaraxaDposClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaraxaDposClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaraxaDposClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaraxaDposClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaraxaDposClientSession struct {
	Contract     *TaraxaDposClient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaraxaDposClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaraxaDposClientCallerSession struct {
	Contract *TaraxaDposClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// TaraxaDposClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaraxaDposClientTransactorSession struct {
	Contract     *TaraxaDposClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TaraxaDposClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaraxaDposClientRaw struct {
	Contract *TaraxaDposClient // Generic contract binding to access the raw methods on
}

// TaraxaDposClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaraxaDposClientCallerRaw struct {
	Contract *TaraxaDposClientCaller // Generic read-only contract binding to access the raw methods on
}

// TaraxaDposClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaraxaDposClientTransactorRaw struct {
	Contract *TaraxaDposClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaraxaDposClient creates a new instance of TaraxaDposClient, bound to a specific deployed contract.
func NewTaraxaDposClient(address common.Address, backend bind.ContractBackend) (*TaraxaDposClient, error) {
	contract, err := bindTaraxaDposClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposClient{TaraxaDposClientCaller: TaraxaDposClientCaller{contract: contract}, TaraxaDposClientTransactor: TaraxaDposClientTransactor{contract: contract}, TaraxaDposClientFilterer: TaraxaDposClientFilterer{contract: contract}}, nil
}

// NewTaraxaDposClientCaller creates a new read-only instance of TaraxaDposClient, bound to a specific deployed contract.
func NewTaraxaDposClientCaller(address common.Address, caller bind.ContractCaller) (*TaraxaDposClientCaller, error) {
	contract, err := bindTaraxaDposClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposClientCaller{contract: contract}, nil
}

// NewTaraxaDposClientTransactor creates a new write-only instance of TaraxaDposClient, bound to a specific deployed contract.
func NewTaraxaDposClientTransactor(address common.Address, transactor bind.ContractTransactor) (*TaraxaDposClientTransactor, error) {
	contract, err := bindTaraxaDposClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposClientTransactor{contract: contract}, nil
}

// NewTaraxaDposClientFilterer creates a new log filterer instance of TaraxaDposClient, bound to a specific deployed contract.
func NewTaraxaDposClientFilterer(address common.Address, filterer bind.ContractFilterer) (*TaraxaDposClientFilterer, error) {
	contract, err := bindTaraxaDposClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposClientFilterer{contract: contract}, nil
}

// bindTaraxaDposClient binds a generic wrapper to an already deployed contract.
func bindTaraxaDposClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaraxaDposClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaraxaDposClient *TaraxaDposClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaraxaDposClient.Contract.TaraxaDposClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaraxaDposClient *TaraxaDposClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.TaraxaDposClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaraxaDposClient *TaraxaDposClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.TaraxaDposClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaraxaDposClient *TaraxaDposClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaraxaDposClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaraxaDposClient *TaraxaDposClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaraxaDposClient *TaraxaDposClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.contract.Transact(opts, method, params...)
}

// GetDelegations is a free data retrieval call binding the contract method 0x8b49d394.
//
// Solidity: function getDelegations(address delegator, uint32 batch) view returns((address,(uint256,uint256))[] delegations, bool end)
func (_TaraxaDposClient *TaraxaDposClientCaller) GetDelegations(opts *bind.CallOpts, delegator common.Address, batch uint32) (struct {
	Delegations []DposInterfaceDelegationData
	End         bool
}, error) {
	var out []interface{}
	err := _TaraxaDposClient.contract.Call(opts, &out, "getDelegations", delegator, batch)

	outstruct := new(struct {
		Delegations []DposInterfaceDelegationData
		End         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Delegations = *abi.ConvertType(out[0], new([]DposInterfaceDelegationData)).(*[]DposInterfaceDelegationData)
	outstruct.End = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetDelegations is a free data retrieval call binding the contract method 0x8b49d394.
//
// Solidity: function getDelegations(address delegator, uint32 batch) view returns((address,(uint256,uint256))[] delegations, bool end)
func (_TaraxaDposClient *TaraxaDposClientSession) GetDelegations(delegator common.Address, batch uint32) (struct {
	Delegations []DposInterfaceDelegationData
	End         bool
}, error) {
	return _TaraxaDposClient.Contract.GetDelegations(&_TaraxaDposClient.CallOpts, delegator, batch)
}

// GetDelegations is a free data retrieval call binding the contract method 0x8b49d394.
//
// Solidity: function getDelegations(address delegator, uint32 batch) view returns((address,(uint256,uint256))[] delegations, bool end)
func (_TaraxaDposClient *TaraxaDposClientCallerSession) GetDelegations(delegator common.Address, batch uint32) (struct {
	Delegations []DposInterfaceDelegationData
	End         bool
}, error) {
	return _TaraxaDposClient.Contract.GetDelegations(&_TaraxaDposClient.CallOpts, delegator, batch)
}

// GetTotalEligibleVotesCount is a free data retrieval call binding the contract method 0xde8e4b50.
//
// Solidity: function getTotalEligibleVotesCount() view returns(uint64)
func (_TaraxaDposClient *TaraxaDposClientCaller) GetTotalEligibleVotesCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _TaraxaDposClient.contract.Call(opts, &out, "getTotalEligibleVotesCount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetTotalEligibleVotesCount is a free data retrieval call binding the contract method 0xde8e4b50.
//
// Solidity: function getTotalEligibleVotesCount() view returns(uint64)
func (_TaraxaDposClient *TaraxaDposClientSession) GetTotalEligibleVotesCount() (uint64, error) {
	return _TaraxaDposClient.Contract.GetTotalEligibleVotesCount(&_TaraxaDposClient.CallOpts)
}

// GetTotalEligibleVotesCount is a free data retrieval call binding the contract method 0xde8e4b50.
//
// Solidity: function getTotalEligibleVotesCount() view returns(uint64)
func (_TaraxaDposClient *TaraxaDposClientCallerSession) GetTotalEligibleVotesCount() (uint64, error) {
	return _TaraxaDposClient.Contract.GetTotalEligibleVotesCount(&_TaraxaDposClient.CallOpts)
}

// GetUndelegations is a free data retrieval call binding the contract method 0x4edd9943.
//
// Solidity: function getUndelegations(address delegator, uint32 batch) view returns((uint256,uint64,address,bool)[] undelegations, bool end)
func (_TaraxaDposClient *TaraxaDposClientCaller) GetUndelegations(opts *bind.CallOpts, delegator common.Address, batch uint32) (struct {
	Undelegations []DposInterfaceUndelegationData
	End           bool
}, error) {
	var out []interface{}
	err := _TaraxaDposClient.contract.Call(opts, &out, "getUndelegations", delegator, batch)

	outstruct := new(struct {
		Undelegations []DposInterfaceUndelegationData
		End           bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Undelegations = *abi.ConvertType(out[0], new([]DposInterfaceUndelegationData)).(*[]DposInterfaceUndelegationData)
	outstruct.End = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetUndelegations is a free data retrieval call binding the contract method 0x4edd9943.
//
// Solidity: function getUndelegations(address delegator, uint32 batch) view returns((uint256,uint64,address,bool)[] undelegations, bool end)
func (_TaraxaDposClient *TaraxaDposClientSession) GetUndelegations(delegator common.Address, batch uint32) (struct {
	Undelegations []DposInterfaceUndelegationData
	End           bool
}, error) {
	return _TaraxaDposClient.Contract.GetUndelegations(&_TaraxaDposClient.CallOpts, delegator, batch)
}

// GetUndelegations is a free data retrieval call binding the contract method 0x4edd9943.
//
// Solidity: function getUndelegations(address delegator, uint32 batch) view returns((uint256,uint64,address,bool)[] undelegations, bool end)
func (_TaraxaDposClient *TaraxaDposClientCallerSession) GetUndelegations(delegator common.Address, batch uint32) (struct {
	Undelegations []DposInterfaceUndelegationData
	End           bool
}, error) {
	return _TaraxaDposClient.Contract.GetUndelegations(&_TaraxaDposClient.CallOpts, delegator, batch)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address validator) view returns((uint256,uint256,uint16,uint64,address,string,string) validator_info)
func (_TaraxaDposClient *TaraxaDposClientCaller) GetValidator(opts *bind.CallOpts, validator common.Address) (DposInterfaceValidatorBasicInfo, error) {
	var out []interface{}
	err := _TaraxaDposClient.contract.Call(opts, &out, "getValidator", validator)

	if err != nil {
		return *new(DposInterfaceValidatorBasicInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(DposInterfaceValidatorBasicInfo)).(*DposInterfaceValidatorBasicInfo)

	return out0, err

}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address validator) view returns((uint256,uint256,uint16,uint64,address,string,string) validator_info)
func (_TaraxaDposClient *TaraxaDposClientSession) GetValidator(validator common.Address) (DposInterfaceValidatorBasicInfo, error) {
	return _TaraxaDposClient.Contract.GetValidator(&_TaraxaDposClient.CallOpts, validator)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address validator) view returns((uint256,uint256,uint16,uint64,address,string,string) validator_info)
func (_TaraxaDposClient *TaraxaDposClientCallerSession) GetValidator(validator common.Address) (DposInterfaceValidatorBasicInfo, error) {
	return _TaraxaDposClient.Contract.GetValidator(&_TaraxaDposClient.CallOpts, validator)
}

// GetValidatorEligibleVotesCount is a free data retrieval call binding the contract method 0x618e3862.
//
// Solidity: function getValidatorEligibleVotesCount(address validator) view returns(uint64)
func (_TaraxaDposClient *TaraxaDposClientCaller) GetValidatorEligibleVotesCount(opts *bind.CallOpts, validator common.Address) (uint64, error) {
	var out []interface{}
	err := _TaraxaDposClient.contract.Call(opts, &out, "getValidatorEligibleVotesCount", validator)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetValidatorEligibleVotesCount is a free data retrieval call binding the contract method 0x618e3862.
//
// Solidity: function getValidatorEligibleVotesCount(address validator) view returns(uint64)
func (_TaraxaDposClient *TaraxaDposClientSession) GetValidatorEligibleVotesCount(validator common.Address) (uint64, error) {
	return _TaraxaDposClient.Contract.GetValidatorEligibleVotesCount(&_TaraxaDposClient.CallOpts, validator)
}

// GetValidatorEligibleVotesCount is a free data retrieval call binding the contract method 0x618e3862.
//
// Solidity: function getValidatorEligibleVotesCount(address validator) view returns(uint64)
func (_TaraxaDposClient *TaraxaDposClientCallerSession) GetValidatorEligibleVotesCount(validator common.Address) (uint64, error) {
	return _TaraxaDposClient.Contract.GetValidatorEligibleVotesCount(&_TaraxaDposClient.CallOpts, validator)
}

// GetValidators is a free data retrieval call binding the contract method 0x19d8024f.
//
// Solidity: function getValidators(uint32 batch) view returns((address,(uint256,uint256,uint16,uint64,address,string,string))[] validators, bool end)
func (_TaraxaDposClient *TaraxaDposClientCaller) GetValidators(opts *bind.CallOpts, batch uint32) (struct {
	Validators []DposInterfaceValidatorData
	End        bool
}, error) {
	var out []interface{}
	err := _TaraxaDposClient.contract.Call(opts, &out, "getValidators", batch)

	outstruct := new(struct {
		Validators []DposInterfaceValidatorData
		End        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validators = *abi.ConvertType(out[0], new([]DposInterfaceValidatorData)).(*[]DposInterfaceValidatorData)
	outstruct.End = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetValidators is a free data retrieval call binding the contract method 0x19d8024f.
//
// Solidity: function getValidators(uint32 batch) view returns((address,(uint256,uint256,uint16,uint64,address,string,string))[] validators, bool end)
func (_TaraxaDposClient *TaraxaDposClientSession) GetValidators(batch uint32) (struct {
	Validators []DposInterfaceValidatorData
	End        bool
}, error) {
	return _TaraxaDposClient.Contract.GetValidators(&_TaraxaDposClient.CallOpts, batch)
}

// GetValidators is a free data retrieval call binding the contract method 0x19d8024f.
//
// Solidity: function getValidators(uint32 batch) view returns((address,(uint256,uint256,uint16,uint64,address,string,string))[] validators, bool end)
func (_TaraxaDposClient *TaraxaDposClientCallerSession) GetValidators(batch uint32) (struct {
	Validators []DposInterfaceValidatorData
	End        bool
}, error) {
	return _TaraxaDposClient.Contract.GetValidators(&_TaraxaDposClient.CallOpts, batch)
}

// GetValidatorsFor is a free data retrieval call binding the contract method 0x724ac6b0.
//
// Solidity: function getValidatorsFor(address owner, uint32 batch) view returns((address,(uint256,uint256,uint16,uint64,address,string,string))[] validators, bool end)
func (_TaraxaDposClient *TaraxaDposClientCaller) GetValidatorsFor(opts *bind.CallOpts, owner common.Address, batch uint32) (struct {
	Validators []DposInterfaceValidatorData
	End        bool
}, error) {
	var out []interface{}
	err := _TaraxaDposClient.contract.Call(opts, &out, "getValidatorsFor", owner, batch)

	outstruct := new(struct {
		Validators []DposInterfaceValidatorData
		End        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validators = *abi.ConvertType(out[0], new([]DposInterfaceValidatorData)).(*[]DposInterfaceValidatorData)
	outstruct.End = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetValidatorsFor is a free data retrieval call binding the contract method 0x724ac6b0.
//
// Solidity: function getValidatorsFor(address owner, uint32 batch) view returns((address,(uint256,uint256,uint16,uint64,address,string,string))[] validators, bool end)
func (_TaraxaDposClient *TaraxaDposClientSession) GetValidatorsFor(owner common.Address, batch uint32) (struct {
	Validators []DposInterfaceValidatorData
	End        bool
}, error) {
	return _TaraxaDposClient.Contract.GetValidatorsFor(&_TaraxaDposClient.CallOpts, owner, batch)
}

// GetValidatorsFor is a free data retrieval call binding the contract method 0x724ac6b0.
//
// Solidity: function getValidatorsFor(address owner, uint32 batch) view returns((address,(uint256,uint256,uint16,uint64,address,string,string))[] validators, bool end)
func (_TaraxaDposClient *TaraxaDposClientCallerSession) GetValidatorsFor(owner common.Address, batch uint32) (struct {
	Validators []DposInterfaceValidatorData
	End        bool
}, error) {
	return _TaraxaDposClient.Contract.GetValidatorsFor(&_TaraxaDposClient.CallOpts, owner, batch)
}

// IsValidatorEligible is a free data retrieval call binding the contract method 0xf3094e90.
//
// Solidity: function isValidatorEligible(address validator) view returns(bool)
func (_TaraxaDposClient *TaraxaDposClientCaller) IsValidatorEligible(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _TaraxaDposClient.contract.Call(opts, &out, "isValidatorEligible", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorEligible is a free data retrieval call binding the contract method 0xf3094e90.
//
// Solidity: function isValidatorEligible(address validator) view returns(bool)
func (_TaraxaDposClient *TaraxaDposClientSession) IsValidatorEligible(validator common.Address) (bool, error) {
	return _TaraxaDposClient.Contract.IsValidatorEligible(&_TaraxaDposClient.CallOpts, validator)
}

// IsValidatorEligible is a free data retrieval call binding the contract method 0xf3094e90.
//
// Solidity: function isValidatorEligible(address validator) view returns(bool)
func (_TaraxaDposClient *TaraxaDposClientCallerSession) IsValidatorEligible(validator common.Address) (bool, error) {
	return _TaraxaDposClient.Contract.IsValidatorEligible(&_TaraxaDposClient.CallOpts, validator)
}

// CancelUndelegate is a paid mutator transaction binding the contract method 0x399ff554.
//
// Solidity: function cancelUndelegate(address validator) returns()
func (_TaraxaDposClient *TaraxaDposClientTransactor) CancelUndelegate(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposClient.contract.Transact(opts, "cancelUndelegate", validator)
}

// CancelUndelegate is a paid mutator transaction binding the contract method 0x399ff554.
//
// Solidity: function cancelUndelegate(address validator) returns()
func (_TaraxaDposClient *TaraxaDposClientSession) CancelUndelegate(validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.CancelUndelegate(&_TaraxaDposClient.TransactOpts, validator)
}

// CancelUndelegate is a paid mutator transaction binding the contract method 0x399ff554.
//
// Solidity: function cancelUndelegate(address validator) returns()
func (_TaraxaDposClient *TaraxaDposClientTransactorSession) CancelUndelegate(validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.CancelUndelegate(&_TaraxaDposClient.TransactOpts, validator)
}

// ClaimAllRewards is a paid mutator transaction binding the contract method 0x09b72e00.
//
// Solidity: function claimAllRewards(uint32 batch) returns(bool end)
func (_TaraxaDposClient *TaraxaDposClientTransactor) ClaimAllRewards(opts *bind.TransactOpts, batch uint32) (*types.Transaction, error) {
	return _TaraxaDposClient.contract.Transact(opts, "claimAllRewards", batch)
}

// ClaimAllRewards is a paid mutator transaction binding the contract method 0x09b72e00.
//
// Solidity: function claimAllRewards(uint32 batch) returns(bool end)
func (_TaraxaDposClient *TaraxaDposClientSession) ClaimAllRewards(batch uint32) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.ClaimAllRewards(&_TaraxaDposClient.TransactOpts, batch)
}

// ClaimAllRewards is a paid mutator transaction binding the contract method 0x09b72e00.
//
// Solidity: function claimAllRewards(uint32 batch) returns(bool end)
func (_TaraxaDposClient *TaraxaDposClientTransactorSession) ClaimAllRewards(batch uint32) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.ClaimAllRewards(&_TaraxaDposClient.TransactOpts, batch)
}

// ClaimCommissionRewards is a paid mutator transaction binding the contract method 0xd0eebfe2.
//
// Solidity: function claimCommissionRewards(address validator) returns()
func (_TaraxaDposClient *TaraxaDposClientTransactor) ClaimCommissionRewards(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposClient.contract.Transact(opts, "claimCommissionRewards", validator)
}

// ClaimCommissionRewards is a paid mutator transaction binding the contract method 0xd0eebfe2.
//
// Solidity: function claimCommissionRewards(address validator) returns()
func (_TaraxaDposClient *TaraxaDposClientSession) ClaimCommissionRewards(validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.ClaimCommissionRewards(&_TaraxaDposClient.TransactOpts, validator)
}

// ClaimCommissionRewards is a paid mutator transaction binding the contract method 0xd0eebfe2.
//
// Solidity: function claimCommissionRewards(address validator) returns()
func (_TaraxaDposClient *TaraxaDposClientTransactorSession) ClaimCommissionRewards(validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.ClaimCommissionRewards(&_TaraxaDposClient.TransactOpts, validator)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address validator) returns()
func (_TaraxaDposClient *TaraxaDposClientTransactor) ClaimRewards(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposClient.contract.Transact(opts, "claimRewards", validator)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address validator) returns()
func (_TaraxaDposClient *TaraxaDposClientSession) ClaimRewards(validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.ClaimRewards(&_TaraxaDposClient.TransactOpts, validator)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address validator) returns()
func (_TaraxaDposClient *TaraxaDposClientTransactorSession) ClaimRewards(validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.ClaimRewards(&_TaraxaDposClient.TransactOpts, validator)
}

// ConfirmUndelegate is a paid mutator transaction binding the contract method 0x45a02561.
//
// Solidity: function confirmUndelegate(address validator) returns()
func (_TaraxaDposClient *TaraxaDposClientTransactor) ConfirmUndelegate(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposClient.contract.Transact(opts, "confirmUndelegate", validator)
}

// ConfirmUndelegate is a paid mutator transaction binding the contract method 0x45a02561.
//
// Solidity: function confirmUndelegate(address validator) returns()
func (_TaraxaDposClient *TaraxaDposClientSession) ConfirmUndelegate(validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.ConfirmUndelegate(&_TaraxaDposClient.TransactOpts, validator)
}

// ConfirmUndelegate is a paid mutator transaction binding the contract method 0x45a02561.
//
// Solidity: function confirmUndelegate(address validator) returns()
func (_TaraxaDposClient *TaraxaDposClientTransactorSession) ConfirmUndelegate(validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.ConfirmUndelegate(&_TaraxaDposClient.TransactOpts, validator)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address validator) payable returns()
func (_TaraxaDposClient *TaraxaDposClientTransactor) Delegate(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposClient.contract.Transact(opts, "delegate", validator)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address validator) payable returns()
func (_TaraxaDposClient *TaraxaDposClientSession) Delegate(validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.Delegate(&_TaraxaDposClient.TransactOpts, validator)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address validator) payable returns()
func (_TaraxaDposClient *TaraxaDposClientTransactorSession) Delegate(validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.Delegate(&_TaraxaDposClient.TransactOpts, validator)
}

// ReDelegate is a paid mutator transaction binding the contract method 0x703812cc.
//
// Solidity: function reDelegate(address validator_from, address validator_to, uint256 amount) returns()
func (_TaraxaDposClient *TaraxaDposClientTransactor) ReDelegate(opts *bind.TransactOpts, validator_from common.Address, validator_to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaraxaDposClient.contract.Transact(opts, "reDelegate", validator_from, validator_to, amount)
}

// ReDelegate is a paid mutator transaction binding the contract method 0x703812cc.
//
// Solidity: function reDelegate(address validator_from, address validator_to, uint256 amount) returns()
func (_TaraxaDposClient *TaraxaDposClientSession) ReDelegate(validator_from common.Address, validator_to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.ReDelegate(&_TaraxaDposClient.TransactOpts, validator_from, validator_to, amount)
}

// ReDelegate is a paid mutator transaction binding the contract method 0x703812cc.
//
// Solidity: function reDelegate(address validator_from, address validator_to, uint256 amount) returns()
func (_TaraxaDposClient *TaraxaDposClientTransactorSession) ReDelegate(validator_from common.Address, validator_to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.ReDelegate(&_TaraxaDposClient.TransactOpts, validator_from, validator_to, amount)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xd6fdc127.
//
// Solidity: function registerValidator(address validator, bytes proof, bytes vrf_key, uint16 commission, string description, string endpoint) payable returns()
func (_TaraxaDposClient *TaraxaDposClientTransactor) RegisterValidator(opts *bind.TransactOpts, validator common.Address, proof []byte, vrf_key []byte, commission uint16, description string, endpoint string) (*types.Transaction, error) {
	return _TaraxaDposClient.contract.Transact(opts, "registerValidator", validator, proof, vrf_key, commission, description, endpoint)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xd6fdc127.
//
// Solidity: function registerValidator(address validator, bytes proof, bytes vrf_key, uint16 commission, string description, string endpoint) payable returns()
func (_TaraxaDposClient *TaraxaDposClientSession) RegisterValidator(validator common.Address, proof []byte, vrf_key []byte, commission uint16, description string, endpoint string) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.RegisterValidator(&_TaraxaDposClient.TransactOpts, validator, proof, vrf_key, commission, description, endpoint)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xd6fdc127.
//
// Solidity: function registerValidator(address validator, bytes proof, bytes vrf_key, uint16 commission, string description, string endpoint) payable returns()
func (_TaraxaDposClient *TaraxaDposClientTransactorSession) RegisterValidator(validator common.Address, proof []byte, vrf_key []byte, commission uint16, description string, endpoint string) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.RegisterValidator(&_TaraxaDposClient.TransactOpts, validator, proof, vrf_key, commission, description, endpoint)
}

// SetCommission is a paid mutator transaction binding the contract method 0xf000322c.
//
// Solidity: function setCommission(address validator, uint16 commission) returns()
func (_TaraxaDposClient *TaraxaDposClientTransactor) SetCommission(opts *bind.TransactOpts, validator common.Address, commission uint16) (*types.Transaction, error) {
	return _TaraxaDposClient.contract.Transact(opts, "setCommission", validator, commission)
}

// SetCommission is a paid mutator transaction binding the contract method 0xf000322c.
//
// Solidity: function setCommission(address validator, uint16 commission) returns()
func (_TaraxaDposClient *TaraxaDposClientSession) SetCommission(validator common.Address, commission uint16) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.SetCommission(&_TaraxaDposClient.TransactOpts, validator, commission)
}

// SetCommission is a paid mutator transaction binding the contract method 0xf000322c.
//
// Solidity: function setCommission(address validator, uint16 commission) returns()
func (_TaraxaDposClient *TaraxaDposClientTransactorSession) SetCommission(validator common.Address, commission uint16) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.SetCommission(&_TaraxaDposClient.TransactOpts, validator, commission)
}

// SetValidatorInfo is a paid mutator transaction binding the contract method 0x0babea4c.
//
// Solidity: function setValidatorInfo(address validator, string description, string endpoint) returns()
func (_TaraxaDposClient *TaraxaDposClientTransactor) SetValidatorInfo(opts *bind.TransactOpts, validator common.Address, description string, endpoint string) (*types.Transaction, error) {
	return _TaraxaDposClient.contract.Transact(opts, "setValidatorInfo", validator, description, endpoint)
}

// SetValidatorInfo is a paid mutator transaction binding the contract method 0x0babea4c.
//
// Solidity: function setValidatorInfo(address validator, string description, string endpoint) returns()
func (_TaraxaDposClient *TaraxaDposClientSession) SetValidatorInfo(validator common.Address, description string, endpoint string) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.SetValidatorInfo(&_TaraxaDposClient.TransactOpts, validator, description, endpoint)
}

// SetValidatorInfo is a paid mutator transaction binding the contract method 0x0babea4c.
//
// Solidity: function setValidatorInfo(address validator, string description, string endpoint) returns()
func (_TaraxaDposClient *TaraxaDposClientTransactorSession) SetValidatorInfo(validator common.Address, description string, endpoint string) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.SetValidatorInfo(&_TaraxaDposClient.TransactOpts, validator, description, endpoint)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validator, uint256 amount) returns()
func (_TaraxaDposClient *TaraxaDposClientTransactor) Undelegate(opts *bind.TransactOpts, validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaraxaDposClient.contract.Transact(opts, "undelegate", validator, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validator, uint256 amount) returns()
func (_TaraxaDposClient *TaraxaDposClientSession) Undelegate(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.Undelegate(&_TaraxaDposClient.TransactOpts, validator, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validator, uint256 amount) returns()
func (_TaraxaDposClient *TaraxaDposClientTransactorSession) Undelegate(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaraxaDposClient.Contract.Undelegate(&_TaraxaDposClient.TransactOpts, validator, amount)
}

// TaraxaDposClientCommissionRewardsClaimedIterator is returned from FilterCommissionRewardsClaimed and is used to iterate over the raw logs and unpacked data for CommissionRewardsClaimed events raised by the TaraxaDposClient contract.
type TaraxaDposClientCommissionRewardsClaimedIterator struct {
	Event *TaraxaDposClientCommissionRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *TaraxaDposClientCommissionRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraxaDposClientCommissionRewardsClaimed)
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
		it.Event = new(TaraxaDposClientCommissionRewardsClaimed)
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
func (it *TaraxaDposClientCommissionRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraxaDposClientCommissionRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraxaDposClientCommissionRewardsClaimed represents a CommissionRewardsClaimed event raised by the TaraxaDposClient contract.
type TaraxaDposClientCommissionRewardsClaimed struct {
	Account   common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCommissionRewardsClaimed is a free log retrieval operation binding the contract event 0xf0ec9e0f6add850a1738c5822244e26ffc3d1f14da7537aa240582b25af12ad0.
//
// Solidity: event CommissionRewardsClaimed(address indexed account, address indexed validator, uint256 amount)
func (_TaraxaDposClient *TaraxaDposClientFilterer) FilterCommissionRewardsClaimed(opts *bind.FilterOpts, account []common.Address, validator []common.Address) (*TaraxaDposClientCommissionRewardsClaimedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposClient.contract.FilterLogs(opts, "CommissionRewardsClaimed", accountRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposClientCommissionRewardsClaimedIterator{contract: _TaraxaDposClient.contract, event: "CommissionRewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchCommissionRewardsClaimed is a free log subscription operation binding the contract event 0xf0ec9e0f6add850a1738c5822244e26ffc3d1f14da7537aa240582b25af12ad0.
//
// Solidity: event CommissionRewardsClaimed(address indexed account, address indexed validator, uint256 amount)
func (_TaraxaDposClient *TaraxaDposClientFilterer) WatchCommissionRewardsClaimed(opts *bind.WatchOpts, sink chan<- *TaraxaDposClientCommissionRewardsClaimed, account []common.Address, validator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposClient.contract.WatchLogs(opts, "CommissionRewardsClaimed", accountRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraxaDposClientCommissionRewardsClaimed)
				if err := _TaraxaDposClient.contract.UnpackLog(event, "CommissionRewardsClaimed", log); err != nil {
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

// ParseCommissionRewardsClaimed is a log parse operation binding the contract event 0xf0ec9e0f6add850a1738c5822244e26ffc3d1f14da7537aa240582b25af12ad0.
//
// Solidity: event CommissionRewardsClaimed(address indexed account, address indexed validator, uint256 amount)
func (_TaraxaDposClient *TaraxaDposClientFilterer) ParseCommissionRewardsClaimed(log types.Log) (*TaraxaDposClientCommissionRewardsClaimed, error) {
	event := new(TaraxaDposClientCommissionRewardsClaimed)
	if err := _TaraxaDposClient.contract.UnpackLog(event, "CommissionRewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraxaDposClientCommissionSetIterator is returned from FilterCommissionSet and is used to iterate over the raw logs and unpacked data for CommissionSet events raised by the TaraxaDposClient contract.
type TaraxaDposClientCommissionSetIterator struct {
	Event *TaraxaDposClientCommissionSet // Event containing the contract specifics and raw log

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
func (it *TaraxaDposClientCommissionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraxaDposClientCommissionSet)
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
		it.Event = new(TaraxaDposClientCommissionSet)
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
func (it *TaraxaDposClientCommissionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraxaDposClientCommissionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraxaDposClientCommissionSet represents a CommissionSet event raised by the TaraxaDposClient contract.
type TaraxaDposClientCommissionSet struct {
	Validator  common.Address
	Commission uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommissionSet is a free log retrieval operation binding the contract event 0xc909daf778d180f43dac53b55d0de934d2f1e0b70412ca274982e4e6e894eb1a.
//
// Solidity: event CommissionSet(address indexed validator, uint16 commission)
func (_TaraxaDposClient *TaraxaDposClientFilterer) FilterCommissionSet(opts *bind.FilterOpts, validator []common.Address) (*TaraxaDposClientCommissionSetIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposClient.contract.FilterLogs(opts, "CommissionSet", validatorRule)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposClientCommissionSetIterator{contract: _TaraxaDposClient.contract, event: "CommissionSet", logs: logs, sub: sub}, nil
}

// WatchCommissionSet is a free log subscription operation binding the contract event 0xc909daf778d180f43dac53b55d0de934d2f1e0b70412ca274982e4e6e894eb1a.
//
// Solidity: event CommissionSet(address indexed validator, uint16 commission)
func (_TaraxaDposClient *TaraxaDposClientFilterer) WatchCommissionSet(opts *bind.WatchOpts, sink chan<- *TaraxaDposClientCommissionSet, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposClient.contract.WatchLogs(opts, "CommissionSet", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraxaDposClientCommissionSet)
				if err := _TaraxaDposClient.contract.UnpackLog(event, "CommissionSet", log); err != nil {
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

// ParseCommissionSet is a log parse operation binding the contract event 0xc909daf778d180f43dac53b55d0de934d2f1e0b70412ca274982e4e6e894eb1a.
//
// Solidity: event CommissionSet(address indexed validator, uint16 commission)
func (_TaraxaDposClient *TaraxaDposClientFilterer) ParseCommissionSet(log types.Log) (*TaraxaDposClientCommissionSet, error) {
	event := new(TaraxaDposClientCommissionSet)
	if err := _TaraxaDposClient.contract.UnpackLog(event, "CommissionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraxaDposClientDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the TaraxaDposClient contract.
type TaraxaDposClientDelegatedIterator struct {
	Event *TaraxaDposClientDelegated // Event containing the contract specifics and raw log

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
func (it *TaraxaDposClientDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraxaDposClientDelegated)
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
		it.Event = new(TaraxaDposClientDelegated)
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
func (it *TaraxaDposClientDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraxaDposClientDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraxaDposClientDelegated represents a Delegated event raised by the TaraxaDposClient contract.
type TaraxaDposClientDelegated struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed delegator, address indexed validator, uint256 amount)
func (_TaraxaDposClient *TaraxaDposClientFilterer) FilterDelegated(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*TaraxaDposClientDelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposClient.contract.FilterLogs(opts, "Delegated", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposClientDelegatedIterator{contract: _TaraxaDposClient.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed delegator, address indexed validator, uint256 amount)
func (_TaraxaDposClient *TaraxaDposClientFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *TaraxaDposClientDelegated, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposClient.contract.WatchLogs(opts, "Delegated", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraxaDposClientDelegated)
				if err := _TaraxaDposClient.contract.UnpackLog(event, "Delegated", log); err != nil {
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

// ParseDelegated is a log parse operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed delegator, address indexed validator, uint256 amount)
func (_TaraxaDposClient *TaraxaDposClientFilterer) ParseDelegated(log types.Log) (*TaraxaDposClientDelegated, error) {
	event := new(TaraxaDposClientDelegated)
	if err := _TaraxaDposClient.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraxaDposClientRedelegatedIterator is returned from FilterRedelegated and is used to iterate over the raw logs and unpacked data for Redelegated events raised by the TaraxaDposClient contract.
type TaraxaDposClientRedelegatedIterator struct {
	Event *TaraxaDposClientRedelegated // Event containing the contract specifics and raw log

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
func (it *TaraxaDposClientRedelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraxaDposClientRedelegated)
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
		it.Event = new(TaraxaDposClientRedelegated)
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
func (it *TaraxaDposClientRedelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraxaDposClientRedelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraxaDposClientRedelegated represents a Redelegated event raised by the TaraxaDposClient contract.
type TaraxaDposClientRedelegated struct {
	Delegator common.Address
	From      common.Address
	To        common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRedelegated is a free log retrieval operation binding the contract event 0x12e144c27d0bad08abc77c66a640b5cf15a03a93f6582f40de6932b033a5fa5e.
//
// Solidity: event Redelegated(address indexed delegator, address indexed from, address indexed to, uint256 amount)
func (_TaraxaDposClient *TaraxaDposClientFilterer) FilterRedelegated(opts *bind.FilterOpts, delegator []common.Address, from []common.Address, to []common.Address) (*TaraxaDposClientRedelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TaraxaDposClient.contract.FilterLogs(opts, "Redelegated", delegatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposClientRedelegatedIterator{contract: _TaraxaDposClient.contract, event: "Redelegated", logs: logs, sub: sub}, nil
}

// WatchRedelegated is a free log subscription operation binding the contract event 0x12e144c27d0bad08abc77c66a640b5cf15a03a93f6582f40de6932b033a5fa5e.
//
// Solidity: event Redelegated(address indexed delegator, address indexed from, address indexed to, uint256 amount)
func (_TaraxaDposClient *TaraxaDposClientFilterer) WatchRedelegated(opts *bind.WatchOpts, sink chan<- *TaraxaDposClientRedelegated, delegator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TaraxaDposClient.contract.WatchLogs(opts, "Redelegated", delegatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraxaDposClientRedelegated)
				if err := _TaraxaDposClient.contract.UnpackLog(event, "Redelegated", log); err != nil {
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

// ParseRedelegated is a log parse operation binding the contract event 0x12e144c27d0bad08abc77c66a640b5cf15a03a93f6582f40de6932b033a5fa5e.
//
// Solidity: event Redelegated(address indexed delegator, address indexed from, address indexed to, uint256 amount)
func (_TaraxaDposClient *TaraxaDposClientFilterer) ParseRedelegated(log types.Log) (*TaraxaDposClientRedelegated, error) {
	event := new(TaraxaDposClientRedelegated)
	if err := _TaraxaDposClient.contract.UnpackLog(event, "Redelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraxaDposClientRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the TaraxaDposClient contract.
type TaraxaDposClientRewardsClaimedIterator struct {
	Event *TaraxaDposClientRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *TaraxaDposClientRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraxaDposClientRewardsClaimed)
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
		it.Event = new(TaraxaDposClientRewardsClaimed)
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
func (it *TaraxaDposClientRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraxaDposClientRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraxaDposClientRewardsClaimed represents a RewardsClaimed event raised by the TaraxaDposClient contract.
type TaraxaDposClientRewardsClaimed struct {
	Account   common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7.
//
// Solidity: event RewardsClaimed(address indexed account, address indexed validator, uint256 amount)
func (_TaraxaDposClient *TaraxaDposClientFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, account []common.Address, validator []common.Address) (*TaraxaDposClientRewardsClaimedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposClient.contract.FilterLogs(opts, "RewardsClaimed", accountRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposClientRewardsClaimedIterator{contract: _TaraxaDposClient.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7.
//
// Solidity: event RewardsClaimed(address indexed account, address indexed validator, uint256 amount)
func (_TaraxaDposClient *TaraxaDposClientFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *TaraxaDposClientRewardsClaimed, account []common.Address, validator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposClient.contract.WatchLogs(opts, "RewardsClaimed", accountRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraxaDposClientRewardsClaimed)
				if err := _TaraxaDposClient.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
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

// ParseRewardsClaimed is a log parse operation binding the contract event 0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7.
//
// Solidity: event RewardsClaimed(address indexed account, address indexed validator, uint256 amount)
func (_TaraxaDposClient *TaraxaDposClientFilterer) ParseRewardsClaimed(log types.Log) (*TaraxaDposClientRewardsClaimed, error) {
	event := new(TaraxaDposClientRewardsClaimed)
	if err := _TaraxaDposClient.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraxaDposClientUndelegateCanceledIterator is returned from FilterUndelegateCanceled and is used to iterate over the raw logs and unpacked data for UndelegateCanceled events raised by the TaraxaDposClient contract.
type TaraxaDposClientUndelegateCanceledIterator struct {
	Event *TaraxaDposClientUndelegateCanceled // Event containing the contract specifics and raw log

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
func (it *TaraxaDposClientUndelegateCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraxaDposClientUndelegateCanceled)
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
		it.Event = new(TaraxaDposClientUndelegateCanceled)
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
func (it *TaraxaDposClientUndelegateCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraxaDposClientUndelegateCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraxaDposClientUndelegateCanceled represents a UndelegateCanceled event raised by the TaraxaDposClient contract.
type TaraxaDposClientUndelegateCanceled struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUndelegateCanceled is a free log retrieval operation binding the contract event 0xfc25f8a919d19f2c2dfce21115718abc9ef2b1e0c9218a488f614c75be4184b7.
//
// Solidity: event UndelegateCanceled(address indexed delegator, address indexed validator, uint256 amount)
func (_TaraxaDposClient *TaraxaDposClientFilterer) FilterUndelegateCanceled(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*TaraxaDposClientUndelegateCanceledIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposClient.contract.FilterLogs(opts, "UndelegateCanceled", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposClientUndelegateCanceledIterator{contract: _TaraxaDposClient.contract, event: "UndelegateCanceled", logs: logs, sub: sub}, nil
}

// WatchUndelegateCanceled is a free log subscription operation binding the contract event 0xfc25f8a919d19f2c2dfce21115718abc9ef2b1e0c9218a488f614c75be4184b7.
//
// Solidity: event UndelegateCanceled(address indexed delegator, address indexed validator, uint256 amount)
func (_TaraxaDposClient *TaraxaDposClientFilterer) WatchUndelegateCanceled(opts *bind.WatchOpts, sink chan<- *TaraxaDposClientUndelegateCanceled, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposClient.contract.WatchLogs(opts, "UndelegateCanceled", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraxaDposClientUndelegateCanceled)
				if err := _TaraxaDposClient.contract.UnpackLog(event, "UndelegateCanceled", log); err != nil {
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

// ParseUndelegateCanceled is a log parse operation binding the contract event 0xfc25f8a919d19f2c2dfce21115718abc9ef2b1e0c9218a488f614c75be4184b7.
//
// Solidity: event UndelegateCanceled(address indexed delegator, address indexed validator, uint256 amount)
func (_TaraxaDposClient *TaraxaDposClientFilterer) ParseUndelegateCanceled(log types.Log) (*TaraxaDposClientUndelegateCanceled, error) {
	event := new(TaraxaDposClientUndelegateCanceled)
	if err := _TaraxaDposClient.contract.UnpackLog(event, "UndelegateCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraxaDposClientUndelegateConfirmedIterator is returned from FilterUndelegateConfirmed and is used to iterate over the raw logs and unpacked data for UndelegateConfirmed events raised by the TaraxaDposClient contract.
type TaraxaDposClientUndelegateConfirmedIterator struct {
	Event *TaraxaDposClientUndelegateConfirmed // Event containing the contract specifics and raw log

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
func (it *TaraxaDposClientUndelegateConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraxaDposClientUndelegateConfirmed)
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
		it.Event = new(TaraxaDposClientUndelegateConfirmed)
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
func (it *TaraxaDposClientUndelegateConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraxaDposClientUndelegateConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraxaDposClientUndelegateConfirmed represents a UndelegateConfirmed event raised by the TaraxaDposClient contract.
type TaraxaDposClientUndelegateConfirmed struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUndelegateConfirmed is a free log retrieval operation binding the contract event 0xf8bef3a6fe3b4c932b5b51c6472a89f171d039f4bacf18cff632208938bf0426.
//
// Solidity: event UndelegateConfirmed(address indexed delegator, address indexed validator, uint256 amount)
func (_TaraxaDposClient *TaraxaDposClientFilterer) FilterUndelegateConfirmed(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*TaraxaDposClientUndelegateConfirmedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposClient.contract.FilterLogs(opts, "UndelegateConfirmed", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposClientUndelegateConfirmedIterator{contract: _TaraxaDposClient.contract, event: "UndelegateConfirmed", logs: logs, sub: sub}, nil
}

// WatchUndelegateConfirmed is a free log subscription operation binding the contract event 0xf8bef3a6fe3b4c932b5b51c6472a89f171d039f4bacf18cff632208938bf0426.
//
// Solidity: event UndelegateConfirmed(address indexed delegator, address indexed validator, uint256 amount)
func (_TaraxaDposClient *TaraxaDposClientFilterer) WatchUndelegateConfirmed(opts *bind.WatchOpts, sink chan<- *TaraxaDposClientUndelegateConfirmed, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposClient.contract.WatchLogs(opts, "UndelegateConfirmed", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraxaDposClientUndelegateConfirmed)
				if err := _TaraxaDposClient.contract.UnpackLog(event, "UndelegateConfirmed", log); err != nil {
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

// ParseUndelegateConfirmed is a log parse operation binding the contract event 0xf8bef3a6fe3b4c932b5b51c6472a89f171d039f4bacf18cff632208938bf0426.
//
// Solidity: event UndelegateConfirmed(address indexed delegator, address indexed validator, uint256 amount)
func (_TaraxaDposClient *TaraxaDposClientFilterer) ParseUndelegateConfirmed(log types.Log) (*TaraxaDposClientUndelegateConfirmed, error) {
	event := new(TaraxaDposClientUndelegateConfirmed)
	if err := _TaraxaDposClient.contract.UnpackLog(event, "UndelegateConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraxaDposClientUndelegatedIterator is returned from FilterUndelegated and is used to iterate over the raw logs and unpacked data for Undelegated events raised by the TaraxaDposClient contract.
type TaraxaDposClientUndelegatedIterator struct {
	Event *TaraxaDposClientUndelegated // Event containing the contract specifics and raw log

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
func (it *TaraxaDposClientUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraxaDposClientUndelegated)
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
		it.Event = new(TaraxaDposClientUndelegated)
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
func (it *TaraxaDposClientUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraxaDposClientUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraxaDposClientUndelegated represents a Undelegated event raised by the TaraxaDposClient contract.
type TaraxaDposClientUndelegated struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUndelegated is a free log retrieval operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed delegator, address indexed validator, uint256 amount)
func (_TaraxaDposClient *TaraxaDposClientFilterer) FilterUndelegated(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*TaraxaDposClientUndelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposClient.contract.FilterLogs(opts, "Undelegated", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposClientUndelegatedIterator{contract: _TaraxaDposClient.contract, event: "Undelegated", logs: logs, sub: sub}, nil
}

// WatchUndelegated is a free log subscription operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed delegator, address indexed validator, uint256 amount)
func (_TaraxaDposClient *TaraxaDposClientFilterer) WatchUndelegated(opts *bind.WatchOpts, sink chan<- *TaraxaDposClientUndelegated, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposClient.contract.WatchLogs(opts, "Undelegated", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraxaDposClientUndelegated)
				if err := _TaraxaDposClient.contract.UnpackLog(event, "Undelegated", log); err != nil {
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

// ParseUndelegated is a log parse operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed delegator, address indexed validator, uint256 amount)
func (_TaraxaDposClient *TaraxaDposClientFilterer) ParseUndelegated(log types.Log) (*TaraxaDposClientUndelegated, error) {
	event := new(TaraxaDposClientUndelegated)
	if err := _TaraxaDposClient.contract.UnpackLog(event, "Undelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraxaDposClientValidatorInfoSetIterator is returned from FilterValidatorInfoSet and is used to iterate over the raw logs and unpacked data for ValidatorInfoSet events raised by the TaraxaDposClient contract.
type TaraxaDposClientValidatorInfoSetIterator struct {
	Event *TaraxaDposClientValidatorInfoSet // Event containing the contract specifics and raw log

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
func (it *TaraxaDposClientValidatorInfoSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraxaDposClientValidatorInfoSet)
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
		it.Event = new(TaraxaDposClientValidatorInfoSet)
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
func (it *TaraxaDposClientValidatorInfoSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraxaDposClientValidatorInfoSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraxaDposClientValidatorInfoSet represents a ValidatorInfoSet event raised by the TaraxaDposClient contract.
type TaraxaDposClientValidatorInfoSet struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorInfoSet is a free log retrieval operation binding the contract event 0x7aa20e1f59764c9066578febd688a51375adbd654aff86cef56593a17a99071d.
//
// Solidity: event ValidatorInfoSet(address indexed validator)
func (_TaraxaDposClient *TaraxaDposClientFilterer) FilterValidatorInfoSet(opts *bind.FilterOpts, validator []common.Address) (*TaraxaDposClientValidatorInfoSetIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposClient.contract.FilterLogs(opts, "ValidatorInfoSet", validatorRule)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposClientValidatorInfoSetIterator{contract: _TaraxaDposClient.contract, event: "ValidatorInfoSet", logs: logs, sub: sub}, nil
}

// WatchValidatorInfoSet is a free log subscription operation binding the contract event 0x7aa20e1f59764c9066578febd688a51375adbd654aff86cef56593a17a99071d.
//
// Solidity: event ValidatorInfoSet(address indexed validator)
func (_TaraxaDposClient *TaraxaDposClientFilterer) WatchValidatorInfoSet(opts *bind.WatchOpts, sink chan<- *TaraxaDposClientValidatorInfoSet, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposClient.contract.WatchLogs(opts, "ValidatorInfoSet", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraxaDposClientValidatorInfoSet)
				if err := _TaraxaDposClient.contract.UnpackLog(event, "ValidatorInfoSet", log); err != nil {
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

// ParseValidatorInfoSet is a log parse operation binding the contract event 0x7aa20e1f59764c9066578febd688a51375adbd654aff86cef56593a17a99071d.
//
// Solidity: event ValidatorInfoSet(address indexed validator)
func (_TaraxaDposClient *TaraxaDposClientFilterer) ParseValidatorInfoSet(log types.Log) (*TaraxaDposClientValidatorInfoSet, error) {
	event := new(TaraxaDposClientValidatorInfoSet)
	if err := _TaraxaDposClient.contract.UnpackLog(event, "ValidatorInfoSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraxaDposClientValidatorRegisteredIterator is returned from FilterValidatorRegistered and is used to iterate over the raw logs and unpacked data for ValidatorRegistered events raised by the TaraxaDposClient contract.
type TaraxaDposClientValidatorRegisteredIterator struct {
	Event *TaraxaDposClientValidatorRegistered // Event containing the contract specifics and raw log

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
func (it *TaraxaDposClientValidatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraxaDposClientValidatorRegistered)
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
		it.Event = new(TaraxaDposClientValidatorRegistered)
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
func (it *TaraxaDposClientValidatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraxaDposClientValidatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraxaDposClientValidatorRegistered represents a ValidatorRegistered event raised by the TaraxaDposClient contract.
type TaraxaDposClientValidatorRegistered struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistered is a free log retrieval operation binding the contract event 0xd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c88853251.
//
// Solidity: event ValidatorRegistered(address indexed validator)
func (_TaraxaDposClient *TaraxaDposClientFilterer) FilterValidatorRegistered(opts *bind.FilterOpts, validator []common.Address) (*TaraxaDposClientValidatorRegisteredIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposClient.contract.FilterLogs(opts, "ValidatorRegistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposClientValidatorRegisteredIterator{contract: _TaraxaDposClient.contract, event: "ValidatorRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistered is a free log subscription operation binding the contract event 0xd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c88853251.
//
// Solidity: event ValidatorRegistered(address indexed validator)
func (_TaraxaDposClient *TaraxaDposClientFilterer) WatchValidatorRegistered(opts *bind.WatchOpts, sink chan<- *TaraxaDposClientValidatorRegistered, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposClient.contract.WatchLogs(opts, "ValidatorRegistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraxaDposClientValidatorRegistered)
				if err := _TaraxaDposClient.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
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

// ParseValidatorRegistered is a log parse operation binding the contract event 0xd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c88853251.
//
// Solidity: event ValidatorRegistered(address indexed validator)
func (_TaraxaDposClient *TaraxaDposClientFilterer) ParseValidatorRegistered(log types.Log) (*TaraxaDposClientValidatorRegistered, error) {
	event := new(TaraxaDposClientValidatorRegistered)
	if err := _TaraxaDposClient.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
