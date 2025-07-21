// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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
)

// StakingFactoryMetaData contains all meta data concerning the StakingFactory contract.
var StakingFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFeeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidIndexRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOwnerAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResourceTokenAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRewardTokenAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStakingAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadySupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ProjectCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakingAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"resourceToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"}],\"name\":\"ResourceStakingCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CREATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_feeTokens\",\"type\":\"address[]\"}],\"name\":\"addFeeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"}],\"name\":\"addRewardToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_rewardTokens\",\"type\":\"address[]\"}],\"name\":\"addRewardTokenBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"totalFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"createProject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_resourceTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint8[]\",\"name\":\"_types\",\"type\":\"uint8[]\"},{\"internalType\":\"string\",\"name\":\"_projectId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endBlock\",\"type\":\"uint256\"}],\"name\":\"createResourceStakingBatch\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"stakingAddresses\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"creatorContractByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"creatorContractCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"creatorOfContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingAddress\",\"type\":\"address\"}],\"name\":\"getResourceStakingType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getStakingByIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakingAddress\",\"type\":\"address\"}],\"name\":\"getStakingIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getStakings\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"stakings\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"isRewardToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakingAddress\",\"type\":\"address\"}],\"name\":\"isValidResourceStaking\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToken\",\"type\":\"address\"}],\"name\":\"removeFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"}],\"name\":\"removeRewardToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sBorrowStakingCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sContractCreator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sFeeTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sMaxStakingCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sNonceUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sResourceStakingCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sRewardTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sStakingCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxStakingCount\",\"type\":\"uint256\"}],\"name\":\"setMaxStakingCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"setProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// StakingFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingFactoryMetaData.ABI instead.
var StakingFactoryABI = StakingFactoryMetaData.ABI

// StakingFactory is an auto generated Go binding around an Ethereum contract.
type StakingFactory struct {
	StakingFactoryCaller     // Read-only binding to the contract
	StakingFactoryTransactor // Write-only binding to the contract
	StakingFactoryFilterer   // Log filterer for contract events
}

// StakingFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingFactorySession struct {
	Contract     *StakingFactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingFactoryCallerSession struct {
	Contract *StakingFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StakingFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingFactoryTransactorSession struct {
	Contract     *StakingFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StakingFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingFactoryRaw struct {
	Contract *StakingFactory // Generic contract binding to access the raw methods on
}

// StakingFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingFactoryCallerRaw struct {
	Contract *StakingFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// StakingFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingFactoryTransactorRaw struct {
	Contract *StakingFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingFactory creates a new instance of StakingFactory, bound to a specific deployed contract.
func NewStakingFactory(address common.Address, backend bind.ContractBackend) (*StakingFactory, error) {
	contract, err := bindStakingFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingFactory{StakingFactoryCaller: StakingFactoryCaller{contract: contract}, StakingFactoryTransactor: StakingFactoryTransactor{contract: contract}, StakingFactoryFilterer: StakingFactoryFilterer{contract: contract}}, nil
}

// NewStakingFactoryCaller creates a new read-only instance of StakingFactory, bound to a specific deployed contract.
func NewStakingFactoryCaller(address common.Address, caller bind.ContractCaller) (*StakingFactoryCaller, error) {
	contract, err := bindStakingFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingFactoryCaller{contract: contract}, nil
}

// NewStakingFactoryTransactor creates a new write-only instance of StakingFactory, bound to a specific deployed contract.
func NewStakingFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingFactoryTransactor, error) {
	contract, err := bindStakingFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingFactoryTransactor{contract: contract}, nil
}

// NewStakingFactoryFilterer creates a new log filterer instance of StakingFactory, bound to a specific deployed contract.
func NewStakingFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFactoryFilterer, error) {
	contract, err := bindStakingFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFactoryFilterer{contract: contract}, nil
}

// bindStakingFactory binds a generic wrapper to an already deployed contract.
func bindStakingFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingFactory *StakingFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingFactory.Contract.StakingFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingFactory *StakingFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingFactory.Contract.StakingFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingFactory *StakingFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingFactory.Contract.StakingFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingFactory *StakingFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingFactory *StakingFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingFactory *StakingFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingFactory.Contract.contract.Transact(opts, method, params...)
}

// CREATORROLE is a free data retrieval call binding the contract method 0x8aeda25a.
//
// Solidity: function CREATOR_ROLE() view returns(bytes32)
func (_StakingFactory *StakingFactoryCaller) CREATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "CREATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CREATORROLE is a free data retrieval call binding the contract method 0x8aeda25a.
//
// Solidity: function CREATOR_ROLE() view returns(bytes32)
func (_StakingFactory *StakingFactorySession) CREATORROLE() ([32]byte, error) {
	return _StakingFactory.Contract.CREATORROLE(&_StakingFactory.CallOpts)
}

// CREATORROLE is a free data retrieval call binding the contract method 0x8aeda25a.
//
// Solidity: function CREATOR_ROLE() view returns(bytes32)
func (_StakingFactory *StakingFactoryCallerSession) CREATORROLE() ([32]byte, error) {
	return _StakingFactory.Contract.CREATORROLE(&_StakingFactory.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingFactory *StakingFactoryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingFactory *StakingFactorySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakingFactory.Contract.DEFAULTADMINROLE(&_StakingFactory.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingFactory *StakingFactoryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakingFactory.Contract.DEFAULTADMINROLE(&_StakingFactory.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_StakingFactory *StakingFactoryCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_StakingFactory *StakingFactorySession) OPERATORROLE() ([32]byte, error) {
	return _StakingFactory.Contract.OPERATORROLE(&_StakingFactory.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_StakingFactory *StakingFactoryCallerSession) OPERATORROLE() ([32]byte, error) {
	return _StakingFactory.Contract.OPERATORROLE(&_StakingFactory.CallOpts)
}

// SIGNERROLE is a free data retrieval call binding the contract method 0xa1ebf35d.
//
// Solidity: function SIGNER_ROLE() view returns(bytes32)
func (_StakingFactory *StakingFactoryCaller) SIGNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "SIGNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNERROLE is a free data retrieval call binding the contract method 0xa1ebf35d.
//
// Solidity: function SIGNER_ROLE() view returns(bytes32)
func (_StakingFactory *StakingFactorySession) SIGNERROLE() ([32]byte, error) {
	return _StakingFactory.Contract.SIGNERROLE(&_StakingFactory.CallOpts)
}

// SIGNERROLE is a free data retrieval call binding the contract method 0xa1ebf35d.
//
// Solidity: function SIGNER_ROLE() view returns(bytes32)
func (_StakingFactory *StakingFactoryCallerSession) SIGNERROLE() ([32]byte, error) {
	return _StakingFactory.Contract.SIGNERROLE(&_StakingFactory.CallOpts)
}

// CreatorContractByIndex is a free data retrieval call binding the contract method 0x69ba679c.
//
// Solidity: function creatorContractByIndex(address creator, uint256 index) view returns(address)
func (_StakingFactory *StakingFactoryCaller) CreatorContractByIndex(opts *bind.CallOpts, creator common.Address, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "creatorContractByIndex", creator, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreatorContractByIndex is a free data retrieval call binding the contract method 0x69ba679c.
//
// Solidity: function creatorContractByIndex(address creator, uint256 index) view returns(address)
func (_StakingFactory *StakingFactorySession) CreatorContractByIndex(creator common.Address, index *big.Int) (common.Address, error) {
	return _StakingFactory.Contract.CreatorContractByIndex(&_StakingFactory.CallOpts, creator, index)
}

// CreatorContractByIndex is a free data retrieval call binding the contract method 0x69ba679c.
//
// Solidity: function creatorContractByIndex(address creator, uint256 index) view returns(address)
func (_StakingFactory *StakingFactoryCallerSession) CreatorContractByIndex(creator common.Address, index *big.Int) (common.Address, error) {
	return _StakingFactory.Contract.CreatorContractByIndex(&_StakingFactory.CallOpts, creator, index)
}

// CreatorContractCount is a free data retrieval call binding the contract method 0x571f7ed4.
//
// Solidity: function creatorContractCount(address creator) view returns(uint256)
func (_StakingFactory *StakingFactoryCaller) CreatorContractCount(opts *bind.CallOpts, creator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "creatorContractCount", creator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreatorContractCount is a free data retrieval call binding the contract method 0x571f7ed4.
//
// Solidity: function creatorContractCount(address creator) view returns(uint256)
func (_StakingFactory *StakingFactorySession) CreatorContractCount(creator common.Address) (*big.Int, error) {
	return _StakingFactory.Contract.CreatorContractCount(&_StakingFactory.CallOpts, creator)
}

// CreatorContractCount is a free data retrieval call binding the contract method 0x571f7ed4.
//
// Solidity: function creatorContractCount(address creator) view returns(uint256)
func (_StakingFactory *StakingFactoryCallerSession) CreatorContractCount(creator common.Address) (*big.Int, error) {
	return _StakingFactory.Contract.CreatorContractCount(&_StakingFactory.CallOpts, creator)
}

// CreatorOfContract is a free data retrieval call binding the contract method 0x3d42bb3b.
//
// Solidity: function creatorOfContract(address contractAddress) view returns(address)
func (_StakingFactory *StakingFactoryCaller) CreatorOfContract(opts *bind.CallOpts, contractAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "creatorOfContract", contractAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreatorOfContract is a free data retrieval call binding the contract method 0x3d42bb3b.
//
// Solidity: function creatorOfContract(address contractAddress) view returns(address)
func (_StakingFactory *StakingFactorySession) CreatorOfContract(contractAddress common.Address) (common.Address, error) {
	return _StakingFactory.Contract.CreatorOfContract(&_StakingFactory.CallOpts, contractAddress)
}

// CreatorOfContract is a free data retrieval call binding the contract method 0x3d42bb3b.
//
// Solidity: function creatorOfContract(address contractAddress) view returns(address)
func (_StakingFactory *StakingFactoryCallerSession) CreatorOfContract(contractAddress common.Address) (common.Address, error) {
	return _StakingFactory.Contract.CreatorOfContract(&_StakingFactory.CallOpts, contractAddress)
}

// GetProtocolRatio is a free data retrieval call binding the contract method 0xc26c4dfe.
//
// Solidity: function getProtocolRatio() view returns(uint256)
func (_StakingFactory *StakingFactoryCaller) GetProtocolRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "getProtocolRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtocolRatio is a free data retrieval call binding the contract method 0xc26c4dfe.
//
// Solidity: function getProtocolRatio() view returns(uint256)
func (_StakingFactory *StakingFactorySession) GetProtocolRatio() (*big.Int, error) {
	return _StakingFactory.Contract.GetProtocolRatio(&_StakingFactory.CallOpts)
}

// GetProtocolRatio is a free data retrieval call binding the contract method 0xc26c4dfe.
//
// Solidity: function getProtocolRatio() view returns(uint256)
func (_StakingFactory *StakingFactoryCallerSession) GetProtocolRatio() (*big.Int, error) {
	return _StakingFactory.Contract.GetProtocolRatio(&_StakingFactory.CallOpts)
}

// GetResourceStakingType is a free data retrieval call binding the contract method 0xa1dc238b.
//
// Solidity: function getResourceStakingType(address _stakingAddress) view returns(uint8)
func (_StakingFactory *StakingFactoryCaller) GetResourceStakingType(opts *bind.CallOpts, _stakingAddress common.Address) (uint8, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "getResourceStakingType", _stakingAddress)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetResourceStakingType is a free data retrieval call binding the contract method 0xa1dc238b.
//
// Solidity: function getResourceStakingType(address _stakingAddress) view returns(uint8)
func (_StakingFactory *StakingFactorySession) GetResourceStakingType(_stakingAddress common.Address) (uint8, error) {
	return _StakingFactory.Contract.GetResourceStakingType(&_StakingFactory.CallOpts, _stakingAddress)
}

// GetResourceStakingType is a free data retrieval call binding the contract method 0xa1dc238b.
//
// Solidity: function getResourceStakingType(address _stakingAddress) view returns(uint8)
func (_StakingFactory *StakingFactoryCallerSession) GetResourceStakingType(_stakingAddress common.Address) (uint8, error) {
	return _StakingFactory.Contract.GetResourceStakingType(&_StakingFactory.CallOpts, _stakingAddress)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingFactory *StakingFactoryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingFactory *StakingFactorySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakingFactory.Contract.GetRoleAdmin(&_StakingFactory.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingFactory *StakingFactoryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakingFactory.Contract.GetRoleAdmin(&_StakingFactory.CallOpts, role)
}

// GetStakingByIndex is a free data retrieval call binding the contract method 0x9834fcf0.
//
// Solidity: function getStakingByIndex(uint256 index) view returns(address)
func (_StakingFactory *StakingFactoryCaller) GetStakingByIndex(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "getStakingByIndex", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakingByIndex is a free data retrieval call binding the contract method 0x9834fcf0.
//
// Solidity: function getStakingByIndex(uint256 index) view returns(address)
func (_StakingFactory *StakingFactorySession) GetStakingByIndex(index *big.Int) (common.Address, error) {
	return _StakingFactory.Contract.GetStakingByIndex(&_StakingFactory.CallOpts, index)
}

// GetStakingByIndex is a free data retrieval call binding the contract method 0x9834fcf0.
//
// Solidity: function getStakingByIndex(uint256 index) view returns(address)
func (_StakingFactory *StakingFactoryCallerSession) GetStakingByIndex(index *big.Int) (common.Address, error) {
	return _StakingFactory.Contract.GetStakingByIndex(&_StakingFactory.CallOpts, index)
}

// GetStakingCount is a free data retrieval call binding the contract method 0xe6572f95.
//
// Solidity: function getStakingCount() view returns(uint256)
func (_StakingFactory *StakingFactoryCaller) GetStakingCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "getStakingCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingCount is a free data retrieval call binding the contract method 0xe6572f95.
//
// Solidity: function getStakingCount() view returns(uint256)
func (_StakingFactory *StakingFactorySession) GetStakingCount() (*big.Int, error) {
	return _StakingFactory.Contract.GetStakingCount(&_StakingFactory.CallOpts)
}

// GetStakingCount is a free data retrieval call binding the contract method 0xe6572f95.
//
// Solidity: function getStakingCount() view returns(uint256)
func (_StakingFactory *StakingFactoryCallerSession) GetStakingCount() (*big.Int, error) {
	return _StakingFactory.Contract.GetStakingCount(&_StakingFactory.CallOpts)
}

// GetStakingIndex is a free data retrieval call binding the contract method 0x90cd7adc.
//
// Solidity: function getStakingIndex(address stakingAddress) view returns(uint256)
func (_StakingFactory *StakingFactoryCaller) GetStakingIndex(opts *bind.CallOpts, stakingAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "getStakingIndex", stakingAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingIndex is a free data retrieval call binding the contract method 0x90cd7adc.
//
// Solidity: function getStakingIndex(address stakingAddress) view returns(uint256)
func (_StakingFactory *StakingFactorySession) GetStakingIndex(stakingAddress common.Address) (*big.Int, error) {
	return _StakingFactory.Contract.GetStakingIndex(&_StakingFactory.CallOpts, stakingAddress)
}

// GetStakingIndex is a free data retrieval call binding the contract method 0x90cd7adc.
//
// Solidity: function getStakingIndex(address stakingAddress) view returns(uint256)
func (_StakingFactory *StakingFactoryCallerSession) GetStakingIndex(stakingAddress common.Address) (*big.Int, error) {
	return _StakingFactory.Contract.GetStakingIndex(&_StakingFactory.CallOpts, stakingAddress)
}

// GetStakings is a free data retrieval call binding the contract method 0x301139fe.
//
// Solidity: function getStakings(uint256 startIndex, uint256 endIndex) view returns(address[] stakings)
func (_StakingFactory *StakingFactoryCaller) GetStakings(opts *bind.CallOpts, startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "getStakings", startIndex, endIndex)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStakings is a free data retrieval call binding the contract method 0x301139fe.
//
// Solidity: function getStakings(uint256 startIndex, uint256 endIndex) view returns(address[] stakings)
func (_StakingFactory *StakingFactorySession) GetStakings(startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _StakingFactory.Contract.GetStakings(&_StakingFactory.CallOpts, startIndex, endIndex)
}

// GetStakings is a free data retrieval call binding the contract method 0x301139fe.
//
// Solidity: function getStakings(uint256 startIndex, uint256 endIndex) view returns(address[] stakings)
func (_StakingFactory *StakingFactoryCallerSession) GetStakings(startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _StakingFactory.Contract.GetStakings(&_StakingFactory.CallOpts, startIndex, endIndex)
}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_StakingFactory *StakingFactoryCaller) GetTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "getTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_StakingFactory *StakingFactorySession) GetTreasury() (common.Address, error) {
	return _StakingFactory.Contract.GetTreasury(&_StakingFactory.CallOpts)
}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_StakingFactory *StakingFactoryCallerSession) GetTreasury() (common.Address, error) {
	return _StakingFactory.Contract.GetTreasury(&_StakingFactory.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingFactory *StakingFactoryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingFactory *StakingFactorySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakingFactory.Contract.HasRole(&_StakingFactory.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingFactory *StakingFactoryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakingFactory.Contract.HasRole(&_StakingFactory.CallOpts, role, account)
}

// IsRewardToken is a free data retrieval call binding the contract method 0xb5fd73f8.
//
// Solidity: function isRewardToken(address tokenAddress) view returns(bool)
func (_StakingFactory *StakingFactoryCaller) IsRewardToken(opts *bind.CallOpts, tokenAddress common.Address) (bool, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "isRewardToken", tokenAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRewardToken is a free data retrieval call binding the contract method 0xb5fd73f8.
//
// Solidity: function isRewardToken(address tokenAddress) view returns(bool)
func (_StakingFactory *StakingFactorySession) IsRewardToken(tokenAddress common.Address) (bool, error) {
	return _StakingFactory.Contract.IsRewardToken(&_StakingFactory.CallOpts, tokenAddress)
}

// IsRewardToken is a free data retrieval call binding the contract method 0xb5fd73f8.
//
// Solidity: function isRewardToken(address tokenAddress) view returns(bool)
func (_StakingFactory *StakingFactoryCallerSession) IsRewardToken(tokenAddress common.Address) (bool, error) {
	return _StakingFactory.Contract.IsRewardToken(&_StakingFactory.CallOpts, tokenAddress)
}

// IsValidResourceStaking is a free data retrieval call binding the contract method 0x5c970733.
//
// Solidity: function isValidResourceStaking(address stakingAddress) view returns(bool)
func (_StakingFactory *StakingFactoryCaller) IsValidResourceStaking(opts *bind.CallOpts, stakingAddress common.Address) (bool, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "isValidResourceStaking", stakingAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidResourceStaking is a free data retrieval call binding the contract method 0x5c970733.
//
// Solidity: function isValidResourceStaking(address stakingAddress) view returns(bool)
func (_StakingFactory *StakingFactorySession) IsValidResourceStaking(stakingAddress common.Address) (bool, error) {
	return _StakingFactory.Contract.IsValidResourceStaking(&_StakingFactory.CallOpts, stakingAddress)
}

// IsValidResourceStaking is a free data retrieval call binding the contract method 0x5c970733.
//
// Solidity: function isValidResourceStaking(address stakingAddress) view returns(bool)
func (_StakingFactory *StakingFactoryCallerSession) IsValidResourceStaking(stakingAddress common.Address) (bool, error) {
	return _StakingFactory.Contract.IsValidResourceStaking(&_StakingFactory.CallOpts, stakingAddress)
}

// SAdmin is a free data retrieval call binding the contract method 0x4c39273e.
//
// Solidity: function sAdmin() view returns(address)
func (_StakingFactory *StakingFactoryCaller) SAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "sAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SAdmin is a free data retrieval call binding the contract method 0x4c39273e.
//
// Solidity: function sAdmin() view returns(address)
func (_StakingFactory *StakingFactorySession) SAdmin() (common.Address, error) {
	return _StakingFactory.Contract.SAdmin(&_StakingFactory.CallOpts)
}

// SAdmin is a free data retrieval call binding the contract method 0x4c39273e.
//
// Solidity: function sAdmin() view returns(address)
func (_StakingFactory *StakingFactoryCallerSession) SAdmin() (common.Address, error) {
	return _StakingFactory.Contract.SAdmin(&_StakingFactory.CallOpts)
}

// SBorrowStakingCount is a free data retrieval call binding the contract method 0x59332dfc.
//
// Solidity: function sBorrowStakingCount() view returns(uint256)
func (_StakingFactory *StakingFactoryCaller) SBorrowStakingCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "sBorrowStakingCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SBorrowStakingCount is a free data retrieval call binding the contract method 0x59332dfc.
//
// Solidity: function sBorrowStakingCount() view returns(uint256)
func (_StakingFactory *StakingFactorySession) SBorrowStakingCount() (*big.Int, error) {
	return _StakingFactory.Contract.SBorrowStakingCount(&_StakingFactory.CallOpts)
}

// SBorrowStakingCount is a free data retrieval call binding the contract method 0x59332dfc.
//
// Solidity: function sBorrowStakingCount() view returns(uint256)
func (_StakingFactory *StakingFactoryCallerSession) SBorrowStakingCount() (*big.Int, error) {
	return _StakingFactory.Contract.SBorrowStakingCount(&_StakingFactory.CallOpts)
}

// SContractCreator is a free data retrieval call binding the contract method 0xd0cdf3e8.
//
// Solidity: function sContractCreator(address ) view returns(address)
func (_StakingFactory *StakingFactoryCaller) SContractCreator(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "sContractCreator", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SContractCreator is a free data retrieval call binding the contract method 0xd0cdf3e8.
//
// Solidity: function sContractCreator(address ) view returns(address)
func (_StakingFactory *StakingFactorySession) SContractCreator(arg0 common.Address) (common.Address, error) {
	return _StakingFactory.Contract.SContractCreator(&_StakingFactory.CallOpts, arg0)
}

// SContractCreator is a free data retrieval call binding the contract method 0xd0cdf3e8.
//
// Solidity: function sContractCreator(address ) view returns(address)
func (_StakingFactory *StakingFactoryCallerSession) SContractCreator(arg0 common.Address) (common.Address, error) {
	return _StakingFactory.Contract.SContractCreator(&_StakingFactory.CallOpts, arg0)
}

// SFee is a free data retrieval call binding the contract method 0xf2512348.
//
// Solidity: function sFee() view returns(uint256)
func (_StakingFactory *StakingFactoryCaller) SFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "sFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SFee is a free data retrieval call binding the contract method 0xf2512348.
//
// Solidity: function sFee() view returns(uint256)
func (_StakingFactory *StakingFactorySession) SFee() (*big.Int, error) {
	return _StakingFactory.Contract.SFee(&_StakingFactory.CallOpts)
}

// SFee is a free data retrieval call binding the contract method 0xf2512348.
//
// Solidity: function sFee() view returns(uint256)
func (_StakingFactory *StakingFactoryCallerSession) SFee() (*big.Int, error) {
	return _StakingFactory.Contract.SFee(&_StakingFactory.CallOpts)
}

// SFeeTokens is a free data retrieval call binding the contract method 0x92ca10ad.
//
// Solidity: function sFeeTokens(uint256 ) view returns(address)
func (_StakingFactory *StakingFactoryCaller) SFeeTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "sFeeTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SFeeTokens is a free data retrieval call binding the contract method 0x92ca10ad.
//
// Solidity: function sFeeTokens(uint256 ) view returns(address)
func (_StakingFactory *StakingFactorySession) SFeeTokens(arg0 *big.Int) (common.Address, error) {
	return _StakingFactory.Contract.SFeeTokens(&_StakingFactory.CallOpts, arg0)
}

// SFeeTokens is a free data retrieval call binding the contract method 0x92ca10ad.
//
// Solidity: function sFeeTokens(uint256 ) view returns(address)
func (_StakingFactory *StakingFactoryCallerSession) SFeeTokens(arg0 *big.Int) (common.Address, error) {
	return _StakingFactory.Contract.SFeeTokens(&_StakingFactory.CallOpts, arg0)
}

// SMaxStakingCount is a free data retrieval call binding the contract method 0x221b8282.
//
// Solidity: function sMaxStakingCount() view returns(uint256)
func (_StakingFactory *StakingFactoryCaller) SMaxStakingCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "sMaxStakingCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SMaxStakingCount is a free data retrieval call binding the contract method 0x221b8282.
//
// Solidity: function sMaxStakingCount() view returns(uint256)
func (_StakingFactory *StakingFactorySession) SMaxStakingCount() (*big.Int, error) {
	return _StakingFactory.Contract.SMaxStakingCount(&_StakingFactory.CallOpts)
}

// SMaxStakingCount is a free data retrieval call binding the contract method 0x221b8282.
//
// Solidity: function sMaxStakingCount() view returns(uint256)
func (_StakingFactory *StakingFactoryCallerSession) SMaxStakingCount() (*big.Int, error) {
	return _StakingFactory.Contract.SMaxStakingCount(&_StakingFactory.CallOpts)
}

// SNonceUsed is a free data retrieval call binding the contract method 0xb6c97f82.
//
// Solidity: function sNonceUsed(uint256 ) view returns(bool)
func (_StakingFactory *StakingFactoryCaller) SNonceUsed(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "sNonceUsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SNonceUsed is a free data retrieval call binding the contract method 0xb6c97f82.
//
// Solidity: function sNonceUsed(uint256 ) view returns(bool)
func (_StakingFactory *StakingFactorySession) SNonceUsed(arg0 *big.Int) (bool, error) {
	return _StakingFactory.Contract.SNonceUsed(&_StakingFactory.CallOpts, arg0)
}

// SNonceUsed is a free data retrieval call binding the contract method 0xb6c97f82.
//
// Solidity: function sNonceUsed(uint256 ) view returns(bool)
func (_StakingFactory *StakingFactoryCallerSession) SNonceUsed(arg0 *big.Int) (bool, error) {
	return _StakingFactory.Contract.SNonceUsed(&_StakingFactory.CallOpts, arg0)
}

// SOperator is a free data retrieval call binding the contract method 0xb0f65bdf.
//
// Solidity: function sOperator() view returns(address)
func (_StakingFactory *StakingFactoryCaller) SOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "sOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SOperator is a free data retrieval call binding the contract method 0xb0f65bdf.
//
// Solidity: function sOperator() view returns(address)
func (_StakingFactory *StakingFactorySession) SOperator() (common.Address, error) {
	return _StakingFactory.Contract.SOperator(&_StakingFactory.CallOpts)
}

// SOperator is a free data retrieval call binding the contract method 0xb0f65bdf.
//
// Solidity: function sOperator() view returns(address)
func (_StakingFactory *StakingFactoryCallerSession) SOperator() (common.Address, error) {
	return _StakingFactory.Contract.SOperator(&_StakingFactory.CallOpts)
}

// SResourceStakingCount is a free data retrieval call binding the contract method 0xcda090f9.
//
// Solidity: function sResourceStakingCount() view returns(uint256)
func (_StakingFactory *StakingFactoryCaller) SResourceStakingCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "sResourceStakingCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SResourceStakingCount is a free data retrieval call binding the contract method 0xcda090f9.
//
// Solidity: function sResourceStakingCount() view returns(uint256)
func (_StakingFactory *StakingFactorySession) SResourceStakingCount() (*big.Int, error) {
	return _StakingFactory.Contract.SResourceStakingCount(&_StakingFactory.CallOpts)
}

// SResourceStakingCount is a free data retrieval call binding the contract method 0xcda090f9.
//
// Solidity: function sResourceStakingCount() view returns(uint256)
func (_StakingFactory *StakingFactoryCallerSession) SResourceStakingCount() (*big.Int, error) {
	return _StakingFactory.Contract.SResourceStakingCount(&_StakingFactory.CallOpts)
}

// SRewardTokenCount is a free data retrieval call binding the contract method 0x27d6c9a1.
//
// Solidity: function sRewardTokenCount() view returns(uint256)
func (_StakingFactory *StakingFactoryCaller) SRewardTokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "sRewardTokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SRewardTokenCount is a free data retrieval call binding the contract method 0x27d6c9a1.
//
// Solidity: function sRewardTokenCount() view returns(uint256)
func (_StakingFactory *StakingFactorySession) SRewardTokenCount() (*big.Int, error) {
	return _StakingFactory.Contract.SRewardTokenCount(&_StakingFactory.CallOpts)
}

// SRewardTokenCount is a free data retrieval call binding the contract method 0x27d6c9a1.
//
// Solidity: function sRewardTokenCount() view returns(uint256)
func (_StakingFactory *StakingFactoryCallerSession) SRewardTokenCount() (*big.Int, error) {
	return _StakingFactory.Contract.SRewardTokenCount(&_StakingFactory.CallOpts)
}

// SStakingCount is a free data retrieval call binding the contract method 0x7dac46ca.
//
// Solidity: function sStakingCount() view returns(uint256)
func (_StakingFactory *StakingFactoryCaller) SStakingCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "sStakingCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SStakingCount is a free data retrieval call binding the contract method 0x7dac46ca.
//
// Solidity: function sStakingCount() view returns(uint256)
func (_StakingFactory *StakingFactorySession) SStakingCount() (*big.Int, error) {
	return _StakingFactory.Contract.SStakingCount(&_StakingFactory.CallOpts)
}

// SStakingCount is a free data retrieval call binding the contract method 0x7dac46ca.
//
// Solidity: function sStakingCount() view returns(uint256)
func (_StakingFactory *StakingFactoryCallerSession) SStakingCount() (*big.Int, error) {
	return _StakingFactory.Contract.SStakingCount(&_StakingFactory.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingFactory *StakingFactoryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StakingFactory.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingFactory *StakingFactorySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakingFactory.Contract.SupportsInterface(&_StakingFactory.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingFactory *StakingFactoryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakingFactory.Contract.SupportsInterface(&_StakingFactory.CallOpts, interfaceId)
}

// AddFeeTokens is a paid mutator transaction binding the contract method 0x7b984e18.
//
// Solidity: function addFeeTokens(address[] _feeTokens) returns()
func (_StakingFactory *StakingFactoryTransactor) AddFeeTokens(opts *bind.TransactOpts, _feeTokens []common.Address) (*types.Transaction, error) {
	return _StakingFactory.contract.Transact(opts, "addFeeTokens", _feeTokens)
}

// AddFeeTokens is a paid mutator transaction binding the contract method 0x7b984e18.
//
// Solidity: function addFeeTokens(address[] _feeTokens) returns()
func (_StakingFactory *StakingFactorySession) AddFeeTokens(_feeTokens []common.Address) (*types.Transaction, error) {
	return _StakingFactory.Contract.AddFeeTokens(&_StakingFactory.TransactOpts, _feeTokens)
}

// AddFeeTokens is a paid mutator transaction binding the contract method 0x7b984e18.
//
// Solidity: function addFeeTokens(address[] _feeTokens) returns()
func (_StakingFactory *StakingFactoryTransactorSession) AddFeeTokens(_feeTokens []common.Address) (*types.Transaction, error) {
	return _StakingFactory.Contract.AddFeeTokens(&_StakingFactory.TransactOpts, _feeTokens)
}

// AddRewardToken is a paid mutator transaction binding the contract method 0x1c03e6cc.
//
// Solidity: function addRewardToken(address _rewardToken) returns()
func (_StakingFactory *StakingFactoryTransactor) AddRewardToken(opts *bind.TransactOpts, _rewardToken common.Address) (*types.Transaction, error) {
	return _StakingFactory.contract.Transact(opts, "addRewardToken", _rewardToken)
}

// AddRewardToken is a paid mutator transaction binding the contract method 0x1c03e6cc.
//
// Solidity: function addRewardToken(address _rewardToken) returns()
func (_StakingFactory *StakingFactorySession) AddRewardToken(_rewardToken common.Address) (*types.Transaction, error) {
	return _StakingFactory.Contract.AddRewardToken(&_StakingFactory.TransactOpts, _rewardToken)
}

// AddRewardToken is a paid mutator transaction binding the contract method 0x1c03e6cc.
//
// Solidity: function addRewardToken(address _rewardToken) returns()
func (_StakingFactory *StakingFactoryTransactorSession) AddRewardToken(_rewardToken common.Address) (*types.Transaction, error) {
	return _StakingFactory.Contract.AddRewardToken(&_StakingFactory.TransactOpts, _rewardToken)
}

// AddRewardTokenBatch is a paid mutator transaction binding the contract method 0x77ce4c1d.
//
// Solidity: function addRewardTokenBatch(address[] _rewardTokens) returns()
func (_StakingFactory *StakingFactoryTransactor) AddRewardTokenBatch(opts *bind.TransactOpts, _rewardTokens []common.Address) (*types.Transaction, error) {
	return _StakingFactory.contract.Transact(opts, "addRewardTokenBatch", _rewardTokens)
}

// AddRewardTokenBatch is a paid mutator transaction binding the contract method 0x77ce4c1d.
//
// Solidity: function addRewardTokenBatch(address[] _rewardTokens) returns()
func (_StakingFactory *StakingFactorySession) AddRewardTokenBatch(_rewardTokens []common.Address) (*types.Transaction, error) {
	return _StakingFactory.Contract.AddRewardTokenBatch(&_StakingFactory.TransactOpts, _rewardTokens)
}

// AddRewardTokenBatch is a paid mutator transaction binding the contract method 0x77ce4c1d.
//
// Solidity: function addRewardTokenBatch(address[] _rewardTokens) returns()
func (_StakingFactory *StakingFactoryTransactorSession) AddRewardTokenBatch(_rewardTokens []common.Address) (*types.Transaction, error) {
	return _StakingFactory.Contract.AddRewardTokenBatch(&_StakingFactory.TransactOpts, _rewardTokens)
}

// CreateProject is a paid mutator transaction binding the contract method 0xf86abdf1.
//
// Solidity: function createProject(bytes data, uint256 totalFee, address feeToken, uint256 nonce, bytes signature) returns()
func (_StakingFactory *StakingFactoryTransactor) CreateProject(opts *bind.TransactOpts, data []byte, totalFee *big.Int, feeToken common.Address, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _StakingFactory.contract.Transact(opts, "createProject", data, totalFee, feeToken, nonce, signature)
}

// CreateProject is a paid mutator transaction binding the contract method 0xf86abdf1.
//
// Solidity: function createProject(bytes data, uint256 totalFee, address feeToken, uint256 nonce, bytes signature) returns()
func (_StakingFactory *StakingFactorySession) CreateProject(data []byte, totalFee *big.Int, feeToken common.Address, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _StakingFactory.Contract.CreateProject(&_StakingFactory.TransactOpts, data, totalFee, feeToken, nonce, signature)
}

// CreateProject is a paid mutator transaction binding the contract method 0xf86abdf1.
//
// Solidity: function createProject(bytes data, uint256 totalFee, address feeToken, uint256 nonce, bytes signature) returns()
func (_StakingFactory *StakingFactoryTransactorSession) CreateProject(data []byte, totalFee *big.Int, feeToken common.Address, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _StakingFactory.Contract.CreateProject(&_StakingFactory.TransactOpts, data, totalFee, feeToken, nonce, signature)
}

// CreateResourceStakingBatch is a paid mutator transaction binding the contract method 0x9965aab2.
//
// Solidity: function createResourceStakingBatch(address[] _resourceTokens, uint8[] _types, string _projectId, uint256 _startBlock, uint256 _endBlock) returns(address[] stakingAddresses)
func (_StakingFactory *StakingFactoryTransactor) CreateResourceStakingBatch(opts *bind.TransactOpts, _resourceTokens []common.Address, _types []uint8, _projectId string, _startBlock *big.Int, _endBlock *big.Int) (*types.Transaction, error) {
	return _StakingFactory.contract.Transact(opts, "createResourceStakingBatch", _resourceTokens, _types, _projectId, _startBlock, _endBlock)
}

// CreateResourceStakingBatch is a paid mutator transaction binding the contract method 0x9965aab2.
//
// Solidity: function createResourceStakingBatch(address[] _resourceTokens, uint8[] _types, string _projectId, uint256 _startBlock, uint256 _endBlock) returns(address[] stakingAddresses)
func (_StakingFactory *StakingFactorySession) CreateResourceStakingBatch(_resourceTokens []common.Address, _types []uint8, _projectId string, _startBlock *big.Int, _endBlock *big.Int) (*types.Transaction, error) {
	return _StakingFactory.Contract.CreateResourceStakingBatch(&_StakingFactory.TransactOpts, _resourceTokens, _types, _projectId, _startBlock, _endBlock)
}

// CreateResourceStakingBatch is a paid mutator transaction binding the contract method 0x9965aab2.
//
// Solidity: function createResourceStakingBatch(address[] _resourceTokens, uint8[] _types, string _projectId, uint256 _startBlock, uint256 _endBlock) returns(address[] stakingAddresses)
func (_StakingFactory *StakingFactoryTransactorSession) CreateResourceStakingBatch(_resourceTokens []common.Address, _types []uint8, _projectId string, _startBlock *big.Int, _endBlock *big.Int) (*types.Transaction, error) {
	return _StakingFactory.Contract.CreateResourceStakingBatch(&_StakingFactory.TransactOpts, _resourceTokens, _types, _projectId, _startBlock, _endBlock)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_StakingFactory *StakingFactoryTransactor) EmergencyWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingFactory.contract.Transact(opts, "emergencyWithdraw")
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_StakingFactory *StakingFactorySession) EmergencyWithdraw() (*types.Transaction, error) {
	return _StakingFactory.Contract.EmergencyWithdraw(&_StakingFactory.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_StakingFactory *StakingFactoryTransactorSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _StakingFactory.Contract.EmergencyWithdraw(&_StakingFactory.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingFactory *StakingFactoryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingFactory.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingFactory *StakingFactorySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingFactory.Contract.GrantRole(&_StakingFactory.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingFactory *StakingFactoryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingFactory.Contract.GrantRole(&_StakingFactory.TransactOpts, role, account)
}

// RemoveFeeToken is a paid mutator transaction binding the contract method 0x67923c45.
//
// Solidity: function removeFeeToken(address _feeToken) returns()
func (_StakingFactory *StakingFactoryTransactor) RemoveFeeToken(opts *bind.TransactOpts, _feeToken common.Address) (*types.Transaction, error) {
	return _StakingFactory.contract.Transact(opts, "removeFeeToken", _feeToken)
}

// RemoveFeeToken is a paid mutator transaction binding the contract method 0x67923c45.
//
// Solidity: function removeFeeToken(address _feeToken) returns()
func (_StakingFactory *StakingFactorySession) RemoveFeeToken(_feeToken common.Address) (*types.Transaction, error) {
	return _StakingFactory.Contract.RemoveFeeToken(&_StakingFactory.TransactOpts, _feeToken)
}

// RemoveFeeToken is a paid mutator transaction binding the contract method 0x67923c45.
//
// Solidity: function removeFeeToken(address _feeToken) returns()
func (_StakingFactory *StakingFactoryTransactorSession) RemoveFeeToken(_feeToken common.Address) (*types.Transaction, error) {
	return _StakingFactory.Contract.RemoveFeeToken(&_StakingFactory.TransactOpts, _feeToken)
}

// RemoveRewardToken is a paid mutator transaction binding the contract method 0x3d509c97.
//
// Solidity: function removeRewardToken(address _rewardToken) returns()
func (_StakingFactory *StakingFactoryTransactor) RemoveRewardToken(opts *bind.TransactOpts, _rewardToken common.Address) (*types.Transaction, error) {
	return _StakingFactory.contract.Transact(opts, "removeRewardToken", _rewardToken)
}

// RemoveRewardToken is a paid mutator transaction binding the contract method 0x3d509c97.
//
// Solidity: function removeRewardToken(address _rewardToken) returns()
func (_StakingFactory *StakingFactorySession) RemoveRewardToken(_rewardToken common.Address) (*types.Transaction, error) {
	return _StakingFactory.Contract.RemoveRewardToken(&_StakingFactory.TransactOpts, _rewardToken)
}

// RemoveRewardToken is a paid mutator transaction binding the contract method 0x3d509c97.
//
// Solidity: function removeRewardToken(address _rewardToken) returns()
func (_StakingFactory *StakingFactoryTransactorSession) RemoveRewardToken(_rewardToken common.Address) (*types.Transaction, error) {
	return _StakingFactory.Contract.RemoveRewardToken(&_StakingFactory.TransactOpts, _rewardToken)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StakingFactory *StakingFactoryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StakingFactory.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StakingFactory *StakingFactorySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StakingFactory.Contract.RenounceRole(&_StakingFactory.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StakingFactory *StakingFactoryTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StakingFactory.Contract.RenounceRole(&_StakingFactory.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingFactory *StakingFactoryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingFactory.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingFactory *StakingFactorySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingFactory.Contract.RevokeRole(&_StakingFactory.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingFactory *StakingFactoryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingFactory.Contract.RevokeRole(&_StakingFactory.TransactOpts, role, account)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address admin) returns()
func (_StakingFactory *StakingFactoryTransactor) SetAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _StakingFactory.contract.Transact(opts, "setAdmin", admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address admin) returns()
func (_StakingFactory *StakingFactorySession) SetAdmin(admin common.Address) (*types.Transaction, error) {
	return _StakingFactory.Contract.SetAdmin(&_StakingFactory.TransactOpts, admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address admin) returns()
func (_StakingFactory *StakingFactoryTransactorSession) SetAdmin(admin common.Address) (*types.Transaction, error) {
	return _StakingFactory.Contract.SetAdmin(&_StakingFactory.TransactOpts, admin)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_StakingFactory *StakingFactoryTransactor) SetFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _StakingFactory.contract.Transact(opts, "setFee", newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_StakingFactory *StakingFactorySession) SetFee(newFee *big.Int) (*types.Transaction, error) {
	return _StakingFactory.Contract.SetFee(&_StakingFactory.TransactOpts, newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_StakingFactory *StakingFactoryTransactorSession) SetFee(newFee *big.Int) (*types.Transaction, error) {
	return _StakingFactory.Contract.SetFee(&_StakingFactory.TransactOpts, newFee)
}

// SetMaxStakingCount is a paid mutator transaction binding the contract method 0x8ee108cc.
//
// Solidity: function setMaxStakingCount(uint256 newMaxStakingCount) returns()
func (_StakingFactory *StakingFactoryTransactor) SetMaxStakingCount(opts *bind.TransactOpts, newMaxStakingCount *big.Int) (*types.Transaction, error) {
	return _StakingFactory.contract.Transact(opts, "setMaxStakingCount", newMaxStakingCount)
}

// SetMaxStakingCount is a paid mutator transaction binding the contract method 0x8ee108cc.
//
// Solidity: function setMaxStakingCount(uint256 newMaxStakingCount) returns()
func (_StakingFactory *StakingFactorySession) SetMaxStakingCount(newMaxStakingCount *big.Int) (*types.Transaction, error) {
	return _StakingFactory.Contract.SetMaxStakingCount(&_StakingFactory.TransactOpts, newMaxStakingCount)
}

// SetMaxStakingCount is a paid mutator transaction binding the contract method 0x8ee108cc.
//
// Solidity: function setMaxStakingCount(uint256 newMaxStakingCount) returns()
func (_StakingFactory *StakingFactoryTransactorSession) SetMaxStakingCount(newMaxStakingCount *big.Int) (*types.Transaction, error) {
	return _StakingFactory.Contract.SetMaxStakingCount(&_StakingFactory.TransactOpts, newMaxStakingCount)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address operator) returns()
func (_StakingFactory *StakingFactoryTransactor) SetOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _StakingFactory.contract.Transact(opts, "setOperator", operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address operator) returns()
func (_StakingFactory *StakingFactorySession) SetOperator(operator common.Address) (*types.Transaction, error) {
	return _StakingFactory.Contract.SetOperator(&_StakingFactory.TransactOpts, operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address operator) returns()
func (_StakingFactory *StakingFactoryTransactorSession) SetOperator(operator common.Address) (*types.Transaction, error) {
	return _StakingFactory.Contract.SetOperator(&_StakingFactory.TransactOpts, operator)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 newFee) returns()
func (_StakingFactory *StakingFactoryTransactor) SetProtocolFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _StakingFactory.contract.Transact(opts, "setProtocolFee", newFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 newFee) returns()
func (_StakingFactory *StakingFactorySession) SetProtocolFee(newFee *big.Int) (*types.Transaction, error) {
	return _StakingFactory.Contract.SetProtocolFee(&_StakingFactory.TransactOpts, newFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 newFee) returns()
func (_StakingFactory *StakingFactoryTransactorSession) SetProtocolFee(newFee *big.Int) (*types.Transaction, error) {
	return _StakingFactory.Contract.SetProtocolFee(&_StakingFactory.TransactOpts, newFee)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newTreasury) returns()
func (_StakingFactory *StakingFactoryTransactor) SetTreasury(opts *bind.TransactOpts, newTreasury common.Address) (*types.Transaction, error) {
	return _StakingFactory.contract.Transact(opts, "setTreasury", newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newTreasury) returns()
func (_StakingFactory *StakingFactorySession) SetTreasury(newTreasury common.Address) (*types.Transaction, error) {
	return _StakingFactory.Contract.SetTreasury(&_StakingFactory.TransactOpts, newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newTreasury) returns()
func (_StakingFactory *StakingFactoryTransactorSession) SetTreasury(newTreasury common.Address) (*types.Transaction, error) {
	return _StakingFactory.Contract.SetTreasury(&_StakingFactory.TransactOpts, newTreasury)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingFactory *StakingFactoryTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingFactory.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingFactory *StakingFactorySession) Receive() (*types.Transaction, error) {
	return _StakingFactory.Contract.Receive(&_StakingFactory.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingFactory *StakingFactoryTransactorSession) Receive() (*types.Transaction, error) {
	return _StakingFactory.Contract.Receive(&_StakingFactory.TransactOpts)
}

// StakingFactoryProjectCreatedIterator is returned from FilterProjectCreated and is used to iterate over the raw logs and unpacked data for ProjectCreated events raised by the StakingFactory contract.
type StakingFactoryProjectCreatedIterator struct {
	Event *StakingFactoryProjectCreated // Event containing the contract specifics and raw log

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
func (it *StakingFactoryProjectCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingFactoryProjectCreated)
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
		it.Event = new(StakingFactoryProjectCreated)
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
func (it *StakingFactoryProjectCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingFactoryProjectCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingFactoryProjectCreated represents a ProjectCreated event raised by the StakingFactory contract.
type StakingFactoryProjectCreated struct {
	Creator common.Address
	Fee     *big.Int
	Data    common.Hash
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterProjectCreated is a free log retrieval operation binding the contract event 0xaa42b0224b61e5eebbec75cc96ff9fccb858d50e7383be09949c4db854b9fa53.
//
// Solidity: event ProjectCreated(address indexed creator, uint256 indexed fee, bytes indexed data)
func (_StakingFactory *StakingFactoryFilterer) FilterProjectCreated(opts *bind.FilterOpts, creator []common.Address, fee []*big.Int, data [][]byte) (*StakingFactoryProjectCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}
	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _StakingFactory.contract.FilterLogs(opts, "ProjectCreated", creatorRule, feeRule, dataRule)
	if err != nil {
		return nil, err
	}
	return &StakingFactoryProjectCreatedIterator{contract: _StakingFactory.contract, event: "ProjectCreated", logs: logs, sub: sub}, nil
}

// WatchProjectCreated is a free log subscription operation binding the contract event 0xaa42b0224b61e5eebbec75cc96ff9fccb858d50e7383be09949c4db854b9fa53.
//
// Solidity: event ProjectCreated(address indexed creator, uint256 indexed fee, bytes indexed data)
func (_StakingFactory *StakingFactoryFilterer) WatchProjectCreated(opts *bind.WatchOpts, sink chan<- *StakingFactoryProjectCreated, creator []common.Address, fee []*big.Int, data [][]byte) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}
	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _StakingFactory.contract.WatchLogs(opts, "ProjectCreated", creatorRule, feeRule, dataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingFactoryProjectCreated)
				if err := _StakingFactory.contract.UnpackLog(event, "ProjectCreated", log); err != nil {
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

// ParseProjectCreated is a log parse operation binding the contract event 0xaa42b0224b61e5eebbec75cc96ff9fccb858d50e7383be09949c4db854b9fa53.
//
// Solidity: event ProjectCreated(address indexed creator, uint256 indexed fee, bytes indexed data)
func (_StakingFactory *StakingFactoryFilterer) ParseProjectCreated(log types.Log) (*StakingFactoryProjectCreated, error) {
	event := new(StakingFactoryProjectCreated)
	if err := _StakingFactory.contract.UnpackLog(event, "ProjectCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingFactoryResourceStakingCreatedIterator is returned from FilterResourceStakingCreated and is used to iterate over the raw logs and unpacked data for ResourceStakingCreated events raised by the StakingFactory contract.
type StakingFactoryResourceStakingCreatedIterator struct {
	Event *StakingFactoryResourceStakingCreated // Event containing the contract specifics and raw log

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
func (it *StakingFactoryResourceStakingCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingFactoryResourceStakingCreated)
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
		it.Event = new(StakingFactoryResourceStakingCreated)
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
func (it *StakingFactoryResourceStakingCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingFactoryResourceStakingCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingFactoryResourceStakingCreated represents a ResourceStakingCreated event raised by the StakingFactory contract.
type StakingFactoryResourceStakingCreated struct {
	StakingAddress common.Address
	Creator        common.Address
	ResourceToken  common.Address
	ProjectId      [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterResourceStakingCreated is a free log retrieval operation binding the contract event 0x386600ee437c3d3907fcae583ccbcfb61a0f5a4c47e2262ec8726dcf44ceacc2.
//
// Solidity: event ResourceStakingCreated(address indexed stakingAddress, address indexed creator, address indexed resourceToken, bytes32 projectId)
func (_StakingFactory *StakingFactoryFilterer) FilterResourceStakingCreated(opts *bind.FilterOpts, stakingAddress []common.Address, creator []common.Address, resourceToken []common.Address) (*StakingFactoryResourceStakingCreatedIterator, error) {

	var stakingAddressRule []interface{}
	for _, stakingAddressItem := range stakingAddress {
		stakingAddressRule = append(stakingAddressRule, stakingAddressItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var resourceTokenRule []interface{}
	for _, resourceTokenItem := range resourceToken {
		resourceTokenRule = append(resourceTokenRule, resourceTokenItem)
	}

	logs, sub, err := _StakingFactory.contract.FilterLogs(opts, "ResourceStakingCreated", stakingAddressRule, creatorRule, resourceTokenRule)
	if err != nil {
		return nil, err
	}
	return &StakingFactoryResourceStakingCreatedIterator{contract: _StakingFactory.contract, event: "ResourceStakingCreated", logs: logs, sub: sub}, nil
}

// WatchResourceStakingCreated is a free log subscription operation binding the contract event 0x386600ee437c3d3907fcae583ccbcfb61a0f5a4c47e2262ec8726dcf44ceacc2.
//
// Solidity: event ResourceStakingCreated(address indexed stakingAddress, address indexed creator, address indexed resourceToken, bytes32 projectId)
func (_StakingFactory *StakingFactoryFilterer) WatchResourceStakingCreated(opts *bind.WatchOpts, sink chan<- *StakingFactoryResourceStakingCreated, stakingAddress []common.Address, creator []common.Address, resourceToken []common.Address) (event.Subscription, error) {

	var stakingAddressRule []interface{}
	for _, stakingAddressItem := range stakingAddress {
		stakingAddressRule = append(stakingAddressRule, stakingAddressItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var resourceTokenRule []interface{}
	for _, resourceTokenItem := range resourceToken {
		resourceTokenRule = append(resourceTokenRule, resourceTokenItem)
	}

	logs, sub, err := _StakingFactory.contract.WatchLogs(opts, "ResourceStakingCreated", stakingAddressRule, creatorRule, resourceTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingFactoryResourceStakingCreated)
				if err := _StakingFactory.contract.UnpackLog(event, "ResourceStakingCreated", log); err != nil {
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

// ParseResourceStakingCreated is a log parse operation binding the contract event 0x386600ee437c3d3907fcae583ccbcfb61a0f5a4c47e2262ec8726dcf44ceacc2.
//
// Solidity: event ResourceStakingCreated(address indexed stakingAddress, address indexed creator, address indexed resourceToken, bytes32 projectId)
func (_StakingFactory *StakingFactoryFilterer) ParseResourceStakingCreated(log types.Log) (*StakingFactoryResourceStakingCreated, error) {
	event := new(StakingFactoryResourceStakingCreated)
	if err := _StakingFactory.contract.UnpackLog(event, "ResourceStakingCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingFactoryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the StakingFactory contract.
type StakingFactoryRoleAdminChangedIterator struct {
	Event *StakingFactoryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakingFactoryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingFactoryRoleAdminChanged)
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
		it.Event = new(StakingFactoryRoleAdminChanged)
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
func (it *StakingFactoryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingFactoryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingFactoryRoleAdminChanged represents a RoleAdminChanged event raised by the StakingFactory contract.
type StakingFactoryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakingFactory *StakingFactoryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StakingFactoryRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _StakingFactory.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StakingFactoryRoleAdminChangedIterator{contract: _StakingFactory.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakingFactory *StakingFactoryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StakingFactoryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _StakingFactory.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingFactoryRoleAdminChanged)
				if err := _StakingFactory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakingFactory *StakingFactoryFilterer) ParseRoleAdminChanged(log types.Log) (*StakingFactoryRoleAdminChanged, error) {
	event := new(StakingFactoryRoleAdminChanged)
	if err := _StakingFactory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingFactoryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the StakingFactory contract.
type StakingFactoryRoleGrantedIterator struct {
	Event *StakingFactoryRoleGranted // Event containing the contract specifics and raw log

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
func (it *StakingFactoryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingFactoryRoleGranted)
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
		it.Event = new(StakingFactoryRoleGranted)
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
func (it *StakingFactoryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingFactoryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingFactoryRoleGranted represents a RoleGranted event raised by the StakingFactory contract.
type StakingFactoryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingFactory *StakingFactoryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingFactoryRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StakingFactory.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingFactoryRoleGrantedIterator{contract: _StakingFactory.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingFactory *StakingFactoryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StakingFactoryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StakingFactory.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingFactoryRoleGranted)
				if err := _StakingFactory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingFactory *StakingFactoryFilterer) ParseRoleGranted(log types.Log) (*StakingFactoryRoleGranted, error) {
	event := new(StakingFactoryRoleGranted)
	if err := _StakingFactory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingFactoryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the StakingFactory contract.
type StakingFactoryRoleRevokedIterator struct {
	Event *StakingFactoryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StakingFactoryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingFactoryRoleRevoked)
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
		it.Event = new(StakingFactoryRoleRevoked)
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
func (it *StakingFactoryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingFactoryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingFactoryRoleRevoked represents a RoleRevoked event raised by the StakingFactory contract.
type StakingFactoryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingFactory *StakingFactoryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingFactoryRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StakingFactory.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingFactoryRoleRevokedIterator{contract: _StakingFactory.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingFactory *StakingFactoryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StakingFactoryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StakingFactory.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingFactoryRoleRevoked)
				if err := _StakingFactory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingFactory *StakingFactoryFilterer) ParseRoleRevoked(log types.Log) (*StakingFactoryRoleRevoked, error) {
	event := new(StakingFactoryRoleRevoked)
	if err := _StakingFactory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
