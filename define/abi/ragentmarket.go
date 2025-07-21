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

// RAgentMarketPlaceAccountInfo is an auto generated low-level Go binding around an user-defined struct.
type RAgentMarketPlaceAccountInfo struct {
	Account common.Address
	TokenId *big.Int
}

// RAgentMarketPlaceBidInfo is an auto generated low-level Go binding around an user-defined struct.
type RAgentMarketPlaceBidInfo struct {
	BidId         *big.Int
	Bidder        common.Address
	Quantity      *big.Int
	TotalPrice    *big.Int
	VestingPeriod *big.Int
	PaymentToken  common.Address
	RAgentToken   common.Address
	CreatedAt     *big.Int
	IsReleased    bool
	IsAgreed      bool
	Seller        common.Address
}

// RAgentMarketMetaData contains all meta data concerning the RAgentMarket contract.
var RAgentMarketMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc6551Registry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_erc6551AccountImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sPaymentToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sTreasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sExecutor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sAdmin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BidAlreadyAgreed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BidAlreadyReleased\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BidNotAgreed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPaymentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVestingPeriod\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonceUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotBidder\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerNotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VestingPeriodNotEnded\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"BidAgreeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vestingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rAgentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isReleased\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isAgreed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structRAgentMarketPlace.BidInfo\",\"name\":\"bid\",\"type\":\"tuple\"}],\"name\":\"BidCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"BidReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"}],\"name\":\"BidRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structRAgentMarketPlace.AccountInfo[]\",\"name\":\"accounts\",\"type\":\"tuple[]\"}],\"name\":\"RAgentBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signer\",\"type\":\"address[]\"}],\"name\":\"SignerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"}],\"name\":\"SignersAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"}],\"name\":\"agreeBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rAgentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vestingPeriod\",\"type\":\"uint256\"}],\"name\":\"bidRagent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rAgentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"buyRagent\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structRAgentMarketPlace.AccountInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc6551AccountImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc6551Registry\",\"outputs\":[{\"internalType\":\"contractIERC6551Registry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBidCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"}],\"name\":\"getBidInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vestingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rAgentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isReleased\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isAgreed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"internalType\":\"structRAgentMarketPlace.BidInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"}],\"name\":\"getBidInfoByBidder\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vestingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rAgentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isReleased\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isAgreed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"internalType\":\"structRAgentMarketPlace.BidInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rAgentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"}],\"name\":\"getBidInfoByRagent\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vestingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rAgentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isReleased\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isAgreed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"internalType\":\"structRAgentMarketPlace.BidInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"getBidInfoByRange\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vestingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rAgentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isReleased\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isAgreed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"internalType\":\"structRAgentMarketPlace.BidInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserBidIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"}],\"name\":\"releaseBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"}],\"name\":\"removeBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"}],\"name\":\"removeSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sAccountInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sBidCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sBidIdIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sBidIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sBidInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vestingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rAgentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isReleased\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isAgreed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sCurrentBidId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sExecutor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sIsSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sNonceUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sPaymentToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sRagentBidCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sRagentBidId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sSellerBidIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sSigners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sUserBidIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sVestingPeriod\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sAdmin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sExecutor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sAdmin\",\"type\":\"address\"}],\"name\":\"setExecutorAndAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sPaymentToken\",\"type\":\"address\"}],\"name\":\"setPaymentToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sTreasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_vestingPeriods\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setVestingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RAgentMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use RAgentMarketMetaData.ABI instead.
var RAgentMarketABI = RAgentMarketMetaData.ABI

// RAgentMarket is an auto generated Go binding around an Ethereum contract.
type RAgentMarket struct {
	RAgentMarketCaller     // Read-only binding to the contract
	RAgentMarketTransactor // Write-only binding to the contract
	RAgentMarketFilterer   // Log filterer for contract events
}

// RAgentMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type RAgentMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RAgentMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RAgentMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RAgentMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RAgentMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RAgentMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RAgentMarketSession struct {
	Contract     *RAgentMarket     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RAgentMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RAgentMarketCallerSession struct {
	Contract *RAgentMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RAgentMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RAgentMarketTransactorSession struct {
	Contract     *RAgentMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RAgentMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type RAgentMarketRaw struct {
	Contract *RAgentMarket // Generic contract binding to access the raw methods on
}

// RAgentMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RAgentMarketCallerRaw struct {
	Contract *RAgentMarketCaller // Generic read-only contract binding to access the raw methods on
}

// RAgentMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RAgentMarketTransactorRaw struct {
	Contract *RAgentMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRAgentMarket creates a new instance of RAgentMarket, bound to a specific deployed contract.
func NewRAgentMarket(address common.Address, backend bind.ContractBackend) (*RAgentMarket, error) {
	contract, err := bindRAgentMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RAgentMarket{RAgentMarketCaller: RAgentMarketCaller{contract: contract}, RAgentMarketTransactor: RAgentMarketTransactor{contract: contract}, RAgentMarketFilterer: RAgentMarketFilterer{contract: contract}}, nil
}

// NewRAgentMarketCaller creates a new read-only instance of RAgentMarket, bound to a specific deployed contract.
func NewRAgentMarketCaller(address common.Address, caller bind.ContractCaller) (*RAgentMarketCaller, error) {
	contract, err := bindRAgentMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RAgentMarketCaller{contract: contract}, nil
}

// NewRAgentMarketTransactor creates a new write-only instance of RAgentMarket, bound to a specific deployed contract.
func NewRAgentMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*RAgentMarketTransactor, error) {
	contract, err := bindRAgentMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RAgentMarketTransactor{contract: contract}, nil
}

// NewRAgentMarketFilterer creates a new log filterer instance of RAgentMarket, bound to a specific deployed contract.
func NewRAgentMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*RAgentMarketFilterer, error) {
	contract, err := bindRAgentMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RAgentMarketFilterer{contract: contract}, nil
}

// bindRAgentMarket binds a generic wrapper to an already deployed contract.
func bindRAgentMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RAgentMarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RAgentMarket *RAgentMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RAgentMarket.Contract.RAgentMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RAgentMarket *RAgentMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RAgentMarket.Contract.RAgentMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RAgentMarket *RAgentMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RAgentMarket.Contract.RAgentMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RAgentMarket *RAgentMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RAgentMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RAgentMarket *RAgentMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RAgentMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RAgentMarket *RAgentMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RAgentMarket.Contract.contract.Transact(opts, method, params...)
}

// Erc6551AccountImplementation is a free data retrieval call binding the contract method 0x918372de.
//
// Solidity: function erc6551AccountImplementation() view returns(address)
func (_RAgentMarket *RAgentMarketCaller) Erc6551AccountImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "erc6551AccountImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc6551AccountImplementation is a free data retrieval call binding the contract method 0x918372de.
//
// Solidity: function erc6551AccountImplementation() view returns(address)
func (_RAgentMarket *RAgentMarketSession) Erc6551AccountImplementation() (common.Address, error) {
	return _RAgentMarket.Contract.Erc6551AccountImplementation(&_RAgentMarket.CallOpts)
}

// Erc6551AccountImplementation is a free data retrieval call binding the contract method 0x918372de.
//
// Solidity: function erc6551AccountImplementation() view returns(address)
func (_RAgentMarket *RAgentMarketCallerSession) Erc6551AccountImplementation() (common.Address, error) {
	return _RAgentMarket.Contract.Erc6551AccountImplementation(&_RAgentMarket.CallOpts)
}

// Erc6551Registry is a free data retrieval call binding the contract method 0x056d5afe.
//
// Solidity: function erc6551Registry() view returns(address)
func (_RAgentMarket *RAgentMarketCaller) Erc6551Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "erc6551Registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc6551Registry is a free data retrieval call binding the contract method 0x056d5afe.
//
// Solidity: function erc6551Registry() view returns(address)
func (_RAgentMarket *RAgentMarketSession) Erc6551Registry() (common.Address, error) {
	return _RAgentMarket.Contract.Erc6551Registry(&_RAgentMarket.CallOpts)
}

// Erc6551Registry is a free data retrieval call binding the contract method 0x056d5afe.
//
// Solidity: function erc6551Registry() view returns(address)
func (_RAgentMarket *RAgentMarketCallerSession) Erc6551Registry() (common.Address, error) {
	return _RAgentMarket.Contract.Erc6551Registry(&_RAgentMarket.CallOpts)
}

// GetBidCount is a free data retrieval call binding the contract method 0x28f6a48a.
//
// Solidity: function getBidCount() view returns(uint256)
func (_RAgentMarket *RAgentMarketCaller) GetBidCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "getBidCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBidCount is a free data retrieval call binding the contract method 0x28f6a48a.
//
// Solidity: function getBidCount() view returns(uint256)
func (_RAgentMarket *RAgentMarketSession) GetBidCount() (*big.Int, error) {
	return _RAgentMarket.Contract.GetBidCount(&_RAgentMarket.CallOpts)
}

// GetBidCount is a free data retrieval call binding the contract method 0x28f6a48a.
//
// Solidity: function getBidCount() view returns(uint256)
func (_RAgentMarket *RAgentMarketCallerSession) GetBidCount() (*big.Int, error) {
	return _RAgentMarket.Contract.GetBidCount(&_RAgentMarket.CallOpts)
}

// GetBidInfo is a free data retrieval call binding the contract method 0x2201db31.
//
// Solidity: function getBidInfo(uint256 bidId) view returns((uint256,address,uint256,uint256,uint256,address,address,uint256,bool,bool,address))
func (_RAgentMarket *RAgentMarketCaller) GetBidInfo(opts *bind.CallOpts, bidId *big.Int) (RAgentMarketPlaceBidInfo, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "getBidInfo", bidId)

	if err != nil {
		return *new(RAgentMarketPlaceBidInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(RAgentMarketPlaceBidInfo)).(*RAgentMarketPlaceBidInfo)

	return out0, err

}

// GetBidInfo is a free data retrieval call binding the contract method 0x2201db31.
//
// Solidity: function getBidInfo(uint256 bidId) view returns((uint256,address,uint256,uint256,uint256,address,address,uint256,bool,bool,address))
func (_RAgentMarket *RAgentMarketSession) GetBidInfo(bidId *big.Int) (RAgentMarketPlaceBidInfo, error) {
	return _RAgentMarket.Contract.GetBidInfo(&_RAgentMarket.CallOpts, bidId)
}

// GetBidInfo is a free data retrieval call binding the contract method 0x2201db31.
//
// Solidity: function getBidInfo(uint256 bidId) view returns((uint256,address,uint256,uint256,uint256,address,address,uint256,bool,bool,address))
func (_RAgentMarket *RAgentMarketCallerSession) GetBidInfo(bidId *big.Int) (RAgentMarketPlaceBidInfo, error) {
	return _RAgentMarket.Contract.GetBidInfo(&_RAgentMarket.CallOpts, bidId)
}

// GetBidInfoByBidder is a free data retrieval call binding the contract method 0x5c3542b2.
//
// Solidity: function getBidInfoByBidder(address bidder) view returns((uint256,address,uint256,uint256,uint256,address,address,uint256,bool,bool,address)[])
func (_RAgentMarket *RAgentMarketCaller) GetBidInfoByBidder(opts *bind.CallOpts, bidder common.Address) ([]RAgentMarketPlaceBidInfo, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "getBidInfoByBidder", bidder)

	if err != nil {
		return *new([]RAgentMarketPlaceBidInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]RAgentMarketPlaceBidInfo)).(*[]RAgentMarketPlaceBidInfo)

	return out0, err

}

// GetBidInfoByBidder is a free data retrieval call binding the contract method 0x5c3542b2.
//
// Solidity: function getBidInfoByBidder(address bidder) view returns((uint256,address,uint256,uint256,uint256,address,address,uint256,bool,bool,address)[])
func (_RAgentMarket *RAgentMarketSession) GetBidInfoByBidder(bidder common.Address) ([]RAgentMarketPlaceBidInfo, error) {
	return _RAgentMarket.Contract.GetBidInfoByBidder(&_RAgentMarket.CallOpts, bidder)
}

// GetBidInfoByBidder is a free data retrieval call binding the contract method 0x5c3542b2.
//
// Solidity: function getBidInfoByBidder(address bidder) view returns((uint256,address,uint256,uint256,uint256,address,address,uint256,bool,bool,address)[])
func (_RAgentMarket *RAgentMarketCallerSession) GetBidInfoByBidder(bidder common.Address) ([]RAgentMarketPlaceBidInfo, error) {
	return _RAgentMarket.Contract.GetBidInfoByBidder(&_RAgentMarket.CallOpts, bidder)
}

// GetBidInfoByRagent is a free data retrieval call binding the contract method 0x2a571313.
//
// Solidity: function getBidInfoByRagent(address rAgentToken, uint256 from, uint256 to) view returns((uint256,address,uint256,uint256,uint256,address,address,uint256,bool,bool,address)[])
func (_RAgentMarket *RAgentMarketCaller) GetBidInfoByRagent(opts *bind.CallOpts, rAgentToken common.Address, from *big.Int, to *big.Int) ([]RAgentMarketPlaceBidInfo, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "getBidInfoByRagent", rAgentToken, from, to)

	if err != nil {
		return *new([]RAgentMarketPlaceBidInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]RAgentMarketPlaceBidInfo)).(*[]RAgentMarketPlaceBidInfo)

	return out0, err

}

// GetBidInfoByRagent is a free data retrieval call binding the contract method 0x2a571313.
//
// Solidity: function getBidInfoByRagent(address rAgentToken, uint256 from, uint256 to) view returns((uint256,address,uint256,uint256,uint256,address,address,uint256,bool,bool,address)[])
func (_RAgentMarket *RAgentMarketSession) GetBidInfoByRagent(rAgentToken common.Address, from *big.Int, to *big.Int) ([]RAgentMarketPlaceBidInfo, error) {
	return _RAgentMarket.Contract.GetBidInfoByRagent(&_RAgentMarket.CallOpts, rAgentToken, from, to)
}

// GetBidInfoByRagent is a free data retrieval call binding the contract method 0x2a571313.
//
// Solidity: function getBidInfoByRagent(address rAgentToken, uint256 from, uint256 to) view returns((uint256,address,uint256,uint256,uint256,address,address,uint256,bool,bool,address)[])
func (_RAgentMarket *RAgentMarketCallerSession) GetBidInfoByRagent(rAgentToken common.Address, from *big.Int, to *big.Int) ([]RAgentMarketPlaceBidInfo, error) {
	return _RAgentMarket.Contract.GetBidInfoByRagent(&_RAgentMarket.CallOpts, rAgentToken, from, to)
}

// GetBidInfoByRange is a free data retrieval call binding the contract method 0x0cc6ecbf.
//
// Solidity: function getBidInfoByRange(uint256 start, uint256 end) view returns((uint256,address,uint256,uint256,uint256,address,address,uint256,bool,bool,address)[])
func (_RAgentMarket *RAgentMarketCaller) GetBidInfoByRange(opts *bind.CallOpts, start *big.Int, end *big.Int) ([]RAgentMarketPlaceBidInfo, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "getBidInfoByRange", start, end)

	if err != nil {
		return *new([]RAgentMarketPlaceBidInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]RAgentMarketPlaceBidInfo)).(*[]RAgentMarketPlaceBidInfo)

	return out0, err

}

// GetBidInfoByRange is a free data retrieval call binding the contract method 0x0cc6ecbf.
//
// Solidity: function getBidInfoByRange(uint256 start, uint256 end) view returns((uint256,address,uint256,uint256,uint256,address,address,uint256,bool,bool,address)[])
func (_RAgentMarket *RAgentMarketSession) GetBidInfoByRange(start *big.Int, end *big.Int) ([]RAgentMarketPlaceBidInfo, error) {
	return _RAgentMarket.Contract.GetBidInfoByRange(&_RAgentMarket.CallOpts, start, end)
}

// GetBidInfoByRange is a free data retrieval call binding the contract method 0x0cc6ecbf.
//
// Solidity: function getBidInfoByRange(uint256 start, uint256 end) view returns((uint256,address,uint256,uint256,uint256,address,address,uint256,bool,bool,address)[])
func (_RAgentMarket *RAgentMarketCallerSession) GetBidInfoByRange(start *big.Int, end *big.Int) ([]RAgentMarketPlaceBidInfo, error) {
	return _RAgentMarket.Contract.GetBidInfoByRange(&_RAgentMarket.CallOpts, start, end)
}

// GetUserBidIds is a free data retrieval call binding the contract method 0xd6c54654.
//
// Solidity: function getUserBidIds(address user) view returns(uint256[])
func (_RAgentMarket *RAgentMarketCaller) GetUserBidIds(opts *bind.CallOpts, user common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "getUserBidIds", user)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserBidIds is a free data retrieval call binding the contract method 0xd6c54654.
//
// Solidity: function getUserBidIds(address user) view returns(uint256[])
func (_RAgentMarket *RAgentMarketSession) GetUserBidIds(user common.Address) ([]*big.Int, error) {
	return _RAgentMarket.Contract.GetUserBidIds(&_RAgentMarket.CallOpts, user)
}

// GetUserBidIds is a free data retrieval call binding the contract method 0xd6c54654.
//
// Solidity: function getUserBidIds(address user) view returns(uint256[])
func (_RAgentMarket *RAgentMarketCallerSession) GetUserBidIds(user common.Address) ([]*big.Int, error) {
	return _RAgentMarket.Contract.GetUserBidIds(&_RAgentMarket.CallOpts, user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RAgentMarket *RAgentMarketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RAgentMarket *RAgentMarketSession) Owner() (common.Address, error) {
	return _RAgentMarket.Contract.Owner(&_RAgentMarket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RAgentMarket *RAgentMarketCallerSession) Owner() (common.Address, error) {
	return _RAgentMarket.Contract.Owner(&_RAgentMarket.CallOpts)
}

// SAccountInfo is a free data retrieval call binding the contract method 0xd20dfe5e.
//
// Solidity: function sAccountInfo(address ) view returns(address account, uint256 tokenId)
func (_RAgentMarket *RAgentMarketCaller) SAccountInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Account common.Address
	TokenId *big.Int
}, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "sAccountInfo", arg0)

	outstruct := new(struct {
		Account common.Address
		TokenId *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Account = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SAccountInfo is a free data retrieval call binding the contract method 0xd20dfe5e.
//
// Solidity: function sAccountInfo(address ) view returns(address account, uint256 tokenId)
func (_RAgentMarket *RAgentMarketSession) SAccountInfo(arg0 common.Address) (struct {
	Account common.Address
	TokenId *big.Int
}, error) {
	return _RAgentMarket.Contract.SAccountInfo(&_RAgentMarket.CallOpts, arg0)
}

// SAccountInfo is a free data retrieval call binding the contract method 0xd20dfe5e.
//
// Solidity: function sAccountInfo(address ) view returns(address account, uint256 tokenId)
func (_RAgentMarket *RAgentMarketCallerSession) SAccountInfo(arg0 common.Address) (struct {
	Account common.Address
	TokenId *big.Int
}, error) {
	return _RAgentMarket.Contract.SAccountInfo(&_RAgentMarket.CallOpts, arg0)
}

// SAdmin is a free data retrieval call binding the contract method 0x4c39273e.
//
// Solidity: function sAdmin() view returns(address)
func (_RAgentMarket *RAgentMarketCaller) SAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "sAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SAdmin is a free data retrieval call binding the contract method 0x4c39273e.
//
// Solidity: function sAdmin() view returns(address)
func (_RAgentMarket *RAgentMarketSession) SAdmin() (common.Address, error) {
	return _RAgentMarket.Contract.SAdmin(&_RAgentMarket.CallOpts)
}

// SAdmin is a free data retrieval call binding the contract method 0x4c39273e.
//
// Solidity: function sAdmin() view returns(address)
func (_RAgentMarket *RAgentMarketCallerSession) SAdmin() (common.Address, error) {
	return _RAgentMarket.Contract.SAdmin(&_RAgentMarket.CallOpts)
}

// SBidCount is a free data retrieval call binding the contract method 0xfb88d98d.
//
// Solidity: function sBidCount() view returns(uint256)
func (_RAgentMarket *RAgentMarketCaller) SBidCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "sBidCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SBidCount is a free data retrieval call binding the contract method 0xfb88d98d.
//
// Solidity: function sBidCount() view returns(uint256)
func (_RAgentMarket *RAgentMarketSession) SBidCount() (*big.Int, error) {
	return _RAgentMarket.Contract.SBidCount(&_RAgentMarket.CallOpts)
}

// SBidCount is a free data retrieval call binding the contract method 0xfb88d98d.
//
// Solidity: function sBidCount() view returns(uint256)
func (_RAgentMarket *RAgentMarketCallerSession) SBidCount() (*big.Int, error) {
	return _RAgentMarket.Contract.SBidCount(&_RAgentMarket.CallOpts)
}

// SBidIdIndex is a free data retrieval call binding the contract method 0xa4a9dc62.
//
// Solidity: function sBidIdIndex(uint256 ) view returns(uint256)
func (_RAgentMarket *RAgentMarketCaller) SBidIdIndex(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "sBidIdIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SBidIdIndex is a free data retrieval call binding the contract method 0xa4a9dc62.
//
// Solidity: function sBidIdIndex(uint256 ) view returns(uint256)
func (_RAgentMarket *RAgentMarketSession) SBidIdIndex(arg0 *big.Int) (*big.Int, error) {
	return _RAgentMarket.Contract.SBidIdIndex(&_RAgentMarket.CallOpts, arg0)
}

// SBidIdIndex is a free data retrieval call binding the contract method 0xa4a9dc62.
//
// Solidity: function sBidIdIndex(uint256 ) view returns(uint256)
func (_RAgentMarket *RAgentMarketCallerSession) SBidIdIndex(arg0 *big.Int) (*big.Int, error) {
	return _RAgentMarket.Contract.SBidIdIndex(&_RAgentMarket.CallOpts, arg0)
}

// SBidIds is a free data retrieval call binding the contract method 0xa05efa69.
//
// Solidity: function sBidIds(uint256 ) view returns(uint256)
func (_RAgentMarket *RAgentMarketCaller) SBidIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "sBidIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SBidIds is a free data retrieval call binding the contract method 0xa05efa69.
//
// Solidity: function sBidIds(uint256 ) view returns(uint256)
func (_RAgentMarket *RAgentMarketSession) SBidIds(arg0 *big.Int) (*big.Int, error) {
	return _RAgentMarket.Contract.SBidIds(&_RAgentMarket.CallOpts, arg0)
}

// SBidIds is a free data retrieval call binding the contract method 0xa05efa69.
//
// Solidity: function sBidIds(uint256 ) view returns(uint256)
func (_RAgentMarket *RAgentMarketCallerSession) SBidIds(arg0 *big.Int) (*big.Int, error) {
	return _RAgentMarket.Contract.SBidIds(&_RAgentMarket.CallOpts, arg0)
}

// SBidInfo is a free data retrieval call binding the contract method 0x2c75a3ed.
//
// Solidity: function sBidInfo(uint256 ) view returns(uint256 bidId, address bidder, uint256 quantity, uint256 totalPrice, uint256 vestingPeriod, address paymentToken, address rAgentToken, uint256 createdAt, bool isReleased, bool isAgreed, address seller)
func (_RAgentMarket *RAgentMarketCaller) SBidInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BidId         *big.Int
	Bidder        common.Address
	Quantity      *big.Int
	TotalPrice    *big.Int
	VestingPeriod *big.Int
	PaymentToken  common.Address
	RAgentToken   common.Address
	CreatedAt     *big.Int
	IsReleased    bool
	IsAgreed      bool
	Seller        common.Address
}, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "sBidInfo", arg0)

	outstruct := new(struct {
		BidId         *big.Int
		Bidder        common.Address
		Quantity      *big.Int
		TotalPrice    *big.Int
		VestingPeriod *big.Int
		PaymentToken  common.Address
		RAgentToken   common.Address
		CreatedAt     *big.Int
		IsReleased    bool
		IsAgreed      bool
		Seller        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BidId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Bidder = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Quantity = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalPrice = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VestingPeriod = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.PaymentToken = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.RAgentToken = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.CreatedAt = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.IsReleased = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.IsAgreed = *abi.ConvertType(out[9], new(bool)).(*bool)
	outstruct.Seller = *abi.ConvertType(out[10], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// SBidInfo is a free data retrieval call binding the contract method 0x2c75a3ed.
//
// Solidity: function sBidInfo(uint256 ) view returns(uint256 bidId, address bidder, uint256 quantity, uint256 totalPrice, uint256 vestingPeriod, address paymentToken, address rAgentToken, uint256 createdAt, bool isReleased, bool isAgreed, address seller)
func (_RAgentMarket *RAgentMarketSession) SBidInfo(arg0 *big.Int) (struct {
	BidId         *big.Int
	Bidder        common.Address
	Quantity      *big.Int
	TotalPrice    *big.Int
	VestingPeriod *big.Int
	PaymentToken  common.Address
	RAgentToken   common.Address
	CreatedAt     *big.Int
	IsReleased    bool
	IsAgreed      bool
	Seller        common.Address
}, error) {
	return _RAgentMarket.Contract.SBidInfo(&_RAgentMarket.CallOpts, arg0)
}

// SBidInfo is a free data retrieval call binding the contract method 0x2c75a3ed.
//
// Solidity: function sBidInfo(uint256 ) view returns(uint256 bidId, address bidder, uint256 quantity, uint256 totalPrice, uint256 vestingPeriod, address paymentToken, address rAgentToken, uint256 createdAt, bool isReleased, bool isAgreed, address seller)
func (_RAgentMarket *RAgentMarketCallerSession) SBidInfo(arg0 *big.Int) (struct {
	BidId         *big.Int
	Bidder        common.Address
	Quantity      *big.Int
	TotalPrice    *big.Int
	VestingPeriod *big.Int
	PaymentToken  common.Address
	RAgentToken   common.Address
	CreatedAt     *big.Int
	IsReleased    bool
	IsAgreed      bool
	Seller        common.Address
}, error) {
	return _RAgentMarket.Contract.SBidInfo(&_RAgentMarket.CallOpts, arg0)
}

// SCurrentBidId is a free data retrieval call binding the contract method 0xb0131889.
//
// Solidity: function sCurrentBidId() view returns(uint256)
func (_RAgentMarket *RAgentMarketCaller) SCurrentBidId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "sCurrentBidId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SCurrentBidId is a free data retrieval call binding the contract method 0xb0131889.
//
// Solidity: function sCurrentBidId() view returns(uint256)
func (_RAgentMarket *RAgentMarketSession) SCurrentBidId() (*big.Int, error) {
	return _RAgentMarket.Contract.SCurrentBidId(&_RAgentMarket.CallOpts)
}

// SCurrentBidId is a free data retrieval call binding the contract method 0xb0131889.
//
// Solidity: function sCurrentBidId() view returns(uint256)
func (_RAgentMarket *RAgentMarketCallerSession) SCurrentBidId() (*big.Int, error) {
	return _RAgentMarket.Contract.SCurrentBidId(&_RAgentMarket.CallOpts)
}

// SExecutor is a free data retrieval call binding the contract method 0x15648ff4.
//
// Solidity: function sExecutor() view returns(address)
func (_RAgentMarket *RAgentMarketCaller) SExecutor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "sExecutor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SExecutor is a free data retrieval call binding the contract method 0x15648ff4.
//
// Solidity: function sExecutor() view returns(address)
func (_RAgentMarket *RAgentMarketSession) SExecutor() (common.Address, error) {
	return _RAgentMarket.Contract.SExecutor(&_RAgentMarket.CallOpts)
}

// SExecutor is a free data retrieval call binding the contract method 0x15648ff4.
//
// Solidity: function sExecutor() view returns(address)
func (_RAgentMarket *RAgentMarketCallerSession) SExecutor() (common.Address, error) {
	return _RAgentMarket.Contract.SExecutor(&_RAgentMarket.CallOpts)
}

// SIsSigner is a free data retrieval call binding the contract method 0xe697b5aa.
//
// Solidity: function sIsSigner(address ) view returns(bool)
func (_RAgentMarket *RAgentMarketCaller) SIsSigner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "sIsSigner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SIsSigner is a free data retrieval call binding the contract method 0xe697b5aa.
//
// Solidity: function sIsSigner(address ) view returns(bool)
func (_RAgentMarket *RAgentMarketSession) SIsSigner(arg0 common.Address) (bool, error) {
	return _RAgentMarket.Contract.SIsSigner(&_RAgentMarket.CallOpts, arg0)
}

// SIsSigner is a free data retrieval call binding the contract method 0xe697b5aa.
//
// Solidity: function sIsSigner(address ) view returns(bool)
func (_RAgentMarket *RAgentMarketCallerSession) SIsSigner(arg0 common.Address) (bool, error) {
	return _RAgentMarket.Contract.SIsSigner(&_RAgentMarket.CallOpts, arg0)
}

// SNonceUsed is a free data retrieval call binding the contract method 0xb6c97f82.
//
// Solidity: function sNonceUsed(uint256 ) view returns(bool)
func (_RAgentMarket *RAgentMarketCaller) SNonceUsed(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "sNonceUsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SNonceUsed is a free data retrieval call binding the contract method 0xb6c97f82.
//
// Solidity: function sNonceUsed(uint256 ) view returns(bool)
func (_RAgentMarket *RAgentMarketSession) SNonceUsed(arg0 *big.Int) (bool, error) {
	return _RAgentMarket.Contract.SNonceUsed(&_RAgentMarket.CallOpts, arg0)
}

// SNonceUsed is a free data retrieval call binding the contract method 0xb6c97f82.
//
// Solidity: function sNonceUsed(uint256 ) view returns(bool)
func (_RAgentMarket *RAgentMarketCallerSession) SNonceUsed(arg0 *big.Int) (bool, error) {
	return _RAgentMarket.Contract.SNonceUsed(&_RAgentMarket.CallOpts, arg0)
}

// SPaymentToken is a free data retrieval call binding the contract method 0x48288f6a.
//
// Solidity: function sPaymentToken() view returns(address)
func (_RAgentMarket *RAgentMarketCaller) SPaymentToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "sPaymentToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SPaymentToken is a free data retrieval call binding the contract method 0x48288f6a.
//
// Solidity: function sPaymentToken() view returns(address)
func (_RAgentMarket *RAgentMarketSession) SPaymentToken() (common.Address, error) {
	return _RAgentMarket.Contract.SPaymentToken(&_RAgentMarket.CallOpts)
}

// SPaymentToken is a free data retrieval call binding the contract method 0x48288f6a.
//
// Solidity: function sPaymentToken() view returns(address)
func (_RAgentMarket *RAgentMarketCallerSession) SPaymentToken() (common.Address, error) {
	return _RAgentMarket.Contract.SPaymentToken(&_RAgentMarket.CallOpts)
}

// SRagentBidCount is a free data retrieval call binding the contract method 0x2cb1145c.
//
// Solidity: function sRagentBidCount(address ) view returns(uint256)
func (_RAgentMarket *RAgentMarketCaller) SRagentBidCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "sRagentBidCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SRagentBidCount is a free data retrieval call binding the contract method 0x2cb1145c.
//
// Solidity: function sRagentBidCount(address ) view returns(uint256)
func (_RAgentMarket *RAgentMarketSession) SRagentBidCount(arg0 common.Address) (*big.Int, error) {
	return _RAgentMarket.Contract.SRagentBidCount(&_RAgentMarket.CallOpts, arg0)
}

// SRagentBidCount is a free data retrieval call binding the contract method 0x2cb1145c.
//
// Solidity: function sRagentBidCount(address ) view returns(uint256)
func (_RAgentMarket *RAgentMarketCallerSession) SRagentBidCount(arg0 common.Address) (*big.Int, error) {
	return _RAgentMarket.Contract.SRagentBidCount(&_RAgentMarket.CallOpts, arg0)
}

// SRagentBidId is a free data retrieval call binding the contract method 0x7370f7bf.
//
// Solidity: function sRagentBidId(address , uint256 ) view returns(uint256)
func (_RAgentMarket *RAgentMarketCaller) SRagentBidId(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "sRagentBidId", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SRagentBidId is a free data retrieval call binding the contract method 0x7370f7bf.
//
// Solidity: function sRagentBidId(address , uint256 ) view returns(uint256)
func (_RAgentMarket *RAgentMarketSession) SRagentBidId(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _RAgentMarket.Contract.SRagentBidId(&_RAgentMarket.CallOpts, arg0, arg1)
}

// SRagentBidId is a free data retrieval call binding the contract method 0x7370f7bf.
//
// Solidity: function sRagentBidId(address , uint256 ) view returns(uint256)
func (_RAgentMarket *RAgentMarketCallerSession) SRagentBidId(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _RAgentMarket.Contract.SRagentBidId(&_RAgentMarket.CallOpts, arg0, arg1)
}

// SSellerBidIds is a free data retrieval call binding the contract method 0x11952fdb.
//
// Solidity: function sSellerBidIds(address , uint256 ) view returns(uint256)
func (_RAgentMarket *RAgentMarketCaller) SSellerBidIds(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "sSellerBidIds", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SSellerBidIds is a free data retrieval call binding the contract method 0x11952fdb.
//
// Solidity: function sSellerBidIds(address , uint256 ) view returns(uint256)
func (_RAgentMarket *RAgentMarketSession) SSellerBidIds(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _RAgentMarket.Contract.SSellerBidIds(&_RAgentMarket.CallOpts, arg0, arg1)
}

// SSellerBidIds is a free data retrieval call binding the contract method 0x11952fdb.
//
// Solidity: function sSellerBidIds(address , uint256 ) view returns(uint256)
func (_RAgentMarket *RAgentMarketCallerSession) SSellerBidIds(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _RAgentMarket.Contract.SSellerBidIds(&_RAgentMarket.CallOpts, arg0, arg1)
}

// SSigners is a free data retrieval call binding the contract method 0x277667eb.
//
// Solidity: function sSigners(uint256 ) view returns(address)
func (_RAgentMarket *RAgentMarketCaller) SSigners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "sSigners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SSigners is a free data retrieval call binding the contract method 0x277667eb.
//
// Solidity: function sSigners(uint256 ) view returns(address)
func (_RAgentMarket *RAgentMarketSession) SSigners(arg0 *big.Int) (common.Address, error) {
	return _RAgentMarket.Contract.SSigners(&_RAgentMarket.CallOpts, arg0)
}

// SSigners is a free data retrieval call binding the contract method 0x277667eb.
//
// Solidity: function sSigners(uint256 ) view returns(address)
func (_RAgentMarket *RAgentMarketCallerSession) SSigners(arg0 *big.Int) (common.Address, error) {
	return _RAgentMarket.Contract.SSigners(&_RAgentMarket.CallOpts, arg0)
}

// STreasury is a free data retrieval call binding the contract method 0xdeea4bb4.
//
// Solidity: function sTreasury() view returns(address)
func (_RAgentMarket *RAgentMarketCaller) STreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "sTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STreasury is a free data retrieval call binding the contract method 0xdeea4bb4.
//
// Solidity: function sTreasury() view returns(address)
func (_RAgentMarket *RAgentMarketSession) STreasury() (common.Address, error) {
	return _RAgentMarket.Contract.STreasury(&_RAgentMarket.CallOpts)
}

// STreasury is a free data retrieval call binding the contract method 0xdeea4bb4.
//
// Solidity: function sTreasury() view returns(address)
func (_RAgentMarket *RAgentMarketCallerSession) STreasury() (common.Address, error) {
	return _RAgentMarket.Contract.STreasury(&_RAgentMarket.CallOpts)
}

// SUserBidIds is a free data retrieval call binding the contract method 0xcc0a80eb.
//
// Solidity: function sUserBidIds(address , uint256 ) view returns(uint256)
func (_RAgentMarket *RAgentMarketCaller) SUserBidIds(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "sUserBidIds", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SUserBidIds is a free data retrieval call binding the contract method 0xcc0a80eb.
//
// Solidity: function sUserBidIds(address , uint256 ) view returns(uint256)
func (_RAgentMarket *RAgentMarketSession) SUserBidIds(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _RAgentMarket.Contract.SUserBidIds(&_RAgentMarket.CallOpts, arg0, arg1)
}

// SUserBidIds is a free data retrieval call binding the contract method 0xcc0a80eb.
//
// Solidity: function sUserBidIds(address , uint256 ) view returns(uint256)
func (_RAgentMarket *RAgentMarketCallerSession) SUserBidIds(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _RAgentMarket.Contract.SUserBidIds(&_RAgentMarket.CallOpts, arg0, arg1)
}

// SVestingPeriod is a free data retrieval call binding the contract method 0x4b30e51a.
//
// Solidity: function sVestingPeriod(uint256 ) view returns(bool)
func (_RAgentMarket *RAgentMarketCaller) SVestingPeriod(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _RAgentMarket.contract.Call(opts, &out, "sVestingPeriod", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SVestingPeriod is a free data retrieval call binding the contract method 0x4b30e51a.
//
// Solidity: function sVestingPeriod(uint256 ) view returns(bool)
func (_RAgentMarket *RAgentMarketSession) SVestingPeriod(arg0 *big.Int) (bool, error) {
	return _RAgentMarket.Contract.SVestingPeriod(&_RAgentMarket.CallOpts, arg0)
}

// SVestingPeriod is a free data retrieval call binding the contract method 0x4b30e51a.
//
// Solidity: function sVestingPeriod(uint256 ) view returns(bool)
func (_RAgentMarket *RAgentMarketCallerSession) SVestingPeriod(arg0 *big.Int) (bool, error) {
	return _RAgentMarket.Contract.SVestingPeriod(&_RAgentMarket.CallOpts, arg0)
}

// AgreeBid is a paid mutator transaction binding the contract method 0x7f988622.
//
// Solidity: function agreeBid(uint256 bidId) returns()
func (_RAgentMarket *RAgentMarketTransactor) AgreeBid(opts *bind.TransactOpts, bidId *big.Int) (*types.Transaction, error) {
	return _RAgentMarket.contract.Transact(opts, "agreeBid", bidId)
}

// AgreeBid is a paid mutator transaction binding the contract method 0x7f988622.
//
// Solidity: function agreeBid(uint256 bidId) returns()
func (_RAgentMarket *RAgentMarketSession) AgreeBid(bidId *big.Int) (*types.Transaction, error) {
	return _RAgentMarket.Contract.AgreeBid(&_RAgentMarket.TransactOpts, bidId)
}

// AgreeBid is a paid mutator transaction binding the contract method 0x7f988622.
//
// Solidity: function agreeBid(uint256 bidId) returns()
func (_RAgentMarket *RAgentMarketTransactorSession) AgreeBid(bidId *big.Int) (*types.Transaction, error) {
	return _RAgentMarket.Contract.AgreeBid(&_RAgentMarket.TransactOpts, bidId)
}

// BidRagent is a paid mutator transaction binding the contract method 0xda95d4ac.
//
// Solidity: function bidRagent(address rAgentToken, uint256 quantity, uint256 totalPrice, address paymentToken, uint256 vestingPeriod) returns()
func (_RAgentMarket *RAgentMarketTransactor) BidRagent(opts *bind.TransactOpts, rAgentToken common.Address, quantity *big.Int, totalPrice *big.Int, paymentToken common.Address, vestingPeriod *big.Int) (*types.Transaction, error) {
	return _RAgentMarket.contract.Transact(opts, "bidRagent", rAgentToken, quantity, totalPrice, paymentToken, vestingPeriod)
}

// BidRagent is a paid mutator transaction binding the contract method 0xda95d4ac.
//
// Solidity: function bidRagent(address rAgentToken, uint256 quantity, uint256 totalPrice, address paymentToken, uint256 vestingPeriod) returns()
func (_RAgentMarket *RAgentMarketSession) BidRagent(rAgentToken common.Address, quantity *big.Int, totalPrice *big.Int, paymentToken common.Address, vestingPeriod *big.Int) (*types.Transaction, error) {
	return _RAgentMarket.Contract.BidRagent(&_RAgentMarket.TransactOpts, rAgentToken, quantity, totalPrice, paymentToken, vestingPeriod)
}

// BidRagent is a paid mutator transaction binding the contract method 0xda95d4ac.
//
// Solidity: function bidRagent(address rAgentToken, uint256 quantity, uint256 totalPrice, address paymentToken, uint256 vestingPeriod) returns()
func (_RAgentMarket *RAgentMarketTransactorSession) BidRagent(rAgentToken common.Address, quantity *big.Int, totalPrice *big.Int, paymentToken common.Address, vestingPeriod *big.Int) (*types.Transaction, error) {
	return _RAgentMarket.Contract.BidRagent(&_RAgentMarket.TransactOpts, rAgentToken, quantity, totalPrice, paymentToken, vestingPeriod)
}

// BuyRagent is a paid mutator transaction binding the contract method 0x7b7ca14e.
//
// Solidity: function buyRagent(address rAgentToken, uint256 quantity, uint256 totalPrice, uint256 nonce, bytes signature) returns((address,uint256)[])
func (_RAgentMarket *RAgentMarketTransactor) BuyRagent(opts *bind.TransactOpts, rAgentToken common.Address, quantity *big.Int, totalPrice *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _RAgentMarket.contract.Transact(opts, "buyRagent", rAgentToken, quantity, totalPrice, nonce, signature)
}

// BuyRagent is a paid mutator transaction binding the contract method 0x7b7ca14e.
//
// Solidity: function buyRagent(address rAgentToken, uint256 quantity, uint256 totalPrice, uint256 nonce, bytes signature) returns((address,uint256)[])
func (_RAgentMarket *RAgentMarketSession) BuyRagent(rAgentToken common.Address, quantity *big.Int, totalPrice *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _RAgentMarket.Contract.BuyRagent(&_RAgentMarket.TransactOpts, rAgentToken, quantity, totalPrice, nonce, signature)
}

// BuyRagent is a paid mutator transaction binding the contract method 0x7b7ca14e.
//
// Solidity: function buyRagent(address rAgentToken, uint256 quantity, uint256 totalPrice, uint256 nonce, bytes signature) returns((address,uint256)[])
func (_RAgentMarket *RAgentMarketTransactorSession) BuyRagent(rAgentToken common.Address, quantity *big.Int, totalPrice *big.Int, nonce *big.Int, signature []byte) (*types.Transaction, error) {
	return _RAgentMarket.Contract.BuyRagent(&_RAgentMarket.TransactOpts, rAgentToken, quantity, totalPrice, nonce, signature)
}

// ReleaseBid is a paid mutator transaction binding the contract method 0x80fd6c6a.
//
// Solidity: function releaseBid(uint256 bidId) returns()
func (_RAgentMarket *RAgentMarketTransactor) ReleaseBid(opts *bind.TransactOpts, bidId *big.Int) (*types.Transaction, error) {
	return _RAgentMarket.contract.Transact(opts, "releaseBid", bidId)
}

// ReleaseBid is a paid mutator transaction binding the contract method 0x80fd6c6a.
//
// Solidity: function releaseBid(uint256 bidId) returns()
func (_RAgentMarket *RAgentMarketSession) ReleaseBid(bidId *big.Int) (*types.Transaction, error) {
	return _RAgentMarket.Contract.ReleaseBid(&_RAgentMarket.TransactOpts, bidId)
}

// ReleaseBid is a paid mutator transaction binding the contract method 0x80fd6c6a.
//
// Solidity: function releaseBid(uint256 bidId) returns()
func (_RAgentMarket *RAgentMarketTransactorSession) ReleaseBid(bidId *big.Int) (*types.Transaction, error) {
	return _RAgentMarket.Contract.ReleaseBid(&_RAgentMarket.TransactOpts, bidId)
}

// RemoveBid is a paid mutator transaction binding the contract method 0xb320f459.
//
// Solidity: function removeBid(uint256 bidId) returns()
func (_RAgentMarket *RAgentMarketTransactor) RemoveBid(opts *bind.TransactOpts, bidId *big.Int) (*types.Transaction, error) {
	return _RAgentMarket.contract.Transact(opts, "removeBid", bidId)
}

// RemoveBid is a paid mutator transaction binding the contract method 0xb320f459.
//
// Solidity: function removeBid(uint256 bidId) returns()
func (_RAgentMarket *RAgentMarketSession) RemoveBid(bidId *big.Int) (*types.Transaction, error) {
	return _RAgentMarket.Contract.RemoveBid(&_RAgentMarket.TransactOpts, bidId)
}

// RemoveBid is a paid mutator transaction binding the contract method 0xb320f459.
//
// Solidity: function removeBid(uint256 bidId) returns()
func (_RAgentMarket *RAgentMarketTransactorSession) RemoveBid(bidId *big.Int) (*types.Transaction, error) {
	return _RAgentMarket.Contract.RemoveBid(&_RAgentMarket.TransactOpts, bidId)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x219674c6.
//
// Solidity: function removeSigner(address[] _signers) returns()
func (_RAgentMarket *RAgentMarketTransactor) RemoveSigner(opts *bind.TransactOpts, _signers []common.Address) (*types.Transaction, error) {
	return _RAgentMarket.contract.Transact(opts, "removeSigner", _signers)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x219674c6.
//
// Solidity: function removeSigner(address[] _signers) returns()
func (_RAgentMarket *RAgentMarketSession) RemoveSigner(_signers []common.Address) (*types.Transaction, error) {
	return _RAgentMarket.Contract.RemoveSigner(&_RAgentMarket.TransactOpts, _signers)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x219674c6.
//
// Solidity: function removeSigner(address[] _signers) returns()
func (_RAgentMarket *RAgentMarketTransactorSession) RemoveSigner(_signers []common.Address) (*types.Transaction, error) {
	return _RAgentMarket.Contract.RemoveSigner(&_RAgentMarket.TransactOpts, _signers)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RAgentMarket *RAgentMarketTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RAgentMarket.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RAgentMarket *RAgentMarketSession) RenounceOwnership() (*types.Transaction, error) {
	return _RAgentMarket.Contract.RenounceOwnership(&_RAgentMarket.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RAgentMarket *RAgentMarketTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RAgentMarket.Contract.RenounceOwnership(&_RAgentMarket.TransactOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _sAdmin) returns()
func (_RAgentMarket *RAgentMarketTransactor) SetAdmin(opts *bind.TransactOpts, _sAdmin common.Address) (*types.Transaction, error) {
	return _RAgentMarket.contract.Transact(opts, "setAdmin", _sAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _sAdmin) returns()
func (_RAgentMarket *RAgentMarketSession) SetAdmin(_sAdmin common.Address) (*types.Transaction, error) {
	return _RAgentMarket.Contract.SetAdmin(&_RAgentMarket.TransactOpts, _sAdmin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _sAdmin) returns()
func (_RAgentMarket *RAgentMarketTransactorSession) SetAdmin(_sAdmin common.Address) (*types.Transaction, error) {
	return _RAgentMarket.Contract.SetAdmin(&_RAgentMarket.TransactOpts, _sAdmin)
}

// SetExecutorAndAdmin is a paid mutator transaction binding the contract method 0x1210d833.
//
// Solidity: function setExecutorAndAdmin(address _sExecutor, address _sAdmin) returns()
func (_RAgentMarket *RAgentMarketTransactor) SetExecutorAndAdmin(opts *bind.TransactOpts, _sExecutor common.Address, _sAdmin common.Address) (*types.Transaction, error) {
	return _RAgentMarket.contract.Transact(opts, "setExecutorAndAdmin", _sExecutor, _sAdmin)
}

// SetExecutorAndAdmin is a paid mutator transaction binding the contract method 0x1210d833.
//
// Solidity: function setExecutorAndAdmin(address _sExecutor, address _sAdmin) returns()
func (_RAgentMarket *RAgentMarketSession) SetExecutorAndAdmin(_sExecutor common.Address, _sAdmin common.Address) (*types.Transaction, error) {
	return _RAgentMarket.Contract.SetExecutorAndAdmin(&_RAgentMarket.TransactOpts, _sExecutor, _sAdmin)
}

// SetExecutorAndAdmin is a paid mutator transaction binding the contract method 0x1210d833.
//
// Solidity: function setExecutorAndAdmin(address _sExecutor, address _sAdmin) returns()
func (_RAgentMarket *RAgentMarketTransactorSession) SetExecutorAndAdmin(_sExecutor common.Address, _sAdmin common.Address) (*types.Transaction, error) {
	return _RAgentMarket.Contract.SetExecutorAndAdmin(&_RAgentMarket.TransactOpts, _sExecutor, _sAdmin)
}

// SetPaymentToken is a paid mutator transaction binding the contract method 0x6a326ab1.
//
// Solidity: function setPaymentToken(address _sPaymentToken) returns()
func (_RAgentMarket *RAgentMarketTransactor) SetPaymentToken(opts *bind.TransactOpts, _sPaymentToken common.Address) (*types.Transaction, error) {
	return _RAgentMarket.contract.Transact(opts, "setPaymentToken", _sPaymentToken)
}

// SetPaymentToken is a paid mutator transaction binding the contract method 0x6a326ab1.
//
// Solidity: function setPaymentToken(address _sPaymentToken) returns()
func (_RAgentMarket *RAgentMarketSession) SetPaymentToken(_sPaymentToken common.Address) (*types.Transaction, error) {
	return _RAgentMarket.Contract.SetPaymentToken(&_RAgentMarket.TransactOpts, _sPaymentToken)
}

// SetPaymentToken is a paid mutator transaction binding the contract method 0x6a326ab1.
//
// Solidity: function setPaymentToken(address _sPaymentToken) returns()
func (_RAgentMarket *RAgentMarketTransactorSession) SetPaymentToken(_sPaymentToken common.Address) (*types.Transaction, error) {
	return _RAgentMarket.Contract.SetPaymentToken(&_RAgentMarket.TransactOpts, _sPaymentToken)
}

// SetSigner is a paid mutator transaction binding the contract method 0x22ee82db.
//
// Solidity: function setSigner(address[] _signers) returns()
func (_RAgentMarket *RAgentMarketTransactor) SetSigner(opts *bind.TransactOpts, _signers []common.Address) (*types.Transaction, error) {
	return _RAgentMarket.contract.Transact(opts, "setSigner", _signers)
}

// SetSigner is a paid mutator transaction binding the contract method 0x22ee82db.
//
// Solidity: function setSigner(address[] _signers) returns()
func (_RAgentMarket *RAgentMarketSession) SetSigner(_signers []common.Address) (*types.Transaction, error) {
	return _RAgentMarket.Contract.SetSigner(&_RAgentMarket.TransactOpts, _signers)
}

// SetSigner is a paid mutator transaction binding the contract method 0x22ee82db.
//
// Solidity: function setSigner(address[] _signers) returns()
func (_RAgentMarket *RAgentMarketTransactorSession) SetSigner(_signers []common.Address) (*types.Transaction, error) {
	return _RAgentMarket.Contract.SetSigner(&_RAgentMarket.TransactOpts, _signers)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _sTreasury) returns()
func (_RAgentMarket *RAgentMarketTransactor) SetTreasury(opts *bind.TransactOpts, _sTreasury common.Address) (*types.Transaction, error) {
	return _RAgentMarket.contract.Transact(opts, "setTreasury", _sTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _sTreasury) returns()
func (_RAgentMarket *RAgentMarketSession) SetTreasury(_sTreasury common.Address) (*types.Transaction, error) {
	return _RAgentMarket.Contract.SetTreasury(&_RAgentMarket.TransactOpts, _sTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _sTreasury) returns()
func (_RAgentMarket *RAgentMarketTransactorSession) SetTreasury(_sTreasury common.Address) (*types.Transaction, error) {
	return _RAgentMarket.Contract.SetTreasury(&_RAgentMarket.TransactOpts, _sTreasury)
}

// SetVestingPeriod is a paid mutator transaction binding the contract method 0x1f65a7e6.
//
// Solidity: function setVestingPeriod(uint256[] _vestingPeriods, bool _isActive) returns()
func (_RAgentMarket *RAgentMarketTransactor) SetVestingPeriod(opts *bind.TransactOpts, _vestingPeriods []*big.Int, _isActive bool) (*types.Transaction, error) {
	return _RAgentMarket.contract.Transact(opts, "setVestingPeriod", _vestingPeriods, _isActive)
}

// SetVestingPeriod is a paid mutator transaction binding the contract method 0x1f65a7e6.
//
// Solidity: function setVestingPeriod(uint256[] _vestingPeriods, bool _isActive) returns()
func (_RAgentMarket *RAgentMarketSession) SetVestingPeriod(_vestingPeriods []*big.Int, _isActive bool) (*types.Transaction, error) {
	return _RAgentMarket.Contract.SetVestingPeriod(&_RAgentMarket.TransactOpts, _vestingPeriods, _isActive)
}

// SetVestingPeriod is a paid mutator transaction binding the contract method 0x1f65a7e6.
//
// Solidity: function setVestingPeriod(uint256[] _vestingPeriods, bool _isActive) returns()
func (_RAgentMarket *RAgentMarketTransactorSession) SetVestingPeriod(_vestingPeriods []*big.Int, _isActive bool) (*types.Transaction, error) {
	return _RAgentMarket.Contract.SetVestingPeriod(&_RAgentMarket.TransactOpts, _vestingPeriods, _isActive)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RAgentMarket *RAgentMarketTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RAgentMarket.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RAgentMarket *RAgentMarketSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RAgentMarket.Contract.TransferOwnership(&_RAgentMarket.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RAgentMarket *RAgentMarketTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RAgentMarket.Contract.TransferOwnership(&_RAgentMarket.TransactOpts, newOwner)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address token, uint256 amount) returns()
func (_RAgentMarket *RAgentMarketTransactor) WithdrawERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RAgentMarket.contract.Transact(opts, "withdrawERC20", token, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address token, uint256 amount) returns()
func (_RAgentMarket *RAgentMarketSession) WithdrawERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RAgentMarket.Contract.WithdrawERC20(&_RAgentMarket.TransactOpts, token, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address token, uint256 amount) returns()
func (_RAgentMarket *RAgentMarketTransactorSession) WithdrawERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _RAgentMarket.Contract.WithdrawERC20(&_RAgentMarket.TransactOpts, token, amount)
}

// RAgentMarketBidAgreededIterator is returned from FilterBidAgreeded and is used to iterate over the raw logs and unpacked data for BidAgreeded events raised by the RAgentMarket contract.
type RAgentMarketBidAgreededIterator struct {
	Event *RAgentMarketBidAgreeded // Event containing the contract specifics and raw log

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
func (it *RAgentMarketBidAgreededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RAgentMarketBidAgreeded)
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
		it.Event = new(RAgentMarketBidAgreeded)
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
func (it *RAgentMarketBidAgreededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RAgentMarketBidAgreededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RAgentMarketBidAgreeded represents a BidAgreeded event raised by the RAgentMarket contract.
type RAgentMarketBidAgreeded struct {
	Bidder common.Address
	BidId  *big.Int
	Seller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBidAgreeded is a free log retrieval operation binding the contract event 0x12ca40d201bbdcab4c301edfeef15880a259c72a2e6eb60a4d595051b7ab155d.
//
// Solidity: event BidAgreeded(address indexed bidder, uint256 indexed bidId, address indexed seller)
func (_RAgentMarket *RAgentMarketFilterer) FilterBidAgreeded(opts *bind.FilterOpts, bidder []common.Address, bidId []*big.Int, seller []common.Address) (*RAgentMarketBidAgreededIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}
	var bidIdRule []interface{}
	for _, bidIdItem := range bidId {
		bidIdRule = append(bidIdRule, bidIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _RAgentMarket.contract.FilterLogs(opts, "BidAgreeded", bidderRule, bidIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &RAgentMarketBidAgreededIterator{contract: _RAgentMarket.contract, event: "BidAgreeded", logs: logs, sub: sub}, nil
}

// WatchBidAgreeded is a free log subscription operation binding the contract event 0x12ca40d201bbdcab4c301edfeef15880a259c72a2e6eb60a4d595051b7ab155d.
//
// Solidity: event BidAgreeded(address indexed bidder, uint256 indexed bidId, address indexed seller)
func (_RAgentMarket *RAgentMarketFilterer) WatchBidAgreeded(opts *bind.WatchOpts, sink chan<- *RAgentMarketBidAgreeded, bidder []common.Address, bidId []*big.Int, seller []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}
	var bidIdRule []interface{}
	for _, bidIdItem := range bidId {
		bidIdRule = append(bidIdRule, bidIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _RAgentMarket.contract.WatchLogs(opts, "BidAgreeded", bidderRule, bidIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RAgentMarketBidAgreeded)
				if err := _RAgentMarket.contract.UnpackLog(event, "BidAgreeded", log); err != nil {
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

// ParseBidAgreeded is a log parse operation binding the contract event 0x12ca40d201bbdcab4c301edfeef15880a259c72a2e6eb60a4d595051b7ab155d.
//
// Solidity: event BidAgreeded(address indexed bidder, uint256 indexed bidId, address indexed seller)
func (_RAgentMarket *RAgentMarketFilterer) ParseBidAgreeded(log types.Log) (*RAgentMarketBidAgreeded, error) {
	event := new(RAgentMarketBidAgreeded)
	if err := _RAgentMarket.contract.UnpackLog(event, "BidAgreeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RAgentMarketBidCreatedIterator is returned from FilterBidCreated and is used to iterate over the raw logs and unpacked data for BidCreated events raised by the RAgentMarket contract.
type RAgentMarketBidCreatedIterator struct {
	Event *RAgentMarketBidCreated // Event containing the contract specifics and raw log

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
func (it *RAgentMarketBidCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RAgentMarketBidCreated)
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
		it.Event = new(RAgentMarketBidCreated)
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
func (it *RAgentMarketBidCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RAgentMarketBidCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RAgentMarketBidCreated represents a BidCreated event raised by the RAgentMarket contract.
type RAgentMarketBidCreated struct {
	Bidder common.Address
	Bid    RAgentMarketPlaceBidInfo
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBidCreated is a free log retrieval operation binding the contract event 0x732e161d31bbd014d1b9e79ea6c18e7151d20d2e6ff90ad2215df55978363089.
//
// Solidity: event BidCreated(address indexed bidder, (uint256,address,uint256,uint256,uint256,address,address,uint256,bool,bool,address) bid)
func (_RAgentMarket *RAgentMarketFilterer) FilterBidCreated(opts *bind.FilterOpts, bidder []common.Address) (*RAgentMarketBidCreatedIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _RAgentMarket.contract.FilterLogs(opts, "BidCreated", bidderRule)
	if err != nil {
		return nil, err
	}
	return &RAgentMarketBidCreatedIterator{contract: _RAgentMarket.contract, event: "BidCreated", logs: logs, sub: sub}, nil
}

// WatchBidCreated is a free log subscription operation binding the contract event 0x732e161d31bbd014d1b9e79ea6c18e7151d20d2e6ff90ad2215df55978363089.
//
// Solidity: event BidCreated(address indexed bidder, (uint256,address,uint256,uint256,uint256,address,address,uint256,bool,bool,address) bid)
func (_RAgentMarket *RAgentMarketFilterer) WatchBidCreated(opts *bind.WatchOpts, sink chan<- *RAgentMarketBidCreated, bidder []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _RAgentMarket.contract.WatchLogs(opts, "BidCreated", bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RAgentMarketBidCreated)
				if err := _RAgentMarket.contract.UnpackLog(event, "BidCreated", log); err != nil {
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

// ParseBidCreated is a log parse operation binding the contract event 0x732e161d31bbd014d1b9e79ea6c18e7151d20d2e6ff90ad2215df55978363089.
//
// Solidity: event BidCreated(address indexed bidder, (uint256,address,uint256,uint256,uint256,address,address,uint256,bool,bool,address) bid)
func (_RAgentMarket *RAgentMarketFilterer) ParseBidCreated(log types.Log) (*RAgentMarketBidCreated, error) {
	event := new(RAgentMarketBidCreated)
	if err := _RAgentMarket.contract.UnpackLog(event, "BidCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RAgentMarketBidReleasedIterator is returned from FilterBidReleased and is used to iterate over the raw logs and unpacked data for BidReleased events raised by the RAgentMarket contract.
type RAgentMarketBidReleasedIterator struct {
	Event *RAgentMarketBidReleased // Event containing the contract specifics and raw log

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
func (it *RAgentMarketBidReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RAgentMarketBidReleased)
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
		it.Event = new(RAgentMarketBidReleased)
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
func (it *RAgentMarketBidReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RAgentMarketBidReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RAgentMarketBidReleased represents a BidReleased event raised by the RAgentMarket contract.
type RAgentMarketBidReleased struct {
	Bidder common.Address
	BidId  *big.Int
	Seller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBidReleased is a free log retrieval operation binding the contract event 0xbb6507c5d47bbab2eb31717fa9b641f289e7091aa3de7298f3d85da411071bbb.
//
// Solidity: event BidReleased(address indexed bidder, uint256 indexed bidId, address indexed seller)
func (_RAgentMarket *RAgentMarketFilterer) FilterBidReleased(opts *bind.FilterOpts, bidder []common.Address, bidId []*big.Int, seller []common.Address) (*RAgentMarketBidReleasedIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}
	var bidIdRule []interface{}
	for _, bidIdItem := range bidId {
		bidIdRule = append(bidIdRule, bidIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _RAgentMarket.contract.FilterLogs(opts, "BidReleased", bidderRule, bidIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &RAgentMarketBidReleasedIterator{contract: _RAgentMarket.contract, event: "BidReleased", logs: logs, sub: sub}, nil
}

// WatchBidReleased is a free log subscription operation binding the contract event 0xbb6507c5d47bbab2eb31717fa9b641f289e7091aa3de7298f3d85da411071bbb.
//
// Solidity: event BidReleased(address indexed bidder, uint256 indexed bidId, address indexed seller)
func (_RAgentMarket *RAgentMarketFilterer) WatchBidReleased(opts *bind.WatchOpts, sink chan<- *RAgentMarketBidReleased, bidder []common.Address, bidId []*big.Int, seller []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}
	var bidIdRule []interface{}
	for _, bidIdItem := range bidId {
		bidIdRule = append(bidIdRule, bidIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _RAgentMarket.contract.WatchLogs(opts, "BidReleased", bidderRule, bidIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RAgentMarketBidReleased)
				if err := _RAgentMarket.contract.UnpackLog(event, "BidReleased", log); err != nil {
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

// ParseBidReleased is a log parse operation binding the contract event 0xbb6507c5d47bbab2eb31717fa9b641f289e7091aa3de7298f3d85da411071bbb.
//
// Solidity: event BidReleased(address indexed bidder, uint256 indexed bidId, address indexed seller)
func (_RAgentMarket *RAgentMarketFilterer) ParseBidReleased(log types.Log) (*RAgentMarketBidReleased, error) {
	event := new(RAgentMarketBidReleased)
	if err := _RAgentMarket.contract.UnpackLog(event, "BidReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RAgentMarketBidRemovedIterator is returned from FilterBidRemoved and is used to iterate over the raw logs and unpacked data for BidRemoved events raised by the RAgentMarket contract.
type RAgentMarketBidRemovedIterator struct {
	Event *RAgentMarketBidRemoved // Event containing the contract specifics and raw log

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
func (it *RAgentMarketBidRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RAgentMarketBidRemoved)
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
		it.Event = new(RAgentMarketBidRemoved)
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
func (it *RAgentMarketBidRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RAgentMarketBidRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RAgentMarketBidRemoved represents a BidRemoved event raised by the RAgentMarket contract.
type RAgentMarketBidRemoved struct {
	Bidder common.Address
	BidId  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBidRemoved is a free log retrieval operation binding the contract event 0x7fd2c6472e4a5cf47dd45d6e616062bb6b2a25c19206eaa78f5ac8ea80cd5f97.
//
// Solidity: event BidRemoved(address indexed bidder, uint256 indexed bidId)
func (_RAgentMarket *RAgentMarketFilterer) FilterBidRemoved(opts *bind.FilterOpts, bidder []common.Address, bidId []*big.Int) (*RAgentMarketBidRemovedIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}
	var bidIdRule []interface{}
	for _, bidIdItem := range bidId {
		bidIdRule = append(bidIdRule, bidIdItem)
	}

	logs, sub, err := _RAgentMarket.contract.FilterLogs(opts, "BidRemoved", bidderRule, bidIdRule)
	if err != nil {
		return nil, err
	}
	return &RAgentMarketBidRemovedIterator{contract: _RAgentMarket.contract, event: "BidRemoved", logs: logs, sub: sub}, nil
}

// WatchBidRemoved is a free log subscription operation binding the contract event 0x7fd2c6472e4a5cf47dd45d6e616062bb6b2a25c19206eaa78f5ac8ea80cd5f97.
//
// Solidity: event BidRemoved(address indexed bidder, uint256 indexed bidId)
func (_RAgentMarket *RAgentMarketFilterer) WatchBidRemoved(opts *bind.WatchOpts, sink chan<- *RAgentMarketBidRemoved, bidder []common.Address, bidId []*big.Int) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}
	var bidIdRule []interface{}
	for _, bidIdItem := range bidId {
		bidIdRule = append(bidIdRule, bidIdItem)
	}

	logs, sub, err := _RAgentMarket.contract.WatchLogs(opts, "BidRemoved", bidderRule, bidIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RAgentMarketBidRemoved)
				if err := _RAgentMarket.contract.UnpackLog(event, "BidRemoved", log); err != nil {
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

// ParseBidRemoved is a log parse operation binding the contract event 0x7fd2c6472e4a5cf47dd45d6e616062bb6b2a25c19206eaa78f5ac8ea80cd5f97.
//
// Solidity: event BidRemoved(address indexed bidder, uint256 indexed bidId)
func (_RAgentMarket *RAgentMarketFilterer) ParseBidRemoved(log types.Log) (*RAgentMarketBidRemoved, error) {
	event := new(RAgentMarketBidRemoved)
	if err := _RAgentMarket.contract.UnpackLog(event, "BidRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RAgentMarketOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RAgentMarket contract.
type RAgentMarketOwnershipTransferredIterator struct {
	Event *RAgentMarketOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RAgentMarketOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RAgentMarketOwnershipTransferred)
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
		it.Event = new(RAgentMarketOwnershipTransferred)
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
func (it *RAgentMarketOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RAgentMarketOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RAgentMarketOwnershipTransferred represents a OwnershipTransferred event raised by the RAgentMarket contract.
type RAgentMarketOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RAgentMarket *RAgentMarketFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RAgentMarketOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RAgentMarket.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RAgentMarketOwnershipTransferredIterator{contract: _RAgentMarket.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RAgentMarket *RAgentMarketFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RAgentMarketOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RAgentMarket.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RAgentMarketOwnershipTransferred)
				if err := _RAgentMarket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RAgentMarket *RAgentMarketFilterer) ParseOwnershipTransferred(log types.Log) (*RAgentMarketOwnershipTransferred, error) {
	event := new(RAgentMarketOwnershipTransferred)
	if err := _RAgentMarket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RAgentMarketRAgentBoughtIterator is returned from FilterRAgentBought and is used to iterate over the raw logs and unpacked data for RAgentBought events raised by the RAgentMarket contract.
type RAgentMarketRAgentBoughtIterator struct {
	Event *RAgentMarketRAgentBought // Event containing the contract specifics and raw log

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
func (it *RAgentMarketRAgentBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RAgentMarketRAgentBought)
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
		it.Event = new(RAgentMarketRAgentBought)
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
func (it *RAgentMarketRAgentBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RAgentMarketRAgentBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RAgentMarketRAgentBought represents a RAgentBought event raised by the RAgentMarket contract.
type RAgentMarketRAgentBought struct {
	To       common.Address
	Accounts []RAgentMarketPlaceAccountInfo
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRAgentBought is a free log retrieval operation binding the contract event 0xc4f4929f6399e6918c75ddd73c9a3110f40bc3d85b5c4ad8aff781063ff867fd.
//
// Solidity: event RAgentBought(address indexed to, (address,uint256)[] accounts)
func (_RAgentMarket *RAgentMarketFilterer) FilterRAgentBought(opts *bind.FilterOpts, to []common.Address) (*RAgentMarketRAgentBoughtIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RAgentMarket.contract.FilterLogs(opts, "RAgentBought", toRule)
	if err != nil {
		return nil, err
	}
	return &RAgentMarketRAgentBoughtIterator{contract: _RAgentMarket.contract, event: "RAgentBought", logs: logs, sub: sub}, nil
}

// WatchRAgentBought is a free log subscription operation binding the contract event 0xc4f4929f6399e6918c75ddd73c9a3110f40bc3d85b5c4ad8aff781063ff867fd.
//
// Solidity: event RAgentBought(address indexed to, (address,uint256)[] accounts)
func (_RAgentMarket *RAgentMarketFilterer) WatchRAgentBought(opts *bind.WatchOpts, sink chan<- *RAgentMarketRAgentBought, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RAgentMarket.contract.WatchLogs(opts, "RAgentBought", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RAgentMarketRAgentBought)
				if err := _RAgentMarket.contract.UnpackLog(event, "RAgentBought", log); err != nil {
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

// ParseRAgentBought is a log parse operation binding the contract event 0xc4f4929f6399e6918c75ddd73c9a3110f40bc3d85b5c4ad8aff781063ff867fd.
//
// Solidity: event RAgentBought(address indexed to, (address,uint256)[] accounts)
func (_RAgentMarket *RAgentMarketFilterer) ParseRAgentBought(log types.Log) (*RAgentMarketRAgentBought, error) {
	event := new(RAgentMarketRAgentBought)
	if err := _RAgentMarket.contract.UnpackLog(event, "RAgentBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RAgentMarketSignerRemovedIterator is returned from FilterSignerRemoved and is used to iterate over the raw logs and unpacked data for SignerRemoved events raised by the RAgentMarket contract.
type RAgentMarketSignerRemovedIterator struct {
	Event *RAgentMarketSignerRemoved // Event containing the contract specifics and raw log

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
func (it *RAgentMarketSignerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RAgentMarketSignerRemoved)
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
		it.Event = new(RAgentMarketSignerRemoved)
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
func (it *RAgentMarketSignerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RAgentMarketSignerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RAgentMarketSignerRemoved represents a SignerRemoved event raised by the RAgentMarket contract.
type RAgentMarketSignerRemoved struct {
	Signer []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignerRemoved is a free log retrieval operation binding the contract event 0x28ef5208f7fe8bc7d0185b9e3c873ffe70677d01ee4d54f4bb1770af14852eaf.
//
// Solidity: event SignerRemoved(address[] signer)
func (_RAgentMarket *RAgentMarketFilterer) FilterSignerRemoved(opts *bind.FilterOpts) (*RAgentMarketSignerRemovedIterator, error) {

	logs, sub, err := _RAgentMarket.contract.FilterLogs(opts, "SignerRemoved")
	if err != nil {
		return nil, err
	}
	return &RAgentMarketSignerRemovedIterator{contract: _RAgentMarket.contract, event: "SignerRemoved", logs: logs, sub: sub}, nil
}

// WatchSignerRemoved is a free log subscription operation binding the contract event 0x28ef5208f7fe8bc7d0185b9e3c873ffe70677d01ee4d54f4bb1770af14852eaf.
//
// Solidity: event SignerRemoved(address[] signer)
func (_RAgentMarket *RAgentMarketFilterer) WatchSignerRemoved(opts *bind.WatchOpts, sink chan<- *RAgentMarketSignerRemoved) (event.Subscription, error) {

	logs, sub, err := _RAgentMarket.contract.WatchLogs(opts, "SignerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RAgentMarketSignerRemoved)
				if err := _RAgentMarket.contract.UnpackLog(event, "SignerRemoved", log); err != nil {
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

// ParseSignerRemoved is a log parse operation binding the contract event 0x28ef5208f7fe8bc7d0185b9e3c873ffe70677d01ee4d54f4bb1770af14852eaf.
//
// Solidity: event SignerRemoved(address[] signer)
func (_RAgentMarket *RAgentMarketFilterer) ParseSignerRemoved(log types.Log) (*RAgentMarketSignerRemoved, error) {
	event := new(RAgentMarketSignerRemoved)
	if err := _RAgentMarket.contract.UnpackLog(event, "SignerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RAgentMarketSignersAddedIterator is returned from FilterSignersAdded and is used to iterate over the raw logs and unpacked data for SignersAdded events raised by the RAgentMarket contract.
type RAgentMarketSignersAddedIterator struct {
	Event *RAgentMarketSignersAdded // Event containing the contract specifics and raw log

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
func (it *RAgentMarketSignersAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RAgentMarketSignersAdded)
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
		it.Event = new(RAgentMarketSignersAdded)
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
func (it *RAgentMarketSignersAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RAgentMarketSignersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RAgentMarketSignersAdded represents a SignersAdded event raised by the RAgentMarket contract.
type RAgentMarketSignersAdded struct {
	Signers []common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignersAdded is a free log retrieval operation binding the contract event 0xab6b3af919a4fa21f9a36b3b9aca5e07fa4426c931e529fa4bde84b5af03f997.
//
// Solidity: event SignersAdded(address[] signers)
func (_RAgentMarket *RAgentMarketFilterer) FilterSignersAdded(opts *bind.FilterOpts) (*RAgentMarketSignersAddedIterator, error) {

	logs, sub, err := _RAgentMarket.contract.FilterLogs(opts, "SignersAdded")
	if err != nil {
		return nil, err
	}
	return &RAgentMarketSignersAddedIterator{contract: _RAgentMarket.contract, event: "SignersAdded", logs: logs, sub: sub}, nil
}

// WatchSignersAdded is a free log subscription operation binding the contract event 0xab6b3af919a4fa21f9a36b3b9aca5e07fa4426c931e529fa4bde84b5af03f997.
//
// Solidity: event SignersAdded(address[] signers)
func (_RAgentMarket *RAgentMarketFilterer) WatchSignersAdded(opts *bind.WatchOpts, sink chan<- *RAgentMarketSignersAdded) (event.Subscription, error) {

	logs, sub, err := _RAgentMarket.contract.WatchLogs(opts, "SignersAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RAgentMarketSignersAdded)
				if err := _RAgentMarket.contract.UnpackLog(event, "SignersAdded", log); err != nil {
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

// ParseSignersAdded is a log parse operation binding the contract event 0xab6b3af919a4fa21f9a36b3b9aca5e07fa4426c931e529fa4bde84b5af03f997.
//
// Solidity: event SignersAdded(address[] signers)
func (_RAgentMarket *RAgentMarketFilterer) ParseSignersAdded(log types.Log) (*RAgentMarketSignersAdded, error) {
	event := new(RAgentMarketSignersAdded)
	if err := _RAgentMarket.contract.UnpackLog(event, "SignersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
