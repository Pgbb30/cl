// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrf_v2plus_sub_owner

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

var VRFV2PlusExternalSubOwnerExampleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vrfCoordinator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"want\",\"type\":\"address\"}],\"name\":\"OnlyCoordinatorCanFulfill\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"want\",\"type\":\"address\"}],\"name\":\"OnlySubOwnerCanSetVRFCoordinator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"requestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"nativePayment\",\"type\":\"bool\"}],\"name\":\"requestRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_randomWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_requestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vrfCoordinator\",\"type\":\"address\"}],\"name\":\"setVRFCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516109ca3803806109ca83398101604081905261002f916100e6565b816001600160a01b0381166100795760405162461bcd60e51b815260206004820152600c60248201526b7a65726f206164647265737360a01b604482015260640160405180910390fd5b600080546001600160a01b039283166001600160a01b031991821617909155600280549483169482169490941790935560038054929091169183169190911790556006805490911633179055610119565b80516001600160a01b03811681146100e157600080fd5b919050565b600080604083850312156100f957600080fd5b610102836100ca565b9150610110602084016100ca565b90509250929050565b6108a2806101286000396000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c8063f2fde38b11610050578063f2fde38b146100ba578063f6eaffc8146100cd578063f793a70e146100e057600080fd5b80631fe543e31461007757806344ff81ce1461008c578063e89e106a1461009f575b600080fd5b61008a61008536600461061f565b6100f3565b005b61008a61009a3660046105b0565b610179565b6100a860055481565b60405190815260200160405180910390f35b61008a6100c83660046105b0565b610233565b6100a86100db3660046105ed565b61029e565b61008a6100ee36600461070e565b6102bf565b60005473ffffffffffffffffffffffffffffffffffffffff16331461016b576000546040517f1cf993f400000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff90911660248201526044015b60405180910390fd5b61017582826103f8565b5050565b60015473ffffffffffffffffffffffffffffffffffffffff1633146101ec576001546040517f4ae338ff00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff9091166024820152604401610162565b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60065473ffffffffffffffffffffffffffffffffffffffff16331461025757600080fd5b600680547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600481815481106102ae57600080fd5b600091825260209091200154905081565b60065473ffffffffffffffffffffffffffffffffffffffff1633146102e357600080fd5b60006040518060c001604052808481526020018867ffffffffffffffff1681526020018661ffff1681526020018763ffffffff1681526020018563ffffffff168152602001610341604051806020016040528086151581525061047b565b90526002546040517f596b8b8800000000000000000000000000000000000000000000000000000000815291925073ffffffffffffffffffffffffffffffffffffffff169063596b8b889061039a90849060040161079c565b602060405180830381600087803b1580156103b457600080fd5b505af11580156103c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103ec9190610606565b60055550505050505050565b6005548214610463576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7265717565737420494420697320696e636f72726563740000000000000000006044820152606401610162565b8051610476906004906020840190610537565b505050565b60607f92fd13387c7fe7befbc38d303d6468778fb9731bc4583f17d92989c6fcfdeaaa826040516024016104b491511515815260200190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915292915050565b828054828255906000526020600020908101928215610572579160200282015b82811115610572578251825591602001919060010190610557565b5061057e929150610582565b5090565b5b8082111561057e5760008155600101610583565b803563ffffffff811681146105ab57600080fd5b919050565b6000602082840312156105c257600080fd5b813573ffffffffffffffffffffffffffffffffffffffff811681146105e657600080fd5b9392505050565b6000602082840312156105ff57600080fd5b5035919050565b60006020828403121561061857600080fd5b5051919050565b6000806040838503121561063257600080fd5b8235915060208084013567ffffffffffffffff8082111561065257600080fd5b818601915086601f83011261066657600080fd5b81358181111561067857610678610866565b8060051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f830116810181811085821117156106bb576106bb610866565b604052828152858101935084860182860187018b10156106da57600080fd5b600095505b838610156106fd5780358552600195909501949386019386016106df565b508096505050505050509250929050565b60008060008060008060c0878903121561072757600080fd5b863567ffffffffffffffff8116811461073f57600080fd5b955061074d60208801610597565b9450604087013561ffff8116811461076457600080fd5b935061077260608801610597565b92506080870135915060a0870135801515811461078e57600080fd5b809150509295509295509295565b6000602080835283518184015267ffffffffffffffff8185015116604084015261ffff6040850151166060840152606084015163ffffffff80821660808601528060808701511660a0860152505060a084015160c08085015280518060e086015260005b8181101561081d5782810184015186820161010001528301610800565b8181111561083057600061010083880101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169390930161010001949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c6343000806000a",
}

var VRFV2PlusExternalSubOwnerExampleABI = VRFV2PlusExternalSubOwnerExampleMetaData.ABI

var VRFV2PlusExternalSubOwnerExampleBin = VRFV2PlusExternalSubOwnerExampleMetaData.Bin

func DeployVRFV2PlusExternalSubOwnerExample(auth *bind.TransactOpts, backend bind.ContractBackend, vrfCoordinator common.Address, link common.Address) (common.Address, *types.Transaction, *VRFV2PlusExternalSubOwnerExample, error) {
	parsed, err := VRFV2PlusExternalSubOwnerExampleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFV2PlusExternalSubOwnerExampleBin), backend, vrfCoordinator, link)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFV2PlusExternalSubOwnerExample{VRFV2PlusExternalSubOwnerExampleCaller: VRFV2PlusExternalSubOwnerExampleCaller{contract: contract}, VRFV2PlusExternalSubOwnerExampleTransactor: VRFV2PlusExternalSubOwnerExampleTransactor{contract: contract}, VRFV2PlusExternalSubOwnerExampleFilterer: VRFV2PlusExternalSubOwnerExampleFilterer{contract: contract}}, nil
}

type VRFV2PlusExternalSubOwnerExample struct {
	address common.Address
	abi     abi.ABI
	VRFV2PlusExternalSubOwnerExampleCaller
	VRFV2PlusExternalSubOwnerExampleTransactor
	VRFV2PlusExternalSubOwnerExampleFilterer
}

type VRFV2PlusExternalSubOwnerExampleCaller struct {
	contract *bind.BoundContract
}

type VRFV2PlusExternalSubOwnerExampleTransactor struct {
	contract *bind.BoundContract
}

type VRFV2PlusExternalSubOwnerExampleFilterer struct {
	contract *bind.BoundContract
}

type VRFV2PlusExternalSubOwnerExampleSession struct {
	Contract     *VRFV2PlusExternalSubOwnerExample
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VRFV2PlusExternalSubOwnerExampleCallerSession struct {
	Contract *VRFV2PlusExternalSubOwnerExampleCaller
	CallOpts bind.CallOpts
}

type VRFV2PlusExternalSubOwnerExampleTransactorSession struct {
	Contract     *VRFV2PlusExternalSubOwnerExampleTransactor
	TransactOpts bind.TransactOpts
}

type VRFV2PlusExternalSubOwnerExampleRaw struct {
	Contract *VRFV2PlusExternalSubOwnerExample
}

type VRFV2PlusExternalSubOwnerExampleCallerRaw struct {
	Contract *VRFV2PlusExternalSubOwnerExampleCaller
}

type VRFV2PlusExternalSubOwnerExampleTransactorRaw struct {
	Contract *VRFV2PlusExternalSubOwnerExampleTransactor
}

func NewVRFV2PlusExternalSubOwnerExample(address common.Address, backend bind.ContractBackend) (*VRFV2PlusExternalSubOwnerExample, error) {
	abi, err := abi.JSON(strings.NewReader(VRFV2PlusExternalSubOwnerExampleABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVRFV2PlusExternalSubOwnerExample(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFV2PlusExternalSubOwnerExample{address: address, abi: abi, VRFV2PlusExternalSubOwnerExampleCaller: VRFV2PlusExternalSubOwnerExampleCaller{contract: contract}, VRFV2PlusExternalSubOwnerExampleTransactor: VRFV2PlusExternalSubOwnerExampleTransactor{contract: contract}, VRFV2PlusExternalSubOwnerExampleFilterer: VRFV2PlusExternalSubOwnerExampleFilterer{contract: contract}}, nil
}

func NewVRFV2PlusExternalSubOwnerExampleCaller(address common.Address, caller bind.ContractCaller) (*VRFV2PlusExternalSubOwnerExampleCaller, error) {
	contract, err := bindVRFV2PlusExternalSubOwnerExample(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFV2PlusExternalSubOwnerExampleCaller{contract: contract}, nil
}

func NewVRFV2PlusExternalSubOwnerExampleTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFV2PlusExternalSubOwnerExampleTransactor, error) {
	contract, err := bindVRFV2PlusExternalSubOwnerExample(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFV2PlusExternalSubOwnerExampleTransactor{contract: contract}, nil
}

func NewVRFV2PlusExternalSubOwnerExampleFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFV2PlusExternalSubOwnerExampleFilterer, error) {
	contract, err := bindVRFV2PlusExternalSubOwnerExample(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFV2PlusExternalSubOwnerExampleFilterer{contract: contract}, nil
}

func bindVRFV2PlusExternalSubOwnerExample(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFV2PlusExternalSubOwnerExampleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExampleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFV2PlusExternalSubOwnerExample.Contract.VRFV2PlusExternalSubOwnerExampleCaller.contract.Call(opts, result, method, params...)
}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExampleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV2PlusExternalSubOwnerExample.Contract.VRFV2PlusExternalSubOwnerExampleTransactor.contract.Transfer(opts)
}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExampleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFV2PlusExternalSubOwnerExample.Contract.VRFV2PlusExternalSubOwnerExampleTransactor.contract.Transact(opts, method, params...)
}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExampleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFV2PlusExternalSubOwnerExample.Contract.contract.Call(opts, result, method, params...)
}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExampleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFV2PlusExternalSubOwnerExample.Contract.contract.Transfer(opts)
}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExampleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFV2PlusExternalSubOwnerExample.Contract.contract.Transact(opts, method, params...)
}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExampleCaller) SRandomWords(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2PlusExternalSubOwnerExample.contract.Call(opts, &out, "s_randomWords", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExampleSession) SRandomWords(arg0 *big.Int) (*big.Int, error) {
	return _VRFV2PlusExternalSubOwnerExample.Contract.SRandomWords(&_VRFV2PlusExternalSubOwnerExample.CallOpts, arg0)
}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExampleCallerSession) SRandomWords(arg0 *big.Int) (*big.Int, error) {
	return _VRFV2PlusExternalSubOwnerExample.Contract.SRandomWords(&_VRFV2PlusExternalSubOwnerExample.CallOpts, arg0)
}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExampleCaller) SRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFV2PlusExternalSubOwnerExample.contract.Call(opts, &out, "s_requestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExampleSession) SRequestId() (*big.Int, error) {
	return _VRFV2PlusExternalSubOwnerExample.Contract.SRequestId(&_VRFV2PlusExternalSubOwnerExample.CallOpts)
}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExampleCallerSession) SRequestId() (*big.Int, error) {
	return _VRFV2PlusExternalSubOwnerExample.Contract.SRequestId(&_VRFV2PlusExternalSubOwnerExample.CallOpts)
}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExampleTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFV2PlusExternalSubOwnerExample.contract.Transact(opts, "rawFulfillRandomWords", requestId, randomWords)
}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExampleSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFV2PlusExternalSubOwnerExample.Contract.RawFulfillRandomWords(&_VRFV2PlusExternalSubOwnerExample.TransactOpts, requestId, randomWords)
}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExampleTransactorSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFV2PlusExternalSubOwnerExample.Contract.RawFulfillRandomWords(&_VRFV2PlusExternalSubOwnerExample.TransactOpts, requestId, randomWords)
}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExampleTransactor) RequestRandomWords(opts *bind.TransactOpts, subId uint64, callbackGasLimit uint32, requestConfirmations uint16, numWords uint32, keyHash [32]byte, nativePayment bool) (*types.Transaction, error) {
	return _VRFV2PlusExternalSubOwnerExample.contract.Transact(opts, "requestRandomWords", subId, callbackGasLimit, requestConfirmations, numWords, keyHash, nativePayment)
}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExampleSession) RequestRandomWords(subId uint64, callbackGasLimit uint32, requestConfirmations uint16, numWords uint32, keyHash [32]byte, nativePayment bool) (*types.Transaction, error) {
	return _VRFV2PlusExternalSubOwnerExample.Contract.RequestRandomWords(&_VRFV2PlusExternalSubOwnerExample.TransactOpts, subId, callbackGasLimit, requestConfirmations, numWords, keyHash, nativePayment)
}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExampleTransactorSession) RequestRandomWords(subId uint64, callbackGasLimit uint32, requestConfirmations uint16, numWords uint32, keyHash [32]byte, nativePayment bool) (*types.Transaction, error) {
	return _VRFV2PlusExternalSubOwnerExample.Contract.RequestRandomWords(&_VRFV2PlusExternalSubOwnerExample.TransactOpts, subId, callbackGasLimit, requestConfirmations, numWords, keyHash, nativePayment)
}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExampleTransactor) SetVRFCoordinator(opts *bind.TransactOpts, _vrfCoordinator common.Address) (*types.Transaction, error) {
	return _VRFV2PlusExternalSubOwnerExample.contract.Transact(opts, "setVRFCoordinator", _vrfCoordinator)
}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExampleSession) SetVRFCoordinator(_vrfCoordinator common.Address) (*types.Transaction, error) {
	return _VRFV2PlusExternalSubOwnerExample.Contract.SetVRFCoordinator(&_VRFV2PlusExternalSubOwnerExample.TransactOpts, _vrfCoordinator)
}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExampleTransactorSession) SetVRFCoordinator(_vrfCoordinator common.Address) (*types.Transaction, error) {
	return _VRFV2PlusExternalSubOwnerExample.Contract.SetVRFCoordinator(&_VRFV2PlusExternalSubOwnerExample.TransactOpts, _vrfCoordinator)
}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExampleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VRFV2PlusExternalSubOwnerExample.contract.Transact(opts, "transferOwnership", newOwner)
}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExampleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VRFV2PlusExternalSubOwnerExample.Contract.TransferOwnership(&_VRFV2PlusExternalSubOwnerExample.TransactOpts, newOwner)
}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExampleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VRFV2PlusExternalSubOwnerExample.Contract.TransferOwnership(&_VRFV2PlusExternalSubOwnerExample.TransactOpts, newOwner)
}

func (_VRFV2PlusExternalSubOwnerExample *VRFV2PlusExternalSubOwnerExample) Address() common.Address {
	return _VRFV2PlusExternalSubOwnerExample.address
}

type VRFV2PlusExternalSubOwnerExampleInterface interface {
	SRandomWords(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error)

	SRequestId(opts *bind.CallOpts) (*big.Int, error)

	RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error)

	RequestRandomWords(opts *bind.TransactOpts, subId uint64, callbackGasLimit uint32, requestConfirmations uint16, numWords uint32, keyHash [32]byte, nativePayment bool) (*types.Transaction, error)

	SetVRFCoordinator(opts *bind.TransactOpts, _vrfCoordinator common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	Address() common.Address
}
