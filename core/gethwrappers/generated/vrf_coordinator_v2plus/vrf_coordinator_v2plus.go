// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrf_coordinator_v2plus

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
)

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

type VRFCoordinatorV2PlusFeeConfig struct {
	FulfillmentFlatFeeLinkPPM uint32
	FulfillmentFlatFeeEthPPM  uint32
}

type VRFCoordinatorV2PlusRequestCommitment struct {
	BlockNum         uint64
	SubId            uint64
	CallbackGasLimit uint32
	NumWords         uint32
	Sender           common.Address
	NativePayment    bool
}

type VRFProof struct {
	Pk            [2]*big.Int
	Gamma         [2]*big.Int
	C             *big.Int
	S             *big.Int
	Seed          *big.Int
	UWitness      common.Address
	CGammaWitness [2]*big.Int
	SHashWitness  [2]*big.Int
	ZInv          *big.Int
}

type VRFV2PlusClientRandomWordsRequest struct {
	KeyHash              [32]byte
	SubId                uint64
	RequestConfirmations uint16
	CallbackGasLimit     uint32
	NumWords             uint32
	ExtraArgs            []byte
}

var VRFCoordinatorV2PlusMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"blockhashStore\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"internalBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"externalBalance\",\"type\":\"uint256\"}],\"name\":\"BalanceInvariantViolated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"}],\"name\":\"BlockhashNotInStore\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToSendEther\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"have\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"want\",\"type\":\"uint32\"}],\"name\":\"GasLimitTooBig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectCommitment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"have\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"InsufficientGasForConsumer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCalldata\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"InvalidConsumer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExtraArgsTag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"linkWei\",\"type\":\"int256\"}],\"name\":\"InvalidLinkWeiPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"have\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"min\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"max\",\"type\":\"uint16\"}],\"name\":\"InvalidRequestConfirmations\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSubscription\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LinkAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedOwner\",\"type\":\"address\"}],\"name\":\"MustBeRequestedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MustBeSubOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoCorrespondingRequest\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"name\":\"NoSuchProvingKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"have\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"want\",\"type\":\"uint32\"}],\"name\":\"NumWordsTooBig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableFromLink\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingRequestExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"name\":\"ProvingKeyAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyConsumers\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"int256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPM\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeEthPPM\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structVRFCoordinatorV2Plus.FeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthFundsRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"ProvingKeyDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"ProvingKeyRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outputSeed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"payment\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"nativePayment\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"preSeed\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"nativePayment\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RandomWordsRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLink\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountEth\",\"type\":\"uint256\"}],\"name\":\"SubscriptionCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"SubscriptionConsumerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"SubscriptionConsumerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SubscriptionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"SubscriptionFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldEthBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newEthBalance\",\"type\":\"uint256\"}],\"name\":\"SubscriptionFundedWithEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SubscriptionOwnerTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SubscriptionOwnerTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLOCKHASH_STORE\",\"outputs\":[{\"internalType\":\"contractBlockhashStoreInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LINK\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LINK_ETH_FEED\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_CONSUMERS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_NUM_WORDS\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_REQUEST_CONFIRMATIONS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"}],\"name\":\"acceptSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"addConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"cancelSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createSubscription\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"publicProvingKey\",\"type\":\"uint256[2]\"}],\"name\":\"deregisterProvingKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"pk\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"gamma\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uWitness\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"cGammaWitness\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"sHashWitness\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"zInv\",\"type\":\"uint256\"}],\"internalType\":\"structVRF.Proof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"nativePayment\",\"type\":\"bool\"}],\"internalType\":\"structVRFCoordinatorV2Plus.RequestCommitment\",\"name\":\"rc\",\"type\":\"tuple\"}],\"name\":\"fulfillRandomWords\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"}],\"name\":\"fundSubscriptionWithEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequestConfig\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"}],\"name\":\"getSubscription\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"balance\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"ethBalance\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"consumers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"publicKey\",\"type\":\"uint256[2]\"}],\"name\":\"hashOfKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"oracleWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"oracleWithdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"}],\"name\":\"ownerCancelSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"}],\"name\":\"pendingRequestExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"recoverEthFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"recoverFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"publicProvingKey\",\"type\":\"uint256[2]\"}],\"name\":\"registerProvingKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"removeConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"requestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structVRFV2PlusClient.RandomWordsRequest\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"requestRandomWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"requestSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_config\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_currentSubId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_fallbackWeiPerUnitLink\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_feeConfig\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPM\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeEthPPM\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_provingKeyHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"s_provingKeys\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_requestCommitments\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_totalBalance\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_totalEthBalance\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint32\"},{\"internalType\":\"int256\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"int256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkPPM\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeEthPPM\",\"type\":\"uint32\"}],\"internalType\":\"structVRFCoordinatorV2Plus.FeeConfig\",\"name\":\"feeConfig\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"}],\"name\":\"setLINK\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"linkEthFeed\",\"type\":\"address\"}],\"name\":\"setLinkEthFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200591e3803806200591e833981016040819052620000349162000188565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000dc565b505060016002555060601b6001600160601b031916608052620001ba565b6001600160a01b038116331415620001375760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000602082840312156200019b57600080fd5b81516001600160a01b0381168114620001b357600080fd5b9392505050565b60805160601c61573e620001e06000396000818161057d01526139b8015261573e6000f3fe6080604052600436106102db5760003560e01c8063827f1c5d11610184578063ad178361116100d6578063da2f26101161008a578063e95704bd11610064578063e95704bd1461091f578063ee9d2d3814610946578063f2fde38b1461097357600080fd5b8063da2f261014610899578063e72f6e30146108cf578063e82ad7d4146108ef57600080fd5b8063caf70c4a116100bb578063caf70c4a14610839578063d7ae1d3014610859578063d98e620e1461087957600080fd5b8063ad178361146107f9578063b08c87951461081957600080fd5b80639f87fad711610138578063a47c769611610112578063a47c769614610789578063a4c0ed36146107b9578063a8cb447b146107d957600080fd5b80639f87fad714610734578063a02e061614610754578063a21a23e41461077457600080fd5b80638741e3ee116101695780638741e3ee146106be5780638da5cb5b146106f6578063981837d51461071457600080fd5b8063827f1c5d1461065a57806386fe91c71461069257600080fd5b806340d6bb821161023d578063689c4517116101f15780637341c10c116101cb5780637341c10c1461060557806379ba509714610625578063823597401461063a57600080fd5b8063689c45171461056b5780636b6feccc1461059f5780636f64f03f146105e557600080fd5b8063596b8b8811610222578063596b8b881461051657806364d51a2a1461053657806366316d8d1461054b57600080fd5b806340d6bb82146104cb57806346d8d486146104f657600080fd5b806308821d5811610294578063181f5a7711610279578063181f5a77146104345780631b6b6d23146104805780633697af8b146104b857600080fd5b806308821d58146103ec57806315c48b841461040c57600080fd5b8063043bd6ae116102c5578063043bd6ae1461032f57806304c357cb14610353578063088070f51461037357600080fd5b8062012291146102e057806302bcc5b61461030d575b600080fd5b3480156102ec57600080fd5b506102f5610993565b60405161030493929190615311565b60405180910390f35b34801561031957600080fd5b5061032d610328366004615151565b610a0f565b005b34801561033b57600080fd5b50610345600e5481565b604051908152602001610304565b34801561035f57600080fd5b5061032d61036e36600461516c565b610a86565b34801561037f57600080fd5b50600f546103bb9061ffff81169063ffffffff620100008204811691660100000000000081048216916a01000000000000000000009091041684565b6040805161ffff909516855263ffffffff938416602086015291831691840191909152166060820152608001610304565b3480156103f857600080fd5b5061032d610407366004614ea9565b610bef565b34801561041857600080fd5b5061042160c881565b60405161ffff9091168152602001610304565b34801561044057600080fd5b50604080518082018252601a81527f565246436f6f7264696e61746f725632506c757320312e302e300000000000006020820152905161030491906152bc565b34801561048c57600080fd5b506003546104a0906001600160a01b031681565b6040516001600160a01b039091168152602001610304565b61032d6104c6366004615151565b610d83565b3480156104d757600080fd5b506104e16101f481565b60405163ffffffff9091168152602001610304565b34801561050257600080fd5b5061032d610511366004614dbf565b610f05565b34801561052257600080fd5b50610345610531366004615043565b6110a0565b34801561054257600080fd5b50610421606481565b34801561055757600080fd5b5061032d610566366004614dbf565b6115b0565b34801561057757600080fd5b506104a07f000000000000000000000000000000000000000000000000000000000000000081565b3480156105ab57600080fd5b506010546105c89063ffffffff8082169164010000000090041682565b6040805163ffffffff938416815292909116602083015201610304565b3480156105f157600080fd5b5061032d610600366004614df4565b61176d565b34801561061157600080fd5b5061032d61062036600461516c565b611885565b34801561063157600080fd5b5061032d611a68565b34801561064657600080fd5b5061032d610655366004615151565b611b19565b34801561066657600080fd5b5061067a610675366004614f7b565b611ca0565b6040516001600160601b039091168152602001610304565b34801561069e57600080fd5b5060075461067a906801000000000000000090046001600160601b031681565b3480156106ca57600080fd5b506007546106de906001600160401b031681565b6040516001600160401b039091168152602001610304565b34801561070257600080fd5b506000546001600160a01b03166104a0565b34801561072057600080fd5b5061032d61072f366004614da2565b612153565b34801561074057600080fd5b5061032d61074f36600461516c565b61217d565b34801561076057600080fd5b5061032d61076f366004614da2565b6124ed565b34801561078057600080fd5b506106de61255a565b34801561079557600080fd5b506107a96107a4366004615151565b612735565b60405161030494939291906153d2565b3480156107c557600080fd5b5061032d6107d4366004614e21565b61282c565b3480156107e557600080fd5b5061032d6107f4366004614da2565b612a28565b34801561080557600080fd5b50600a546104a0906001600160a01b031681565b34801561082557600080fd5b5061032d610834366004615098565b612b44565b34801561084557600080fd5b50610345610854366004614ec5565b612cce565b34801561086557600080fd5b5061032d61087436600461516c565b612cfe565b34801561088557600080fd5b50610345610894366004614efe565b612de9565b3480156108a557600080fd5b506104a06108b4366004614efe565b600b602052600090815260409020546001600160a01b031681565b3480156108db57600080fd5b5061032d6108ea366004614da2565b612e0a565b3480156108fb57600080fd5b5061090f61090a366004615151565b612fd2565b6040519015158152602001610304565b34801561092b57600080fd5b5060075461067a90600160a01b90046001600160601b031681565b34801561095257600080fd5b50610345610961366004614efe565b600d6020526000908152604090205481565b34801561097f57600080fd5b5061032d61098e366004614da2565b6131f2565b600f54600c805460408051602080840282018101909252828152600094859460609461ffff8316946201000090930463ffffffff169391928391908301828280156109fd57602002820191906000526020600020905b8154815260200190600101908083116109e9575b50505050509050925092509250909192565b610a17613203565b6001600160401b0381166000908152600560205260409020546001600160a01b0316610a5657604051630fb532db60e11b815260040160405180910390fd5b6001600160401b038116600090815260056020526040902054610a839082906001600160a01b031661325f565b50565b6001600160401b03821660009081526005602052604090205482906001600160a01b031680610ac857604051630fb532db60e11b815260040160405180910390fd5b336001600160a01b03821614610b0157604051636c51fda960e11b81526001600160a01b03821660048201526024015b60405180910390fd5b600280541415610b415760405162461bcd60e51b815260206004820152601f60248201526000805160206157128339815191526044820152606401610af8565b600280556001600160401b0384166000908152600560205260409020600101546001600160a01b03848116911614610be4576001600160401b03841660008181526005602090815260409182902060010180546001600160a01b0319166001600160a01b0388169081179091558251338152918201527f69436ea6df009049404f564eff6622cd00522b0bd6a89efd9e52a355c4a879be91015b60405180910390a25b505060016002555050565b610bf7613203565b604080518082018252600091610c26919084906002908390839080828437600092019190915250612cce915050565b6000818152600b60205260409020549091506001600160a01b031680610c6257604051631dfd6e1360e21b815260048101839052602401610af8565b6000828152600b6020526040812080546001600160a01b03191690555b600c54811015610d3a5782600c8281548110610c9d57610c9d6156c2565b90600052602060002001541415610d2857600c805460009190610cc2906001906155a3565b81548110610cd257610cd26156c2565b9060005260206000200154905080600c8381548110610cf357610cf36156c2565b600091825260209091200155600c805480610d1057610d106156ac565b60019003818190600052602060002001600090559055505b80610d328161562a565b915050610c7f565b50806001600160a01b03167f72be339577868f868798bac2c93e52d6f034fef4689a9848996c14ebb7416c0d83604051610d7691815260200190565b60405180910390a2505050565b600280541415610dc35760405162461bcd60e51b815260206004820152601f60248201526000805160206157128339815191526044820152606401610af8565b600280556001600160401b0381166000908152600560205260409020546001600160a01b0316610e0657604051630fb532db60e11b815260040160405180910390fd5b6001600160401b03811660009081526006602052604090208054600160601b90046001600160601b0316903490600c610e3f838561554e565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555034600760148282829054906101000a90046001600160601b0316610e87919061554e565b92506101000a8154816001600160601b0302191690836001600160601b03160217905550816001600160401b03167f4e4421a74ab9b76003de9d9527153d002f1338c36c1bbc180069671631cf0958823484610ee3919061550b565b604080519283526020830191909152015b60405180910390a250506001600255565b600280541415610f455760405162461bcd60e51b815260206004820152601f60248201526000805160206157128339815191526044820152606401610af8565b60028055336000908152600960205260409020546001600160601b0380831691161015610f8557604051631e9acf1760e31b815260040160405180910390fd5b3360009081526009602052604081208054839290610fad9084906001600160601b03166155ba565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555080600760148282829054906101000a90046001600160601b0316610ff591906155ba565b92506101000a8154816001600160601b0302191690836001600160601b031602179055506000826001600160a01b0316826001600160601b031660405160006040518083038185875af1925050503d806000811461106f576040519150601f19603f3d011682016040523d82523d6000602084013e611074565b606091505b505090508061109657604051630dcf35db60e41b815260040160405180910390fd5b5050600160025550565b60006002805414156110e25760405162461bcd60e51b815260206004820152601f60248201526000805160206157128339815191526044820152606401610af8565b6002805560006005816110fb6040860160208701615151565b6001600160401b031681526020810191909152604001600020546001600160a01b0316141561113d57604051630fb532db60e11b815260040160405180910390fd5b33600090815260046020908152604080832091839161116191908701908701615151565b6001600160401b039081168252602082019290925260400160002054169050806111c0576111956040840160208501615151565b604051637800cff360e11b81526001600160401b039091166004820152336024820152604401610af8565b600f5461ffff166111d7606085016040860161507d565b61ffff1610806111fa575060c86111f4606085016040860161507d565b61ffff16115b156112405761120f606084016040850161507d565b600f5460405163539c34bb60e11b815261ffff92831660048201529116602482015260c86044820152606401610af8565b600f5462010000900463ffffffff1661125f6080850160608601615136565b63ffffffff1611156112c85761127b6080840160608501615136565b600f546040517ff5d7e01e00000000000000000000000000000000000000000000000000000000815263ffffffff9283166004820152620100009091049091166024820152604401610af8565b6101f46112db60a0850160808601615136565b63ffffffff16111561133a576112f760a0840160808501615136565b6040517f47386bec00000000000000000000000000000000000000000000000000000000815263ffffffff90911660048201526101f46024820152604401610af8565b6000611347826001615523565b90506000806113da86353361136260408a0160208b01615151565b6040805160208082018690526001600160a01b0394909416818301526001600160401b0392831660608201529188166080808401919091528151808403909101815260a08301825280519084012060c083019490945260e0808301859052815180840390910181526101009092019052805191012091565b909250905060006113f66113f160a0890189615449565b61365d565b5190508261140261372d565b61141260408a0160208b01615151565b61142260808b0160608c01615136565b61143260a08c0160808d01615136565b6040805160208101969096528501939093526001600160401b03909116606084015263ffffffff90811660808401521660a08201523360c082015260e00160408051601f1981840301815291815281516020928301206000868152600d84528290205533916114a5918a01908a01615151565b6001600160401b031688357fad8d0b36e0dd1bc233432f9e3c28d376b9713fc9d26ac2dc972077be458d86fe86866114e360608e0160408f0161507d565b8d60600160208101906114f69190615136565b8e60800160208101906115099190615136565b60408051958652602086019490945261ffff929092168484015263ffffffff908116606085015216608083015286151560a0830152519081900360c00190a433600090815260046020908152604080832087939092909161156f91908c01908c01615151565b6001600160401b0390811682526020820192909252604001600020805467ffffffffffffffff19169290911691909117905550506001600255949350505050565b6002805414156115f05760405162461bcd60e51b815260206004820152601f60248201526000805160206157128339815191526044820152606401610af8565b60028055336000908152600860205260409020546001600160601b038083169116101561163057604051631e9acf1760e31b815260040160405180910390fd5b33600090815260086020526040812080548392906116589084906001600160601b03166155ba565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555080600760088282829054906101000a90046001600160601b03166116a091906155ba565b82546101009290920a6001600160601b0381810219909316918316021790915560035460405163a9059cbb60e01b81526001600160a01b03868116600483015292851660248201529116915063a9059cbb90604401602060405180830381600087803b15801561170f57600080fd5b505af1158015611723573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117479190614ee1565b61176457604051631e9acf1760e31b815260040160405180910390fd5b50506001600255565b611775613203565b6040805180820182526000916117a4919084906002908390839080828437600092019190915250612cce915050565b6000818152600b60205260409020549091506001600160a01b0316156117f9576040517f4a0b8fa700000000000000000000000000000000000000000000000000000000815260048101829052602401610af8565b6000818152600b6020908152604080832080546001600160a01b0319166001600160a01b038816908117909155600c805460018101825594527fdf6966c971051c3d54ec59162606531493a51404a002842f56009d7e5cf4a8c7909301849055518381527fe729ae16526293f74ade739043022254f1489f616295a25bf72dfb4511ed73b89101610d76565b6001600160401b03821660009081526005602052604090205482906001600160a01b0316806118c757604051630fb532db60e11b815260040160405180910390fd5b336001600160a01b038216146118fb57604051636c51fda960e11b81526001600160a01b0382166004820152602401610af8565b60028054141561193b5760405162461bcd60e51b815260206004820152601f60248201526000805160206157128339815191526044820152606401610af8565b60028080556001600160401b0385166000908152600560205260409020015460641415611994576040517f05a48e0f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03831660009081526004602090815260408083206001600160401b03808916855292529091205416156119cd57610be4565b6001600160a01b03831660008181526004602090815260408083206001600160401b038916808552908352818420805467ffffffffffffffff19166001908117909155600584528285206002018054918201815585529383902090930180546001600160a01b031916851790555192835290917f43dc749a04ac8fb825cbd514f7c0e13f13bc6f2ee66043b76629d51776cff8e09101610bdb565b6001546001600160a01b03163314611ac25760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610af8565b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b600280541415611b595760405162461bcd60e51b815260206004820152601f60248201526000805160206157128339815191526044820152606401610af8565b600280556001600160401b0381166000908152600560205260409020546001600160a01b0316611b9c57604051630fb532db60e11b815260040160405180910390fd5b6001600160401b0381166000908152600560205260409020600101546001600160a01b03163314611c22576001600160401b038116600090815260056020526040908190206001015490517fd084e9750000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610af8565b6001600160401b0381166000818152600560209081526040918290208054336001600160a01b0319808316821784556001909301805490931690925583516001600160a01b03909116808252928101919091529092917f6f1dc65165ffffedfd8e507b4a0f1fcfdada045ed11f6c26ba27cedfe87802f09101610ef4565b6000600280541415611ce25760405162461bcd60e51b815260206004820152601f60248201526000805160206157128339815191526044820152606401610af8565b6002805560005a90506000611cf785856137c6565b90506000846060015163ffffffff166001600160401b03811115611d1d57611d1d6156d8565b604051908082528060200260200182016040528015611d46578160200160208202803683370190505b50905060005b856060015163ffffffff16811015611dc657826040015181604051602001611d7e929190918252602082015260400190565b6040516020818303038152906040528051906020012060001c828281518110611da957611da96156c2565b602090810291909101015280611dbe8161562a565b915050611d4c565b50602080830180516000908152600d90925260408083208390559051905182917f1fe543e30000000000000000000000000000000000000000000000000000000091611e1791908690602401615384565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090506000611e91886040015163ffffffff16896080015184613af3565b90506000611ec087600f600001600a9054906101000a900463ffffffff1663ffffffff163a8c60a00151613b41565b90508860a0015115611fdf576020808a01516001600160401b03166000908152600690915260409020546001600160601b03808316600160601b909204161015611f1d57604051631e9acf1760e31b815260040160405180910390fd5b6020898101516001600160401b031660009081526006909152604090208054829190600c90611f5d908490600160601b90046001600160601b03166155ba565b82546101009290920a6001600160601b0381810219909316918316021790915587516000908152600b60209081526040808320546001600160a01b031683526009909152812080548594509092611fb69185911661554e565b92506101000a8154816001600160601b0302191690836001600160601b031602179055506120dd565b6020808a01516001600160401b03166000908152600690915260409020546001600160601b038083169116101561202957604051631e9acf1760e31b815260040160405180910390fd5b6020808a01516001600160401b03166000908152600690915260408120805483929061205f9084906001600160601b03166155ba565b82546101009290920a6001600160601b0381810219909316918316021790915587516000908152600b60209081526040808320546001600160a01b0316835260089091528120805485945090926120b89185911661554e565b92506101000a8154816001600160601b0302191690836001600160601b031602179055505b60208087015160408089015160a08d015182519182526001600160601b03861694820194909452921515908301528315156060830152907f04624b1adcd18b732732ed3009cbdcc7dfc5f0340105149d0c193e8361d955e19060800160405180910390a260016002559998505050505050505050565b61215b613203565b600a80546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160401b03821660009081526005602052604090205482906001600160a01b0316806121bf57604051630fb532db60e11b815260040160405180910390fd5b336001600160a01b038216146121f357604051636c51fda960e11b81526001600160a01b0382166004820152602401610af8565b6002805414156122335760405162461bcd60e51b815260206004820152601f60248201526000805160206157128339815191526044820152606401610af8565b6002805561224084612fd2565b1561225e57604051631685ecdd60e31b815260040160405180910390fd5b6001600160a01b03831660009081526004602090815260408083206001600160401b038089168552925290912054166122c457604051637800cff360e11b81526001600160401b03851660048201526001600160a01b0384166024820152604401610af8565b6001600160401b03841660009081526005602090815260408083206002018054825181850281018501909352808352919290919083018282801561233157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612313575b5050505050905060006001825161234891906155a3565b905060005b825181101561246f57856001600160a01b0316838281518110612372576123726156c2565b60200260200101516001600160a01b0316141561245d57600083838151811061239d5761239d6156c2565b6020026020010151905080600560008a6001600160401b03166001600160401b0316815260200190815260200160002060020183815481106123e1576123e16156c2565b600091825260208083209190910180546001600160a01b0319166001600160a01b0394909416939093179092556001600160401b038a168152600590915260409020600201805480612435576124356156ac565b600082815260209020810160001990810180546001600160a01b03191690550190555061246f565b806124678161562a565b91505061234d565b506001600160a01b03851660008181526004602090815260408083206001600160401b038b1680855290835292819020805467ffffffffffffffff191690555192835290917f182bff9831466789164ca77075fffd84916d35a8180ba73c27e45634549b445b91015b60405180910390a25050600160025550505050565b6124f5613203565b6003546001600160a01b031615612538576040517f2d118a6e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380546001600160a01b0319166001600160a01b0392909216919091179055565b600060028054141561259c5760405162461bcd60e51b815260206004820152601f60248201526000805160206157128339815191526044820152606401610af8565b60028055600780546001600160401b03169060006125b983615645565b82546101009290920a6001600160401b0381810219909316918316021790915560075416905060008060405190808252806020026020018201604052801561260b578160200160208202803683370190505b50604080518082018252600080825260208083018281526001600160401b038816808452600683528584209451855492516001600160601b03908116600160601b027fffffffffffffffff0000000000000000000000000000000000000000000000009094169116179190911790935583516060810185523381528082018381528186018781529484526005835294909220825181546001600160a01b039182166001600160a01b03199182161783559551600183018054919092169616959095179094559151805194955090936126e99260028501920190614b28565b50506040513381526001600160401b03841691507f464722b4166576d3dcbba877b999bc35cf911f4eaf434b7eba68fa113951d0bf9060200160405180910390a2509050600160025590565b6001600160401b038116600090815260056020526040812054819081906060906001600160a01b031661277b57604051630fb532db60e11b815260040160405180910390fd5b6001600160401b03851660009081526006602090815260408083205460058352928190208054600290910180548351818602810186019094528084526001600160601b0380871696600160601b900416946001600160a01b0390931693919283919083018282801561281657602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116127f8575b5050505050905093509350935093509193509193565b60028054141561286c5760405162461bcd60e51b815260206004820152601f60248201526000805160206157128339815191526044820152606401610af8565b600280556003546001600160a01b031633146128b4576040517f44b0e3c300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b602081146128ee576040517f8129bbcd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006128fc82840184615151565b6001600160401b0381166000908152600560205260409020549091506001600160a01b031661293e57604051630fb532db60e11b815260040160405180910390fd5b6001600160401b038116600090815260066020526040812080546001600160601b03169186919061296f838561554e565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555084600760088282829054906101000a90046001600160601b03166129b7919061554e565b92506101000a8154816001600160601b0302191690836001600160601b03160217905550816001600160401b03167fd39ec07f4e209f627a4c427971473820dc129761ba28de8906bd56f57101d4f8828784612a13919061550b565b604080519283526020830191909152016124d8565b612a30613203565b6007544790600160a01b90046001600160601b031681811115612a70576040516354ced18160e11b81526004810182905260248101839052604401610af8565b81811015612b3f576000612a8482846155a3565b90506000846001600160a01b03168260405160006040518083038185875af1925050503d8060008114612ad3576040519150601f19603f3d011682016040523d82523d6000602084013e612ad8565b606091505b5050905080612afa57604051630dcf35db60e41b815260040160405180910390fd5b604080516001600160a01b0387168152602081018490527f879c9ea2b9d5345b84ccd12610b032602808517cebdb795007f3dcb4df377317910160405180910390a150505b505050565b612b4c613203565b60c861ffff87161115612b865760405163539c34bb60e11b815261ffff871660048201819052602482015260c86044820152606401610af8565b60008213612baa576040516321ea67b360e11b815260048101839052602401610af8565b604080516080808201835261ffff891680835263ffffffff89811660208086018290528a83168688018190528a84166060978801819052600f805465ffffffffffff19168717620100008602176dffffffffffffffff0000000000001916660100000000000084026dffffffff000000000000000000001916176a010000000000000000000083021790558951601080548c86015192881667ffffffffffffffff1990911617640100000000928816929092029190911790819055600e8c9055895196875286840194909452978501529483019590955291810186905283821660a08201529290911c1660c08201527f777357bb93f63d088f18112d3dba38457aec633eb8f1341e1d418380ad328e789060e00160405180910390a1505050505050565b600081604051602001612ce191906152ae565b604051602081830303815290604052805190602001209050919050565b6001600160401b03821660009081526005602052604090205482906001600160a01b031680612d4057604051630fb532db60e11b815260040160405180910390fd5b336001600160a01b03821614612d7457604051636c51fda960e11b81526001600160a01b0382166004820152602401610af8565b600280541415612db45760405162461bcd60e51b815260206004820152601f60248201526000805160206157128339815191526044820152606401610af8565b60028055612dc184612fd2565b15612ddf57604051631685ecdd60e31b815260040160405180910390fd5b610be4848461325f565b600c8181548110612df957600080fd5b600091825260209091200154905081565b612e12613203565b6003546040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000916001600160a01b0316906370a082319060240160206040518083038186803b158015612e6f57600080fd5b505afa158015612e83573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ea79190614f17565b6007549091506801000000000000000090046001600160601b031681811115612eed576040516354ced18160e11b81526004810182905260248101839052604401610af8565b81811015612b3f576000612f0182846155a3565b60035460405163a9059cbb60e01b81526001600160a01b0387811660048301526024820184905292935091169063a9059cbb90604401602060405180830381600087803b158015612f5157600080fd5b505af1158015612f65573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f899190614ee1565b50604080516001600160a01b0386168152602081018390527f59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b436600910160405180910390a150505050565b6001600160401b0381166000908152600560209081526040808320815160608101835281546001600160a01b039081168252600183015416818501526002820180548451818702810187018652818152879693958601939092919083018282801561306657602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613048575b505050505081525050905060005b8160400151518110156131e85760005b600c548110156131d557600061319e600c83815481106130a6576130a66156c2565b9060005260206000200154856040015185815181106130c7576130c76156c2565b60200260200101518860046000896040015189815181106130ea576130ea6156c2565b6020908102919091018101516001600160a01b0316825281810192909252604090810160009081206001600160401b03808f16835293522054166040805160208082018790526001600160a01b0395909516818301526001600160401b039384166060820152919092166080808301919091528251808303909101815260a08201835280519084012060c082019490945260e080820185905282518083039091018152610100909101909152805191012091565b506000818152600d6020526040902054909150156131c25750600195945050505050565b50806131cd8161562a565b915050613084565b50806131e08161562a565b915050613074565b5060009392505050565b6131fa613203565b610a8381613b91565b6000546001600160a01b0316331461325d5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610af8565b565b6001600160401b0382166000908152600560209081526040808320815160608101835281546001600160a01b039081168252600183015416818501526002820180548451818702810187018652818152929593948601938301828280156132ef57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116132d1575b505050919092525050506001600160401b03841660009081526006602090815260408083208151808301909252546001600160601b03808216808452600160601b9092041692820183905293945092915b8460400151518110156133c1576004600086604001518381518110613367576133676156c2565b6020908102919091018101516001600160a01b0316825281810192909252604090810160009081206001600160401b038b1682529092529020805467ffffffffffffffff19169055806133b98161562a565b915050613340565b506001600160401b038616600090815260056020526040812080546001600160a01b031990811682556001820180549091169055906134036002830182614b8d565b50506001600160401b038616600090815260066020526040902080547fffffffffffffffff0000000000000000000000000000000000000000000000001690556007805483919060089061346d9084906801000000000000000090046001600160601b03166155ba565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555080600760148282829054906101000a90046001600160601b03166134b591906155ba565b82546101009290920a6001600160601b0381810219909316918316021790915560035460405163a9059cbb60e01b81526001600160a01b03898116600483015292861660248201529116915063a9059cbb90604401602060405180830381600087803b15801561352457600080fd5b505af1158015613538573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061355c9190614ee1565b61357957604051631e9acf1760e31b815260040160405180910390fd5b6000856001600160a01b0316826001600160601b031660405160006040518083038185875af1925050503d80600081146135cf576040519150601f19603f3d011682016040523d82523d6000602084013e6135d4565b606091505b50509050806135f657604051630dcf35db60e41b815260040160405180910390fd5b604080516001600160a01b03881681526001600160601b03858116602083015284168183015290516001600160401b038916917fb3a76e2c52bb7a484e325118e4cadc4346fdaa16a7df0621c5525e1f7ab4d495919081900360600190a250505050505050565b604080516020810190915260008152816136865750604080516020810190915260008152613727565b7f92fd1338000000000000000000000000000000000000000000000000000000006136b183856155e2565b7fffffffff00000000000000000000000000000000000000000000000000000000161461370a576040517f5247fdce00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61371782600481866154e1565b8101906137249190614f30565b90505b92915050565b60004661a4b1811480613742575062066eed81145b156137bf5760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561378157600080fd5b505afa158015613795573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137b99190614f17565b91505090565b4391505090565b604080516060810182526000808252602082018190529181019190915260006137f28460000151612cce565b6000818152600b60205260409020549091506001600160a01b03168061382e57604051631dfd6e1360e21b815260048101839052602401610af8565b6000828660800151604051602001613850929190918252602082015260400190565b60408051601f1981840301815291815281516020928301206000818152600d909352912054909150806138af576040517f3688124a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b85516020808801516040808a015160608b015160808c0151925161391a96899690959491019586526001600160401b03948516602087015292909316604085015263ffffffff90811660608501529190911660808301526001600160a01b031660a082015260c00190565b604051602081830303815290604052805190602001208114613968576040517fd529142c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006139778760000151613c3b565b905080613a815786516040517fe9413d380000000000000000000000000000000000000000000000000000000081526001600160401b0390911660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063e9413d389060240160206040518083038186803b158015613a0257600080fd5b505afa158015613a16573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a3a9190614f17565b905080613a815786516040517f175dadad0000000000000000000000000000000000000000000000000000000081526001600160401b039091166004820152602401610af8565b6000886080015182604051602001613aa3929190918252602082015260400190565b6040516020818303038152906040528051906020012060001c90506000613aca8a83613d3e565b604080516060810182529889526020890196909652948701949094525093979650505050505050565b60005a611388811015613b0557600080fd5b611388810390508460408204820311613b1d57600080fd5b50823b613b2957600080fd5b60008083516020850160008789f190505b9392505050565b60008115613b6f57601054613b689086908690640100000000900463ffffffff1686613da9565b9050613b89565b601054613b86908690869063ffffffff1686613e13565b90505b949350505050565b6001600160a01b038116331415613bea5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610af8565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60004661a4b1811480613c50575062066eed81145b15613d2f57610100836001600160401b0316613c6a61372d565b613c7491906155a3565b1180613c905750613c8361372d565b836001600160401b031610155b15613c9e5750600092915050565b6040517f2b407a820000000000000000000000000000000000000000000000000000000081526001600160401b0384166004820152606490632b407a829060240160206040518083038186803b158015613cf757600080fd5b505afa158015613d0b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b3a9190614f17565b50506001600160401b03164090565b6000613d728360000151846020015185604001518660600151868860a001518960c001518a60e001518b6101000151613f1a565b60038360200151604051602001613d8a929190615370565b60408051601f1981840301815291905280516020909101209392505050565b600080613db4614155565b905060005a613dc3888861550b565b613dcd91906155a3565b613dd79085615584565b90506000613df063ffffffff871664e8d4a51000615584565b905082613dfd828461550b565b613e07919061550b565b98975050505050505050565b600080613e1e6141b1565b905060008113613e44576040516321ea67b360e11b815260048101829052602401610af8565b6000613e4e614155565b9050600082825a613e5f8b8b61550b565b613e6991906155a3565b613e739088615584565b613e7d919061550b565b613e8f90670de0b6b3a7640000615584565b613e999190615570565b90506000613eb263ffffffff881664e8d4a51000615584565b9050613eca816b033b2e3c9fd0803ce80000006155a3565b821115613f03576040517fe80fa38100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613f0d818361550b565b9998505050505050505050565b613f2389614298565b613f6f5760405162461bcd60e51b815260206004820152601a60248201527f7075626c6963206b6579206973206e6f74206f6e2063757276650000000000006044820152606401610af8565b613f7888614298565b613fc45760405162461bcd60e51b815260206004820152601560248201527f67616d6d61206973206e6f74206f6e20637572766500000000000000000000006044820152606401610af8565b613fcd83614298565b6140195760405162461bcd60e51b815260206004820152601d60248201527f6347616d6d615769746e657373206973206e6f74206f6e2063757276650000006044820152606401610af8565b61402282614298565b61406e5760405162461bcd60e51b815260206004820152601c60248201527f73486173685769746e657373206973206e6f74206f6e206375727665000000006044820152606401610af8565b61407a878a8887614371565b6140c65760405162461bcd60e51b815260206004820152601960248201527f6164647228632a706b2b732a6729213d5f755769746e657373000000000000006044820152606401610af8565b60006140d28a876144a6565b905060006140e5898b878b86898961450a565b905060006140f6838d8d8a8661462a565b9050808a146141475760405162461bcd60e51b815260206004820152600d60248201527f696e76616c69642070726f6f66000000000000000000000000000000000000006044820152606401610af8565b505050505050505050505050565b60004661a4b181148061416a575062066eed81145b156141a957606c6001600160a01b031663c6f7de0e6040518163ffffffff1660e01b815260040160206040518083038186803b15801561378157600080fd5b600091505090565b600f54600a54604080517ffeaf968c00000000000000000000000000000000000000000000000000000000815290516000936601000000000000900463ffffffff169283151592859283926001600160a01b03169163feaf968c9160048083019260a0929190829003018186803b15801561422b57600080fd5b505afa15801561423f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061426391906151a3565b509450909250849150508015614287575061427e82426155a3565b8463ffffffff16105b15613b895750600e54949350505050565b80516000906401000003d019116142f15760405162461bcd60e51b815260206004820152601260248201527f696e76616c696420782d6f7264696e61746500000000000000000000000000006044820152606401610af8565b60208201516401000003d0191161434a5760405162461bcd60e51b815260206004820152601260248201527f696e76616c696420792d6f7264696e61746500000000000000000000000000006044820152606401610af8565b60208201516401000003d01990800961436a8360005b602002015161466a565b1492915050565b60006001600160a01b0382166143c95760405162461bcd60e51b815260206004820152600b60248201527f626164207769746e6573730000000000000000000000000000000000000000006044820152606401610af8565b6020840151600090600116156143e057601c6143e3565b601b5b9050600070014551231950b75fc4402da1732fc9bebe1985876000602002015109865170014551231950b75fc4402da1732fc9bebe19918203925060009190890987516040805160008082526020820180845287905260ff88169282019290925260608101929092526080820183905291925060019060a0016020604051602081039080840390855afa15801561447e573d6000803e3d6000fd5b5050604051601f1901516001600160a01b039081169088161495505050505050949350505050565b6144ae614bab565b6144db600184846040516020016144c79392919061528d565b60405160208183030381529060405261468e565b90505b6144e781614298565b61372757805160408051602081019290925261450391016144c7565b90506144de565b614512614bab565b825186516401000003d01990819006910614156145715760405162461bcd60e51b815260206004820152601e60248201527f706f696e747320696e2073756d206d7573742062652064697374696e637400006044820152606401610af8565b61457c8789886146dd565b6145c85760405162461bcd60e51b815260206004820152601660248201527f4669727374206d756c20636865636b206661696c6564000000000000000000006044820152606401610af8565b6145d38486856146dd565b61461f5760405162461bcd60e51b815260206004820152601760248201527f5365636f6e64206d756c20636865636b206661696c65640000000000000000006044820152606401610af8565b613e07868484614817565b6000600286868685876040516020016146489695949392919061522e565b60408051601f1981840301815291905280516020909101209695505050505050565b6000806401000003d01980848509840990506401000003d019600782089392505050565b614696614bab565b61469f826148de565b81526146b46146af826000614360565b614919565b6020820181905260029006600114156146d8576020810180516401000003d0190390525b919050565b60008261472c5760405162461bcd60e51b815260206004820152600b60248201527f7a65726f207363616c61720000000000000000000000000000000000000000006044820152606401610af8565b835160208501516000906147429060029061566c565b1561474e57601c614751565b601b5b9050600070014551231950b75fc4402da1732fc9bebe198387096040805160008082526020820180845281905260ff86169282019290925260608101869052608081018390529192509060019060a0016020604051602081039080840390855afa1580156147c3573d6000803e3d6000fd5b5050506020604051035190506000866040516020016147e2919061521c565b60408051601f1981840301815291905280516020909101206001600160a01b0392831692169190911498975050505050505050565b61481f614bab565b83516020808601518551918601516000938493849361484093909190614939565b919450925090506401000003d0198582096001146148a05760405162461bcd60e51b815260206004820152601960248201527f696e765a206d75737420626520696e7665727365206f66207a000000000000006044820152606401610af8565b60405180604001604052806401000003d019806148bf576148bf615696565b87860981526020016401000003d0198785099052979650505050505050565b805160208201205b6401000003d01981106146d8576040805160208082019390935281518082038401815290820190915280519101206148e6565b60006137278260026149326401000003d019600161550b565b901c614a19565b60008080600180826401000003d019896401000003d019038808905060006401000003d0198b6401000003d019038a089050600061497983838585614abb565b909850905061498a88828e88614adf565b909850905061499b88828c87614adf565b909850905060006149ae8d878b85614adf565b90985090506149bf88828686614abb565b90985090506149d088828e89614adf565b9098509050818114614a05576401000003d019818a0998506401000003d01982890997506401000003d0198183099650614a09565b8196505b5050505050509450945094915050565b600080614a24614bc9565b6020808252818101819052604082015260608101859052608081018490526401000003d01960a0820152614a56614be7565b60208160c0846005600019fa925082614ab15760405162461bcd60e51b815260206004820152601260248201527f6269674d6f64457870206661696c7572652100000000000000000000000000006044820152606401610af8565b5195945050505050565b6000806401000003d0198487096401000003d0198487099097909650945050505050565b600080806401000003d019878509905060006401000003d01987876401000003d019030990506401000003d0198183086401000003d01986890990999098509650505050505050565b828054828255906000526020600020908101928215614b7d579160200282015b82811115614b7d57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190614b48565b50614b89929150614c05565b5090565b5080546000825590600052602060002090810190610a839190614c05565b60405180604001604052806002906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b5b80821115614b895760008155600101614c06565b80356146d8816156ee565b806040810183101561372757600080fd5b600082601f830112614c4757600080fd5b614c4f615496565b808385604086011115614c6157600080fd5b60005b6002811015614c83578135845260209384019390910190600101614c64565b509095945050505050565b600060c08284031215614ca057600080fd5b60405160c081018181106001600160401b0382111715614cc257614cc26156d8565b604052905080614cd183614d5a565b8152614cdf60208401614d5a565b6020820152614cf060408401614d46565b6040820152614d0160608401614d46565b60608201526080830135614d14816156ee565b608082015260a0830135614d2781615703565b60a0919091015292915050565b803561ffff811681146146d857600080fd5b803563ffffffff811681146146d857600080fd5b80356001600160401b03811681146146d857600080fd5b805169ffffffffffffffffffff811681146146d857600080fd5b80356001600160601b03811681146146d857600080fd5b600060208284031215614db457600080fd5b8135613b3a816156ee565b60008060408385031215614dd257600080fd5b8235614ddd816156ee565b9150614deb60208401614d8b565b90509250929050565b60008060608385031215614e0757600080fd5b8235614e12816156ee565b9150614deb8460208501614c25565b60008060008060608587031215614e3757600080fd5b8435614e42816156ee565b93506020850135925060408501356001600160401b0380821115614e6557600080fd5b818701915087601f830112614e7957600080fd5b813581811115614e8857600080fd5b886020828501011115614e9a57600080fd5b95989497505060200194505050565b600060408284031215614ebb57600080fd5b6137248383614c25565b600060408284031215614ed757600080fd5b6137248383614c36565b600060208284031215614ef357600080fd5b8151613b3a81615703565b600060208284031215614f1057600080fd5b5035919050565b600060208284031215614f2957600080fd5b5051919050565b600060208284031215614f4257600080fd5b604051602081018181106001600160401b0382111715614f6457614f646156d8565b6040528235614f7281615703565b81529392505050565b600080828403610260811215614f9057600080fd5b6101a080821215614fa057600080fd5b614fa86154be565b9150614fb48686614c36565b8252614fc38660408701614c36565b60208301526080850135604083015260a0850135606083015260c08501356080830152614ff260e08601614c1a565b60a083015261010061500687828801614c36565b60c0840152615019876101408801614c36565b60e0840152610180860135818401525081935061503886828701614c8e565b925050509250929050565b60006020828403121561505557600080fd5b81356001600160401b0381111561506b57600080fd5b820160c08185031215613b3a57600080fd5b60006020828403121561508f57600080fd5b61372482614d34565b60008060008060008086880360e08112156150b257600080fd5b6150bb88614d34565b96506150c960208901614d46565b95506150d760408901614d46565b94506150e560608901614d46565b9350608088013592506040609f198201121561510057600080fd5b50615109615496565b61511560a08901614d46565b815261512360c08901614d46565b6020820152809150509295509295509295565b60006020828403121561514857600080fd5b61372482614d46565b60006020828403121561516357600080fd5b61372482614d5a565b6000806040838503121561517f57600080fd5b61518883614d5a565b91506020830135615198816156ee565b809150509250929050565b600080600080600060a086880312156151bb57600080fd5b6151c486614d71565b94506020860151935060408601519250606086015191506151e760808701614d71565b90509295509295909350565b8060005b60028110156152165781518452602093840193909101906001016151f7565b50505050565b61522681836151f3565b604001919050565b86815261523e60208201876151f3565b61524b60608201866151f3565b61525860a08201856151f3565b61526560e08201846151f3565b60609190911b6bffffffffffffffffffffffff19166101208201526101340195945050505050565b83815261529d60208201846151f3565b606081019190915260800192915050565b6040810161372782846151f3565b600060208083528351808285015260005b818110156152e9578581018301518582016040015282016152cd565b818111156152fb576000604083870101525b50601f01601f1916929092016040019392505050565b60006060820161ffff86168352602063ffffffff86168185015260606040850152818551808452608086019150828701935060005b8181101561536257845183529383019391830191600101615346565b509098975050505050505050565b82815260608101613b3a60208301846151f3565b6000604082018483526020604081850152818551808452606086019150828701935060005b818110156153c5578451835293830193918301916001016153a9565b5090979650505050505050565b6000608082016001600160601b0380881684526020818816818601526001600160a01b03915081871660408601526080606086015282865180855260a087019150828801945060005b8181101561543957855185168352948301949183019160010161541b565b50909a9950505050505050505050565b6000808335601e1984360301811261546057600080fd5b8301803591506001600160401b0382111561547a57600080fd5b60200191503681900382131561548f57600080fd5b9250929050565b604080519081016001600160401b03811182821017156154b8576154b86156d8565b60405290565b60405161012081016001600160401b03811182821017156154b8576154b86156d8565b600080858511156154f157600080fd5b838611156154fe57600080fd5b5050820193919092039150565b6000821982111561551e5761551e615680565b500190565b60006001600160401b0380831681851680830382111561554557615545615680565b01949350505050565b60006001600160601b0380831681851680830382111561554557615545615680565b60008261557f5761557f615696565b500490565b600081600019048311821515161561559e5761559e615680565b500290565b6000828210156155b5576155b5615680565b500390565b60006001600160601b03838116908316818110156155da576155da615680565b039392505050565b7fffffffff0000000000000000000000000000000000000000000000000000000081358181169160048510156156225780818660040360031b1b83161692505b505092915050565b600060001982141561563e5761563e615680565b5060010190565b60006001600160401b038083168181141561566257615662615680565b6001019392505050565b60008261567b5761567b615696565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114610a8357600080fd5b8015158114610a8357600080fdfe5265656e7472616e637947756172643a207265656e7472616e742063616c6c00a164736f6c6343000806000a",
}

var VRFCoordinatorV2PlusABI = VRFCoordinatorV2PlusMetaData.ABI

var VRFCoordinatorV2PlusBin = VRFCoordinatorV2PlusMetaData.Bin

func DeployVRFCoordinatorV2Plus(auth *bind.TransactOpts, backend bind.ContractBackend, blockhashStore common.Address) (common.Address, *types.Transaction, *VRFCoordinatorV2Plus, error) {
	parsed, err := VRFCoordinatorV2PlusMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFCoordinatorV2PlusBin), backend, blockhashStore)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFCoordinatorV2Plus{VRFCoordinatorV2PlusCaller: VRFCoordinatorV2PlusCaller{contract: contract}, VRFCoordinatorV2PlusTransactor: VRFCoordinatorV2PlusTransactor{contract: contract}, VRFCoordinatorV2PlusFilterer: VRFCoordinatorV2PlusFilterer{contract: contract}}, nil
}

type VRFCoordinatorV2Plus struct {
	address common.Address
	abi     abi.ABI
	VRFCoordinatorV2PlusCaller
	VRFCoordinatorV2PlusTransactor
	VRFCoordinatorV2PlusFilterer
}

type VRFCoordinatorV2PlusCaller struct {
	contract *bind.BoundContract
}

type VRFCoordinatorV2PlusTransactor struct {
	contract *bind.BoundContract
}

type VRFCoordinatorV2PlusFilterer struct {
	contract *bind.BoundContract
}

type VRFCoordinatorV2PlusSession struct {
	Contract     *VRFCoordinatorV2Plus
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFCoordinatorV2PlusCallerSession struct {
	Contract *VRFCoordinatorV2PlusCaller
	CallOpts bind.CallOpts
}

type VRFCoordinatorV2PlusTransactorSession struct {
	Contract     *VRFCoordinatorV2PlusTransactor
	TransactOpts bind.TransactOpts
}

type VRFCoordinatorV2PlusRaw struct {
	Contract *VRFCoordinatorV2Plus
}

type VRFCoordinatorV2PlusCallerRaw struct {
	Contract *VRFCoordinatorV2PlusCaller
}

type VRFCoordinatorV2PlusTransactorRaw struct {
	Contract *VRFCoordinatorV2PlusTransactor
}

func NewVRFCoordinatorV2Plus(address common.Address, backend bind.ContractBackend) (*VRFCoordinatorV2Plus, error) {
	abi, err := abi.JSON(strings.NewReader(VRFCoordinatorV2PlusABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVRFCoordinatorV2Plus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2Plus{address: address, abi: abi, VRFCoordinatorV2PlusCaller: VRFCoordinatorV2PlusCaller{contract: contract}, VRFCoordinatorV2PlusTransactor: VRFCoordinatorV2PlusTransactor{contract: contract}, VRFCoordinatorV2PlusFilterer: VRFCoordinatorV2PlusFilterer{contract: contract}}, nil
}

func NewVRFCoordinatorV2PlusCaller(address common.Address, caller bind.ContractCaller) (*VRFCoordinatorV2PlusCaller, error) {
	contract, err := bindVRFCoordinatorV2Plus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusCaller{contract: contract}, nil
}

func NewVRFCoordinatorV2PlusTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFCoordinatorV2PlusTransactor, error) {
	contract, err := bindVRFCoordinatorV2Plus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusTransactor{contract: contract}, nil
}

func NewVRFCoordinatorV2PlusFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFCoordinatorV2PlusFilterer, error) {
	contract, err := bindVRFCoordinatorV2Plus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusFilterer{contract: contract}, nil
}

func bindVRFCoordinatorV2Plus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFCoordinatorV2PlusMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFCoordinatorV2Plus.Contract.VRFCoordinatorV2PlusCaller.contract.Call(opts, result, method, params...)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.VRFCoordinatorV2PlusTransactor.contract.Transfer(opts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.VRFCoordinatorV2PlusTransactor.contract.Transact(opts, method, params...)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFCoordinatorV2Plus.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.contract.Transfer(opts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.contract.Transact(opts, method, params...)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) BLOCKHASHSTORE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "BLOCKHASH_STORE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) BLOCKHASHSTORE() (common.Address, error) {
	return _VRFCoordinatorV2Plus.Contract.BLOCKHASHSTORE(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) BLOCKHASHSTORE() (common.Address, error) {
	return _VRFCoordinatorV2Plus.Contract.BLOCKHASHSTORE(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) LINK(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "LINK")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) LINK() (common.Address, error) {
	return _VRFCoordinatorV2Plus.Contract.LINK(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) LINK() (common.Address, error) {
	return _VRFCoordinatorV2Plus.Contract.LINK(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) LINKETHFEED(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "LINK_ETH_FEED")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) LINKETHFEED() (common.Address, error) {
	return _VRFCoordinatorV2Plus.Contract.LINKETHFEED(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) LINKETHFEED() (common.Address, error) {
	return _VRFCoordinatorV2Plus.Contract.LINKETHFEED(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) MAXCONSUMERS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "MAX_CONSUMERS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) MAXCONSUMERS() (uint16, error) {
	return _VRFCoordinatorV2Plus.Contract.MAXCONSUMERS(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) MAXCONSUMERS() (uint16, error) {
	return _VRFCoordinatorV2Plus.Contract.MAXCONSUMERS(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) MAXNUMWORDS(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "MAX_NUM_WORDS")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) MAXNUMWORDS() (uint32, error) {
	return _VRFCoordinatorV2Plus.Contract.MAXNUMWORDS(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) MAXNUMWORDS() (uint32, error) {
	return _VRFCoordinatorV2Plus.Contract.MAXNUMWORDS(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) MAXREQUESTCONFIRMATIONS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "MAX_REQUEST_CONFIRMATIONS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) MAXREQUESTCONFIRMATIONS() (uint16, error) {
	return _VRFCoordinatorV2Plus.Contract.MAXREQUESTCONFIRMATIONS(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) MAXREQUESTCONFIRMATIONS() (uint16, error) {
	return _VRFCoordinatorV2Plus.Contract.MAXREQUESTCONFIRMATIONS(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) GetRequestConfig(opts *bind.CallOpts) (uint16, uint32, [][32]byte, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "getRequestConfig")

	if err != nil {
		return *new(uint16), *new(uint32), *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)
	out2 := *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)

	return out0, out1, out2, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) GetRequestConfig() (uint16, uint32, [][32]byte, error) {
	return _VRFCoordinatorV2Plus.Contract.GetRequestConfig(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) GetRequestConfig() (uint16, uint32, [][32]byte, error) {
	return _VRFCoordinatorV2Plus.Contract.GetRequestConfig(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) GetSubscription(opts *bind.CallOpts, subId uint64) (GetSubscription,

	error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "getSubscription", subId)

	outstruct := new(GetSubscription)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EthBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Consumers = *abi.ConvertType(out[3], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) GetSubscription(subId uint64) (GetSubscription,

	error) {
	return _VRFCoordinatorV2Plus.Contract.GetSubscription(&_VRFCoordinatorV2Plus.CallOpts, subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) GetSubscription(subId uint64) (GetSubscription,

	error) {
	return _VRFCoordinatorV2Plus.Contract.GetSubscription(&_VRFCoordinatorV2Plus.CallOpts, subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) HashOfKey(opts *bind.CallOpts, publicKey [2]*big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "hashOfKey", publicKey)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) HashOfKey(publicKey [2]*big.Int) ([32]byte, error) {
	return _VRFCoordinatorV2Plus.Contract.HashOfKey(&_VRFCoordinatorV2Plus.CallOpts, publicKey)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) HashOfKey(publicKey [2]*big.Int) ([32]byte, error) {
	return _VRFCoordinatorV2Plus.Contract.HashOfKey(&_VRFCoordinatorV2Plus.CallOpts, publicKey)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) Owner() (common.Address, error) {
	return _VRFCoordinatorV2Plus.Contract.Owner(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) Owner() (common.Address, error) {
	return _VRFCoordinatorV2Plus.Contract.Owner(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) PendingRequestExists(opts *bind.CallOpts, subId uint64) (bool, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "pendingRequestExists", subId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) PendingRequestExists(subId uint64) (bool, error) {
	return _VRFCoordinatorV2Plus.Contract.PendingRequestExists(&_VRFCoordinatorV2Plus.CallOpts, subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) PendingRequestExists(subId uint64) (bool, error) {
	return _VRFCoordinatorV2Plus.Contract.PendingRequestExists(&_VRFCoordinatorV2Plus.CallOpts, subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) SConfig(opts *bind.CallOpts) (SConfig,

	error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "s_config")

	outstruct := new(SConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinimumRequestConfirmations = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.MaxGasLimit = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.StalenessSeconds = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.GasAfterPaymentCalculation = *abi.ConvertType(out[3], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) SConfig() (SConfig,

	error) {
	return _VRFCoordinatorV2Plus.Contract.SConfig(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) SConfig() (SConfig,

	error) {
	return _VRFCoordinatorV2Plus.Contract.SConfig(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) SCurrentSubId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "s_currentSubId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) SCurrentSubId() (uint64, error) {
	return _VRFCoordinatorV2Plus.Contract.SCurrentSubId(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) SCurrentSubId() (uint64, error) {
	return _VRFCoordinatorV2Plus.Contract.SCurrentSubId(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) SFallbackWeiPerUnitLink(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "s_fallbackWeiPerUnitLink")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) SFallbackWeiPerUnitLink() (*big.Int, error) {
	return _VRFCoordinatorV2Plus.Contract.SFallbackWeiPerUnitLink(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) SFallbackWeiPerUnitLink() (*big.Int, error) {
	return _VRFCoordinatorV2Plus.Contract.SFallbackWeiPerUnitLink(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) SFeeConfig(opts *bind.CallOpts) (SFeeConfig,

	error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "s_feeConfig")

	outstruct := new(SFeeConfig)
	if err != nil {
		return *outstruct, err
	}

	outstruct.FulfillmentFlatFeeLinkPPM = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.FulfillmentFlatFeeEthPPM = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) SFeeConfig() (SFeeConfig,

	error) {
	return _VRFCoordinatorV2Plus.Contract.SFeeConfig(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) SFeeConfig() (SFeeConfig,

	error) {
	return _VRFCoordinatorV2Plus.Contract.SFeeConfig(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) SProvingKeyHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "s_provingKeyHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) SProvingKeyHashes(arg0 *big.Int) ([32]byte, error) {
	return _VRFCoordinatorV2Plus.Contract.SProvingKeyHashes(&_VRFCoordinatorV2Plus.CallOpts, arg0)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) SProvingKeyHashes(arg0 *big.Int) ([32]byte, error) {
	return _VRFCoordinatorV2Plus.Contract.SProvingKeyHashes(&_VRFCoordinatorV2Plus.CallOpts, arg0)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) SProvingKeys(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "s_provingKeys", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) SProvingKeys(arg0 [32]byte) (common.Address, error) {
	return _VRFCoordinatorV2Plus.Contract.SProvingKeys(&_VRFCoordinatorV2Plus.CallOpts, arg0)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) SProvingKeys(arg0 [32]byte) (common.Address, error) {
	return _VRFCoordinatorV2Plus.Contract.SProvingKeys(&_VRFCoordinatorV2Plus.CallOpts, arg0)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) SRequestCommitments(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "s_requestCommitments", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) SRequestCommitments(arg0 *big.Int) ([32]byte, error) {
	return _VRFCoordinatorV2Plus.Contract.SRequestCommitments(&_VRFCoordinatorV2Plus.CallOpts, arg0)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) SRequestCommitments(arg0 *big.Int) ([32]byte, error) {
	return _VRFCoordinatorV2Plus.Contract.SRequestCommitments(&_VRFCoordinatorV2Plus.CallOpts, arg0)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) STotalBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "s_totalBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) STotalBalance() (*big.Int, error) {
	return _VRFCoordinatorV2Plus.Contract.STotalBalance(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) STotalBalance() (*big.Int, error) {
	return _VRFCoordinatorV2Plus.Contract.STotalBalance(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) STotalEthBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "s_totalEthBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) STotalEthBalance() (*big.Int, error) {
	return _VRFCoordinatorV2Plus.Contract.STotalEthBalance(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) STotalEthBalance() (*big.Int, error) {
	return _VRFCoordinatorV2Plus.Contract.STotalEthBalance(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Plus.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) TypeAndVersion() (string, error) {
	return _VRFCoordinatorV2Plus.Contract.TypeAndVersion(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusCallerSession) TypeAndVersion() (string, error) {
	return _VRFCoordinatorV2Plus.Contract.TypeAndVersion(&_VRFCoordinatorV2Plus.CallOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "acceptOwnership")
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.AcceptOwnership(&_VRFCoordinatorV2Plus.TransactOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.AcceptOwnership(&_VRFCoordinatorV2Plus.TransactOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) AcceptSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "acceptSubscriptionOwnerTransfer", subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) AcceptSubscriptionOwnerTransfer(subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.AcceptSubscriptionOwnerTransfer(&_VRFCoordinatorV2Plus.TransactOpts, subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) AcceptSubscriptionOwnerTransfer(subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.AcceptSubscriptionOwnerTransfer(&_VRFCoordinatorV2Plus.TransactOpts, subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) AddConsumer(opts *bind.TransactOpts, subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "addConsumer", subId, consumer)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) AddConsumer(subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.AddConsumer(&_VRFCoordinatorV2Plus.TransactOpts, subId, consumer)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) AddConsumer(subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.AddConsumer(&_VRFCoordinatorV2Plus.TransactOpts, subId, consumer)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) CancelSubscription(opts *bind.TransactOpts, subId uint64, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "cancelSubscription", subId, to)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) CancelSubscription(subId uint64, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.CancelSubscription(&_VRFCoordinatorV2Plus.TransactOpts, subId, to)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) CancelSubscription(subId uint64, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.CancelSubscription(&_VRFCoordinatorV2Plus.TransactOpts, subId, to)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) CreateSubscription(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "createSubscription")
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) CreateSubscription() (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.CreateSubscription(&_VRFCoordinatorV2Plus.TransactOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) CreateSubscription() (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.CreateSubscription(&_VRFCoordinatorV2Plus.TransactOpts)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) DeregisterProvingKey(opts *bind.TransactOpts, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "deregisterProvingKey", publicProvingKey)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) DeregisterProvingKey(publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.DeregisterProvingKey(&_VRFCoordinatorV2Plus.TransactOpts, publicProvingKey)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) DeregisterProvingKey(publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.DeregisterProvingKey(&_VRFCoordinatorV2Plus.TransactOpts, publicProvingKey)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) FulfillRandomWords(opts *bind.TransactOpts, proof VRFProof, rc VRFCoordinatorV2PlusRequestCommitment) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "fulfillRandomWords", proof, rc)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) FulfillRandomWords(proof VRFProof, rc VRFCoordinatorV2PlusRequestCommitment) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.FulfillRandomWords(&_VRFCoordinatorV2Plus.TransactOpts, proof, rc)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) FulfillRandomWords(proof VRFProof, rc VRFCoordinatorV2PlusRequestCommitment) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.FulfillRandomWords(&_VRFCoordinatorV2Plus.TransactOpts, proof, rc)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) FundSubscriptionWithEth(opts *bind.TransactOpts, subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "fundSubscriptionWithEth", subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) FundSubscriptionWithEth(subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.FundSubscriptionWithEth(&_VRFCoordinatorV2Plus.TransactOpts, subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) FundSubscriptionWithEth(subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.FundSubscriptionWithEth(&_VRFCoordinatorV2Plus.TransactOpts, subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) OnTokenTransfer(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "onTokenTransfer", arg0, amount, data)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) OnTokenTransfer(arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.OnTokenTransfer(&_VRFCoordinatorV2Plus.TransactOpts, arg0, amount, data)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) OnTokenTransfer(arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.OnTokenTransfer(&_VRFCoordinatorV2Plus.TransactOpts, arg0, amount, data)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) OracleWithdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "oracleWithdraw", recipient, amount)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) OracleWithdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.OracleWithdraw(&_VRFCoordinatorV2Plus.TransactOpts, recipient, amount)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) OracleWithdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.OracleWithdraw(&_VRFCoordinatorV2Plus.TransactOpts, recipient, amount)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) OracleWithdrawEth(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "oracleWithdrawEth", recipient, amount)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) OracleWithdrawEth(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.OracleWithdrawEth(&_VRFCoordinatorV2Plus.TransactOpts, recipient, amount)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) OracleWithdrawEth(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.OracleWithdrawEth(&_VRFCoordinatorV2Plus.TransactOpts, recipient, amount)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) OwnerCancelSubscription(opts *bind.TransactOpts, subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "ownerCancelSubscription", subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) OwnerCancelSubscription(subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.OwnerCancelSubscription(&_VRFCoordinatorV2Plus.TransactOpts, subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) OwnerCancelSubscription(subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.OwnerCancelSubscription(&_VRFCoordinatorV2Plus.TransactOpts, subId)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) RecoverEthFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "recoverEthFunds", to)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) RecoverEthFunds(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RecoverEthFunds(&_VRFCoordinatorV2Plus.TransactOpts, to)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) RecoverEthFunds(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RecoverEthFunds(&_VRFCoordinatorV2Plus.TransactOpts, to)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) RecoverFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "recoverFunds", to)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) RecoverFunds(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RecoverFunds(&_VRFCoordinatorV2Plus.TransactOpts, to)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) RecoverFunds(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RecoverFunds(&_VRFCoordinatorV2Plus.TransactOpts, to)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) RegisterProvingKey(opts *bind.TransactOpts, oracle common.Address, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "registerProvingKey", oracle, publicProvingKey)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) RegisterProvingKey(oracle common.Address, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RegisterProvingKey(&_VRFCoordinatorV2Plus.TransactOpts, oracle, publicProvingKey)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) RegisterProvingKey(oracle common.Address, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RegisterProvingKey(&_VRFCoordinatorV2Plus.TransactOpts, oracle, publicProvingKey)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) RemoveConsumer(opts *bind.TransactOpts, subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "removeConsumer", subId, consumer)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) RemoveConsumer(subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RemoveConsumer(&_VRFCoordinatorV2Plus.TransactOpts, subId, consumer)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) RemoveConsumer(subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RemoveConsumer(&_VRFCoordinatorV2Plus.TransactOpts, subId, consumer)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) RequestRandomWords(opts *bind.TransactOpts, req VRFV2PlusClientRandomWordsRequest) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "requestRandomWords", req)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) RequestRandomWords(req VRFV2PlusClientRandomWordsRequest) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RequestRandomWords(&_VRFCoordinatorV2Plus.TransactOpts, req)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) RequestRandomWords(req VRFV2PlusClientRandomWordsRequest) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RequestRandomWords(&_VRFCoordinatorV2Plus.TransactOpts, req)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) RequestSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId uint64, newOwner common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "requestSubscriptionOwnerTransfer", subId, newOwner)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) RequestSubscriptionOwnerTransfer(subId uint64, newOwner common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RequestSubscriptionOwnerTransfer(&_VRFCoordinatorV2Plus.TransactOpts, subId, newOwner)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) RequestSubscriptionOwnerTransfer(subId uint64, newOwner common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.RequestSubscriptionOwnerTransfer(&_VRFCoordinatorV2Plus.TransactOpts, subId, newOwner)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) SetConfig(opts *bind.TransactOpts, minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, feeConfig VRFCoordinatorV2PlusFeeConfig) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "setConfig", minimumRequestConfirmations, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, feeConfig)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) SetConfig(minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, feeConfig VRFCoordinatorV2PlusFeeConfig) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.SetConfig(&_VRFCoordinatorV2Plus.TransactOpts, minimumRequestConfirmations, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, feeConfig)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) SetConfig(minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, feeConfig VRFCoordinatorV2PlusFeeConfig) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.SetConfig(&_VRFCoordinatorV2Plus.TransactOpts, minimumRequestConfirmations, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, feeConfig)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) SetLINK(opts *bind.TransactOpts, link common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "setLINK", link)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) SetLINK(link common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.SetLINK(&_VRFCoordinatorV2Plus.TransactOpts, link)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) SetLINK(link common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.SetLINK(&_VRFCoordinatorV2Plus.TransactOpts, link)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) SetLinkEthFeed(opts *bind.TransactOpts, linkEthFeed common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "setLinkEthFeed", linkEthFeed)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) SetLinkEthFeed(linkEthFeed common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.SetLinkEthFeed(&_VRFCoordinatorV2Plus.TransactOpts, linkEthFeed)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) SetLinkEthFeed(linkEthFeed common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.SetLinkEthFeed(&_VRFCoordinatorV2Plus.TransactOpts, linkEthFeed)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.contract.Transact(opts, "transferOwnership", to)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.TransferOwnership(&_VRFCoordinatorV2Plus.TransactOpts, to)
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Plus.Contract.TransferOwnership(&_VRFCoordinatorV2Plus.TransactOpts, to)
}

type VRFCoordinatorV2PlusConfigSetIterator struct {
	Event *VRFCoordinatorV2PlusConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusConfigSet)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusConfigSet)
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

func (it *VRFCoordinatorV2PlusConfigSetIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusConfigSet struct {
	MinimumRequestConfirmations uint16
	MaxGasLimit                 uint32
	StalenessSeconds            uint32
	GasAfterPaymentCalculation  uint32
	FallbackWeiPerUnitLink      *big.Int
	FeeConfig                   VRFCoordinatorV2PlusFeeConfig
	Raw                         types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterConfigSet(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusConfigSetIterator, error) {

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusConfigSetIterator{contract: _VRFCoordinatorV2Plus.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusConfigSet) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusConfigSet)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseConfigSet(log types.Log) (*VRFCoordinatorV2PlusConfigSet, error) {
	event := new(VRFCoordinatorV2PlusConfigSet)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusEthFundsRecoveredIterator struct {
	Event *VRFCoordinatorV2PlusEthFundsRecovered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusEthFundsRecoveredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusEthFundsRecovered)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusEthFundsRecovered)
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

func (it *VRFCoordinatorV2PlusEthFundsRecoveredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusEthFundsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusEthFundsRecovered struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterEthFundsRecovered(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusEthFundsRecoveredIterator, error) {

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "EthFundsRecovered")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusEthFundsRecoveredIterator{contract: _VRFCoordinatorV2Plus.contract, event: "EthFundsRecovered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchEthFundsRecovered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusEthFundsRecovered) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "EthFundsRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusEthFundsRecovered)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "EthFundsRecovered", log); err != nil {
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

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseEthFundsRecovered(log types.Log) (*VRFCoordinatorV2PlusEthFundsRecovered, error) {
	event := new(VRFCoordinatorV2PlusEthFundsRecovered)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "EthFundsRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusFundsRecoveredIterator struct {
	Event *VRFCoordinatorV2PlusFundsRecovered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusFundsRecoveredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusFundsRecovered)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusFundsRecovered)
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

func (it *VRFCoordinatorV2PlusFundsRecoveredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusFundsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusFundsRecovered struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterFundsRecovered(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusFundsRecoveredIterator, error) {

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "FundsRecovered")
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusFundsRecoveredIterator{contract: _VRFCoordinatorV2Plus.contract, event: "FundsRecovered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchFundsRecovered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusFundsRecovered) (event.Subscription, error) {

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "FundsRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusFundsRecovered)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "FundsRecovered", log); err != nil {
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

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseFundsRecovered(log types.Log) (*VRFCoordinatorV2PlusFundsRecovered, error) {
	event := new(VRFCoordinatorV2PlusFundsRecovered)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "FundsRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusOwnershipTransferRequestedIterator struct {
	Event *VRFCoordinatorV2PlusOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusOwnershipTransferRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusOwnershipTransferRequested)
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

func (it *VRFCoordinatorV2PlusOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV2PlusOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusOwnershipTransferRequestedIterator{contract: _VRFCoordinatorV2Plus.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusOwnershipTransferRequested)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseOwnershipTransferRequested(log types.Log) (*VRFCoordinatorV2PlusOwnershipTransferRequested, error) {
	event := new(VRFCoordinatorV2PlusOwnershipTransferRequested)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusOwnershipTransferredIterator struct {
	Event *VRFCoordinatorV2PlusOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusOwnershipTransferred)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusOwnershipTransferred)
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

func (it *VRFCoordinatorV2PlusOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV2PlusOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusOwnershipTransferredIterator{contract: _VRFCoordinatorV2Plus.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusOwnershipTransferred)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseOwnershipTransferred(log types.Log) (*VRFCoordinatorV2PlusOwnershipTransferred, error) {
	event := new(VRFCoordinatorV2PlusOwnershipTransferred)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusProvingKeyDeregisteredIterator struct {
	Event *VRFCoordinatorV2PlusProvingKeyDeregistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusProvingKeyDeregisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusProvingKeyDeregistered)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusProvingKeyDeregistered)
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

func (it *VRFCoordinatorV2PlusProvingKeyDeregisteredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusProvingKeyDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusProvingKeyDeregistered struct {
	KeyHash [32]byte
	Oracle  common.Address
	Raw     types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterProvingKeyDeregistered(opts *bind.FilterOpts, oracle []common.Address) (*VRFCoordinatorV2PlusProvingKeyDeregisteredIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "ProvingKeyDeregistered", oracleRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusProvingKeyDeregisteredIterator{contract: _VRFCoordinatorV2Plus.contract, event: "ProvingKeyDeregistered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchProvingKeyDeregistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusProvingKeyDeregistered, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "ProvingKeyDeregistered", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusProvingKeyDeregistered)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "ProvingKeyDeregistered", log); err != nil {
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

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseProvingKeyDeregistered(log types.Log) (*VRFCoordinatorV2PlusProvingKeyDeregistered, error) {
	event := new(VRFCoordinatorV2PlusProvingKeyDeregistered)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "ProvingKeyDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusProvingKeyRegisteredIterator struct {
	Event *VRFCoordinatorV2PlusProvingKeyRegistered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusProvingKeyRegisteredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusProvingKeyRegistered)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusProvingKeyRegistered)
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

func (it *VRFCoordinatorV2PlusProvingKeyRegisteredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusProvingKeyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusProvingKeyRegistered struct {
	KeyHash [32]byte
	Oracle  common.Address
	Raw     types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterProvingKeyRegistered(opts *bind.FilterOpts, oracle []common.Address) (*VRFCoordinatorV2PlusProvingKeyRegisteredIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "ProvingKeyRegistered", oracleRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusProvingKeyRegisteredIterator{contract: _VRFCoordinatorV2Plus.contract, event: "ProvingKeyRegistered", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchProvingKeyRegistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusProvingKeyRegistered, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "ProvingKeyRegistered", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusProvingKeyRegistered)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "ProvingKeyRegistered", log); err != nil {
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

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseProvingKeyRegistered(log types.Log) (*VRFCoordinatorV2PlusProvingKeyRegistered, error) {
	event := new(VRFCoordinatorV2PlusProvingKeyRegistered)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "ProvingKeyRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusRandomWordsFulfilledIterator struct {
	Event *VRFCoordinatorV2PlusRandomWordsFulfilled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusRandomWordsFulfilledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusRandomWordsFulfilled)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusRandomWordsFulfilled)
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

func (it *VRFCoordinatorV2PlusRandomWordsFulfilledIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusRandomWordsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusRandomWordsFulfilled struct {
	RequestId     *big.Int
	OutputSeed    *big.Int
	Payment       *big.Int
	NativePayment bool
	Success       bool
	Raw           types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterRandomWordsFulfilled(opts *bind.FilterOpts, requestId []*big.Int) (*VRFCoordinatorV2PlusRandomWordsFulfilledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "RandomWordsFulfilled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusRandomWordsFulfilledIterator{contract: _VRFCoordinatorV2Plus.contract, event: "RandomWordsFulfilled", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusRandomWordsFulfilled, requestId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "RandomWordsFulfilled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusRandomWordsFulfilled)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
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

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseRandomWordsFulfilled(log types.Log) (*VRFCoordinatorV2PlusRandomWordsFulfilled, error) {
	event := new(VRFCoordinatorV2PlusRandomWordsFulfilled)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusRandomWordsRequestedIterator struct {
	Event *VRFCoordinatorV2PlusRandomWordsRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusRandomWordsRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusRandomWordsRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusRandomWordsRequested)
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

func (it *VRFCoordinatorV2PlusRandomWordsRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusRandomWordsRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusRandomWordsRequested struct {
	KeyHash                     [32]byte
	RequestId                   *big.Int
	PreSeed                     *big.Int
	SubId                       uint64
	MinimumRequestConfirmations uint16
	CallbackGasLimit            uint32
	NumWords                    uint32
	NativePayment               bool
	Sender                      common.Address
	Raw                         types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterRandomWordsRequested(opts *bind.FilterOpts, keyHash [][32]byte, subId []uint64, sender []common.Address) (*VRFCoordinatorV2PlusRandomWordsRequestedIterator, error) {

	var keyHashRule []interface{}
	for _, keyHashItem := range keyHash {
		keyHashRule = append(keyHashRule, keyHashItem)
	}

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "RandomWordsRequested", keyHashRule, subIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusRandomWordsRequestedIterator{contract: _VRFCoordinatorV2Plus.contract, event: "RandomWordsRequested", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchRandomWordsRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusRandomWordsRequested, keyHash [][32]byte, subId []uint64, sender []common.Address) (event.Subscription, error) {

	var keyHashRule []interface{}
	for _, keyHashItem := range keyHash {
		keyHashRule = append(keyHashRule, keyHashItem)
	}

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "RandomWordsRequested", keyHashRule, subIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusRandomWordsRequested)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
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

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseRandomWordsRequested(log types.Log) (*VRFCoordinatorV2PlusRandomWordsRequested, error) {
	event := new(VRFCoordinatorV2PlusRandomWordsRequested)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusSubscriptionCanceledIterator struct {
	Event *VRFCoordinatorV2PlusSubscriptionCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusSubscriptionCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusSubscriptionCanceled)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusSubscriptionCanceled)
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

func (it *VRFCoordinatorV2PlusSubscriptionCanceledIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusSubscriptionCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusSubscriptionCanceled struct {
	SubId      uint64
	To         common.Address
	AmountLink *big.Int
	AmountEth  *big.Int
	Raw        types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterSubscriptionCanceled(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2PlusSubscriptionCanceledIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "SubscriptionCanceled", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusSubscriptionCanceledIterator{contract: _VRFCoordinatorV2Plus.contract, event: "SubscriptionCanceled", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchSubscriptionCanceled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionCanceled, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "SubscriptionCanceled", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusSubscriptionCanceled)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
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

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseSubscriptionCanceled(log types.Log) (*VRFCoordinatorV2PlusSubscriptionCanceled, error) {
	event := new(VRFCoordinatorV2PlusSubscriptionCanceled)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusSubscriptionConsumerAddedIterator struct {
	Event *VRFCoordinatorV2PlusSubscriptionConsumerAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusSubscriptionConsumerAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusSubscriptionConsumerAdded)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusSubscriptionConsumerAdded)
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

func (it *VRFCoordinatorV2PlusSubscriptionConsumerAddedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusSubscriptionConsumerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusSubscriptionConsumerAdded struct {
	SubId    uint64
	Consumer common.Address
	Raw      types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterSubscriptionConsumerAdded(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2PlusSubscriptionConsumerAddedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "SubscriptionConsumerAdded", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusSubscriptionConsumerAddedIterator{contract: _VRFCoordinatorV2Plus.contract, event: "SubscriptionConsumerAdded", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchSubscriptionConsumerAdded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionConsumerAdded, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "SubscriptionConsumerAdded", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusSubscriptionConsumerAdded)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
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

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseSubscriptionConsumerAdded(log types.Log) (*VRFCoordinatorV2PlusSubscriptionConsumerAdded, error) {
	event := new(VRFCoordinatorV2PlusSubscriptionConsumerAdded)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusSubscriptionConsumerRemovedIterator struct {
	Event *VRFCoordinatorV2PlusSubscriptionConsumerRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusSubscriptionConsumerRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusSubscriptionConsumerRemoved)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusSubscriptionConsumerRemoved)
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

func (it *VRFCoordinatorV2PlusSubscriptionConsumerRemovedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusSubscriptionConsumerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusSubscriptionConsumerRemoved struct {
	SubId    uint64
	Consumer common.Address
	Raw      types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterSubscriptionConsumerRemoved(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2PlusSubscriptionConsumerRemovedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "SubscriptionConsumerRemoved", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusSubscriptionConsumerRemovedIterator{contract: _VRFCoordinatorV2Plus.contract, event: "SubscriptionConsumerRemoved", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchSubscriptionConsumerRemoved(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionConsumerRemoved, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "SubscriptionConsumerRemoved", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusSubscriptionConsumerRemoved)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
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

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseSubscriptionConsumerRemoved(log types.Log) (*VRFCoordinatorV2PlusSubscriptionConsumerRemoved, error) {
	event := new(VRFCoordinatorV2PlusSubscriptionConsumerRemoved)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusSubscriptionCreatedIterator struct {
	Event *VRFCoordinatorV2PlusSubscriptionCreated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusSubscriptionCreatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusSubscriptionCreated)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusSubscriptionCreated)
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

func (it *VRFCoordinatorV2PlusSubscriptionCreatedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusSubscriptionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusSubscriptionCreated struct {
	SubId uint64
	Owner common.Address
	Raw   types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterSubscriptionCreated(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2PlusSubscriptionCreatedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "SubscriptionCreated", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusSubscriptionCreatedIterator{contract: _VRFCoordinatorV2Plus.contract, event: "SubscriptionCreated", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchSubscriptionCreated(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionCreated, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "SubscriptionCreated", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusSubscriptionCreated)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
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

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseSubscriptionCreated(log types.Log) (*VRFCoordinatorV2PlusSubscriptionCreated, error) {
	event := new(VRFCoordinatorV2PlusSubscriptionCreated)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusSubscriptionFundedIterator struct {
	Event *VRFCoordinatorV2PlusSubscriptionFunded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusSubscriptionFundedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusSubscriptionFunded)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusSubscriptionFunded)
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

func (it *VRFCoordinatorV2PlusSubscriptionFundedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusSubscriptionFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusSubscriptionFunded struct {
	SubId      uint64
	OldBalance *big.Int
	NewBalance *big.Int
	Raw        types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterSubscriptionFunded(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2PlusSubscriptionFundedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "SubscriptionFunded", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusSubscriptionFundedIterator{contract: _VRFCoordinatorV2Plus.contract, event: "SubscriptionFunded", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchSubscriptionFunded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionFunded, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "SubscriptionFunded", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusSubscriptionFunded)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
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

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseSubscriptionFunded(log types.Log) (*VRFCoordinatorV2PlusSubscriptionFunded, error) {
	event := new(VRFCoordinatorV2PlusSubscriptionFunded)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusSubscriptionFundedWithEthIterator struct {
	Event *VRFCoordinatorV2PlusSubscriptionFundedWithEth

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusSubscriptionFundedWithEthIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusSubscriptionFundedWithEth)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusSubscriptionFundedWithEth)
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

func (it *VRFCoordinatorV2PlusSubscriptionFundedWithEthIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusSubscriptionFundedWithEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusSubscriptionFundedWithEth struct {
	SubId         uint64
	OldEthBalance *big.Int
	NewEthBalance *big.Int
	Raw           types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterSubscriptionFundedWithEth(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2PlusSubscriptionFundedWithEthIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "SubscriptionFundedWithEth", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusSubscriptionFundedWithEthIterator{contract: _VRFCoordinatorV2Plus.contract, event: "SubscriptionFundedWithEth", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchSubscriptionFundedWithEth(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionFundedWithEth, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "SubscriptionFundedWithEth", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusSubscriptionFundedWithEth)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionFundedWithEth", log); err != nil {
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

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseSubscriptionFundedWithEth(log types.Log) (*VRFCoordinatorV2PlusSubscriptionFundedWithEth, error) {
	event := new(VRFCoordinatorV2PlusSubscriptionFundedWithEth)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionFundedWithEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusSubscriptionOwnerTransferRequestedIterator struct {
	Event *VRFCoordinatorV2PlusSubscriptionOwnerTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusSubscriptionOwnerTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusSubscriptionOwnerTransferRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusSubscriptionOwnerTransferRequested)
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

func (it *VRFCoordinatorV2PlusSubscriptionOwnerTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusSubscriptionOwnerTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusSubscriptionOwnerTransferRequested struct {
	SubId uint64
	From  common.Address
	To    common.Address
	Raw   types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterSubscriptionOwnerTransferRequested(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2PlusSubscriptionOwnerTransferRequestedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "SubscriptionOwnerTransferRequested", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusSubscriptionOwnerTransferRequestedIterator{contract: _VRFCoordinatorV2Plus.contract, event: "SubscriptionOwnerTransferRequested", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchSubscriptionOwnerTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionOwnerTransferRequested, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "SubscriptionOwnerTransferRequested", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusSubscriptionOwnerTransferRequested)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionOwnerTransferRequested", log); err != nil {
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

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseSubscriptionOwnerTransferRequested(log types.Log) (*VRFCoordinatorV2PlusSubscriptionOwnerTransferRequested, error) {
	event := new(VRFCoordinatorV2PlusSubscriptionOwnerTransferRequested)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionOwnerTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VRFCoordinatorV2PlusSubscriptionOwnerTransferredIterator struct {
	Event *VRFCoordinatorV2PlusSubscriptionOwnerTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VRFCoordinatorV2PlusSubscriptionOwnerTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFCoordinatorV2PlusSubscriptionOwnerTransferred)
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

	select {
	case log := <-it.logs:
		it.Event = new(VRFCoordinatorV2PlusSubscriptionOwnerTransferred)
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

func (it *VRFCoordinatorV2PlusSubscriptionOwnerTransferredIterator) Error() error {
	return it.fail
}

func (it *VRFCoordinatorV2PlusSubscriptionOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VRFCoordinatorV2PlusSubscriptionOwnerTransferred struct {
	SubId uint64
	From  common.Address
	To    common.Address
	Raw   types.Log
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) FilterSubscriptionOwnerTransferred(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2PlusSubscriptionOwnerTransferredIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.FilterLogs(opts, "SubscriptionOwnerTransferred", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2PlusSubscriptionOwnerTransferredIterator{contract: _VRFCoordinatorV2Plus.contract, event: "SubscriptionOwnerTransferred", logs: logs, sub: sub}, nil
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) WatchSubscriptionOwnerTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionOwnerTransferred, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VRFCoordinatorV2Plus.contract.WatchLogs(opts, "SubscriptionOwnerTransferred", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VRFCoordinatorV2PlusSubscriptionOwnerTransferred)
				if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionOwnerTransferred", log); err != nil {
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

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2PlusFilterer) ParseSubscriptionOwnerTransferred(log types.Log) (*VRFCoordinatorV2PlusSubscriptionOwnerTransferred, error) {
	event := new(VRFCoordinatorV2PlusSubscriptionOwnerTransferred)
	if err := _VRFCoordinatorV2Plus.contract.UnpackLog(event, "SubscriptionOwnerTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetSubscription struct {
	Balance    *big.Int
	EthBalance *big.Int
	Owner      common.Address
	Consumers  []common.Address
}
type SConfig struct {
	MinimumRequestConfirmations uint16
	MaxGasLimit                 uint32
	StalenessSeconds            uint32
	GasAfterPaymentCalculation  uint32
}
type SFeeConfig struct {
	FulfillmentFlatFeeLinkPPM uint32
	FulfillmentFlatFeeEthPPM  uint32
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2Plus) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VRFCoordinatorV2Plus.abi.Events["ConfigSet"].ID:
		return _VRFCoordinatorV2Plus.ParseConfigSet(log)
	case _VRFCoordinatorV2Plus.abi.Events["EthFundsRecovered"].ID:
		return _VRFCoordinatorV2Plus.ParseEthFundsRecovered(log)
	case _VRFCoordinatorV2Plus.abi.Events["FundsRecovered"].ID:
		return _VRFCoordinatorV2Plus.ParseFundsRecovered(log)
	case _VRFCoordinatorV2Plus.abi.Events["OwnershipTransferRequested"].ID:
		return _VRFCoordinatorV2Plus.ParseOwnershipTransferRequested(log)
	case _VRFCoordinatorV2Plus.abi.Events["OwnershipTransferred"].ID:
		return _VRFCoordinatorV2Plus.ParseOwnershipTransferred(log)
	case _VRFCoordinatorV2Plus.abi.Events["ProvingKeyDeregistered"].ID:
		return _VRFCoordinatorV2Plus.ParseProvingKeyDeregistered(log)
	case _VRFCoordinatorV2Plus.abi.Events["ProvingKeyRegistered"].ID:
		return _VRFCoordinatorV2Plus.ParseProvingKeyRegistered(log)
	case _VRFCoordinatorV2Plus.abi.Events["RandomWordsFulfilled"].ID:
		return _VRFCoordinatorV2Plus.ParseRandomWordsFulfilled(log)
	case _VRFCoordinatorV2Plus.abi.Events["RandomWordsRequested"].ID:
		return _VRFCoordinatorV2Plus.ParseRandomWordsRequested(log)
	case _VRFCoordinatorV2Plus.abi.Events["SubscriptionCanceled"].ID:
		return _VRFCoordinatorV2Plus.ParseSubscriptionCanceled(log)
	case _VRFCoordinatorV2Plus.abi.Events["SubscriptionConsumerAdded"].ID:
		return _VRFCoordinatorV2Plus.ParseSubscriptionConsumerAdded(log)
	case _VRFCoordinatorV2Plus.abi.Events["SubscriptionConsumerRemoved"].ID:
		return _VRFCoordinatorV2Plus.ParseSubscriptionConsumerRemoved(log)
	case _VRFCoordinatorV2Plus.abi.Events["SubscriptionCreated"].ID:
		return _VRFCoordinatorV2Plus.ParseSubscriptionCreated(log)
	case _VRFCoordinatorV2Plus.abi.Events["SubscriptionFunded"].ID:
		return _VRFCoordinatorV2Plus.ParseSubscriptionFunded(log)
	case _VRFCoordinatorV2Plus.abi.Events["SubscriptionFundedWithEth"].ID:
		return _VRFCoordinatorV2Plus.ParseSubscriptionFundedWithEth(log)
	case _VRFCoordinatorV2Plus.abi.Events["SubscriptionOwnerTransferRequested"].ID:
		return _VRFCoordinatorV2Plus.ParseSubscriptionOwnerTransferRequested(log)
	case _VRFCoordinatorV2Plus.abi.Events["SubscriptionOwnerTransferred"].ID:
		return _VRFCoordinatorV2Plus.ParseSubscriptionOwnerTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VRFCoordinatorV2PlusConfigSet) Topic() common.Hash {
	return common.HexToHash("0x777357bb93f63d088f18112d3dba38457aec633eb8f1341e1d418380ad328e78")
}

func (VRFCoordinatorV2PlusEthFundsRecovered) Topic() common.Hash {
	return common.HexToHash("0x879c9ea2b9d5345b84ccd12610b032602808517cebdb795007f3dcb4df377317")
}

func (VRFCoordinatorV2PlusFundsRecovered) Topic() common.Hash {
	return common.HexToHash("0x59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b436600")
}

func (VRFCoordinatorV2PlusOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VRFCoordinatorV2PlusOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (VRFCoordinatorV2PlusProvingKeyDeregistered) Topic() common.Hash {
	return common.HexToHash("0x72be339577868f868798bac2c93e52d6f034fef4689a9848996c14ebb7416c0d")
}

func (VRFCoordinatorV2PlusProvingKeyRegistered) Topic() common.Hash {
	return common.HexToHash("0xe729ae16526293f74ade739043022254f1489f616295a25bf72dfb4511ed73b8")
}

func (VRFCoordinatorV2PlusRandomWordsFulfilled) Topic() common.Hash {
	return common.HexToHash("0x04624b1adcd18b732732ed3009cbdcc7dfc5f0340105149d0c193e8361d955e1")
}

func (VRFCoordinatorV2PlusRandomWordsRequested) Topic() common.Hash {
	return common.HexToHash("0xad8d0b36e0dd1bc233432f9e3c28d376b9713fc9d26ac2dc972077be458d86fe")
}

func (VRFCoordinatorV2PlusSubscriptionCanceled) Topic() common.Hash {
	return common.HexToHash("0xb3a76e2c52bb7a484e325118e4cadc4346fdaa16a7df0621c5525e1f7ab4d495")
}

func (VRFCoordinatorV2PlusSubscriptionConsumerAdded) Topic() common.Hash {
	return common.HexToHash("0x43dc749a04ac8fb825cbd514f7c0e13f13bc6f2ee66043b76629d51776cff8e0")
}

func (VRFCoordinatorV2PlusSubscriptionConsumerRemoved) Topic() common.Hash {
	return common.HexToHash("0x182bff9831466789164ca77075fffd84916d35a8180ba73c27e45634549b445b")
}

func (VRFCoordinatorV2PlusSubscriptionCreated) Topic() common.Hash {
	return common.HexToHash("0x464722b4166576d3dcbba877b999bc35cf911f4eaf434b7eba68fa113951d0bf")
}

func (VRFCoordinatorV2PlusSubscriptionFunded) Topic() common.Hash {
	return common.HexToHash("0xd39ec07f4e209f627a4c427971473820dc129761ba28de8906bd56f57101d4f8")
}

func (VRFCoordinatorV2PlusSubscriptionFundedWithEth) Topic() common.Hash {
	return common.HexToHash("0x4e4421a74ab9b76003de9d9527153d002f1338c36c1bbc180069671631cf0958")
}

func (VRFCoordinatorV2PlusSubscriptionOwnerTransferRequested) Topic() common.Hash {
	return common.HexToHash("0x69436ea6df009049404f564eff6622cd00522b0bd6a89efd9e52a355c4a879be")
}

func (VRFCoordinatorV2PlusSubscriptionOwnerTransferred) Topic() common.Hash {
	return common.HexToHash("0x6f1dc65165ffffedfd8e507b4a0f1fcfdada045ed11f6c26ba27cedfe87802f0")
}

func (_VRFCoordinatorV2Plus *VRFCoordinatorV2Plus) Address() common.Address {
	return _VRFCoordinatorV2Plus.address
}

type VRFCoordinatorV2PlusInterface interface {
	BLOCKHASHSTORE(opts *bind.CallOpts) (common.Address, error)

	LINK(opts *bind.CallOpts) (common.Address, error)

	LINKETHFEED(opts *bind.CallOpts) (common.Address, error)

	MAXCONSUMERS(opts *bind.CallOpts) (uint16, error)

	MAXNUMWORDS(opts *bind.CallOpts) (uint32, error)

	MAXREQUESTCONFIRMATIONS(opts *bind.CallOpts) (uint16, error)

	GetRequestConfig(opts *bind.CallOpts) (uint16, uint32, [][32]byte, error)

	GetSubscription(opts *bind.CallOpts, subId uint64) (GetSubscription,

		error)

	HashOfKey(opts *bind.CallOpts, publicKey [2]*big.Int) ([32]byte, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	PendingRequestExists(opts *bind.CallOpts, subId uint64) (bool, error)

	SConfig(opts *bind.CallOpts) (SConfig,

		error)

	SCurrentSubId(opts *bind.CallOpts) (uint64, error)

	SFallbackWeiPerUnitLink(opts *bind.CallOpts) (*big.Int, error)

	SFeeConfig(opts *bind.CallOpts) (SFeeConfig,

		error)

	SProvingKeyHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error)

	SProvingKeys(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error)

	SRequestCommitments(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error)

	STotalBalance(opts *bind.CallOpts) (*big.Int, error)

	STotalEthBalance(opts *bind.CallOpts) (*big.Int, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AcceptSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId uint64) (*types.Transaction, error)

	AddConsumer(opts *bind.TransactOpts, subId uint64, consumer common.Address) (*types.Transaction, error)

	CancelSubscription(opts *bind.TransactOpts, subId uint64, to common.Address) (*types.Transaction, error)

	CreateSubscription(opts *bind.TransactOpts) (*types.Transaction, error)

	DeregisterProvingKey(opts *bind.TransactOpts, publicProvingKey [2]*big.Int) (*types.Transaction, error)

	FulfillRandomWords(opts *bind.TransactOpts, proof VRFProof, rc VRFCoordinatorV2PlusRequestCommitment) (*types.Transaction, error)

	FundSubscriptionWithEth(opts *bind.TransactOpts, subId uint64) (*types.Transaction, error)

	OnTokenTransfer(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error)

	OracleWithdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	OracleWithdrawEth(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	OwnerCancelSubscription(opts *bind.TransactOpts, subId uint64) (*types.Transaction, error)

	RecoverEthFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	RecoverFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	RegisterProvingKey(opts *bind.TransactOpts, oracle common.Address, publicProvingKey [2]*big.Int) (*types.Transaction, error)

	RemoveConsumer(opts *bind.TransactOpts, subId uint64, consumer common.Address) (*types.Transaction, error)

	RequestRandomWords(opts *bind.TransactOpts, req VRFV2PlusClientRandomWordsRequest) (*types.Transaction, error)

	RequestSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId uint64, newOwner common.Address) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, feeConfig VRFCoordinatorV2PlusFeeConfig) (*types.Transaction, error)

	SetLINK(opts *bind.TransactOpts, link common.Address) (*types.Transaction, error)

	SetLinkEthFeed(opts *bind.TransactOpts, linkEthFeed common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterConfigSet(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusConfigSet) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*VRFCoordinatorV2PlusConfigSet, error)

	FilterEthFundsRecovered(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusEthFundsRecoveredIterator, error)

	WatchEthFundsRecovered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusEthFundsRecovered) (event.Subscription, error)

	ParseEthFundsRecovered(log types.Log) (*VRFCoordinatorV2PlusEthFundsRecovered, error)

	FilterFundsRecovered(opts *bind.FilterOpts) (*VRFCoordinatorV2PlusFundsRecoveredIterator, error)

	WatchFundsRecovered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusFundsRecovered) (event.Subscription, error)

	ParseFundsRecovered(log types.Log) (*VRFCoordinatorV2PlusFundsRecovered, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV2PlusOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VRFCoordinatorV2PlusOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VRFCoordinatorV2PlusOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VRFCoordinatorV2PlusOwnershipTransferred, error)

	FilterProvingKeyDeregistered(opts *bind.FilterOpts, oracle []common.Address) (*VRFCoordinatorV2PlusProvingKeyDeregisteredIterator, error)

	WatchProvingKeyDeregistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusProvingKeyDeregistered, oracle []common.Address) (event.Subscription, error)

	ParseProvingKeyDeregistered(log types.Log) (*VRFCoordinatorV2PlusProvingKeyDeregistered, error)

	FilterProvingKeyRegistered(opts *bind.FilterOpts, oracle []common.Address) (*VRFCoordinatorV2PlusProvingKeyRegisteredIterator, error)

	WatchProvingKeyRegistered(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusProvingKeyRegistered, oracle []common.Address) (event.Subscription, error)

	ParseProvingKeyRegistered(log types.Log) (*VRFCoordinatorV2PlusProvingKeyRegistered, error)

	FilterRandomWordsFulfilled(opts *bind.FilterOpts, requestId []*big.Int) (*VRFCoordinatorV2PlusRandomWordsFulfilledIterator, error)

	WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusRandomWordsFulfilled, requestId []*big.Int) (event.Subscription, error)

	ParseRandomWordsFulfilled(log types.Log) (*VRFCoordinatorV2PlusRandomWordsFulfilled, error)

	FilterRandomWordsRequested(opts *bind.FilterOpts, keyHash [][32]byte, subId []uint64, sender []common.Address) (*VRFCoordinatorV2PlusRandomWordsRequestedIterator, error)

	WatchRandomWordsRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusRandomWordsRequested, keyHash [][32]byte, subId []uint64, sender []common.Address) (event.Subscription, error)

	ParseRandomWordsRequested(log types.Log) (*VRFCoordinatorV2PlusRandomWordsRequested, error)

	FilterSubscriptionCanceled(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2PlusSubscriptionCanceledIterator, error)

	WatchSubscriptionCanceled(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionCanceled, subId []uint64) (event.Subscription, error)

	ParseSubscriptionCanceled(log types.Log) (*VRFCoordinatorV2PlusSubscriptionCanceled, error)

	FilterSubscriptionConsumerAdded(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2PlusSubscriptionConsumerAddedIterator, error)

	WatchSubscriptionConsumerAdded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionConsumerAdded, subId []uint64) (event.Subscription, error)

	ParseSubscriptionConsumerAdded(log types.Log) (*VRFCoordinatorV2PlusSubscriptionConsumerAdded, error)

	FilterSubscriptionConsumerRemoved(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2PlusSubscriptionConsumerRemovedIterator, error)

	WatchSubscriptionConsumerRemoved(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionConsumerRemoved, subId []uint64) (event.Subscription, error)

	ParseSubscriptionConsumerRemoved(log types.Log) (*VRFCoordinatorV2PlusSubscriptionConsumerRemoved, error)

	FilterSubscriptionCreated(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2PlusSubscriptionCreatedIterator, error)

	WatchSubscriptionCreated(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionCreated, subId []uint64) (event.Subscription, error)

	ParseSubscriptionCreated(log types.Log) (*VRFCoordinatorV2PlusSubscriptionCreated, error)

	FilterSubscriptionFunded(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2PlusSubscriptionFundedIterator, error)

	WatchSubscriptionFunded(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionFunded, subId []uint64) (event.Subscription, error)

	ParseSubscriptionFunded(log types.Log) (*VRFCoordinatorV2PlusSubscriptionFunded, error)

	FilterSubscriptionFundedWithEth(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2PlusSubscriptionFundedWithEthIterator, error)

	WatchSubscriptionFundedWithEth(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionFundedWithEth, subId []uint64) (event.Subscription, error)

	ParseSubscriptionFundedWithEth(log types.Log) (*VRFCoordinatorV2PlusSubscriptionFundedWithEth, error)

	FilterSubscriptionOwnerTransferRequested(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2PlusSubscriptionOwnerTransferRequestedIterator, error)

	WatchSubscriptionOwnerTransferRequested(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionOwnerTransferRequested, subId []uint64) (event.Subscription, error)

	ParseSubscriptionOwnerTransferRequested(log types.Log) (*VRFCoordinatorV2PlusSubscriptionOwnerTransferRequested, error)

	FilterSubscriptionOwnerTransferred(opts *bind.FilterOpts, subId []uint64) (*VRFCoordinatorV2PlusSubscriptionOwnerTransferredIterator, error)

	WatchSubscriptionOwnerTransferred(opts *bind.WatchOpts, sink chan<- *VRFCoordinatorV2PlusSubscriptionOwnerTransferred, subId []uint64) (event.Subscription, error)

	ParseSubscriptionOwnerTransferred(log types.Log) (*VRFCoordinatorV2PlusSubscriptionOwnerTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
