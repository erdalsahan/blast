// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// SafeProxyFactoryMetaData contains all meta data concerning the SafeProxyFactory contract.
var SafeProxyFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractSafeProxy\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"singleton\",\"type\":\"address\"}],\"name\":\"ProxyCreation\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_singleton\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initializer\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"saltNonce\",\"type\":\"uint256\"}],\"name\":\"createChainSpecificProxyWithNonce\",\"outputs\":[{\"internalType\":\"contractSafeProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_singleton\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initializer\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"saltNonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIProxyCreationCallback\",\"name\":\"callback\",\"type\":\"address\"}],\"name\":\"createProxyWithCallback\",\"outputs\":[{\"internalType\":\"contractSafeProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_singleton\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initializer\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"saltNonce\",\"type\":\"uint256\"}],\"name\":\"createProxyWithNonce\",\"outputs\":[{\"internalType\":\"contractSafeProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyCreationCode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061077a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80631688f0b91461005c5780633408e4701461008c57806353e5d9351461009a578063d18af54d146100af578063ec9e80bb146100c2575b600080fd5b61006f61006a36600461048a565b6100d5565b6040516001600160a01b0390911681526020015b60405180910390f35b604051468152602001610083565b6100a261016a565b604051610083919061053f565b61006f6100bd366004610559565b610194565b61006f6100d036600461048a565b61026a565b6000808380519060200120836040516020016100fb929190918252602082015260400190565b60405160208183030381529060405280519060200120905061011e85858361029c565b6040516001600160a01b038781168252919350908316907f4f51faf6c4561ff95f067657e43439f0f856d97c04d9ec9070a6199ad418e2359060200160405180910390a2509392505050565b60606040518060200161017c906103c2565b601f1982820381018352601f90910116604052919050565b60008083836040516020016101c592919091825260601b6bffffffffffffffffffffffff1916602082015260340190565b6040516020818303038152906040528051906020012060001c90506101eb8686836100d5565b91506001600160a01b03831615610261576040516303ca56a360e31b81526001600160a01b03841690631e52b5189061022e9085908a908a908a906004016105c5565b600060405180830381600087803b15801561024857600080fd5b505af115801561025c573d6000803e3d6000fd5b505050505b50949350505050565b60008083805190602001208361027d4690565b60408051602081019490945283019190915260608201526080016100fb565b6000833b6102f15760405162461bcd60e51b815260206004820152601f60248201527f53696e676c65746f6e20636f6e7472616374206e6f74206465706c6f7965640060448201526064015b60405180910390fd5b600060405180602001610303906103c2565b601f1982820381018352601f90910116604081905261033091906001600160a01b03881690602001610602565b6040516020818303038152906040529050828151826020016000f591506001600160a01b0382166103995760405162461bcd60e51b815260206004820152601360248201527210dc99585d194c8818d85b1b0819985a5b1959606a1b60448201526064016102e8565b8351156103ba5760008060008651602088016000875af1036103ba57600080fd5b509392505050565b6101498061062583390190565b6001600160a01b03811681146103e457600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261040e57600080fd5b813567ffffffffffffffff80821115610429576104296103e7565b604051601f8301601f19908116603f01168101908282118183101715610451576104516103e7565b8160405283815286602085880101111561046a57600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060006060848603121561049f57600080fd5b83356104aa816103cf565b9250602084013567ffffffffffffffff8111156104c657600080fd5b6104d2868287016103fd565b925050604084013590509250925092565b60005b838110156104fe5781810151838201526020016104e6565b8381111561050d576000848401525b50505050565b6000815180845261052b8160208601602086016104e3565b601f01601f19169290920160200192915050565b6020815260006105526020830184610513565b9392505050565b6000806000806080858703121561056f57600080fd5b843561057a816103cf565b9350602085013567ffffffffffffffff81111561059657600080fd5b6105a2878288016103fd565b9350506040850135915060608501356105ba816103cf565b939692955090935050565b6001600160a01b038581168252841660208201526080604082018190526000906105f190830185610513565b905082606083015295945050505050565b600083516106148184602088016104e3565b919091019182525060200191905056fe608060405234801561001057600080fd5b5060405161014938038061014983398101604081905261002f916100b9565b6001600160a01b0381166100945760405162461bcd60e51b815260206004820152602260248201527f496e76616c69642073696e676c65746f6e20616464726573732070726f766964604482015261195960f21b606482015260840160405180910390fd5b600080546001600160a01b0319166001600160a01b03929092169190911790556100e9565b6000602082840312156100cb57600080fd5b81516001600160a01b03811681146100e257600080fd5b9392505050565b6052806100f76000396000f3fe6080604052600080546001600160a01b0316632cf35bc960e11b823501602757808252602082f35b3682833781823684845af490503d82833e806040573d82fd5b503d81f3fea164736f6c634300080f000aa164736f6c634300080f000a",
}

// SafeProxyFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeProxyFactoryMetaData.ABI instead.
var SafeProxyFactoryABI = SafeProxyFactoryMetaData.ABI

// SafeProxyFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeProxyFactoryMetaData.Bin instead.
var SafeProxyFactoryBin = SafeProxyFactoryMetaData.Bin

// DeploySafeProxyFactory deploys a new Ethereum contract, binding an instance of SafeProxyFactory to it.
func DeploySafeProxyFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeProxyFactory, error) {
	parsed, err := SafeProxyFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeProxyFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeProxyFactory{SafeProxyFactoryCaller: SafeProxyFactoryCaller{contract: contract}, SafeProxyFactoryTransactor: SafeProxyFactoryTransactor{contract: contract}, SafeProxyFactoryFilterer: SafeProxyFactoryFilterer{contract: contract}}, nil
}

// SafeProxyFactory is an auto generated Go binding around an Ethereum contract.
type SafeProxyFactory struct {
	SafeProxyFactoryCaller     // Read-only binding to the contract
	SafeProxyFactoryTransactor // Write-only binding to the contract
	SafeProxyFactoryFilterer   // Log filterer for contract events
}

// SafeProxyFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeProxyFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeProxyFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeProxyFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeProxyFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeProxyFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeProxyFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeProxyFactorySession struct {
	Contract     *SafeProxyFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeProxyFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeProxyFactoryCallerSession struct {
	Contract *SafeProxyFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SafeProxyFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeProxyFactoryTransactorSession struct {
	Contract     *SafeProxyFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SafeProxyFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeProxyFactoryRaw struct {
	Contract *SafeProxyFactory // Generic contract binding to access the raw methods on
}

// SafeProxyFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeProxyFactoryCallerRaw struct {
	Contract *SafeProxyFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// SafeProxyFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeProxyFactoryTransactorRaw struct {
	Contract *SafeProxyFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeProxyFactory creates a new instance of SafeProxyFactory, bound to a specific deployed contract.
func NewSafeProxyFactory(address common.Address, backend bind.ContractBackend) (*SafeProxyFactory, error) {
	contract, err := bindSafeProxyFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactory{SafeProxyFactoryCaller: SafeProxyFactoryCaller{contract: contract}, SafeProxyFactoryTransactor: SafeProxyFactoryTransactor{contract: contract}, SafeProxyFactoryFilterer: SafeProxyFactoryFilterer{contract: contract}}, nil
}

// NewSafeProxyFactoryCaller creates a new read-only instance of SafeProxyFactory, bound to a specific deployed contract.
func NewSafeProxyFactoryCaller(address common.Address, caller bind.ContractCaller) (*SafeProxyFactoryCaller, error) {
	contract, err := bindSafeProxyFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactoryCaller{contract: contract}, nil
}

// NewSafeProxyFactoryTransactor creates a new write-only instance of SafeProxyFactory, bound to a specific deployed contract.
func NewSafeProxyFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeProxyFactoryTransactor, error) {
	contract, err := bindSafeProxyFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactoryTransactor{contract: contract}, nil
}

// NewSafeProxyFactoryFilterer creates a new log filterer instance of SafeProxyFactory, bound to a specific deployed contract.
func NewSafeProxyFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeProxyFactoryFilterer, error) {
	contract, err := bindSafeProxyFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactoryFilterer{contract: contract}, nil
}

// bindSafeProxyFactory binds a generic wrapper to an already deployed contract.
func bindSafeProxyFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeProxyFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeProxyFactory *SafeProxyFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeProxyFactory.Contract.SafeProxyFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeProxyFactory *SafeProxyFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.SafeProxyFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeProxyFactory *SafeProxyFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.SafeProxyFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeProxyFactory *SafeProxyFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeProxyFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeProxyFactory *SafeProxyFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeProxyFactory *SafeProxyFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.contract.Transact(opts, method, params...)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_SafeProxyFactory *SafeProxyFactoryCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SafeProxyFactory.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_SafeProxyFactory *SafeProxyFactorySession) GetChainId() (*big.Int, error) {
	return _SafeProxyFactory.Contract.GetChainId(&_SafeProxyFactory.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_SafeProxyFactory *SafeProxyFactoryCallerSession) GetChainId() (*big.Int, error) {
	return _SafeProxyFactory.Contract.GetChainId(&_SafeProxyFactory.CallOpts)
}

// ProxyCreationCode is a free data retrieval call binding the contract method 0x53e5d935.
//
// Solidity: function proxyCreationCode() pure returns(bytes)
func (_SafeProxyFactory *SafeProxyFactoryCaller) ProxyCreationCode(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _SafeProxyFactory.contract.Call(opts, &out, "proxyCreationCode")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ProxyCreationCode is a free data retrieval call binding the contract method 0x53e5d935.
//
// Solidity: function proxyCreationCode() pure returns(bytes)
func (_SafeProxyFactory *SafeProxyFactorySession) ProxyCreationCode() ([]byte, error) {
	return _SafeProxyFactory.Contract.ProxyCreationCode(&_SafeProxyFactory.CallOpts)
}

// ProxyCreationCode is a free data retrieval call binding the contract method 0x53e5d935.
//
// Solidity: function proxyCreationCode() pure returns(bytes)
func (_SafeProxyFactory *SafeProxyFactoryCallerSession) ProxyCreationCode() ([]byte, error) {
	return _SafeProxyFactory.Contract.ProxyCreationCode(&_SafeProxyFactory.CallOpts)
}

// CreateChainSpecificProxyWithNonce is a paid mutator transaction binding the contract method 0xec9e80bb.
//
// Solidity: function createChainSpecificProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactory *SafeProxyFactoryTransactor) CreateChainSpecificProxyWithNonce(opts *bind.TransactOpts, _singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactory.contract.Transact(opts, "createChainSpecificProxyWithNonce", _singleton, initializer, saltNonce)
}

// CreateChainSpecificProxyWithNonce is a paid mutator transaction binding the contract method 0xec9e80bb.
//
// Solidity: function createChainSpecificProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactory *SafeProxyFactorySession) CreateChainSpecificProxyWithNonce(_singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.CreateChainSpecificProxyWithNonce(&_SafeProxyFactory.TransactOpts, _singleton, initializer, saltNonce)
}

// CreateChainSpecificProxyWithNonce is a paid mutator transaction binding the contract method 0xec9e80bb.
//
// Solidity: function createChainSpecificProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactory *SafeProxyFactoryTransactorSession) CreateChainSpecificProxyWithNonce(_singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.CreateChainSpecificProxyWithNonce(&_SafeProxyFactory.TransactOpts, _singleton, initializer, saltNonce)
}

// CreateProxyWithCallback is a paid mutator transaction binding the contract method 0xd18af54d.
//
// Solidity: function createProxyWithCallback(address _singleton, bytes initializer, uint256 saltNonce, address callback) returns(address proxy)
func (_SafeProxyFactory *SafeProxyFactoryTransactor) CreateProxyWithCallback(opts *bind.TransactOpts, _singleton common.Address, initializer []byte, saltNonce *big.Int, callback common.Address) (*types.Transaction, error) {
	return _SafeProxyFactory.contract.Transact(opts, "createProxyWithCallback", _singleton, initializer, saltNonce, callback)
}

// CreateProxyWithCallback is a paid mutator transaction binding the contract method 0xd18af54d.
//
// Solidity: function createProxyWithCallback(address _singleton, bytes initializer, uint256 saltNonce, address callback) returns(address proxy)
func (_SafeProxyFactory *SafeProxyFactorySession) CreateProxyWithCallback(_singleton common.Address, initializer []byte, saltNonce *big.Int, callback common.Address) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.CreateProxyWithCallback(&_SafeProxyFactory.TransactOpts, _singleton, initializer, saltNonce, callback)
}

// CreateProxyWithCallback is a paid mutator transaction binding the contract method 0xd18af54d.
//
// Solidity: function createProxyWithCallback(address _singleton, bytes initializer, uint256 saltNonce, address callback) returns(address proxy)
func (_SafeProxyFactory *SafeProxyFactoryTransactorSession) CreateProxyWithCallback(_singleton common.Address, initializer []byte, saltNonce *big.Int, callback common.Address) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.CreateProxyWithCallback(&_SafeProxyFactory.TransactOpts, _singleton, initializer, saltNonce, callback)
}

// CreateProxyWithNonce is a paid mutator transaction binding the contract method 0x1688f0b9.
//
// Solidity: function createProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactory *SafeProxyFactoryTransactor) CreateProxyWithNonce(opts *bind.TransactOpts, _singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactory.contract.Transact(opts, "createProxyWithNonce", _singleton, initializer, saltNonce)
}

// CreateProxyWithNonce is a paid mutator transaction binding the contract method 0x1688f0b9.
//
// Solidity: function createProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactory *SafeProxyFactorySession) CreateProxyWithNonce(_singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.CreateProxyWithNonce(&_SafeProxyFactory.TransactOpts, _singleton, initializer, saltNonce)
}

// CreateProxyWithNonce is a paid mutator transaction binding the contract method 0x1688f0b9.
//
// Solidity: function createProxyWithNonce(address _singleton, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_SafeProxyFactory *SafeProxyFactoryTransactorSession) CreateProxyWithNonce(_singleton common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _SafeProxyFactory.Contract.CreateProxyWithNonce(&_SafeProxyFactory.TransactOpts, _singleton, initializer, saltNonce)
}

// SafeProxyFactoryProxyCreationIterator is returned from FilterProxyCreation and is used to iterate over the raw logs and unpacked data for ProxyCreation events raised by the SafeProxyFactory contract.
type SafeProxyFactoryProxyCreationIterator struct {
	Event *SafeProxyFactoryProxyCreation // Event containing the contract specifics and raw log

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
func (it *SafeProxyFactoryProxyCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeProxyFactoryProxyCreation)
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
		it.Event = new(SafeProxyFactoryProxyCreation)
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
func (it *SafeProxyFactoryProxyCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeProxyFactoryProxyCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeProxyFactoryProxyCreation represents a ProxyCreation event raised by the SafeProxyFactory contract.
type SafeProxyFactoryProxyCreation struct {
	Proxy     common.Address
	Singleton common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProxyCreation is a free log retrieval operation binding the contract event 0x4f51faf6c4561ff95f067657e43439f0f856d97c04d9ec9070a6199ad418e235.
//
// Solidity: event ProxyCreation(address indexed proxy, address singleton)
func (_SafeProxyFactory *SafeProxyFactoryFilterer) FilterProxyCreation(opts *bind.FilterOpts, proxy []common.Address) (*SafeProxyFactoryProxyCreationIterator, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _SafeProxyFactory.contract.FilterLogs(opts, "ProxyCreation", proxyRule)
	if err != nil {
		return nil, err
	}
	return &SafeProxyFactoryProxyCreationIterator{contract: _SafeProxyFactory.contract, event: "ProxyCreation", logs: logs, sub: sub}, nil
}

// WatchProxyCreation is a free log subscription operation binding the contract event 0x4f51faf6c4561ff95f067657e43439f0f856d97c04d9ec9070a6199ad418e235.
//
// Solidity: event ProxyCreation(address indexed proxy, address singleton)
func (_SafeProxyFactory *SafeProxyFactoryFilterer) WatchProxyCreation(opts *bind.WatchOpts, sink chan<- *SafeProxyFactoryProxyCreation, proxy []common.Address) (event.Subscription, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _SafeProxyFactory.contract.WatchLogs(opts, "ProxyCreation", proxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeProxyFactoryProxyCreation)
				if err := _SafeProxyFactory.contract.UnpackLog(event, "ProxyCreation", log); err != nil {
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

// ParseProxyCreation is a log parse operation binding the contract event 0x4f51faf6c4561ff95f067657e43439f0f856d97c04d9ec9070a6199ad418e235.
//
// Solidity: event ProxyCreation(address indexed proxy, address singleton)
func (_SafeProxyFactory *SafeProxyFactoryFilterer) ParseProxyCreation(log types.Log) (*SafeProxyFactoryProxyCreation, error) {
	event := new(SafeProxyFactoryProxyCreation)
	if err := _SafeProxyFactory.contract.UnpackLog(event, "ProxyCreation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
