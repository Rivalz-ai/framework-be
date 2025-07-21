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
	_ = abi.ConvertType
)

// WRizClaimMetaData contains all meta data concerning the WRizClaim contract.
var WRizClaimMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wRizAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientPayment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonceUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerExists\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SignerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SignerRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newSigners\",\"type\":\"address[]\"}],\"name\":\"addSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signersToRemove\",\"type\":\"address[]\"}],\"name\":\"removeSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usedNonces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wRiz\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// WRizClaimABI is the input ABI used to generate the binding from.
// Deprecated: Use WRizClaimMetaData.ABI instead.
var WRizClaimABI = WRizClaimMetaData.ABI

// WRizClaim is an auto generated Go binding around an Ethereum contract.
type WRizClaim struct {
	WRizClaimCaller     // Read-only binding to the contract
	WRizClaimTransactor // Write-only binding to the contract
	WRizClaimFilterer   // Log filterer for contract events
}

// WRizClaimCaller is an auto generated read-only Go binding around an Ethereum contract.
type WRizClaimCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WRizClaimTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WRizClaimTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WRizClaimFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WRizClaimFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WRizClaimSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WRizClaimSession struct {
	Contract     *WRizClaim        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WRizClaimCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WRizClaimCallerSession struct {
	Contract *WRizClaimCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// WRizClaimTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WRizClaimTransactorSession struct {
	Contract     *WRizClaimTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// WRizClaimRaw is an auto generated low-level Go binding around an Ethereum contract.
type WRizClaimRaw struct {
	Contract *WRizClaim // Generic contract binding to access the raw methods on
}

// WRizClaimCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WRizClaimCallerRaw struct {
	Contract *WRizClaimCaller // Generic read-only contract binding to access the raw methods on
}

// WRizClaimTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WRizClaimTransactorRaw struct {
	Contract *WRizClaimTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWRizClaim creates a new instance of WRizClaim, bound to a specific deployed contract.
func NewWRizClaim(address common.Address, backend bind.ContractBackend) (*WRizClaim, error) {
	contract, err := bindWRizClaim(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WRizClaim{WRizClaimCaller: WRizClaimCaller{contract: contract}, WRizClaimTransactor: WRizClaimTransactor{contract: contract}, WRizClaimFilterer: WRizClaimFilterer{contract: contract}}, nil
}

// NewWRizClaimCaller creates a new read-only instance of WRizClaim, bound to a specific deployed contract.
func NewWRizClaimCaller(address common.Address, caller bind.ContractCaller) (*WRizClaimCaller, error) {
	contract, err := bindWRizClaim(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WRizClaimCaller{contract: contract}, nil
}

// NewWRizClaimTransactor creates a new write-only instance of WRizClaim, bound to a specific deployed contract.
func NewWRizClaimTransactor(address common.Address, transactor bind.ContractTransactor) (*WRizClaimTransactor, error) {
	contract, err := bindWRizClaim(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WRizClaimTransactor{contract: contract}, nil
}

// NewWRizClaimFilterer creates a new log filterer instance of WRizClaim, bound to a specific deployed contract.
func NewWRizClaimFilterer(address common.Address, filterer bind.ContractFilterer) (*WRizClaimFilterer, error) {
	contract, err := bindWRizClaim(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WRizClaimFilterer{contract: contract}, nil
}

// bindWRizClaim binds a generic wrapper to an already deployed contract.
func bindWRizClaim(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WRizClaimMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WRizClaim *WRizClaimRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WRizClaim.Contract.WRizClaimCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WRizClaim *WRizClaimRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WRizClaim.Contract.WRizClaimTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WRizClaim *WRizClaimRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WRizClaim.Contract.WRizClaimTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WRizClaim *WRizClaimCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WRizClaim.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WRizClaim *WRizClaimTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WRizClaim.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WRizClaim *WRizClaimTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WRizClaim.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_WRizClaim *WRizClaimCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WRizClaim.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_WRizClaim *WRizClaimSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _WRizClaim.Contract.DEFAULTADMINROLE(&_WRizClaim.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_WRizClaim *WRizClaimCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _WRizClaim.Contract.DEFAULTADMINROLE(&_WRizClaim.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_WRizClaim *WRizClaimCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WRizClaim.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_WRizClaim *WRizClaimSession) OPERATORROLE() ([32]byte, error) {
	return _WRizClaim.Contract.OPERATORROLE(&_WRizClaim.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_WRizClaim *WRizClaimCallerSession) OPERATORROLE() ([32]byte, error) {
	return _WRizClaim.Contract.OPERATORROLE(&_WRizClaim.CallOpts)
}

// ClaimedAmount is a free data retrieval call binding the contract method 0x04e86903.
//
// Solidity: function claimedAmount(address ) view returns(uint256)
func (_WRizClaim *WRizClaimCaller) ClaimedAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WRizClaim.contract.Call(opts, &out, "claimedAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedAmount is a free data retrieval call binding the contract method 0x04e86903.
//
// Solidity: function claimedAmount(address ) view returns(uint256)
func (_WRizClaim *WRizClaimSession) ClaimedAmount(arg0 common.Address) (*big.Int, error) {
	return _WRizClaim.Contract.ClaimedAmount(&_WRizClaim.CallOpts, arg0)
}

// ClaimedAmount is a free data retrieval call binding the contract method 0x04e86903.
//
// Solidity: function claimedAmount(address ) view returns(uint256)
func (_WRizClaim *WRizClaimCallerSession) ClaimedAmount(arg0 common.Address) (*big.Int, error) {
	return _WRizClaim.Contract.ClaimedAmount(&_WRizClaim.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_WRizClaim *WRizClaimCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _WRizClaim.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_WRizClaim *WRizClaimSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _WRizClaim.Contract.GetRoleAdmin(&_WRizClaim.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_WRizClaim *WRizClaimCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _WRizClaim.Contract.GetRoleAdmin(&_WRizClaim.CallOpts, role)
}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_WRizClaim *WRizClaimCaller) GetSigners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _WRizClaim.contract.Call(opts, &out, "getSigners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_WRizClaim *WRizClaimSession) GetSigners() ([]common.Address, error) {
	return _WRizClaim.Contract.GetSigners(&_WRizClaim.CallOpts)
}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[])
func (_WRizClaim *WRizClaimCallerSession) GetSigners() ([]common.Address, error) {
	return _WRizClaim.Contract.GetSigners(&_WRizClaim.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_WRizClaim *WRizClaimCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _WRizClaim.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_WRizClaim *WRizClaimSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _WRizClaim.Contract.HasRole(&_WRizClaim.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_WRizClaim *WRizClaimCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _WRizClaim.Contract.HasRole(&_WRizClaim.CallOpts, role, account)
}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) view returns(address)
func (_WRizClaim *WRizClaimCaller) Signers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _WRizClaim.contract.Call(opts, &out, "signers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) view returns(address)
func (_WRizClaim *WRizClaimSession) Signers(arg0 *big.Int) (common.Address, error) {
	return _WRizClaim.Contract.Signers(&_WRizClaim.CallOpts, arg0)
}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) view returns(address)
func (_WRizClaim *WRizClaimCallerSession) Signers(arg0 *big.Int) (common.Address, error) {
	return _WRizClaim.Contract.Signers(&_WRizClaim.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_WRizClaim *WRizClaimCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _WRizClaim.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_WRizClaim *WRizClaimSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WRizClaim.Contract.SupportsInterface(&_WRizClaim.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_WRizClaim *WRizClaimCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WRizClaim.Contract.SupportsInterface(&_WRizClaim.CallOpts, interfaceId)
}

// UsedNonces is a free data retrieval call binding the contract method 0x6717e41c.
//
// Solidity: function usedNonces(uint256 ) view returns(bool)
func (_WRizClaim *WRizClaimCaller) UsedNonces(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _WRizClaim.contract.Call(opts, &out, "usedNonces", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedNonces is a free data retrieval call binding the contract method 0x6717e41c.
//
// Solidity: function usedNonces(uint256 ) view returns(bool)
func (_WRizClaim *WRizClaimSession) UsedNonces(arg0 *big.Int) (bool, error) {
	return _WRizClaim.Contract.UsedNonces(&_WRizClaim.CallOpts, arg0)
}

// UsedNonces is a free data retrieval call binding the contract method 0x6717e41c.
//
// Solidity: function usedNonces(uint256 ) view returns(bool)
func (_WRizClaim *WRizClaimCallerSession) UsedNonces(arg0 *big.Int) (bool, error) {
	return _WRizClaim.Contract.UsedNonces(&_WRizClaim.CallOpts, arg0)
}

// WRiz is a free data retrieval call binding the contract method 0x812ccf7c.
//
// Solidity: function wRiz() view returns(address)
func (_WRizClaim *WRizClaimCaller) WRiz(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WRizClaim.contract.Call(opts, &out, "wRiz")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WRiz is a free data retrieval call binding the contract method 0x812ccf7c.
//
// Solidity: function wRiz() view returns(address)
func (_WRizClaim *WRizClaimSession) WRiz() (common.Address, error) {
	return _WRizClaim.Contract.WRiz(&_WRizClaim.CallOpts)
}

// WRiz is a free data retrieval call binding the contract method 0x812ccf7c.
//
// Solidity: function wRiz() view returns(address)
func (_WRizClaim *WRizClaimCallerSession) WRiz() (common.Address, error) {
	return _WRizClaim.Contract.WRiz(&_WRizClaim.CallOpts)
}

// AddSigners is a paid mutator transaction binding the contract method 0xe8906a2d.
//
// Solidity: function addSigners(address[] newSigners) returns()
func (_WRizClaim *WRizClaimTransactor) AddSigners(opts *bind.TransactOpts, newSigners []common.Address) (*types.Transaction, error) {
	return _WRizClaim.contract.Transact(opts, "addSigners", newSigners)
}

// AddSigners is a paid mutator transaction binding the contract method 0xe8906a2d.
//
// Solidity: function addSigners(address[] newSigners) returns()
func (_WRizClaim *WRizClaimSession) AddSigners(newSigners []common.Address) (*types.Transaction, error) {
	return _WRizClaim.Contract.AddSigners(&_WRizClaim.TransactOpts, newSigners)
}

// AddSigners is a paid mutator transaction binding the contract method 0xe8906a2d.
//
// Solidity: function addSigners(address[] newSigners) returns()
func (_WRizClaim *WRizClaimTransactorSession) AddSigners(newSigners []common.Address) (*types.Transaction, error) {
	return _WRizClaim.Contract.AddSigners(&_WRizClaim.TransactOpts, newSigners)
}

// Claim is a paid mutator transaction binding the contract method 0x5eddd157.
//
// Solidity: function claim(uint256 amount, uint256 nonce, bytes signature) returns()
func (_WRizClaim *WRizClaimTransactor) Claim(opts *bind.TransactOpts, amount *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _WRizClaim.contract.Transact(opts, "claim", amount, nonce, signature)
}

// Claim is a paid mutator transaction binding the contract method 0x5eddd157.
//
// Solidity: function claim(uint256 amount, uint256 nonce, bytes signature) returns()
func (_WRizClaim *WRizClaimSession) Claim(amount *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _WRizClaim.Contract.Claim(&_WRizClaim.TransactOpts, amount, nonce, signature)
}

// Claim is a paid mutator transaction binding the contract method 0x5eddd157.
//
// Solidity: function claim(uint256 amount, uint256 nonce, bytes signature) returns()
func (_WRizClaim *WRizClaimTransactorSession) Claim(amount *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _WRizClaim.Contract.Claim(&_WRizClaim.TransactOpts, amount, nonce, signature)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x6ff1c9bc.
//
// Solidity: function emergencyWithdraw(address token) returns()
func (_WRizClaim *WRizClaimTransactor) EmergencyWithdraw(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _WRizClaim.contract.Transact(opts, "emergencyWithdraw", token)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x6ff1c9bc.
//
// Solidity: function emergencyWithdraw(address token) returns()
func (_WRizClaim *WRizClaimSession) EmergencyWithdraw(token common.Address) (*types.Transaction, error) {
	return _WRizClaim.Contract.EmergencyWithdraw(&_WRizClaim.TransactOpts, token)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x6ff1c9bc.
//
// Solidity: function emergencyWithdraw(address token) returns()
func (_WRizClaim *WRizClaimTransactorSession) EmergencyWithdraw(token common.Address) (*types.Transaction, error) {
	return _WRizClaim.Contract.EmergencyWithdraw(&_WRizClaim.TransactOpts, token)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_WRizClaim *WRizClaimTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WRizClaim.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_WRizClaim *WRizClaimSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WRizClaim.Contract.GrantRole(&_WRizClaim.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_WRizClaim *WRizClaimTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WRizClaim.Contract.GrantRole(&_WRizClaim.TransactOpts, role, account)
}

// RemoveSigners is a paid mutator transaction binding the contract method 0x8d361e43.
//
// Solidity: function removeSigners(address[] signersToRemove) returns()
func (_WRizClaim *WRizClaimTransactor) RemoveSigners(opts *bind.TransactOpts, signersToRemove []common.Address) (*types.Transaction, error) {
	return _WRizClaim.contract.Transact(opts, "removeSigners", signersToRemove)
}

// RemoveSigners is a paid mutator transaction binding the contract method 0x8d361e43.
//
// Solidity: function removeSigners(address[] signersToRemove) returns()
func (_WRizClaim *WRizClaimSession) RemoveSigners(signersToRemove []common.Address) (*types.Transaction, error) {
	return _WRizClaim.Contract.RemoveSigners(&_WRizClaim.TransactOpts, signersToRemove)
}

// RemoveSigners is a paid mutator transaction binding the contract method 0x8d361e43.
//
// Solidity: function removeSigners(address[] signersToRemove) returns()
func (_WRizClaim *WRizClaimTransactorSession) RemoveSigners(signersToRemove []common.Address) (*types.Transaction, error) {
	return _WRizClaim.Contract.RemoveSigners(&_WRizClaim.TransactOpts, signersToRemove)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_WRizClaim *WRizClaimTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _WRizClaim.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_WRizClaim *WRizClaimSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _WRizClaim.Contract.RenounceRole(&_WRizClaim.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_WRizClaim *WRizClaimTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _WRizClaim.Contract.RenounceRole(&_WRizClaim.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_WRizClaim *WRizClaimTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WRizClaim.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_WRizClaim *WRizClaimSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WRizClaim.Contract.RevokeRole(&_WRizClaim.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_WRizClaim *WRizClaimTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _WRizClaim.Contract.RevokeRole(&_WRizClaim.TransactOpts, role, account)
}

// WRizClaimClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the WRizClaim contract.
type WRizClaimClaimedIterator struct {
	Event *WRizClaimClaimed // Event containing the contract specifics and raw log

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
func (it *WRizClaimClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRizClaimClaimed)
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
		it.Event = new(WRizClaimClaimed)
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
func (it *WRizClaimClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRizClaimClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRizClaimClaimed represents a Claimed event raised by the WRizClaim contract.
type WRizClaimClaimed struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed user, uint256 indexed amount)
func (_WRizClaim *WRizClaimFilterer) FilterClaimed(opts *bind.FilterOpts, user []common.Address, amount []*big.Int) (*WRizClaimClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _WRizClaim.contract.FilterLogs(opts, "Claimed", userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &WRizClaimClaimedIterator{contract: _WRizClaim.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed user, uint256 indexed amount)
func (_WRizClaim *WRizClaimFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *WRizClaimClaimed, user []common.Address, amount []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _WRizClaim.contract.WatchLogs(opts, "Claimed", userRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRizClaimClaimed)
				if err := _WRizClaim.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed user, uint256 indexed amount)
func (_WRizClaim *WRizClaimFilterer) ParseClaimed(log types.Log) (*WRizClaimClaimed, error) {
	event := new(WRizClaimClaimed)
	if err := _WRizClaim.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WRizClaimRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the WRizClaim contract.
type WRizClaimRoleAdminChangedIterator struct {
	Event *WRizClaimRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *WRizClaimRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRizClaimRoleAdminChanged)
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
		it.Event = new(WRizClaimRoleAdminChanged)
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
func (it *WRizClaimRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRizClaimRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRizClaimRoleAdminChanged represents a RoleAdminChanged event raised by the WRizClaim contract.
type WRizClaimRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_WRizClaim *WRizClaimFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*WRizClaimRoleAdminChangedIterator, error) {

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

	logs, sub, err := _WRizClaim.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &WRizClaimRoleAdminChangedIterator{contract: _WRizClaim.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_WRizClaim *WRizClaimFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *WRizClaimRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _WRizClaim.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRizClaimRoleAdminChanged)
				if err := _WRizClaim.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_WRizClaim *WRizClaimFilterer) ParseRoleAdminChanged(log types.Log) (*WRizClaimRoleAdminChanged, error) {
	event := new(WRizClaimRoleAdminChanged)
	if err := _WRizClaim.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WRizClaimRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the WRizClaim contract.
type WRizClaimRoleGrantedIterator struct {
	Event *WRizClaimRoleGranted // Event containing the contract specifics and raw log

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
func (it *WRizClaimRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRizClaimRoleGranted)
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
		it.Event = new(WRizClaimRoleGranted)
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
func (it *WRizClaimRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRizClaimRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRizClaimRoleGranted represents a RoleGranted event raised by the WRizClaim contract.
type WRizClaimRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_WRizClaim *WRizClaimFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*WRizClaimRoleGrantedIterator, error) {

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

	logs, sub, err := _WRizClaim.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &WRizClaimRoleGrantedIterator{contract: _WRizClaim.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_WRizClaim *WRizClaimFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *WRizClaimRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _WRizClaim.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRizClaimRoleGranted)
				if err := _WRizClaim.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_WRizClaim *WRizClaimFilterer) ParseRoleGranted(log types.Log) (*WRizClaimRoleGranted, error) {
	event := new(WRizClaimRoleGranted)
	if err := _WRizClaim.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WRizClaimRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the WRizClaim contract.
type WRizClaimRoleRevokedIterator struct {
	Event *WRizClaimRoleRevoked // Event containing the contract specifics and raw log

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
func (it *WRizClaimRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRizClaimRoleRevoked)
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
		it.Event = new(WRizClaimRoleRevoked)
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
func (it *WRizClaimRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRizClaimRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRizClaimRoleRevoked represents a RoleRevoked event raised by the WRizClaim contract.
type WRizClaimRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_WRizClaim *WRizClaimFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*WRizClaimRoleRevokedIterator, error) {

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

	logs, sub, err := _WRizClaim.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &WRizClaimRoleRevokedIterator{contract: _WRizClaim.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_WRizClaim *WRizClaimFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *WRizClaimRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _WRizClaim.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRizClaimRoleRevoked)
				if err := _WRizClaim.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_WRizClaim *WRizClaimFilterer) ParseRoleRevoked(log types.Log) (*WRizClaimRoleRevoked, error) {
	event := new(WRizClaimRoleRevoked)
	if err := _WRizClaim.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WRizClaimSignerAddedIterator is returned from FilterSignerAdded and is used to iterate over the raw logs and unpacked data for SignerAdded events raised by the WRizClaim contract.
type WRizClaimSignerAddedIterator struct {
	Event *WRizClaimSignerAdded // Event containing the contract specifics and raw log

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
func (it *WRizClaimSignerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRizClaimSignerAdded)
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
		it.Event = new(WRizClaimSignerAdded)
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
func (it *WRizClaimSignerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRizClaimSignerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRizClaimSignerAdded represents a SignerAdded event raised by the WRizClaim contract.
type WRizClaimSignerAdded struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignerAdded is a free log retrieval operation binding the contract event 0x47d1c22a25bb3a5d4e481b9b1e6944c2eade3181a0a20b495ed61d35b5323f24.
//
// Solidity: event SignerAdded(address indexed signer)
func (_WRizClaim *WRizClaimFilterer) FilterSignerAdded(opts *bind.FilterOpts, signer []common.Address) (*WRizClaimSignerAddedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _WRizClaim.contract.FilterLogs(opts, "SignerAdded", signerRule)
	if err != nil {
		return nil, err
	}
	return &WRizClaimSignerAddedIterator{contract: _WRizClaim.contract, event: "SignerAdded", logs: logs, sub: sub}, nil
}

// WatchSignerAdded is a free log subscription operation binding the contract event 0x47d1c22a25bb3a5d4e481b9b1e6944c2eade3181a0a20b495ed61d35b5323f24.
//
// Solidity: event SignerAdded(address indexed signer)
func (_WRizClaim *WRizClaimFilterer) WatchSignerAdded(opts *bind.WatchOpts, sink chan<- *WRizClaimSignerAdded, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _WRizClaim.contract.WatchLogs(opts, "SignerAdded", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRizClaimSignerAdded)
				if err := _WRizClaim.contract.UnpackLog(event, "SignerAdded", log); err != nil {
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

// ParseSignerAdded is a log parse operation binding the contract event 0x47d1c22a25bb3a5d4e481b9b1e6944c2eade3181a0a20b495ed61d35b5323f24.
//
// Solidity: event SignerAdded(address indexed signer)
func (_WRizClaim *WRizClaimFilterer) ParseSignerAdded(log types.Log) (*WRizClaimSignerAdded, error) {
	event := new(WRizClaimSignerAdded)
	if err := _WRizClaim.contract.UnpackLog(event, "SignerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WRizClaimSignerRemovedIterator is returned from FilterSignerRemoved and is used to iterate over the raw logs and unpacked data for SignerRemoved events raised by the WRizClaim contract.
type WRizClaimSignerRemovedIterator struct {
	Event *WRizClaimSignerRemoved // Event containing the contract specifics and raw log

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
func (it *WRizClaimSignerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WRizClaimSignerRemoved)
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
		it.Event = new(WRizClaimSignerRemoved)
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
func (it *WRizClaimSignerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WRizClaimSignerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WRizClaimSignerRemoved represents a SignerRemoved event raised by the WRizClaim contract.
type WRizClaimSignerRemoved struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignerRemoved is a free log retrieval operation binding the contract event 0x3525e22824a8a7df2c9a6029941c824cf95b6447f1e13d5128fd3826d35afe8b.
//
// Solidity: event SignerRemoved(address indexed signer)
func (_WRizClaim *WRizClaimFilterer) FilterSignerRemoved(opts *bind.FilterOpts, signer []common.Address) (*WRizClaimSignerRemovedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _WRizClaim.contract.FilterLogs(opts, "SignerRemoved", signerRule)
	if err != nil {
		return nil, err
	}
	return &WRizClaimSignerRemovedIterator{contract: _WRizClaim.contract, event: "SignerRemoved", logs: logs, sub: sub}, nil
}

// WatchSignerRemoved is a free log subscription operation binding the contract event 0x3525e22824a8a7df2c9a6029941c824cf95b6447f1e13d5128fd3826d35afe8b.
//
// Solidity: event SignerRemoved(address indexed signer)
func (_WRizClaim *WRizClaimFilterer) WatchSignerRemoved(opts *bind.WatchOpts, sink chan<- *WRizClaimSignerRemoved, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _WRizClaim.contract.WatchLogs(opts, "SignerRemoved", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WRizClaimSignerRemoved)
				if err := _WRizClaim.contract.UnpackLog(event, "SignerRemoved", log); err != nil {
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

// ParseSignerRemoved is a log parse operation binding the contract event 0x3525e22824a8a7df2c9a6029941c824cf95b6447f1e13d5128fd3826d35afe8b.
//
// Solidity: event SignerRemoved(address indexed signer)
func (_WRizClaim *WRizClaimFilterer) ParseSignerRemoved(log types.Log) (*WRizClaimSignerRemoved, error) {
	event := new(WRizClaimSignerRemoved)
	if err := _WRizClaim.contract.UnpackLog(event, "SignerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
