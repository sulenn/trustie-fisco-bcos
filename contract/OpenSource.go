// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package opensource

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/event"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OpenSourceABI is the input ABI used to generate the binding from.
const OpenSourceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"token_name\",\"type\":\"string\"},{\"name\":\"owner\",\"type\":\"string\"},{\"name\":\"total_supply\",\"type\":\"int256\"}],\"name\":\"insert\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OpenSourceBin is the compiled bytecode used for deploying new contracts.
var OpenSourceBin = "0x60806040523480156200001157600080fd5b506110016000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166356004b6a6040805190810160405280600481526020017f7265706f000000000000000000000000000000000000000000000000000000008152506040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016200010191906200040a565b602060405180830381600087803b1580156200011c57600080fd5b505af115801562000131573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525062000157919081019062000278565b506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166356004b6a6040805190810160405280600481526020017f75736572000000000000000000000000000000000000000000000000000000008152506040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401620002059190620003bc565b602060405180830381600087803b1580156200022057600080fd5b505af115801562000235573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506200025b919081019062000278565b50620004b4565b600062000270825162000463565b905092915050565b6000602082840312156200028b57600080fd5b60006200029b8482850162000262565b91505092915050565b6000620002b18262000458565b808452620002c78160208601602086016200046d565b620002d281620004a3565b602085010191505092915050565b6000601282527f6f776e65722c746f74616c5f737570706c7900000000000000000000000000006020830152604082019050919050565b6000601282527f746f6b656e5f6e616d652c62616c616e636500000000000000000000000000006020830152604082019050919050565b6000600482527f75736572000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000600a82527f746f6b656e5f6e616d65000000000000000000000000000000000000000000006020830152604082019050919050565b60006060820190508181036000830152620003d88184620002a4565b90508181036020830152620003ed816200034e565b90508181036040830152620004028162000317565b905092915050565b60006060820190508181036000830152620004268184620002a4565b905081810360208301526200043b8162000385565b905081810360408301526200045081620002e0565b905092915050565b600081519050919050565b6000819050919050565b60005b838110156200048d57808201518184015260208101905062000470565b838111156200049d576000848401525b50505050565b6000601f19601f8301169050919050565b610aaf80620004c46000396000f300608060405260043610610041576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806333575c7f14610046575b600080fd5b34801561005257600080fd5b5061006d60048036036100689190810190610696565b610083565b60405161007a9190610844565b60405180910390f35b60008060008060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f23f63c96040805190810160405280600481526020017f7265706f000000000000000000000000000000000000000000000000000000008152506040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401610136919061085f565b602060405180830381600087803b15801561015057600080fd5b505af1158015610164573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506101889190810190610644565b93506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f23f63c96040805190810160405280600481526020017f75736572000000000000000000000000000000000000000000000000000000008152506040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401610235919061085f565b602060405180830381600087803b15801561024f57600080fd5b505af1158015610263573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506102879190810190610644565b92508373ffffffffffffffffffffffffffffffffffffffff166313db93466040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156102ed57600080fd5b505af1158015610301573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610325919081019061061b565b91508173ffffffffffffffffffffffffffffffffffffffff1663e942b516896040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040161037c9190610914565b600060405180830381600087803b15801561039657600080fd5b505af11580156103aa573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff1663e942b516886040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040161040391906108b1565b600060405180830381600087803b15801561041d57600080fd5b505af1158015610431573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff16632ef8ba74876040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040161048a91906108e6565b600060405180830381600087803b1580156104a457600080fd5b505af11580156104b8573d6000803e3d6000fd5b505050508373ffffffffffffffffffffffffffffffffffffffff166331afac3689846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401610513929190610881565b602060405180830381600087803b15801561052d57600080fd5b505af1158015610541573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610565919081019061066d565b9050809450505050509392505050565b600061058182516109e2565b905092915050565b600061059582516109f4565b905092915050565b60006105a98235610a06565b905092915050565b60006105bd8251610a06565b905092915050565b600082601f83011215156105d857600080fd5b81356105eb6105e682610976565b610949565b9150808252602083016020830185838301111561060757600080fd5b610612838284610a22565b50505092915050565b60006020828403121561062d57600080fd5b600061063b84828501610575565b91505092915050565b60006020828403121561065657600080fd5b600061066484828501610589565b91505092915050565b60006020828403121561067f57600080fd5b600061068d848285016105b1565b91505092915050565b6000806000606084860312156106ab57600080fd5b600084013567ffffffffffffffff8111156106c557600080fd5b6106d1868287016105c5565b935050602084013567ffffffffffffffff8111156106ee57600080fd5b6106fa868287016105c5565b925050604061070b8682870161059d565b9150509250925092565b61071e81610a10565b82525050565b61072d816109d8565b82525050565b600061073e826109ad565b808452610752816020860160208601610a31565b61075b81610a64565b602085010191505092915050565b6000610774826109a2565b808452610788816020860160208601610a31565b61079181610a64565b602085010191505092915050565b6000600582527f6f776e65720000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000600c82527f746f74616c5f737570706c7900000000000000000000000000000000000000006020830152604082019050919050565b6000600a82527f746f6b656e5f6e616d65000000000000000000000000000000000000000000006020830152604082019050919050565b60006020820190506108596000830184610724565b92915050565b600060208201905081810360008301526108798184610769565b905092915050565b6000604082019050818103600083015261089b8185610733565b90506108aa6020830184610715565b9392505050565b600060408201905081810360008301526108ca8161079f565b905081810360208301526108de8184610733565b905092915050565b600060408201905081810360008301526108ff816107d6565b905061090e6020830184610724565b92915050565b6000604082019050818103600083015261092d8161080d565b905081810360208301526109418184610733565b905092915050565b6000604051905081810181811067ffffffffffffffff8211171561096c57600080fd5b8060405250919050565b600067ffffffffffffffff82111561098d57600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006109ed826109b8565b9050919050565b60006109ff826109b8565b9050919050565b6000819050919050565b6000610a1b826109b8565b9050919050565b82818337600083830152505050565b60005b83811015610a4f578082015181840152602081019050610a34565b83811115610a5e576000848401525b50505050565b6000601f19601f83011690509190505600a265627a7a7230582053527d6058a42d781413dd109ae89e05bc61b6908a695cd763cfc215f7d605856c6578706572696d656e74616cf50037"

// DeployOpenSource deploys a new contract, binding an instance of OpenSource to it.
func DeployOpenSource(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OpenSource, error) {
	parsed, err := abi.JSON(strings.NewReader(OpenSourceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OpenSourceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OpenSource{OpenSourceCaller: OpenSourceCaller{contract: contract}, OpenSourceTransactor: OpenSourceTransactor{contract: contract}, OpenSourceFilterer: OpenSourceFilterer{contract: contract}}, nil
}

func AsyncDeployOpenSource(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(OpenSourceABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(OpenSourceBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// OpenSource is an auto generated Go binding around a Solidity contract.
type OpenSource struct {
	OpenSourceCaller     // Read-only binding to the contract
	OpenSourceTransactor // Write-only binding to the contract
	OpenSourceFilterer   // Log filterer for contract events
}

// OpenSourceCaller is an auto generated read-only Go binding around a Solidity contract.
type OpenSourceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenSourceTransactor is an auto generated write-only Go binding around a Solidity contract.
type OpenSourceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenSourceFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type OpenSourceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpenSourceSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type OpenSourceSession struct {
	Contract     *OpenSource       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpenSourceCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type OpenSourceCallerSession struct {
	Contract *OpenSourceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// OpenSourceTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type OpenSourceTransactorSession struct {
	Contract     *OpenSourceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OpenSourceRaw is an auto generated low-level Go binding around a Solidity contract.
type OpenSourceRaw struct {
	Contract *OpenSource // Generic contract binding to access the raw methods on
}

// OpenSourceCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type OpenSourceCallerRaw struct {
	Contract *OpenSourceCaller // Generic read-only contract binding to access the raw methods on
}

// OpenSourceTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type OpenSourceTransactorRaw struct {
	Contract *OpenSourceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOpenSource creates a new instance of OpenSource, bound to a specific deployed contract.
func NewOpenSource(address common.Address, backend bind.ContractBackend) (*OpenSource, error) {
	contract, err := bindOpenSource(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OpenSource{OpenSourceCaller: OpenSourceCaller{contract: contract}, OpenSourceTransactor: OpenSourceTransactor{contract: contract}, OpenSourceFilterer: OpenSourceFilterer{contract: contract}}, nil
}

// NewOpenSourceCaller creates a new read-only instance of OpenSource, bound to a specific deployed contract.
func NewOpenSourceCaller(address common.Address, caller bind.ContractCaller) (*OpenSourceCaller, error) {
	contract, err := bindOpenSource(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpenSourceCaller{contract: contract}, nil
}

// NewOpenSourceTransactor creates a new write-only instance of OpenSource, bound to a specific deployed contract.
func NewOpenSourceTransactor(address common.Address, transactor bind.ContractTransactor) (*OpenSourceTransactor, error) {
	contract, err := bindOpenSource(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpenSourceTransactor{contract: contract}, nil
}

// NewOpenSourceFilterer creates a new log filterer instance of OpenSource, bound to a specific deployed contract.
func NewOpenSourceFilterer(address common.Address, filterer bind.ContractFilterer) (*OpenSourceFilterer, error) {
	contract, err := bindOpenSource(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpenSourceFilterer{contract: contract}, nil
}

// bindOpenSource binds a generic wrapper to an already deployed contract.
func bindOpenSource(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OpenSourceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpenSource *OpenSourceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OpenSource.Contract.OpenSourceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpenSource *OpenSourceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _OpenSource.Contract.OpenSourceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpenSource *OpenSourceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _OpenSource.Contract.OpenSourceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpenSource *OpenSourceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OpenSource.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpenSource *OpenSourceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _OpenSource.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpenSource *OpenSourceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _OpenSource.Contract.contract.Transact(opts, method, params...)
}

// Insert is a paid mutator transaction binding the contract method 0x33575c7f.
//
// Solidity: function insert(string token_name, string owner, int256 total_supply) returns(int256)
func (_OpenSource *OpenSourceTransactor) Insert(opts *bind.TransactOpts, token_name string, owner string, total_supply *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _OpenSource.contract.Transact(opts, "insert", token_name, owner, total_supply)
}

func (_OpenSource *OpenSourceTransactor) AsyncInsert(handler func(*types.Receipt, error), opts *bind.TransactOpts, token_name string, owner string, total_supply *big.Int) (*types.Transaction, error) {
	return _OpenSource.contract.AsyncTransact(opts, handler, "insert", token_name, owner, total_supply)
}

// Insert is a paid mutator transaction binding the contract method 0x33575c7f.
//
// Solidity: function insert(string token_name, string owner, int256 total_supply) returns(int256)
func (_OpenSource *OpenSourceSession) Insert(token_name string, owner string, total_supply *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _OpenSource.Contract.Insert(&_OpenSource.TransactOpts, token_name, owner, total_supply)
}

func (_OpenSource *OpenSourceSession) AsyncInsert(handler func(*types.Receipt, error), token_name string, owner string, total_supply *big.Int) (*types.Transaction, error) {
	return _OpenSource.Contract.AsyncInsert(handler, &_OpenSource.TransactOpts, token_name, owner, total_supply)
}

// Insert is a paid mutator transaction binding the contract method 0x33575c7f.
//
// Solidity: function insert(string token_name, string owner, int256 total_supply) returns(int256)
func (_OpenSource *OpenSourceTransactorSession) Insert(token_name string, owner string, total_supply *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _OpenSource.Contract.Insert(&_OpenSource.TransactOpts, token_name, owner, total_supply)
}

func (_OpenSource *OpenSourceTransactorSession) AsyncInsert(handler func(*types.Receipt, error), token_name string, owner string, total_supply *big.Int) (*types.Transaction, error) {
	return _OpenSource.Contract.AsyncInsert(handler, &_OpenSource.TransactOpts, token_name, owner, total_supply)
}
