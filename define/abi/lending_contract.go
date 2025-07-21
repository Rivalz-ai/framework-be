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

// ILendingBorrowInfo is an auto generated low-level Go binding around an user-defined struct.
type ILendingBorrowInfo struct {
	Borrower    common.Address
	StartBlock  *big.Int
	EndBlock    *big.Int
	FeePerBlock *big.Int
}

// ILendingLendInfo is an auto generated low-level Go binding around an user-defined struct.
type ILendingLendInfo struct {
	TokenId    *big.Int
	StartBlock *big.Int
	IsLent     bool
	RewardDebt *big.Int
}

// LendingMetaData contains all meta data concerning the Lending contract.
var LendingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lendingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_borrowToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BorrowPeriodExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotRecoverRewardToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ContractPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmergencyUnstakeNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedsBorrowLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDecimals\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDurationDays\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotBorrower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotTokenLender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotTokenOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyBorrowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenCurrentlyBorrowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotLent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotOverdue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBlocksPerDay\",\"type\":\"uint256\"}],\"name\":\"BlockTimeConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"EmergencyModeEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"EmergencyUnstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ForceUnstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRewardPerBlock\",\"type\":\"uint256\"}],\"name\":\"RewardPerBlockUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePerBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"TokenBorrowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TokenLent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TokenReturned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TokenWithdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRECISION_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_availableDurations\",\"type\":\"uint256\"}],\"name\":\"addAvailableDurations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"durationInDays\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"borrowedTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"borrowedTokensOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"durationInDays\",\"type\":\"uint256\"}],\"name\":\"calculateBorrowFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"durationInDays\",\"type\":\"uint256\"}],\"name\":\"checkAvailableDuration\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"checkOverdueToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"days_\",\"type\":\"uint256\"}],\"name\":\"daysToBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"emergencyUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"emergencyUnstakeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyUnstakeEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"forceUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAvailableDurations\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getBorrowInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePerBlock\",\"type\":\"uint256\"}],\"internalType\":\"structILending.BorrowInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"getExpiredBorrowedTokensOf\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLendInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isLent\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"internalType\":\"structILending.LendInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getPendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastActionBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"lend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"lendBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"}],\"name\":\"lentTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"}],\"name\":\"lentTokensOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"returnOverdueToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"returnOverdueTokenBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sAccRewardPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sAvailableDurations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sBaseFeePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sBlocksPerDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sBorrowToken\",\"outputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sBorrowerTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sCurrentRewardPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sLastUpdateBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sLenderTokenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sLendingToken\",\"outputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sProtocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sRewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sTokenBorrows\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePerBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sTokenLender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sTotalBorrowedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sTotalFeePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sTotalLentTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sTotalReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sTotalRewardClaimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sUserLends\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isLent\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sUserReserveBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseFeePerBlock\",\"type\":\"uint256\"}],\"name\":\"setBaseFeePerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setEmergencyUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_protocolFee\",\"type\":\"uint256\"}],\"name\":\"setProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfBorrowerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfLenderByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"unstakeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blocksPerDay\",\"type\":\"uint256\"}],\"name\":\"updateBlockTimeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validateState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAllReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LendingABI is the input ABI used to generate the binding from.
// Deprecated: Use LendingMetaData.ABI instead.
var LendingABI = LendingMetaData.ABI

// Lending is an auto generated Go binding around an Ethereum contract.
type Lending struct {
	LendingCaller     // Read-only binding to the contract
	LendingTransactor // Write-only binding to the contract
	LendingFilterer   // Log filterer for contract events
}

// LendingCaller is an auto generated read-only Go binding around an Ethereum contract.
type LendingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LendingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LendingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LendingSession struct {
	Contract     *Lending          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LendingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LendingCallerSession struct {
	Contract *LendingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LendingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LendingTransactorSession struct {
	Contract     *LendingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LendingRaw is an auto generated low-level Go binding around an Ethereum contract.
type LendingRaw struct {
	Contract *Lending // Generic contract binding to access the raw methods on
}

// LendingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LendingCallerRaw struct {
	Contract *LendingCaller // Generic read-only contract binding to access the raw methods on
}

// LendingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LendingTransactorRaw struct {
	Contract *LendingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLending creates a new instance of Lending, bound to a specific deployed contract.
func NewLending(address common.Address, backend bind.ContractBackend) (*Lending, error) {
	contract, err := bindLending(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lending{LendingCaller: LendingCaller{contract: contract}, LendingTransactor: LendingTransactor{contract: contract}, LendingFilterer: LendingFilterer{contract: contract}}, nil
}

// NewLendingCaller creates a new read-only instance of Lending, bound to a specific deployed contract.
func NewLendingCaller(address common.Address, caller bind.ContractCaller) (*LendingCaller, error) {
	contract, err := bindLending(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LendingCaller{contract: contract}, nil
}

// NewLendingTransactor creates a new write-only instance of Lending, bound to a specific deployed contract.
func NewLendingTransactor(address common.Address, transactor bind.ContractTransactor) (*LendingTransactor, error) {
	contract, err := bindLending(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LendingTransactor{contract: contract}, nil
}

// NewLendingFilterer creates a new log filterer instance of Lending, bound to a specific deployed contract.
func NewLendingFilterer(address common.Address, filterer bind.ContractFilterer) (*LendingFilterer, error) {
	contract, err := bindLending(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LendingFilterer{contract: contract}, nil
}

// bindLending binds a generic wrapper to an already deployed contract.
func bindLending(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LendingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lending *LendingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lending.Contract.LendingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lending *LendingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lending.Contract.LendingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lending *LendingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lending.Contract.LendingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lending *LendingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lending.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lending *LendingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lending.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lending *LendingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lending.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Lending *LendingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Lending *LendingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Lending.Contract.DEFAULTADMINROLE(&_Lending.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Lending *LendingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Lending.Contract.DEFAULTADMINROLE(&_Lending.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Lending *LendingCaller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Lending *LendingSession) OPERATORROLE() ([32]byte, error) {
	return _Lending.Contract.OPERATORROLE(&_Lending.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_Lending *LendingCallerSession) OPERATORROLE() ([32]byte, error) {
	return _Lending.Contract.OPERATORROLE(&_Lending.CallOpts)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_Lending *LendingCaller) PRECISIONFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "PRECISION_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_Lending *LendingSession) PRECISIONFACTOR() (*big.Int, error) {
	return _Lending.Contract.PRECISIONFACTOR(&_Lending.CallOpts)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_Lending *LendingCallerSession) PRECISIONFACTOR() (*big.Int, error) {
	return _Lending.Contract.PRECISIONFACTOR(&_Lending.CallOpts)
}

// BorrowedTokenCount is a free data retrieval call binding the contract method 0x121841c4.
//
// Solidity: function borrowedTokenCount(address borrower) view returns(uint256)
func (_Lending *LendingCaller) BorrowedTokenCount(opts *bind.CallOpts, borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "borrowedTokenCount", borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowedTokenCount is a free data retrieval call binding the contract method 0x121841c4.
//
// Solidity: function borrowedTokenCount(address borrower) view returns(uint256)
func (_Lending *LendingSession) BorrowedTokenCount(borrower common.Address) (*big.Int, error) {
	return _Lending.Contract.BorrowedTokenCount(&_Lending.CallOpts, borrower)
}

// BorrowedTokenCount is a free data retrieval call binding the contract method 0x121841c4.
//
// Solidity: function borrowedTokenCount(address borrower) view returns(uint256)
func (_Lending *LendingCallerSession) BorrowedTokenCount(borrower common.Address) (*big.Int, error) {
	return _Lending.Contract.BorrowedTokenCount(&_Lending.CallOpts, borrower)
}

// BorrowedTokensOfOwner is a free data retrieval call binding the contract method 0x80d24e2f.
//
// Solidity: function borrowedTokensOfOwner(address borrower) view returns(uint256[])
func (_Lending *LendingCaller) BorrowedTokensOfOwner(opts *bind.CallOpts, borrower common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "borrowedTokensOfOwner", borrower)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BorrowedTokensOfOwner is a free data retrieval call binding the contract method 0x80d24e2f.
//
// Solidity: function borrowedTokensOfOwner(address borrower) view returns(uint256[])
func (_Lending *LendingSession) BorrowedTokensOfOwner(borrower common.Address) ([]*big.Int, error) {
	return _Lending.Contract.BorrowedTokensOfOwner(&_Lending.CallOpts, borrower)
}

// BorrowedTokensOfOwner is a free data retrieval call binding the contract method 0x80d24e2f.
//
// Solidity: function borrowedTokensOfOwner(address borrower) view returns(uint256[])
func (_Lending *LendingCallerSession) BorrowedTokensOfOwner(borrower common.Address) ([]*big.Int, error) {
	return _Lending.Contract.BorrowedTokensOfOwner(&_Lending.CallOpts, borrower)
}

// CalculateBorrowFee is a free data retrieval call binding the contract method 0xae0f1c7f.
//
// Solidity: function calculateBorrowFee(uint256 durationInDays) view returns(uint256)
func (_Lending *LendingCaller) CalculateBorrowFee(opts *bind.CallOpts, durationInDays *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "calculateBorrowFee", durationInDays)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateBorrowFee is a free data retrieval call binding the contract method 0xae0f1c7f.
//
// Solidity: function calculateBorrowFee(uint256 durationInDays) view returns(uint256)
func (_Lending *LendingSession) CalculateBorrowFee(durationInDays *big.Int) (*big.Int, error) {
	return _Lending.Contract.CalculateBorrowFee(&_Lending.CallOpts, durationInDays)
}

// CalculateBorrowFee is a free data retrieval call binding the contract method 0xae0f1c7f.
//
// Solidity: function calculateBorrowFee(uint256 durationInDays) view returns(uint256)
func (_Lending *LendingCallerSession) CalculateBorrowFee(durationInDays *big.Int) (*big.Int, error) {
	return _Lending.Contract.CalculateBorrowFee(&_Lending.CallOpts, durationInDays)
}

// CheckAvailableDuration is a free data retrieval call binding the contract method 0x6323f3b5.
//
// Solidity: function checkAvailableDuration(uint256 durationInDays) view returns(bool)
func (_Lending *LendingCaller) CheckAvailableDuration(opts *bind.CallOpts, durationInDays *big.Int) (bool, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "checkAvailableDuration", durationInDays)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckAvailableDuration is a free data retrieval call binding the contract method 0x6323f3b5.
//
// Solidity: function checkAvailableDuration(uint256 durationInDays) view returns(bool)
func (_Lending *LendingSession) CheckAvailableDuration(durationInDays *big.Int) (bool, error) {
	return _Lending.Contract.CheckAvailableDuration(&_Lending.CallOpts, durationInDays)
}

// CheckAvailableDuration is a free data retrieval call binding the contract method 0x6323f3b5.
//
// Solidity: function checkAvailableDuration(uint256 durationInDays) view returns(bool)
func (_Lending *LendingCallerSession) CheckAvailableDuration(durationInDays *big.Int) (bool, error) {
	return _Lending.Contract.CheckAvailableDuration(&_Lending.CallOpts, durationInDays)
}

// CheckOverdueToken is a free data retrieval call binding the contract method 0x3b9eefe0.
//
// Solidity: function checkOverdueToken(uint256 tokenId) view returns(bool)
func (_Lending *LendingCaller) CheckOverdueToken(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "checkOverdueToken", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckOverdueToken is a free data retrieval call binding the contract method 0x3b9eefe0.
//
// Solidity: function checkOverdueToken(uint256 tokenId) view returns(bool)
func (_Lending *LendingSession) CheckOverdueToken(tokenId *big.Int) (bool, error) {
	return _Lending.Contract.CheckOverdueToken(&_Lending.CallOpts, tokenId)
}

// CheckOverdueToken is a free data retrieval call binding the contract method 0x3b9eefe0.
//
// Solidity: function checkOverdueToken(uint256 tokenId) view returns(bool)
func (_Lending *LendingCallerSession) CheckOverdueToken(tokenId *big.Int) (bool, error) {
	return _Lending.Contract.CheckOverdueToken(&_Lending.CallOpts, tokenId)
}

// DaysToBlocks is a free data retrieval call binding the contract method 0x13ad3fda.
//
// Solidity: function daysToBlocks(uint256 days_) view returns(uint256)
func (_Lending *LendingCaller) DaysToBlocks(opts *bind.CallOpts, days_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "daysToBlocks", days_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DaysToBlocks is a free data retrieval call binding the contract method 0x13ad3fda.
//
// Solidity: function daysToBlocks(uint256 days_) view returns(uint256)
func (_Lending *LendingSession) DaysToBlocks(days_ *big.Int) (*big.Int, error) {
	return _Lending.Contract.DaysToBlocks(&_Lending.CallOpts, days_)
}

// DaysToBlocks is a free data retrieval call binding the contract method 0x13ad3fda.
//
// Solidity: function daysToBlocks(uint256 days_) view returns(uint256)
func (_Lending *LendingCallerSession) DaysToBlocks(days_ *big.Int) (*big.Int, error) {
	return _Lending.Contract.DaysToBlocks(&_Lending.CallOpts, days_)
}

// EmergencyUnstakeEnabled is a free data retrieval call binding the contract method 0x0d6b6da0.
//
// Solidity: function emergencyUnstakeEnabled() view returns(bool)
func (_Lending *LendingCaller) EmergencyUnstakeEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "emergencyUnstakeEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EmergencyUnstakeEnabled is a free data retrieval call binding the contract method 0x0d6b6da0.
//
// Solidity: function emergencyUnstakeEnabled() view returns(bool)
func (_Lending *LendingSession) EmergencyUnstakeEnabled() (bool, error) {
	return _Lending.Contract.EmergencyUnstakeEnabled(&_Lending.CallOpts)
}

// EmergencyUnstakeEnabled is a free data retrieval call binding the contract method 0x0d6b6da0.
//
// Solidity: function emergencyUnstakeEnabled() view returns(bool)
func (_Lending *LendingCallerSession) EmergencyUnstakeEnabled() (bool, error) {
	return _Lending.Contract.EmergencyUnstakeEnabled(&_Lending.CallOpts)
}

// GetAvailableDurations is a free data retrieval call binding the contract method 0xb6b365c2.
//
// Solidity: function getAvailableDurations() view returns(uint256[])
func (_Lending *LendingCaller) GetAvailableDurations(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "getAvailableDurations")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAvailableDurations is a free data retrieval call binding the contract method 0xb6b365c2.
//
// Solidity: function getAvailableDurations() view returns(uint256[])
func (_Lending *LendingSession) GetAvailableDurations() ([]*big.Int, error) {
	return _Lending.Contract.GetAvailableDurations(&_Lending.CallOpts)
}

// GetAvailableDurations is a free data retrieval call binding the contract method 0xb6b365c2.
//
// Solidity: function getAvailableDurations() view returns(uint256[])
func (_Lending *LendingCallerSession) GetAvailableDurations() ([]*big.Int, error) {
	return _Lending.Contract.GetAvailableDurations(&_Lending.CallOpts)
}

// GetBorrowInfo is a free data retrieval call binding the contract method 0x0e137368.
//
// Solidity: function getBorrowInfo(uint256 tokenId) view returns((address,uint256,uint256,uint256))
func (_Lending *LendingCaller) GetBorrowInfo(opts *bind.CallOpts, tokenId *big.Int) (ILendingBorrowInfo, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "getBorrowInfo", tokenId)

	if err != nil {
		return *new(ILendingBorrowInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ILendingBorrowInfo)).(*ILendingBorrowInfo)

	return out0, err

}

// GetBorrowInfo is a free data retrieval call binding the contract method 0x0e137368.
//
// Solidity: function getBorrowInfo(uint256 tokenId) view returns((address,uint256,uint256,uint256))
func (_Lending *LendingSession) GetBorrowInfo(tokenId *big.Int) (ILendingBorrowInfo, error) {
	return _Lending.Contract.GetBorrowInfo(&_Lending.CallOpts, tokenId)
}

// GetBorrowInfo is a free data retrieval call binding the contract method 0x0e137368.
//
// Solidity: function getBorrowInfo(uint256 tokenId) view returns((address,uint256,uint256,uint256))
func (_Lending *LendingCallerSession) GetBorrowInfo(tokenId *big.Int) (ILendingBorrowInfo, error) {
	return _Lending.Contract.GetBorrowInfo(&_Lending.CallOpts, tokenId)
}

// GetExpiredBorrowedTokensOf is a free data retrieval call binding the contract method 0xe984778a.
//
// Solidity: function getExpiredBorrowedTokensOf(address borrower) view returns(uint256[])
func (_Lending *LendingCaller) GetExpiredBorrowedTokensOf(opts *bind.CallOpts, borrower common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "getExpiredBorrowedTokensOf", borrower)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetExpiredBorrowedTokensOf is a free data retrieval call binding the contract method 0xe984778a.
//
// Solidity: function getExpiredBorrowedTokensOf(address borrower) view returns(uint256[])
func (_Lending *LendingSession) GetExpiredBorrowedTokensOf(borrower common.Address) ([]*big.Int, error) {
	return _Lending.Contract.GetExpiredBorrowedTokensOf(&_Lending.CallOpts, borrower)
}

// GetExpiredBorrowedTokensOf is a free data retrieval call binding the contract method 0xe984778a.
//
// Solidity: function getExpiredBorrowedTokensOf(address borrower) view returns(uint256[])
func (_Lending *LendingCallerSession) GetExpiredBorrowedTokensOf(borrower common.Address) ([]*big.Int, error) {
	return _Lending.Contract.GetExpiredBorrowedTokensOf(&_Lending.CallOpts, borrower)
}

// GetLendInfo is a free data retrieval call binding the contract method 0x4e932e72.
//
// Solidity: function getLendInfo(uint256 tokenId) view returns((uint256,uint256,bool,uint256))
func (_Lending *LendingCaller) GetLendInfo(opts *bind.CallOpts, tokenId *big.Int) (ILendingLendInfo, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "getLendInfo", tokenId)

	if err != nil {
		return *new(ILendingLendInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ILendingLendInfo)).(*ILendingLendInfo)

	return out0, err

}

// GetLendInfo is a free data retrieval call binding the contract method 0x4e932e72.
//
// Solidity: function getLendInfo(uint256 tokenId) view returns((uint256,uint256,bool,uint256))
func (_Lending *LendingSession) GetLendInfo(tokenId *big.Int) (ILendingLendInfo, error) {
	return _Lending.Contract.GetLendInfo(&_Lending.CallOpts, tokenId)
}

// GetLendInfo is a free data retrieval call binding the contract method 0x4e932e72.
//
// Solidity: function getLendInfo(uint256 tokenId) view returns((uint256,uint256,bool,uint256))
func (_Lending *LendingCallerSession) GetLendInfo(tokenId *big.Int) (ILendingLendInfo, error) {
	return _Lending.Contract.GetLendInfo(&_Lending.CallOpts, tokenId)
}

// GetPendingRewards is a free data retrieval call binding the contract method 0xf6ed2017.
//
// Solidity: function getPendingRewards(address user) view returns(uint256)
func (_Lending *LendingCaller) GetPendingRewards(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "getPendingRewards", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingRewards is a free data retrieval call binding the contract method 0xf6ed2017.
//
// Solidity: function getPendingRewards(address user) view returns(uint256)
func (_Lending *LendingSession) GetPendingRewards(user common.Address) (*big.Int, error) {
	return _Lending.Contract.GetPendingRewards(&_Lending.CallOpts, user)
}

// GetPendingRewards is a free data retrieval call binding the contract method 0xf6ed2017.
//
// Solidity: function getPendingRewards(address user) view returns(uint256)
func (_Lending *LendingCallerSession) GetPendingRewards(user common.Address) (*big.Int, error) {
	return _Lending.Contract.GetPendingRewards(&_Lending.CallOpts, user)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Lending *LendingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Lending *LendingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Lending.Contract.GetRoleAdmin(&_Lending.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Lending *LendingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Lending.Contract.GetRoleAdmin(&_Lending.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Lending *LendingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Lending *LendingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Lending.Contract.HasRole(&_Lending.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Lending *LendingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Lending.Contract.HasRole(&_Lending.CallOpts, role, account)
}

// LastActionBlock is a free data retrieval call binding the contract method 0xe5d84527.
//
// Solidity: function lastActionBlock(address ) view returns(uint256)
func (_Lending *LendingCaller) LastActionBlock(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "lastActionBlock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastActionBlock is a free data retrieval call binding the contract method 0xe5d84527.
//
// Solidity: function lastActionBlock(address ) view returns(uint256)
func (_Lending *LendingSession) LastActionBlock(arg0 common.Address) (*big.Int, error) {
	return _Lending.Contract.LastActionBlock(&_Lending.CallOpts, arg0)
}

// LastActionBlock is a free data retrieval call binding the contract method 0xe5d84527.
//
// Solidity: function lastActionBlock(address ) view returns(uint256)
func (_Lending *LendingCallerSession) LastActionBlock(arg0 common.Address) (*big.Int, error) {
	return _Lending.Contract.LastActionBlock(&_Lending.CallOpts, arg0)
}

// LentTokenCount is a free data retrieval call binding the contract method 0xe2a6b02e.
//
// Solidity: function lentTokenCount(address lender) view returns(uint256)
func (_Lending *LendingCaller) LentTokenCount(opts *bind.CallOpts, lender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "lentTokenCount", lender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LentTokenCount is a free data retrieval call binding the contract method 0xe2a6b02e.
//
// Solidity: function lentTokenCount(address lender) view returns(uint256)
func (_Lending *LendingSession) LentTokenCount(lender common.Address) (*big.Int, error) {
	return _Lending.Contract.LentTokenCount(&_Lending.CallOpts, lender)
}

// LentTokenCount is a free data retrieval call binding the contract method 0xe2a6b02e.
//
// Solidity: function lentTokenCount(address lender) view returns(uint256)
func (_Lending *LendingCallerSession) LentTokenCount(lender common.Address) (*big.Int, error) {
	return _Lending.Contract.LentTokenCount(&_Lending.CallOpts, lender)
}

// LentTokensOfOwner is a free data retrieval call binding the contract method 0xe3c9c4aa.
//
// Solidity: function lentTokensOfOwner(address lender) view returns(uint256[])
func (_Lending *LendingCaller) LentTokensOfOwner(opts *bind.CallOpts, lender common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "lentTokensOfOwner", lender)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// LentTokensOfOwner is a free data retrieval call binding the contract method 0xe3c9c4aa.
//
// Solidity: function lentTokensOfOwner(address lender) view returns(uint256[])
func (_Lending *LendingSession) LentTokensOfOwner(lender common.Address) ([]*big.Int, error) {
	return _Lending.Contract.LentTokensOfOwner(&_Lending.CallOpts, lender)
}

// LentTokensOfOwner is a free data retrieval call binding the contract method 0xe3c9c4aa.
//
// Solidity: function lentTokensOfOwner(address lender) view returns(uint256[])
func (_Lending *LendingCallerSession) LentTokensOfOwner(lender common.Address) ([]*big.Int, error) {
	return _Lending.Contract.LentTokensOfOwner(&_Lending.CallOpts, lender)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Lending *LendingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Lending *LendingSession) Paused() (bool, error) {
	return _Lending.Contract.Paused(&_Lending.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Lending *LendingCallerSession) Paused() (bool, error) {
	return _Lending.Contract.Paused(&_Lending.CallOpts)
}

// SAccRewardPerShare is a free data retrieval call binding the contract method 0xfd4e5cde.
//
// Solidity: function sAccRewardPerShare() view returns(uint256)
func (_Lending *LendingCaller) SAccRewardPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "sAccRewardPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SAccRewardPerShare is a free data retrieval call binding the contract method 0xfd4e5cde.
//
// Solidity: function sAccRewardPerShare() view returns(uint256)
func (_Lending *LendingSession) SAccRewardPerShare() (*big.Int, error) {
	return _Lending.Contract.SAccRewardPerShare(&_Lending.CallOpts)
}

// SAccRewardPerShare is a free data retrieval call binding the contract method 0xfd4e5cde.
//
// Solidity: function sAccRewardPerShare() view returns(uint256)
func (_Lending *LendingCallerSession) SAccRewardPerShare() (*big.Int, error) {
	return _Lending.Contract.SAccRewardPerShare(&_Lending.CallOpts)
}

// SAvailableDurations is a free data retrieval call binding the contract method 0xa8387fce.
//
// Solidity: function sAvailableDurations(uint256 ) view returns(uint256)
func (_Lending *LendingCaller) SAvailableDurations(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "sAvailableDurations", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SAvailableDurations is a free data retrieval call binding the contract method 0xa8387fce.
//
// Solidity: function sAvailableDurations(uint256 ) view returns(uint256)
func (_Lending *LendingSession) SAvailableDurations(arg0 *big.Int) (*big.Int, error) {
	return _Lending.Contract.SAvailableDurations(&_Lending.CallOpts, arg0)
}

// SAvailableDurations is a free data retrieval call binding the contract method 0xa8387fce.
//
// Solidity: function sAvailableDurations(uint256 ) view returns(uint256)
func (_Lending *LendingCallerSession) SAvailableDurations(arg0 *big.Int) (*big.Int, error) {
	return _Lending.Contract.SAvailableDurations(&_Lending.CallOpts, arg0)
}

// SBaseFeePerBlock is a free data retrieval call binding the contract method 0xfe66c349.
//
// Solidity: function sBaseFeePerBlock() view returns(uint256)
func (_Lending *LendingCaller) SBaseFeePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "sBaseFeePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SBaseFeePerBlock is a free data retrieval call binding the contract method 0xfe66c349.
//
// Solidity: function sBaseFeePerBlock() view returns(uint256)
func (_Lending *LendingSession) SBaseFeePerBlock() (*big.Int, error) {
	return _Lending.Contract.SBaseFeePerBlock(&_Lending.CallOpts)
}

// SBaseFeePerBlock is a free data retrieval call binding the contract method 0xfe66c349.
//
// Solidity: function sBaseFeePerBlock() view returns(uint256)
func (_Lending *LendingCallerSession) SBaseFeePerBlock() (*big.Int, error) {
	return _Lending.Contract.SBaseFeePerBlock(&_Lending.CallOpts)
}

// SBlocksPerDay is a free data retrieval call binding the contract method 0xcbdf5781.
//
// Solidity: function sBlocksPerDay() view returns(uint256)
func (_Lending *LendingCaller) SBlocksPerDay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "sBlocksPerDay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SBlocksPerDay is a free data retrieval call binding the contract method 0xcbdf5781.
//
// Solidity: function sBlocksPerDay() view returns(uint256)
func (_Lending *LendingSession) SBlocksPerDay() (*big.Int, error) {
	return _Lending.Contract.SBlocksPerDay(&_Lending.CallOpts)
}

// SBlocksPerDay is a free data retrieval call binding the contract method 0xcbdf5781.
//
// Solidity: function sBlocksPerDay() view returns(uint256)
func (_Lending *LendingCallerSession) SBlocksPerDay() (*big.Int, error) {
	return _Lending.Contract.SBlocksPerDay(&_Lending.CallOpts)
}

// SBorrowToken is a free data retrieval call binding the contract method 0xc41dbdbc.
//
// Solidity: function sBorrowToken() view returns(address)
func (_Lending *LendingCaller) SBorrowToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "sBorrowToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SBorrowToken is a free data retrieval call binding the contract method 0xc41dbdbc.
//
// Solidity: function sBorrowToken() view returns(address)
func (_Lending *LendingSession) SBorrowToken() (common.Address, error) {
	return _Lending.Contract.SBorrowToken(&_Lending.CallOpts)
}

// SBorrowToken is a free data retrieval call binding the contract method 0xc41dbdbc.
//
// Solidity: function sBorrowToken() view returns(address)
func (_Lending *LendingCallerSession) SBorrowToken() (common.Address, error) {
	return _Lending.Contract.SBorrowToken(&_Lending.CallOpts)
}

// SBorrowerTokenCount is a free data retrieval call binding the contract method 0xa8464b31.
//
// Solidity: function sBorrowerTokenCount(address ) view returns(uint256)
func (_Lending *LendingCaller) SBorrowerTokenCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "sBorrowerTokenCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SBorrowerTokenCount is a free data retrieval call binding the contract method 0xa8464b31.
//
// Solidity: function sBorrowerTokenCount(address ) view returns(uint256)
func (_Lending *LendingSession) SBorrowerTokenCount(arg0 common.Address) (*big.Int, error) {
	return _Lending.Contract.SBorrowerTokenCount(&_Lending.CallOpts, arg0)
}

// SBorrowerTokenCount is a free data retrieval call binding the contract method 0xa8464b31.
//
// Solidity: function sBorrowerTokenCount(address ) view returns(uint256)
func (_Lending *LendingCallerSession) SBorrowerTokenCount(arg0 common.Address) (*big.Int, error) {
	return _Lending.Contract.SBorrowerTokenCount(&_Lending.CallOpts, arg0)
}

// SCurrentRewardPerBlock is a free data retrieval call binding the contract method 0x4e38e9fd.
//
// Solidity: function sCurrentRewardPerBlock() view returns(uint256)
func (_Lending *LendingCaller) SCurrentRewardPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "sCurrentRewardPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SCurrentRewardPerBlock is a free data retrieval call binding the contract method 0x4e38e9fd.
//
// Solidity: function sCurrentRewardPerBlock() view returns(uint256)
func (_Lending *LendingSession) SCurrentRewardPerBlock() (*big.Int, error) {
	return _Lending.Contract.SCurrentRewardPerBlock(&_Lending.CallOpts)
}

// SCurrentRewardPerBlock is a free data retrieval call binding the contract method 0x4e38e9fd.
//
// Solidity: function sCurrentRewardPerBlock() view returns(uint256)
func (_Lending *LendingCallerSession) SCurrentRewardPerBlock() (*big.Int, error) {
	return _Lending.Contract.SCurrentRewardPerBlock(&_Lending.CallOpts)
}

// SLastUpdateBlock is a free data retrieval call binding the contract method 0x053bc9f1.
//
// Solidity: function sLastUpdateBlock() view returns(uint256)
func (_Lending *LendingCaller) SLastUpdateBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "sLastUpdateBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SLastUpdateBlock is a free data retrieval call binding the contract method 0x053bc9f1.
//
// Solidity: function sLastUpdateBlock() view returns(uint256)
func (_Lending *LendingSession) SLastUpdateBlock() (*big.Int, error) {
	return _Lending.Contract.SLastUpdateBlock(&_Lending.CallOpts)
}

// SLastUpdateBlock is a free data retrieval call binding the contract method 0x053bc9f1.
//
// Solidity: function sLastUpdateBlock() view returns(uint256)
func (_Lending *LendingCallerSession) SLastUpdateBlock() (*big.Int, error) {
	return _Lending.Contract.SLastUpdateBlock(&_Lending.CallOpts)
}

// SLenderTokenCount is a free data retrieval call binding the contract method 0xa1245b5d.
//
// Solidity: function sLenderTokenCount(address ) view returns(uint256)
func (_Lending *LendingCaller) SLenderTokenCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "sLenderTokenCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SLenderTokenCount is a free data retrieval call binding the contract method 0xa1245b5d.
//
// Solidity: function sLenderTokenCount(address ) view returns(uint256)
func (_Lending *LendingSession) SLenderTokenCount(arg0 common.Address) (*big.Int, error) {
	return _Lending.Contract.SLenderTokenCount(&_Lending.CallOpts, arg0)
}

// SLenderTokenCount is a free data retrieval call binding the contract method 0xa1245b5d.
//
// Solidity: function sLenderTokenCount(address ) view returns(uint256)
func (_Lending *LendingCallerSession) SLenderTokenCount(arg0 common.Address) (*big.Int, error) {
	return _Lending.Contract.SLenderTokenCount(&_Lending.CallOpts, arg0)
}

// SLendingToken is a free data retrieval call binding the contract method 0x9c30eeff.
//
// Solidity: function sLendingToken() view returns(address)
func (_Lending *LendingCaller) SLendingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "sLendingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SLendingToken is a free data retrieval call binding the contract method 0x9c30eeff.
//
// Solidity: function sLendingToken() view returns(address)
func (_Lending *LendingSession) SLendingToken() (common.Address, error) {
	return _Lending.Contract.SLendingToken(&_Lending.CallOpts)
}

// SLendingToken is a free data retrieval call binding the contract method 0x9c30eeff.
//
// Solidity: function sLendingToken() view returns(address)
func (_Lending *LendingCallerSession) SLendingToken() (common.Address, error) {
	return _Lending.Contract.SLendingToken(&_Lending.CallOpts)
}

// SProtocolFee is a free data retrieval call binding the contract method 0xc33731c2.
//
// Solidity: function sProtocolFee() view returns(uint256)
func (_Lending *LendingCaller) SProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "sProtocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SProtocolFee is a free data retrieval call binding the contract method 0xc33731c2.
//
// Solidity: function sProtocolFee() view returns(uint256)
func (_Lending *LendingSession) SProtocolFee() (*big.Int, error) {
	return _Lending.Contract.SProtocolFee(&_Lending.CallOpts)
}

// SProtocolFee is a free data retrieval call binding the contract method 0xc33731c2.
//
// Solidity: function sProtocolFee() view returns(uint256)
func (_Lending *LendingCallerSession) SProtocolFee() (*big.Int, error) {
	return _Lending.Contract.SProtocolFee(&_Lending.CallOpts)
}

// SRewardToken is a free data retrieval call binding the contract method 0x304dc809.
//
// Solidity: function sRewardToken() view returns(address)
func (_Lending *LendingCaller) SRewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "sRewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SRewardToken is a free data retrieval call binding the contract method 0x304dc809.
//
// Solidity: function sRewardToken() view returns(address)
func (_Lending *LendingSession) SRewardToken() (common.Address, error) {
	return _Lending.Contract.SRewardToken(&_Lending.CallOpts)
}

// SRewardToken is a free data retrieval call binding the contract method 0x304dc809.
//
// Solidity: function sRewardToken() view returns(address)
func (_Lending *LendingCallerSession) SRewardToken() (common.Address, error) {
	return _Lending.Contract.SRewardToken(&_Lending.CallOpts)
}

// STokenBorrows is a free data retrieval call binding the contract method 0xc23f4e24.
//
// Solidity: function sTokenBorrows(uint256 ) view returns(address borrower, uint256 startBlock, uint256 endBlock, uint256 feePerBlock)
func (_Lending *LendingCaller) STokenBorrows(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Borrower    common.Address
	StartBlock  *big.Int
	EndBlock    *big.Int
	FeePerBlock *big.Int
}, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "sTokenBorrows", arg0)

	outstruct := new(struct {
		Borrower    common.Address
		StartBlock  *big.Int
		EndBlock    *big.Int
		FeePerBlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Borrower = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StartBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FeePerBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// STokenBorrows is a free data retrieval call binding the contract method 0xc23f4e24.
//
// Solidity: function sTokenBorrows(uint256 ) view returns(address borrower, uint256 startBlock, uint256 endBlock, uint256 feePerBlock)
func (_Lending *LendingSession) STokenBorrows(arg0 *big.Int) (struct {
	Borrower    common.Address
	StartBlock  *big.Int
	EndBlock    *big.Int
	FeePerBlock *big.Int
}, error) {
	return _Lending.Contract.STokenBorrows(&_Lending.CallOpts, arg0)
}

// STokenBorrows is a free data retrieval call binding the contract method 0xc23f4e24.
//
// Solidity: function sTokenBorrows(uint256 ) view returns(address borrower, uint256 startBlock, uint256 endBlock, uint256 feePerBlock)
func (_Lending *LendingCallerSession) STokenBorrows(arg0 *big.Int) (struct {
	Borrower    common.Address
	StartBlock  *big.Int
	EndBlock    *big.Int
	FeePerBlock *big.Int
}, error) {
	return _Lending.Contract.STokenBorrows(&_Lending.CallOpts, arg0)
}

// STokenLender is a free data retrieval call binding the contract method 0x854d8fd6.
//
// Solidity: function sTokenLender(uint256 ) view returns(address)
func (_Lending *LendingCaller) STokenLender(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "sTokenLender", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STokenLender is a free data retrieval call binding the contract method 0x854d8fd6.
//
// Solidity: function sTokenLender(uint256 ) view returns(address)
func (_Lending *LendingSession) STokenLender(arg0 *big.Int) (common.Address, error) {
	return _Lending.Contract.STokenLender(&_Lending.CallOpts, arg0)
}

// STokenLender is a free data retrieval call binding the contract method 0x854d8fd6.
//
// Solidity: function sTokenLender(uint256 ) view returns(address)
func (_Lending *LendingCallerSession) STokenLender(arg0 *big.Int) (common.Address, error) {
	return _Lending.Contract.STokenLender(&_Lending.CallOpts, arg0)
}

// STotalBorrowedTokens is a free data retrieval call binding the contract method 0x1c905e8d.
//
// Solidity: function sTotalBorrowedTokens() view returns(uint256)
func (_Lending *LendingCaller) STotalBorrowedTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "sTotalBorrowedTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STotalBorrowedTokens is a free data retrieval call binding the contract method 0x1c905e8d.
//
// Solidity: function sTotalBorrowedTokens() view returns(uint256)
func (_Lending *LendingSession) STotalBorrowedTokens() (*big.Int, error) {
	return _Lending.Contract.STotalBorrowedTokens(&_Lending.CallOpts)
}

// STotalBorrowedTokens is a free data retrieval call binding the contract method 0x1c905e8d.
//
// Solidity: function sTotalBorrowedTokens() view returns(uint256)
func (_Lending *LendingCallerSession) STotalBorrowedTokens() (*big.Int, error) {
	return _Lending.Contract.STotalBorrowedTokens(&_Lending.CallOpts)
}

// STotalFeePerBlock is a free data retrieval call binding the contract method 0x1f41c5f8.
//
// Solidity: function sTotalFeePerBlock() view returns(uint256)
func (_Lending *LendingCaller) STotalFeePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "sTotalFeePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STotalFeePerBlock is a free data retrieval call binding the contract method 0x1f41c5f8.
//
// Solidity: function sTotalFeePerBlock() view returns(uint256)
func (_Lending *LendingSession) STotalFeePerBlock() (*big.Int, error) {
	return _Lending.Contract.STotalFeePerBlock(&_Lending.CallOpts)
}

// STotalFeePerBlock is a free data retrieval call binding the contract method 0x1f41c5f8.
//
// Solidity: function sTotalFeePerBlock() view returns(uint256)
func (_Lending *LendingCallerSession) STotalFeePerBlock() (*big.Int, error) {
	return _Lending.Contract.STotalFeePerBlock(&_Lending.CallOpts)
}

// STotalLentTokens is a free data retrieval call binding the contract method 0x07196718.
//
// Solidity: function sTotalLentTokens() view returns(uint256)
func (_Lending *LendingCaller) STotalLentTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "sTotalLentTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STotalLentTokens is a free data retrieval call binding the contract method 0x07196718.
//
// Solidity: function sTotalLentTokens() view returns(uint256)
func (_Lending *LendingSession) STotalLentTokens() (*big.Int, error) {
	return _Lending.Contract.STotalLentTokens(&_Lending.CallOpts)
}

// STotalLentTokens is a free data retrieval call binding the contract method 0x07196718.
//
// Solidity: function sTotalLentTokens() view returns(uint256)
func (_Lending *LendingCallerSession) STotalLentTokens() (*big.Int, error) {
	return _Lending.Contract.STotalLentTokens(&_Lending.CallOpts)
}

// STotalReward is a free data retrieval call binding the contract method 0xebefc994.
//
// Solidity: function sTotalReward() view returns(uint256)
func (_Lending *LendingCaller) STotalReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "sTotalReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STotalReward is a free data retrieval call binding the contract method 0xebefc994.
//
// Solidity: function sTotalReward() view returns(uint256)
func (_Lending *LendingSession) STotalReward() (*big.Int, error) {
	return _Lending.Contract.STotalReward(&_Lending.CallOpts)
}

// STotalReward is a free data retrieval call binding the contract method 0xebefc994.
//
// Solidity: function sTotalReward() view returns(uint256)
func (_Lending *LendingCallerSession) STotalReward() (*big.Int, error) {
	return _Lending.Contract.STotalReward(&_Lending.CallOpts)
}

// STotalRewardClaimed is a free data retrieval call binding the contract method 0x90a51480.
//
// Solidity: function sTotalRewardClaimed() view returns(uint256)
func (_Lending *LendingCaller) STotalRewardClaimed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "sTotalRewardClaimed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STotalRewardClaimed is a free data retrieval call binding the contract method 0x90a51480.
//
// Solidity: function sTotalRewardClaimed() view returns(uint256)
func (_Lending *LendingSession) STotalRewardClaimed() (*big.Int, error) {
	return _Lending.Contract.STotalRewardClaimed(&_Lending.CallOpts)
}

// STotalRewardClaimed is a free data retrieval call binding the contract method 0x90a51480.
//
// Solidity: function sTotalRewardClaimed() view returns(uint256)
func (_Lending *LendingCallerSession) STotalRewardClaimed() (*big.Int, error) {
	return _Lending.Contract.STotalRewardClaimed(&_Lending.CallOpts)
}

// SUserLends is a free data retrieval call binding the contract method 0xf2e826a5.
//
// Solidity: function sUserLends(address , uint256 ) view returns(uint256 tokenId, uint256 startBlock, bool isLent, uint256 rewardDebt)
func (_Lending *LendingCaller) SUserLends(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	TokenId    *big.Int
	StartBlock *big.Int
	IsLent     bool
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "sUserLends", arg0, arg1)

	outstruct := new(struct {
		TokenId    *big.Int
		StartBlock *big.Int
		IsLent     bool
		RewardDebt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsLent = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.RewardDebt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SUserLends is a free data retrieval call binding the contract method 0xf2e826a5.
//
// Solidity: function sUserLends(address , uint256 ) view returns(uint256 tokenId, uint256 startBlock, bool isLent, uint256 rewardDebt)
func (_Lending *LendingSession) SUserLends(arg0 common.Address, arg1 *big.Int) (struct {
	TokenId    *big.Int
	StartBlock *big.Int
	IsLent     bool
	RewardDebt *big.Int
}, error) {
	return _Lending.Contract.SUserLends(&_Lending.CallOpts, arg0, arg1)
}

// SUserLends is a free data retrieval call binding the contract method 0xf2e826a5.
//
// Solidity: function sUserLends(address , uint256 ) view returns(uint256 tokenId, uint256 startBlock, bool isLent, uint256 rewardDebt)
func (_Lending *LendingCallerSession) SUserLends(arg0 common.Address, arg1 *big.Int) (struct {
	TokenId    *big.Int
	StartBlock *big.Int
	IsLent     bool
	RewardDebt *big.Int
}, error) {
	return _Lending.Contract.SUserLends(&_Lending.CallOpts, arg0, arg1)
}

// SUserReserveBalance is a free data retrieval call binding the contract method 0x40328f6e.
//
// Solidity: function sUserReserveBalance(address ) view returns(uint256)
func (_Lending *LendingCaller) SUserReserveBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "sUserReserveBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SUserReserveBalance is a free data retrieval call binding the contract method 0x40328f6e.
//
// Solidity: function sUserReserveBalance(address ) view returns(uint256)
func (_Lending *LendingSession) SUserReserveBalance(arg0 common.Address) (*big.Int, error) {
	return _Lending.Contract.SUserReserveBalance(&_Lending.CallOpts, arg0)
}

// SUserReserveBalance is a free data retrieval call binding the contract method 0x40328f6e.
//
// Solidity: function sUserReserveBalance(address ) view returns(uint256)
func (_Lending *LendingCallerSession) SUserReserveBalance(arg0 common.Address) (*big.Int, error) {
	return _Lending.Contract.SUserReserveBalance(&_Lending.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Lending *LendingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Lending *LendingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Lending.Contract.SupportsInterface(&_Lending.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Lending *LendingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Lending.Contract.SupportsInterface(&_Lending.CallOpts, interfaceId)
}

// TokenOfBorrowerByIndex is a free data retrieval call binding the contract method 0x8c222a97.
//
// Solidity: function tokenOfBorrowerByIndex(address borrower, uint256 index) view returns(uint256)
func (_Lending *LendingCaller) TokenOfBorrowerByIndex(opts *bind.CallOpts, borrower common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "tokenOfBorrowerByIndex", borrower, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfBorrowerByIndex is a free data retrieval call binding the contract method 0x8c222a97.
//
// Solidity: function tokenOfBorrowerByIndex(address borrower, uint256 index) view returns(uint256)
func (_Lending *LendingSession) TokenOfBorrowerByIndex(borrower common.Address, index *big.Int) (*big.Int, error) {
	return _Lending.Contract.TokenOfBorrowerByIndex(&_Lending.CallOpts, borrower, index)
}

// TokenOfBorrowerByIndex is a free data retrieval call binding the contract method 0x8c222a97.
//
// Solidity: function tokenOfBorrowerByIndex(address borrower, uint256 index) view returns(uint256)
func (_Lending *LendingCallerSession) TokenOfBorrowerByIndex(borrower common.Address, index *big.Int) (*big.Int, error) {
	return _Lending.Contract.TokenOfBorrowerByIndex(&_Lending.CallOpts, borrower, index)
}

// TokenOfLenderByIndex is a free data retrieval call binding the contract method 0xc672f058.
//
// Solidity: function tokenOfLenderByIndex(address lender, uint256 index) view returns(uint256)
func (_Lending *LendingCaller) TokenOfLenderByIndex(opts *bind.CallOpts, lender common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "tokenOfLenderByIndex", lender, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfLenderByIndex is a free data retrieval call binding the contract method 0xc672f058.
//
// Solidity: function tokenOfLenderByIndex(address lender, uint256 index) view returns(uint256)
func (_Lending *LendingSession) TokenOfLenderByIndex(lender common.Address, index *big.Int) (*big.Int, error) {
	return _Lending.Contract.TokenOfLenderByIndex(&_Lending.CallOpts, lender, index)
}

// TokenOfLenderByIndex is a free data retrieval call binding the contract method 0xc672f058.
//
// Solidity: function tokenOfLenderByIndex(address lender, uint256 index) view returns(uint256)
func (_Lending *LendingCallerSession) TokenOfLenderByIndex(lender common.Address, index *big.Int) (*big.Int, error) {
	return _Lending.Contract.TokenOfLenderByIndex(&_Lending.CallOpts, lender, index)
}

// ValidateState is a free data retrieval call binding the contract method 0x929bb56b.
//
// Solidity: function validateState() view returns(bool)
func (_Lending *LendingCaller) ValidateState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lending.contract.Call(opts, &out, "validateState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateState is a free data retrieval call binding the contract method 0x929bb56b.
//
// Solidity: function validateState() view returns(bool)
func (_Lending *LendingSession) ValidateState() (bool, error) {
	return _Lending.Contract.ValidateState(&_Lending.CallOpts)
}

// ValidateState is a free data retrieval call binding the contract method 0x929bb56b.
//
// Solidity: function validateState() view returns(bool)
func (_Lending *LendingCallerSession) ValidateState() (bool, error) {
	return _Lending.Contract.ValidateState(&_Lending.CallOpts)
}

// AddAvailableDurations is a paid mutator transaction binding the contract method 0x86ecae3a.
//
// Solidity: function addAvailableDurations(uint256 _availableDurations) returns()
func (_Lending *LendingTransactor) AddAvailableDurations(opts *bind.TransactOpts, _availableDurations *big.Int) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "addAvailableDurations", _availableDurations)
}

// AddAvailableDurations is a paid mutator transaction binding the contract method 0x86ecae3a.
//
// Solidity: function addAvailableDurations(uint256 _availableDurations) returns()
func (_Lending *LendingSession) AddAvailableDurations(_availableDurations *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.AddAvailableDurations(&_Lending.TransactOpts, _availableDurations)
}

// AddAvailableDurations is a paid mutator transaction binding the contract method 0x86ecae3a.
//
// Solidity: function addAvailableDurations(uint256 _availableDurations) returns()
func (_Lending *LendingTransactorSession) AddAvailableDurations(_availableDurations *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.AddAvailableDurations(&_Lending.TransactOpts, _availableDurations)
}

// Borrow is a paid mutator transaction binding the contract method 0x0ecbcdab.
//
// Solidity: function borrow(uint256 amount, uint256 durationInDays) returns()
func (_Lending *LendingTransactor) Borrow(opts *bind.TransactOpts, amount *big.Int, durationInDays *big.Int) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "borrow", amount, durationInDays)
}

// Borrow is a paid mutator transaction binding the contract method 0x0ecbcdab.
//
// Solidity: function borrow(uint256 amount, uint256 durationInDays) returns()
func (_Lending *LendingSession) Borrow(amount *big.Int, durationInDays *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.Borrow(&_Lending.TransactOpts, amount, durationInDays)
}

// Borrow is a paid mutator transaction binding the contract method 0x0ecbcdab.
//
// Solidity: function borrow(uint256 amount, uint256 durationInDays) returns()
func (_Lending *LendingTransactorSession) Borrow(amount *big.Int, durationInDays *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.Borrow(&_Lending.TransactOpts, amount, durationInDays)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_Lending *LendingTransactor) ClaimRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "claimRewards")
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_Lending *LendingSession) ClaimRewards() (*types.Transaction, error) {
	return _Lending.Contract.ClaimRewards(&_Lending.TransactOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_Lending *LendingTransactorSession) ClaimRewards() (*types.Transaction, error) {
	return _Lending.Contract.ClaimRewards(&_Lending.TransactOpts)
}

// EmergencyUnstake is a paid mutator transaction binding the contract method 0x012ce501.
//
// Solidity: function emergencyUnstake(uint256 tokenId) returns()
func (_Lending *LendingTransactor) EmergencyUnstake(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "emergencyUnstake", tokenId)
}

// EmergencyUnstake is a paid mutator transaction binding the contract method 0x012ce501.
//
// Solidity: function emergencyUnstake(uint256 tokenId) returns()
func (_Lending *LendingSession) EmergencyUnstake(tokenId *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.EmergencyUnstake(&_Lending.TransactOpts, tokenId)
}

// EmergencyUnstake is a paid mutator transaction binding the contract method 0x012ce501.
//
// Solidity: function emergencyUnstake(uint256 tokenId) returns()
func (_Lending *LendingTransactorSession) EmergencyUnstake(tokenId *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.EmergencyUnstake(&_Lending.TransactOpts, tokenId)
}

// EmergencyUnstakeBatch is a paid mutator transaction binding the contract method 0xb1224ab3.
//
// Solidity: function emergencyUnstakeBatch(uint256[] tokenIds) returns()
func (_Lending *LendingTransactor) EmergencyUnstakeBatch(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "emergencyUnstakeBatch", tokenIds)
}

// EmergencyUnstakeBatch is a paid mutator transaction binding the contract method 0xb1224ab3.
//
// Solidity: function emergencyUnstakeBatch(uint256[] tokenIds) returns()
func (_Lending *LendingSession) EmergencyUnstakeBatch(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Lending.Contract.EmergencyUnstakeBatch(&_Lending.TransactOpts, tokenIds)
}

// EmergencyUnstakeBatch is a paid mutator transaction binding the contract method 0xb1224ab3.
//
// Solidity: function emergencyUnstakeBatch(uint256[] tokenIds) returns()
func (_Lending *LendingTransactorSession) EmergencyUnstakeBatch(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Lending.Contract.EmergencyUnstakeBatch(&_Lending.TransactOpts, tokenIds)
}

// ForceUnstake is a paid mutator transaction binding the contract method 0x91460149.
//
// Solidity: function forceUnstake(uint256 tokenId) returns()
func (_Lending *LendingTransactor) ForceUnstake(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "forceUnstake", tokenId)
}

// ForceUnstake is a paid mutator transaction binding the contract method 0x91460149.
//
// Solidity: function forceUnstake(uint256 tokenId) returns()
func (_Lending *LendingSession) ForceUnstake(tokenId *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.ForceUnstake(&_Lending.TransactOpts, tokenId)
}

// ForceUnstake is a paid mutator transaction binding the contract method 0x91460149.
//
// Solidity: function forceUnstake(uint256 tokenId) returns()
func (_Lending *LendingTransactorSession) ForceUnstake(tokenId *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.ForceUnstake(&_Lending.TransactOpts, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Lending *LendingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Lending *LendingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Lending.Contract.GrantRole(&_Lending.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Lending *LendingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Lending.Contract.GrantRole(&_Lending.TransactOpts, role, account)
}

// Lend is a paid mutator transaction binding the contract method 0xa6aa57ce.
//
// Solidity: function lend(uint256 tokenId) returns()
func (_Lending *LendingTransactor) Lend(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "lend", tokenId)
}

// Lend is a paid mutator transaction binding the contract method 0xa6aa57ce.
//
// Solidity: function lend(uint256 tokenId) returns()
func (_Lending *LendingSession) Lend(tokenId *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.Lend(&_Lending.TransactOpts, tokenId)
}

// Lend is a paid mutator transaction binding the contract method 0xa6aa57ce.
//
// Solidity: function lend(uint256 tokenId) returns()
func (_Lending *LendingTransactorSession) Lend(tokenId *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.Lend(&_Lending.TransactOpts, tokenId)
}

// LendBatch is a paid mutator transaction binding the contract method 0x0e7ed818.
//
// Solidity: function lendBatch(uint256[] tokenIds) returns()
func (_Lending *LendingTransactor) LendBatch(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "lendBatch", tokenIds)
}

// LendBatch is a paid mutator transaction binding the contract method 0x0e7ed818.
//
// Solidity: function lendBatch(uint256[] tokenIds) returns()
func (_Lending *LendingSession) LendBatch(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Lending.Contract.LendBatch(&_Lending.TransactOpts, tokenIds)
}

// LendBatch is a paid mutator transaction binding the contract method 0x0e7ed818.
//
// Solidity: function lendBatch(uint256[] tokenIds) returns()
func (_Lending *LendingTransactorSession) LendBatch(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Lending.Contract.LendBatch(&_Lending.TransactOpts, tokenIds)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Lending *LendingTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Lending *LendingSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Lending.Contract.OnERC721Received(&_Lending.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Lending *LendingTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Lending.Contract.OnERC721Received(&_Lending.TransactOpts, arg0, arg1, arg2, arg3)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_Lending *LendingTransactor) RecoverERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "recoverERC20", token, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_Lending *LendingSession) RecoverERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.RecoverERC20(&_Lending.TransactOpts, token, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_Lending *LendingTransactorSession) RecoverERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.RecoverERC20(&_Lending.TransactOpts, token, amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Lending *LendingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Lending *LendingSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Lending.Contract.RenounceRole(&_Lending.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Lending *LendingTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Lending.Contract.RenounceRole(&_Lending.TransactOpts, role, callerConfirmation)
}

// ReturnOverdueToken is a paid mutator transaction binding the contract method 0xa19d10f8.
//
// Solidity: function returnOverdueToken(uint256 tokenId) returns()
func (_Lending *LendingTransactor) ReturnOverdueToken(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "returnOverdueToken", tokenId)
}

// ReturnOverdueToken is a paid mutator transaction binding the contract method 0xa19d10f8.
//
// Solidity: function returnOverdueToken(uint256 tokenId) returns()
func (_Lending *LendingSession) ReturnOverdueToken(tokenId *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.ReturnOverdueToken(&_Lending.TransactOpts, tokenId)
}

// ReturnOverdueToken is a paid mutator transaction binding the contract method 0xa19d10f8.
//
// Solidity: function returnOverdueToken(uint256 tokenId) returns()
func (_Lending *LendingTransactorSession) ReturnOverdueToken(tokenId *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.ReturnOverdueToken(&_Lending.TransactOpts, tokenId)
}

// ReturnOverdueTokenBatch is a paid mutator transaction binding the contract method 0xef62cb0d.
//
// Solidity: function returnOverdueTokenBatch(uint256[] tokenIds) returns()
func (_Lending *LendingTransactor) ReturnOverdueTokenBatch(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "returnOverdueTokenBatch", tokenIds)
}

// ReturnOverdueTokenBatch is a paid mutator transaction binding the contract method 0xef62cb0d.
//
// Solidity: function returnOverdueTokenBatch(uint256[] tokenIds) returns()
func (_Lending *LendingSession) ReturnOverdueTokenBatch(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Lending.Contract.ReturnOverdueTokenBatch(&_Lending.TransactOpts, tokenIds)
}

// ReturnOverdueTokenBatch is a paid mutator transaction binding the contract method 0xef62cb0d.
//
// Solidity: function returnOverdueTokenBatch(uint256[] tokenIds) returns()
func (_Lending *LendingTransactorSession) ReturnOverdueTokenBatch(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Lending.Contract.ReturnOverdueTokenBatch(&_Lending.TransactOpts, tokenIds)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Lending *LendingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Lending *LendingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Lending.Contract.RevokeRole(&_Lending.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Lending *LendingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Lending.Contract.RevokeRole(&_Lending.TransactOpts, role, account)
}

// SetBaseFeePerBlock is a paid mutator transaction binding the contract method 0x18506a48.
//
// Solidity: function setBaseFeePerBlock(uint256 _baseFeePerBlock) returns()
func (_Lending *LendingTransactor) SetBaseFeePerBlock(opts *bind.TransactOpts, _baseFeePerBlock *big.Int) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "setBaseFeePerBlock", _baseFeePerBlock)
}

// SetBaseFeePerBlock is a paid mutator transaction binding the contract method 0x18506a48.
//
// Solidity: function setBaseFeePerBlock(uint256 _baseFeePerBlock) returns()
func (_Lending *LendingSession) SetBaseFeePerBlock(_baseFeePerBlock *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.SetBaseFeePerBlock(&_Lending.TransactOpts, _baseFeePerBlock)
}

// SetBaseFeePerBlock is a paid mutator transaction binding the contract method 0x18506a48.
//
// Solidity: function setBaseFeePerBlock(uint256 _baseFeePerBlock) returns()
func (_Lending *LendingTransactorSession) SetBaseFeePerBlock(_baseFeePerBlock *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.SetBaseFeePerBlock(&_Lending.TransactOpts, _baseFeePerBlock)
}

// SetEmergencyUnstake is a paid mutator transaction binding the contract method 0x86e39d72.
//
// Solidity: function setEmergencyUnstake(bool _enabled) returns()
func (_Lending *LendingTransactor) SetEmergencyUnstake(opts *bind.TransactOpts, _enabled bool) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "setEmergencyUnstake", _enabled)
}

// SetEmergencyUnstake is a paid mutator transaction binding the contract method 0x86e39d72.
//
// Solidity: function setEmergencyUnstake(bool _enabled) returns()
func (_Lending *LendingSession) SetEmergencyUnstake(_enabled bool) (*types.Transaction, error) {
	return _Lending.Contract.SetEmergencyUnstake(&_Lending.TransactOpts, _enabled)
}

// SetEmergencyUnstake is a paid mutator transaction binding the contract method 0x86e39d72.
//
// Solidity: function setEmergencyUnstake(bool _enabled) returns()
func (_Lending *LendingTransactorSession) SetEmergencyUnstake(_enabled bool) (*types.Transaction, error) {
	return _Lending.Contract.SetEmergencyUnstake(&_Lending.TransactOpts, _enabled)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _paused) returns()
func (_Lending *LendingTransactor) SetPaused(opts *bind.TransactOpts, _paused bool) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "setPaused", _paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _paused) returns()
func (_Lending *LendingSession) SetPaused(_paused bool) (*types.Transaction, error) {
	return _Lending.Contract.SetPaused(&_Lending.TransactOpts, _paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _paused) returns()
func (_Lending *LendingTransactorSession) SetPaused(_paused bool) (*types.Transaction, error) {
	return _Lending.Contract.SetPaused(&_Lending.TransactOpts, _paused)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 _protocolFee) returns()
func (_Lending *LendingTransactor) SetProtocolFee(opts *bind.TransactOpts, _protocolFee *big.Int) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "setProtocolFee", _protocolFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 _protocolFee) returns()
func (_Lending *LendingSession) SetProtocolFee(_protocolFee *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.SetProtocolFee(&_Lending.TransactOpts, _protocolFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 _protocolFee) returns()
func (_Lending *LendingTransactorSession) SetProtocolFee(_protocolFee *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.SetProtocolFee(&_Lending.TransactOpts, _protocolFee)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 tokenId) returns()
func (_Lending *LendingTransactor) Unstake(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "unstake", tokenId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 tokenId) returns()
func (_Lending *LendingSession) Unstake(tokenId *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.Unstake(&_Lending.TransactOpts, tokenId)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 tokenId) returns()
func (_Lending *LendingTransactorSession) Unstake(tokenId *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.Unstake(&_Lending.TransactOpts, tokenId)
}

// UnstakeBatch is a paid mutator transaction binding the contract method 0xe64a21f3.
//
// Solidity: function unstakeBatch(uint256[] tokenIds) returns()
func (_Lending *LendingTransactor) UnstakeBatch(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "unstakeBatch", tokenIds)
}

// UnstakeBatch is a paid mutator transaction binding the contract method 0xe64a21f3.
//
// Solidity: function unstakeBatch(uint256[] tokenIds) returns()
func (_Lending *LendingSession) UnstakeBatch(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Lending.Contract.UnstakeBatch(&_Lending.TransactOpts, tokenIds)
}

// UnstakeBatch is a paid mutator transaction binding the contract method 0xe64a21f3.
//
// Solidity: function unstakeBatch(uint256[] tokenIds) returns()
func (_Lending *LendingTransactorSession) UnstakeBatch(tokenIds []*big.Int) (*types.Transaction, error) {
	return _Lending.Contract.UnstakeBatch(&_Lending.TransactOpts, tokenIds)
}

// UpdateBlockTimeConfig is a paid mutator transaction binding the contract method 0x965b6346.
//
// Solidity: function updateBlockTimeConfig(uint256 _blocksPerDay) returns()
func (_Lending *LendingTransactor) UpdateBlockTimeConfig(opts *bind.TransactOpts, _blocksPerDay *big.Int) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "updateBlockTimeConfig", _blocksPerDay)
}

// UpdateBlockTimeConfig is a paid mutator transaction binding the contract method 0x965b6346.
//
// Solidity: function updateBlockTimeConfig(uint256 _blocksPerDay) returns()
func (_Lending *LendingSession) UpdateBlockTimeConfig(_blocksPerDay *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.UpdateBlockTimeConfig(&_Lending.TransactOpts, _blocksPerDay)
}

// UpdateBlockTimeConfig is a paid mutator transaction binding the contract method 0x965b6346.
//
// Solidity: function updateBlockTimeConfig(uint256 _blocksPerDay) returns()
func (_Lending *LendingTransactorSession) UpdateBlockTimeConfig(_blocksPerDay *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.UpdateBlockTimeConfig(&_Lending.TransactOpts, _blocksPerDay)
}

// WithdrawAllReward is a paid mutator transaction binding the contract method 0x6b37fbaa.
//
// Solidity: function withdrawAllReward() returns()
func (_Lending *LendingTransactor) WithdrawAllReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "withdrawAllReward")
}

// WithdrawAllReward is a paid mutator transaction binding the contract method 0x6b37fbaa.
//
// Solidity: function withdrawAllReward() returns()
func (_Lending *LendingSession) WithdrawAllReward() (*types.Transaction, error) {
	return _Lending.Contract.WithdrawAllReward(&_Lending.TransactOpts)
}

// WithdrawAllReward is a paid mutator transaction binding the contract method 0x6b37fbaa.
//
// Solidity: function withdrawAllReward() returns()
func (_Lending *LendingTransactorSession) WithdrawAllReward() (*types.Transaction, error) {
	return _Lending.Contract.WithdrawAllReward(&_Lending.TransactOpts)
}

// WithdrawReward is a paid mutator transaction binding the contract method 0x523a3f08.
//
// Solidity: function withdrawReward(uint256 amount) returns()
func (_Lending *LendingTransactor) WithdrawReward(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "withdrawReward", amount)
}

// WithdrawReward is a paid mutator transaction binding the contract method 0x523a3f08.
//
// Solidity: function withdrawReward(uint256 amount) returns()
func (_Lending *LendingSession) WithdrawReward(amount *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.WithdrawReward(&_Lending.TransactOpts, amount)
}

// WithdrawReward is a paid mutator transaction binding the contract method 0x523a3f08.
//
// Solidity: function withdrawReward(uint256 amount) returns()
func (_Lending *LendingTransactorSession) WithdrawReward(amount *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.WithdrawReward(&_Lending.TransactOpts, amount)
}

// LendingBlockTimeConfigUpdatedIterator is returned from FilterBlockTimeConfigUpdated and is used to iterate over the raw logs and unpacked data for BlockTimeConfigUpdated events raised by the Lending contract.
type LendingBlockTimeConfigUpdatedIterator struct {
	Event *LendingBlockTimeConfigUpdated // Event containing the contract specifics and raw log

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
func (it *LendingBlockTimeConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingBlockTimeConfigUpdated)
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
		it.Event = new(LendingBlockTimeConfigUpdated)
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
func (it *LendingBlockTimeConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingBlockTimeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingBlockTimeConfigUpdated represents a BlockTimeConfigUpdated event raised by the Lending contract.
type LendingBlockTimeConfigUpdated struct {
	NewBlocksPerDay *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBlockTimeConfigUpdated is a free log retrieval operation binding the contract event 0xcdb1cb3fa7b28ae929986bba7f827f6aa59fbbb4033f4fcb3ff8281fb28ad6aa.
//
// Solidity: event BlockTimeConfigUpdated(uint256 newBlocksPerDay)
func (_Lending *LendingFilterer) FilterBlockTimeConfigUpdated(opts *bind.FilterOpts) (*LendingBlockTimeConfigUpdatedIterator, error) {

	logs, sub, err := _Lending.contract.FilterLogs(opts, "BlockTimeConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &LendingBlockTimeConfigUpdatedIterator{contract: _Lending.contract, event: "BlockTimeConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchBlockTimeConfigUpdated is a free log subscription operation binding the contract event 0xcdb1cb3fa7b28ae929986bba7f827f6aa59fbbb4033f4fcb3ff8281fb28ad6aa.
//
// Solidity: event BlockTimeConfigUpdated(uint256 newBlocksPerDay)
func (_Lending *LendingFilterer) WatchBlockTimeConfigUpdated(opts *bind.WatchOpts, sink chan<- *LendingBlockTimeConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _Lending.contract.WatchLogs(opts, "BlockTimeConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingBlockTimeConfigUpdated)
				if err := _Lending.contract.UnpackLog(event, "BlockTimeConfigUpdated", log); err != nil {
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

// ParseBlockTimeConfigUpdated is a log parse operation binding the contract event 0xcdb1cb3fa7b28ae929986bba7f827f6aa59fbbb4033f4fcb3ff8281fb28ad6aa.
//
// Solidity: event BlockTimeConfigUpdated(uint256 newBlocksPerDay)
func (_Lending *LendingFilterer) ParseBlockTimeConfigUpdated(log types.Log) (*LendingBlockTimeConfigUpdated, error) {
	event := new(LendingBlockTimeConfigUpdated)
	if err := _Lending.contract.UnpackLog(event, "BlockTimeConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingEmergencyModeEnabledIterator is returned from FilterEmergencyModeEnabled and is used to iterate over the raw logs and unpacked data for EmergencyModeEnabled events raised by the Lending contract.
type LendingEmergencyModeEnabledIterator struct {
	Event *LendingEmergencyModeEnabled // Event containing the contract specifics and raw log

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
func (it *LendingEmergencyModeEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingEmergencyModeEnabled)
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
		it.Event = new(LendingEmergencyModeEnabled)
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
func (it *LendingEmergencyModeEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingEmergencyModeEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingEmergencyModeEnabled represents a EmergencyModeEnabled event raised by the Lending contract.
type LendingEmergencyModeEnabled struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEmergencyModeEnabled is a free log retrieval operation binding the contract event 0xb5c4f1a6d88ab40b649568d5bc704f15499f2c01d2e984cca5fd4fd3d8507244.
//
// Solidity: event EmergencyModeEnabled(bool enabled)
func (_Lending *LendingFilterer) FilterEmergencyModeEnabled(opts *bind.FilterOpts) (*LendingEmergencyModeEnabledIterator, error) {

	logs, sub, err := _Lending.contract.FilterLogs(opts, "EmergencyModeEnabled")
	if err != nil {
		return nil, err
	}
	return &LendingEmergencyModeEnabledIterator{contract: _Lending.contract, event: "EmergencyModeEnabled", logs: logs, sub: sub}, nil
}

// WatchEmergencyModeEnabled is a free log subscription operation binding the contract event 0xb5c4f1a6d88ab40b649568d5bc704f15499f2c01d2e984cca5fd4fd3d8507244.
//
// Solidity: event EmergencyModeEnabled(bool enabled)
func (_Lending *LendingFilterer) WatchEmergencyModeEnabled(opts *bind.WatchOpts, sink chan<- *LendingEmergencyModeEnabled) (event.Subscription, error) {

	logs, sub, err := _Lending.contract.WatchLogs(opts, "EmergencyModeEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingEmergencyModeEnabled)
				if err := _Lending.contract.UnpackLog(event, "EmergencyModeEnabled", log); err != nil {
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

// ParseEmergencyModeEnabled is a log parse operation binding the contract event 0xb5c4f1a6d88ab40b649568d5bc704f15499f2c01d2e984cca5fd4fd3d8507244.
//
// Solidity: event EmergencyModeEnabled(bool enabled)
func (_Lending *LendingFilterer) ParseEmergencyModeEnabled(log types.Log) (*LendingEmergencyModeEnabled, error) {
	event := new(LendingEmergencyModeEnabled)
	if err := _Lending.contract.UnpackLog(event, "EmergencyModeEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingEmergencyUnstakeIterator is returned from FilterEmergencyUnstake and is used to iterate over the raw logs and unpacked data for EmergencyUnstake events raised by the Lending contract.
type LendingEmergencyUnstakeIterator struct {
	Event *LendingEmergencyUnstake // Event containing the contract specifics and raw log

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
func (it *LendingEmergencyUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingEmergencyUnstake)
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
		it.Event = new(LendingEmergencyUnstake)
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
func (it *LendingEmergencyUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingEmergencyUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingEmergencyUnstake represents a EmergencyUnstake event raised by the Lending contract.
type LendingEmergencyUnstake struct {
	User    common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEmergencyUnstake is a free log retrieval operation binding the contract event 0x571394674ec9d9e81517060110f8f894ce912af2b2febc091bee0cdea68adf00.
//
// Solidity: event EmergencyUnstake(address indexed user, uint256 tokenId)
func (_Lending *LendingFilterer) FilterEmergencyUnstake(opts *bind.FilterOpts, user []common.Address) (*LendingEmergencyUnstakeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lending.contract.FilterLogs(opts, "EmergencyUnstake", userRule)
	if err != nil {
		return nil, err
	}
	return &LendingEmergencyUnstakeIterator{contract: _Lending.contract, event: "EmergencyUnstake", logs: logs, sub: sub}, nil
}

// WatchEmergencyUnstake is a free log subscription operation binding the contract event 0x571394674ec9d9e81517060110f8f894ce912af2b2febc091bee0cdea68adf00.
//
// Solidity: event EmergencyUnstake(address indexed user, uint256 tokenId)
func (_Lending *LendingFilterer) WatchEmergencyUnstake(opts *bind.WatchOpts, sink chan<- *LendingEmergencyUnstake, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lending.contract.WatchLogs(opts, "EmergencyUnstake", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingEmergencyUnstake)
				if err := _Lending.contract.UnpackLog(event, "EmergencyUnstake", log); err != nil {
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

// ParseEmergencyUnstake is a log parse operation binding the contract event 0x571394674ec9d9e81517060110f8f894ce912af2b2febc091bee0cdea68adf00.
//
// Solidity: event EmergencyUnstake(address indexed user, uint256 tokenId)
func (_Lending *LendingFilterer) ParseEmergencyUnstake(log types.Log) (*LendingEmergencyUnstake, error) {
	event := new(LendingEmergencyUnstake)
	if err := _Lending.contract.UnpackLog(event, "EmergencyUnstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingForceUnstakeIterator is returned from FilterForceUnstake and is used to iterate over the raw logs and unpacked data for ForceUnstake events raised by the Lending contract.
type LendingForceUnstakeIterator struct {
	Event *LendingForceUnstake // Event containing the contract specifics and raw log

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
func (it *LendingForceUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingForceUnstake)
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
		it.Event = new(LendingForceUnstake)
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
func (it *LendingForceUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingForceUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingForceUnstake represents a ForceUnstake event raised by the Lending contract.
type LendingForceUnstake struct {
	User    common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterForceUnstake is a free log retrieval operation binding the contract event 0x5ff34d37e2d8080f35765155ee896c6103f108a2c402597f80265511ec8659e6.
//
// Solidity: event ForceUnstake(address indexed user, uint256 tokenId)
func (_Lending *LendingFilterer) FilterForceUnstake(opts *bind.FilterOpts, user []common.Address) (*LendingForceUnstakeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lending.contract.FilterLogs(opts, "ForceUnstake", userRule)
	if err != nil {
		return nil, err
	}
	return &LendingForceUnstakeIterator{contract: _Lending.contract, event: "ForceUnstake", logs: logs, sub: sub}, nil
}

// WatchForceUnstake is a free log subscription operation binding the contract event 0x5ff34d37e2d8080f35765155ee896c6103f108a2c402597f80265511ec8659e6.
//
// Solidity: event ForceUnstake(address indexed user, uint256 tokenId)
func (_Lending *LendingFilterer) WatchForceUnstake(opts *bind.WatchOpts, sink chan<- *LendingForceUnstake, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lending.contract.WatchLogs(opts, "ForceUnstake", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingForceUnstake)
				if err := _Lending.contract.UnpackLog(event, "ForceUnstake", log); err != nil {
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

// ParseForceUnstake is a log parse operation binding the contract event 0x5ff34d37e2d8080f35765155ee896c6103f108a2c402597f80265511ec8659e6.
//
// Solidity: event ForceUnstake(address indexed user, uint256 tokenId)
func (_Lending *LendingFilterer) ParseForceUnstake(log types.Log) (*LendingForceUnstake, error) {
	event := new(LendingForceUnstake)
	if err := _Lending.contract.UnpackLog(event, "ForceUnstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingRewardPerBlockUpdatedIterator is returned from FilterRewardPerBlockUpdated and is used to iterate over the raw logs and unpacked data for RewardPerBlockUpdated events raised by the Lending contract.
type LendingRewardPerBlockUpdatedIterator struct {
	Event *LendingRewardPerBlockUpdated // Event containing the contract specifics and raw log

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
func (it *LendingRewardPerBlockUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingRewardPerBlockUpdated)
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
		it.Event = new(LendingRewardPerBlockUpdated)
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
func (it *LendingRewardPerBlockUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingRewardPerBlockUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingRewardPerBlockUpdated represents a RewardPerBlockUpdated event raised by the Lending contract.
type LendingRewardPerBlockUpdated struct {
	NewRewardPerBlock *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRewardPerBlockUpdated is a free log retrieval operation binding the contract event 0x5ed0ffa545a9eae0edd36b74378d16454cf385281383c7632ad5b2ebf3ab2b92.
//
// Solidity: event RewardPerBlockUpdated(uint256 newRewardPerBlock)
func (_Lending *LendingFilterer) FilterRewardPerBlockUpdated(opts *bind.FilterOpts) (*LendingRewardPerBlockUpdatedIterator, error) {

	logs, sub, err := _Lending.contract.FilterLogs(opts, "RewardPerBlockUpdated")
	if err != nil {
		return nil, err
	}
	return &LendingRewardPerBlockUpdatedIterator{contract: _Lending.contract, event: "RewardPerBlockUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardPerBlockUpdated is a free log subscription operation binding the contract event 0x5ed0ffa545a9eae0edd36b74378d16454cf385281383c7632ad5b2ebf3ab2b92.
//
// Solidity: event RewardPerBlockUpdated(uint256 newRewardPerBlock)
func (_Lending *LendingFilterer) WatchRewardPerBlockUpdated(opts *bind.WatchOpts, sink chan<- *LendingRewardPerBlockUpdated) (event.Subscription, error) {

	logs, sub, err := _Lending.contract.WatchLogs(opts, "RewardPerBlockUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingRewardPerBlockUpdated)
				if err := _Lending.contract.UnpackLog(event, "RewardPerBlockUpdated", log); err != nil {
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

// ParseRewardPerBlockUpdated is a log parse operation binding the contract event 0x5ed0ffa545a9eae0edd36b74378d16454cf385281383c7632ad5b2ebf3ab2b92.
//
// Solidity: event RewardPerBlockUpdated(uint256 newRewardPerBlock)
func (_Lending *LendingFilterer) ParseRewardPerBlockUpdated(log types.Log) (*LendingRewardPerBlockUpdated, error) {
	event := new(LendingRewardPerBlockUpdated)
	if err := _Lending.contract.UnpackLog(event, "RewardPerBlockUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the Lending contract.
type LendingRewardsClaimedIterator struct {
	Event *LendingRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *LendingRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingRewardsClaimed)
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
		it.Event = new(LendingRewardsClaimed)
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
func (it *LendingRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingRewardsClaimed represents a RewardsClaimed event raised by the Lending contract.
type LendingRewardsClaimed struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0xfc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe.
//
// Solidity: event RewardsClaimed(address indexed user, uint256 amount)
func (_Lending *LendingFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, user []common.Address) (*LendingRewardsClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lending.contract.FilterLogs(opts, "RewardsClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &LendingRewardsClaimedIterator{contract: _Lending.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0xfc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe.
//
// Solidity: event RewardsClaimed(address indexed user, uint256 amount)
func (_Lending *LendingFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *LendingRewardsClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lending.contract.WatchLogs(opts, "RewardsClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingRewardsClaimed)
				if err := _Lending.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
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

// ParseRewardsClaimed is a log parse operation binding the contract event 0xfc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe.
//
// Solidity: event RewardsClaimed(address indexed user, uint256 amount)
func (_Lending *LendingFilterer) ParseRewardsClaimed(log types.Log) (*LendingRewardsClaimed, error) {
	event := new(LendingRewardsClaimed)
	if err := _Lending.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Lending contract.
type LendingRoleAdminChangedIterator struct {
	Event *LendingRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LendingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingRoleAdminChanged)
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
		it.Event = new(LendingRoleAdminChanged)
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
func (it *LendingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingRoleAdminChanged represents a RoleAdminChanged event raised by the Lending contract.
type LendingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Lending *LendingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LendingRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Lending.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LendingRoleAdminChangedIterator{contract: _Lending.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Lending *LendingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LendingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Lending.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingRoleAdminChanged)
				if err := _Lending.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Lending *LendingFilterer) ParseRoleAdminChanged(log types.Log) (*LendingRoleAdminChanged, error) {
	event := new(LendingRoleAdminChanged)
	if err := _Lending.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Lending contract.
type LendingRoleGrantedIterator struct {
	Event *LendingRoleGranted // Event containing the contract specifics and raw log

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
func (it *LendingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingRoleGranted)
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
		it.Event = new(LendingRoleGranted)
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
func (it *LendingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingRoleGranted represents a RoleGranted event raised by the Lending contract.
type LendingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Lending *LendingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LendingRoleGrantedIterator, error) {

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

	logs, sub, err := _Lending.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LendingRoleGrantedIterator{contract: _Lending.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Lending *LendingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LendingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Lending.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingRoleGranted)
				if err := _Lending.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Lending *LendingFilterer) ParseRoleGranted(log types.Log) (*LendingRoleGranted, error) {
	event := new(LendingRoleGranted)
	if err := _Lending.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Lending contract.
type LendingRoleRevokedIterator struct {
	Event *LendingRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LendingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingRoleRevoked)
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
		it.Event = new(LendingRoleRevoked)
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
func (it *LendingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingRoleRevoked represents a RoleRevoked event raised by the Lending contract.
type LendingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Lending *LendingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LendingRoleRevokedIterator, error) {

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

	logs, sub, err := _Lending.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LendingRoleRevokedIterator{contract: _Lending.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Lending *LendingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LendingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Lending.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingRoleRevoked)
				if err := _Lending.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Lending *LendingFilterer) ParseRoleRevoked(log types.Log) (*LendingRoleRevoked, error) {
	event := new(LendingRoleRevoked)
	if err := _Lending.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingTokenBorrowedIterator is returned from FilterTokenBorrowed and is used to iterate over the raw logs and unpacked data for TokenBorrowed events raised by the Lending contract.
type LendingTokenBorrowedIterator struct {
	Event *LendingTokenBorrowed // Event containing the contract specifics and raw log

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
func (it *LendingTokenBorrowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingTokenBorrowed)
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
		it.Event = new(LendingTokenBorrowed)
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
func (it *LendingTokenBorrowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingTokenBorrowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingTokenBorrowed represents a TokenBorrowed event raised by the Lending contract.
type LendingTokenBorrowed struct {
	Borrower    common.Address
	TokenId     *big.Int
	FeePerBlock *big.Int
	Duration    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenBorrowed is a free log retrieval operation binding the contract event 0xc60cb8faabff3d0b881a70b4f73d2548d4f68e17d6dea66fd097980126f2b398.
//
// Solidity: event TokenBorrowed(address indexed borrower, uint256 tokenId, uint256 feePerBlock, uint256 duration)
func (_Lending *LendingFilterer) FilterTokenBorrowed(opts *bind.FilterOpts, borrower []common.Address) (*LendingTokenBorrowedIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Lending.contract.FilterLogs(opts, "TokenBorrowed", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &LendingTokenBorrowedIterator{contract: _Lending.contract, event: "TokenBorrowed", logs: logs, sub: sub}, nil
}

// WatchTokenBorrowed is a free log subscription operation binding the contract event 0xc60cb8faabff3d0b881a70b4f73d2548d4f68e17d6dea66fd097980126f2b398.
//
// Solidity: event TokenBorrowed(address indexed borrower, uint256 tokenId, uint256 feePerBlock, uint256 duration)
func (_Lending *LendingFilterer) WatchTokenBorrowed(opts *bind.WatchOpts, sink chan<- *LendingTokenBorrowed, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Lending.contract.WatchLogs(opts, "TokenBorrowed", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingTokenBorrowed)
				if err := _Lending.contract.UnpackLog(event, "TokenBorrowed", log); err != nil {
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

// ParseTokenBorrowed is a log parse operation binding the contract event 0xc60cb8faabff3d0b881a70b4f73d2548d4f68e17d6dea66fd097980126f2b398.
//
// Solidity: event TokenBorrowed(address indexed borrower, uint256 tokenId, uint256 feePerBlock, uint256 duration)
func (_Lending *LendingFilterer) ParseTokenBorrowed(log types.Log) (*LendingTokenBorrowed, error) {
	event := new(LendingTokenBorrowed)
	if err := _Lending.contract.UnpackLog(event, "TokenBorrowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingTokenLentIterator is returned from FilterTokenLent and is used to iterate over the raw logs and unpacked data for TokenLent events raised by the Lending contract.
type LendingTokenLentIterator struct {
	Event *LendingTokenLent // Event containing the contract specifics and raw log

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
func (it *LendingTokenLentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingTokenLent)
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
		it.Event = new(LendingTokenLent)
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
func (it *LendingTokenLentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingTokenLentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingTokenLent represents a TokenLent event raised by the Lending contract.
type LendingTokenLent struct {
	User    common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenLent is a free log retrieval operation binding the contract event 0x88647e6f7327f9338472aa17121c80a128d944eb771198644371a4a665b4279d.
//
// Solidity: event TokenLent(address indexed user, uint256 tokenId)
func (_Lending *LendingFilterer) FilterTokenLent(opts *bind.FilterOpts, user []common.Address) (*LendingTokenLentIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lending.contract.FilterLogs(opts, "TokenLent", userRule)
	if err != nil {
		return nil, err
	}
	return &LendingTokenLentIterator{contract: _Lending.contract, event: "TokenLent", logs: logs, sub: sub}, nil
}

// WatchTokenLent is a free log subscription operation binding the contract event 0x88647e6f7327f9338472aa17121c80a128d944eb771198644371a4a665b4279d.
//
// Solidity: event TokenLent(address indexed user, uint256 tokenId)
func (_Lending *LendingFilterer) WatchTokenLent(opts *bind.WatchOpts, sink chan<- *LendingTokenLent, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lending.contract.WatchLogs(opts, "TokenLent", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingTokenLent)
				if err := _Lending.contract.UnpackLog(event, "TokenLent", log); err != nil {
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

// ParseTokenLent is a log parse operation binding the contract event 0x88647e6f7327f9338472aa17121c80a128d944eb771198644371a4a665b4279d.
//
// Solidity: event TokenLent(address indexed user, uint256 tokenId)
func (_Lending *LendingFilterer) ParseTokenLent(log types.Log) (*LendingTokenLent, error) {
	event := new(LendingTokenLent)
	if err := _Lending.contract.UnpackLog(event, "TokenLent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingTokenReturnedIterator is returned from FilterTokenReturned and is used to iterate over the raw logs and unpacked data for TokenReturned events raised by the Lending contract.
type LendingTokenReturnedIterator struct {
	Event *LendingTokenReturned // Event containing the contract specifics and raw log

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
func (it *LendingTokenReturnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingTokenReturned)
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
		it.Event = new(LendingTokenReturned)
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
func (it *LendingTokenReturnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingTokenReturnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingTokenReturned represents a TokenReturned event raised by the Lending contract.
type LendingTokenReturned struct {
	Borrower common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenReturned is a free log retrieval operation binding the contract event 0xeaf449319c042c9ba3474fa0c5329eb58cd1f23be110cdbf9d697b8d303dac15.
//
// Solidity: event TokenReturned(address indexed borrower, uint256 tokenId)
func (_Lending *LendingFilterer) FilterTokenReturned(opts *bind.FilterOpts, borrower []common.Address) (*LendingTokenReturnedIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Lending.contract.FilterLogs(opts, "TokenReturned", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &LendingTokenReturnedIterator{contract: _Lending.contract, event: "TokenReturned", logs: logs, sub: sub}, nil
}

// WatchTokenReturned is a free log subscription operation binding the contract event 0xeaf449319c042c9ba3474fa0c5329eb58cd1f23be110cdbf9d697b8d303dac15.
//
// Solidity: event TokenReturned(address indexed borrower, uint256 tokenId)
func (_Lending *LendingFilterer) WatchTokenReturned(opts *bind.WatchOpts, sink chan<- *LendingTokenReturned, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Lending.contract.WatchLogs(opts, "TokenReturned", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingTokenReturned)
				if err := _Lending.contract.UnpackLog(event, "TokenReturned", log); err != nil {
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

// ParseTokenReturned is a log parse operation binding the contract event 0xeaf449319c042c9ba3474fa0c5329eb58cd1f23be110cdbf9d697b8d303dac15.
//
// Solidity: event TokenReturned(address indexed borrower, uint256 tokenId)
func (_Lending *LendingFilterer) ParseTokenReturned(log types.Log) (*LendingTokenReturned, error) {
	event := new(LendingTokenReturned)
	if err := _Lending.contract.UnpackLog(event, "TokenReturned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendingTokenWithdrawnIterator is returned from FilterTokenWithdrawn and is used to iterate over the raw logs and unpacked data for TokenWithdrawn events raised by the Lending contract.
type LendingTokenWithdrawnIterator struct {
	Event *LendingTokenWithdrawn // Event containing the contract specifics and raw log

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
func (it *LendingTokenWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingTokenWithdrawn)
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
		it.Event = new(LendingTokenWithdrawn)
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
func (it *LendingTokenWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingTokenWithdrawn represents a TokenWithdrawn event raised by the Lending contract.
type LendingTokenWithdrawn struct {
	User    common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdrawn is a free log retrieval operation binding the contract event 0xa2bd9fcfcdba69f52bcd9a520846ad4bd685b187483f53efc42d035b2ddebff0.
//
// Solidity: event TokenWithdrawn(address indexed user, uint256 tokenId)
func (_Lending *LendingFilterer) FilterTokenWithdrawn(opts *bind.FilterOpts, user []common.Address) (*LendingTokenWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lending.contract.FilterLogs(opts, "TokenWithdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &LendingTokenWithdrawnIterator{contract: _Lending.contract, event: "TokenWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokenWithdrawn is a free log subscription operation binding the contract event 0xa2bd9fcfcdba69f52bcd9a520846ad4bd685b187483f53efc42d035b2ddebff0.
//
// Solidity: event TokenWithdrawn(address indexed user, uint256 tokenId)
func (_Lending *LendingFilterer) WatchTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *LendingTokenWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lending.contract.WatchLogs(opts, "TokenWithdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingTokenWithdrawn)
				if err := _Lending.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
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

// ParseTokenWithdrawn is a log parse operation binding the contract event 0xa2bd9fcfcdba69f52bcd9a520846ad4bd685b187483f53efc42d035b2ddebff0.
//
// Solidity: event TokenWithdrawn(address indexed user, uint256 tokenId)
func (_Lending *LendingFilterer) ParseTokenWithdrawn(log types.Log) (*LendingTokenWithdrawn, error) {
	event := new(LendingTokenWithdrawn)
	if err := _Lending.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
