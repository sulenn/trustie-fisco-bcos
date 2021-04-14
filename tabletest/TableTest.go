// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tabletest

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

// TableTestABI is the input ABI used to generate the binding from.
const TableTestABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"item_id\",\"type\":\"int256\"},{\"name\":\"item_name\",\"type\":\"string\"}],\"name\":\"update\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"item_id\",\"type\":\"int256\"}],\"name\":\"remove\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"item_id\",\"type\":\"int256\"},{\"name\":\"item_name\",\"type\":\"string\"}],\"name\":\"insert\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"select\",\"outputs\":[{\"name\":\"\",\"type\":\"string[]\"},{\"name\":\"\",\"type\":\"int256[]\"},{\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"count\",\"type\":\"int256\"}],\"name\":\"CreateResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"count\",\"type\":\"int256\"}],\"name\":\"InsertResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"count\",\"type\":\"int256\"}],\"name\":\"UpdateResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"count\",\"type\":\"int256\"}],\"name\":\"RemoveResult\",\"type\":\"event\"}]"

// TableTestBin is the compiled bytecode used for deploying new contracts.
var TableTestBin = "0x60806040523480156200001157600080fd5b506110016000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166356004b6a6040805190810160405280600681526020017f745f7465737400000000000000000000000000000000000000000000000000008152506040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016200010191906200024a565b602060405180830381600087803b1580156200011c57600080fd5b505af115801562000131573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525062000157919081019062000174565b50620002f4565b60006200016c8251620002a3565b905092915050565b6000602082840312156200018757600080fd5b600062000197848285016200015e565b91505092915050565b6000620001ad8262000298565b808452620001c3816020860160208601620002ad565b620001ce81620002e3565b602085010191505092915050565b6000600482527f6e616d65000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000601182527f6974656d5f69642c6974656d5f6e616d650000000000000000000000000000006020830152604082019050919050565b60006060820190508181036000830152620002668184620001a0565b905081810360208301526200027b81620001dc565b90508181036040830152620002908162000213565b905092915050565b600081519050919050565b6000819050919050565b60005b83811015620002cd578082015181840152602081019050620002b0565b83811115620002dd576000848401525b50505050565b6000601f19601f8301169050919050565b611fc980620003046000396000f300608060405260043610610062576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063487a5a1014610067578063c4f41ab3146100a4578063ebf3b24f146100e1578063fcd7e3c11461011e575b600080fd5b34801561007357600080fd5b5061008e6004803603610089919081019061193e565b61015d565b60405161009b9190611c10565b60405180910390f35b3480156100b057600080fd5b506100cb60048036036100c691908101906118ea565b610627565b6040516100d89190611c10565b60405180910390f35b3480156100ed57600080fd5b506101086004803603610103919081019061193e565b6109c6565b6040516101159190611c10565b60405180910390f35b34801561012a57600080fd5b5061014560048036036101409190810190611868565b610ded565b60405161015493929190611bc4565b60405180910390f35b60008060008060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f23f63c96040805190810160405280600681526020017f745f7465737400000000000000000000000000000000000000000000000000008152506040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016102109190611c2b565b602060405180830381600087803b15801561022a57600080fd5b505af115801561023e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506102629190810190611816565b93508373ffffffffffffffffffffffffffffffffffffffff166313db93466040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156102c857600080fd5b505af11580156102dc573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061030091908101906117ed565b92508273ffffffffffffffffffffffffffffffffffffffff1663e942b516876040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016103579190611d60565b600060405180830381600087803b15801561037157600080fd5b505af1158015610385573d6000803e3d6000fd5b505050508373ffffffffffffffffffffffffffffffffffffffff16637857d7c96040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156103ed57600080fd5b505af1158015610401573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610425919081019061179b565b91508173ffffffffffffffffffffffffffffffffffffffff1663cd30a1d1896040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040161047c9190611d0b565b600060405180830381600087803b15801561049657600080fd5b505af11580156104aa573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff1663e44594b9886040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016105039190611db5565b600060405180830381600087803b15801561051d57600080fd5b505af1158015610531573d6000803e3d6000fd5b505050508373ffffffffffffffffffffffffffffffffffffffff1663bf2b70a18985856040518463ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040161058e93929190611cad565b602060405180830381600087803b1580156105a857600080fd5b505af11580156105bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506105e0919081019061183f565b90507f8e5890af40fc24a059396aca2f83d6ce41fcef086876548fa4fb8ec27e9d292a816040516106119190611c10565b60405180910390a1809450505050509392505050565b6000806000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f23f63c96040805190810160405280600681526020017f745f7465737400000000000000000000000000000000000000000000000000008152506040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016106d89190611c2b565b602060405180830381600087803b1580156106f257600080fd5b505af1158015610706573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061072a9190810190611816565b92508273ffffffffffffffffffffffffffffffffffffffff16637857d7c96040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561079057600080fd5b505af11580156107a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506107c8919081019061179b565b91508173ffffffffffffffffffffffffffffffffffffffff1663cd30a1d1876040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040161081f9190611d0b565b600060405180830381600087803b15801561083957600080fd5b505af115801561084d573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff1663e44594b9866040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016108a69190611db5565b600060405180830381600087803b1580156108c057600080fd5b505af11580156108d4573d6000803e3d6000fd5b505050508273ffffffffffffffffffffffffffffffffffffffff166328bb211787846040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040161092f929190611c4d565b602060405180830381600087803b15801561094957600080fd5b505af115801561095d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610981919081019061183f565b90507f4b930e280fe29620bdff00c88155d46d6d82a39f45dd5c3ea114dc3157358112816040516109b29190611c10565b60405180910390a180935050505092915050565b6000806000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f23f63c96040805190810160405280600681526020017f745f7465737400000000000000000000000000000000000000000000000000008152506040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401610a779190611c2b565b602060405180830381600087803b158015610a9157600080fd5b505af1158015610aa5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610ac99190810190611816565b92508273ffffffffffffffffffffffffffffffffffffffff166313db93466040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015610b2f57600080fd5b505af1158015610b43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610b6791908101906117ed565b91508173ffffffffffffffffffffffffffffffffffffffff1663e942b516886040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401610bbe9190611d0b565b600060405180830381600087803b158015610bd857600080fd5b505af1158015610bec573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff16632ef8ba74876040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401610c459190611db5565b600060405180830381600087803b158015610c5f57600080fd5b505af1158015610c73573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff1663e942b516866040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401610ccc9190611d60565b600060405180830381600087803b158015610ce657600080fd5b505af1158015610cfa573d6000803e3d6000fd5b505050508273ffffffffffffffffffffffffffffffffffffffff166331afac3688846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401610d55929190611c7d565b602060405180830381600087803b158015610d6f57600080fd5b505af1158015610d83573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610da7919081019061183f565b90507fc57b01fa77f41df77eaab79a0e2623fab2e7ae3e9530d9b1cab225ad65f2b7ce81604051610dd89190611c10565b60405180910390a18093505050509392505050565b6060806060600080600060608060606000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f23f63c96040805190810160405280600681526020017f745f7465737400000000000000000000000000000000000000000000000000008152506040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401610eaa9190611c2b565b602060405180830381600087803b158015610ec457600080fd5b505af1158015610ed8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610efc9190810190611816565b97508773ffffffffffffffffffffffffffffffffffffffff16637857d7c96040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015610f6257600080fd5b505af1158015610f76573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610f9a919081019061179b565b96508773ffffffffffffffffffffffffffffffffffffffff1663e8434e398d896040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401610ff3929190611c4d565b602060405180830381600087803b15801561100d57600080fd5b505af1158015611021573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061104591908101906117c4565b95508573ffffffffffffffffffffffffffffffffffffffff1663949d225d6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156110ab57600080fd5b505af11580156110bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506110e3919081019061183f565b60405190808252806020026020018201604052801561111657816020015b60608152602001906001900390816111015790505b5094508573ffffffffffffffffffffffffffffffffffffffff1663949d225d6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561117d57600080fd5b505af1158015611191573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506111b5919081019061183f565b6040519080825280602002602001820160405280156111e35781602001602082028038833980820191505090505b5093508573ffffffffffffffffffffffffffffffffffffffff1663949d225d6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561124a57600080fd5b505af115801561125e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250611282919081019061183f565b6040519080825280602002602001820160405280156112b557816020015b60608152602001906001900390816112a05790505b509250600091505b8573ffffffffffffffffffffffffffffffffffffffff1663949d225d6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561132157600080fd5b505af1158015611335573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250611359919081019061183f565b82121561165f578573ffffffffffffffffffffffffffffffffffffffff1663846719e0836040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016113b59190611c10565b602060405180830381600087803b1580156113cf57600080fd5b505af11580156113e3573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061140791908101906117ed565b90508073ffffffffffffffffffffffffffffffffffffffff16639c981fcb6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040161145c90611ceb565b600060405180830381600087803b15801561147657600080fd5b505af115801561148a573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506114b391908101906118a9565b85838151811015156114c157fe5b906020019060200201819052508073ffffffffffffffffffffffffffffffffffffffff1663fda69fae6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040161152190611d95565b602060405180830381600087803b15801561153b57600080fd5b505af115801561154f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250611573919081019061183f565b848381518110151561158157fe5b90602001906020020181815250508073ffffffffffffffffffffffffffffffffffffffff16639c981fcb6040518163ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016115e290611d40565b600060405180830381600087803b1580156115fc57600080fd5b505af1158015611610573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525061163991908101906118a9565b838381518110151561164757fe5b906020019060200201819052508160010191506112bd565b8484849a509a509a5050505050505050509193909250565b60006116838251611ec6565b905092915050565b60006116978251611ed8565b905092915050565b60006116ab8251611eea565b905092915050565b60006116bf8251611efc565b905092915050565b60006116d38235611f0e565b905092915050565b60006116e78251611f0e565b905092915050565b600082601f830112151561170257600080fd5b813561171561171082611e10565b611de3565b9150808252602083016020830185838301111561173157600080fd5b61173c838284611f3c565b50505092915050565b600082601f830112151561175857600080fd5b815161176b61176682611e10565b611de3565b9150808252602083016020830185838301111561178757600080fd5b611792838284611f4b565b50505092915050565b6000602082840312156117ad57600080fd5b60006117bb84828501611677565b91505092915050565b6000602082840312156117d657600080fd5b60006117e48482850161168b565b91505092915050565b6000602082840312156117ff57600080fd5b600061180d8482850161169f565b91505092915050565b60006020828403121561182857600080fd5b6000611836848285016116b3565b91505092915050565b60006020828403121561185157600080fd5b600061185f848285016116db565b91505092915050565b60006020828403121561187a57600080fd5b600082013567ffffffffffffffff81111561189457600080fd5b6118a0848285016116ef565b91505092915050565b6000602082840312156118bb57600080fd5b600082015167ffffffffffffffff8111156118d557600080fd5b6118e184828501611745565b91505092915050565b600080604083850312156118fd57600080fd5b600083013567ffffffffffffffff81111561191757600080fd5b611923858286016116ef565b9250506020611934858286016116c7565b9150509250929050565b60008060006060848603121561195357600080fd5b600084013567ffffffffffffffff81111561196d57600080fd5b611979868287016116ef565b935050602061198a868287016116c7565b925050604084013567ffffffffffffffff8111156119a757600080fd5b6119b3868287016116ef565b9150509250925092565b60006119c882611e56565b8084526020840193506119da83611e3c565b60005b82811015611a0c576119f0868351611aa4565b6119f982611e82565b91506020860195506001810190506119dd565b50849250505092915050565b6000611a2382611e61565b80845260208401935083602082028501611a3c85611e49565b60005b84811015611a75578383038852611a57838351611ae9565b9250611a6282611e8f565b9150602088019750600181019050611a3f565b508196508694505050505092915050565b611a8f81611f18565b82525050565b611a9e81611f2a565b82525050565b611aad81611ebc565b82525050565b6000611abe82611e77565b808452611ad2816020860160208601611f4b565b611adb81611f7e565b602085010191505092915050565b6000611af482611e6c565b808452611b08816020860160208601611f4b565b611b1181611f7e565b602085010191505092915050565b6000600482527f6e616d65000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000600982527f6974656d5f6e616d6500000000000000000000000000000000000000000000006020830152604082019050919050565b6000600782527f6974656d5f6964000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006060820190508181036000830152611bde8186611a18565b90508181036020830152611bf281856119bd565b90508181036040830152611c068184611a18565b9050949350505050565b6000602082019050611c256000830184611aa4565b92915050565b60006020820190508181036000830152611c458184611ae9565b905092915050565b60006040820190508181036000830152611c678185611ab3565b9050611c766020830184611a86565b9392505050565b60006040820190508181036000830152611c978185611ab3565b9050611ca66020830184611a95565b9392505050565b60006060820190508181036000830152611cc78186611ab3565b9050611cd66020830185611a95565b611ce36040830184611a86565b949350505050565b60006020820190508181036000830152611d0481611b1f565b9050919050565b60006040820190508181036000830152611d2481611b1f565b90508181036020830152611d388184611ab3565b905092915050565b60006020820190508181036000830152611d5981611b56565b9050919050565b60006040820190508181036000830152611d7981611b56565b90508181036020830152611d8d8184611ab3565b905092915050565b60006020820190508181036000830152611dae81611b8d565b9050919050565b60006040820190508181036000830152611dce81611b8d565b9050611ddd6020830184611aa4565b92915050565b6000604051905081810181811067ffffffffffffffff82111715611e0657600080fd5b8060405250919050565b600067ffffffffffffffff821115611e2757600080fd5b601f19601f8301169050602081019050919050565b6000602082019050919050565b6000602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611ed182611e9c565b9050919050565b6000611ee382611e9c565b9050919050565b6000611ef582611e9c565b9050919050565b6000611f0782611e9c565b9050919050565b6000819050919050565b6000611f2382611e9c565b9050919050565b6000611f3582611e9c565b9050919050565b82818337600083830152505050565b60005b83811015611f69578082015181840152602081019050611f4e565b83811115611f78576000848401525b50505050565b6000601f19601f83011690509190505600a265627a7a723058200414277651c746f53a4cd8d1f6db146a8bbf13994323f708650eba14a4963c016c6578706572696d656e74616cf50037"

// DeployTableTest deploys a new contract, binding an instance of TableTest to it.
func DeployTableTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TableTest, error) {
	parsed, err := abi.JSON(strings.NewReader(TableTestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TableTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TableTest{TableTestCaller: TableTestCaller{contract: contract}, TableTestTransactor: TableTestTransactor{contract: contract}, TableTestFilterer: TableTestFilterer{contract: contract}}, nil
}

func AsyncDeployTableTest(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(TableTestABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(TableTestBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// TableTest is an auto generated Go binding around a Solidity contract.
type TableTest struct {
	TableTestCaller     // Read-only binding to the contract
	TableTestTransactor // Write-only binding to the contract
	TableTestFilterer   // Log filterer for contract events
}

// TableTestCaller is an auto generated read-only Go binding around a Solidity contract.
type TableTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TableTestTransactor is an auto generated write-only Go binding around a Solidity contract.
type TableTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TableTestFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type TableTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TableTestSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type TableTestSession struct {
	Contract     *TableTest        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TableTestCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type TableTestCallerSession struct {
	Contract *TableTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TableTestTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type TableTestTransactorSession struct {
	Contract     *TableTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TableTestRaw is an auto generated low-level Go binding around a Solidity contract.
type TableTestRaw struct {
	Contract *TableTest // Generic contract binding to access the raw methods on
}

// TableTestCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type TableTestCallerRaw struct {
	Contract *TableTestCaller // Generic read-only contract binding to access the raw methods on
}

// TableTestTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type TableTestTransactorRaw struct {
	Contract *TableTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTableTest creates a new instance of TableTest, bound to a specific deployed contract.
func NewTableTest(address common.Address, backend bind.ContractBackend) (*TableTest, error) {
	contract, err := bindTableTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TableTest{TableTestCaller: TableTestCaller{contract: contract}, TableTestTransactor: TableTestTransactor{contract: contract}, TableTestFilterer: TableTestFilterer{contract: contract}}, nil
}

// NewTableTestCaller creates a new read-only instance of TableTest, bound to a specific deployed contract.
func NewTableTestCaller(address common.Address, caller bind.ContractCaller) (*TableTestCaller, error) {
	contract, err := bindTableTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TableTestCaller{contract: contract}, nil
}

// NewTableTestTransactor creates a new write-only instance of TableTest, bound to a specific deployed contract.
func NewTableTestTransactor(address common.Address, transactor bind.ContractTransactor) (*TableTestTransactor, error) {
	contract, err := bindTableTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TableTestTransactor{contract: contract}, nil
}

// NewTableTestFilterer creates a new log filterer instance of TableTest, bound to a specific deployed contract.
func NewTableTestFilterer(address common.Address, filterer bind.ContractFilterer) (*TableTestFilterer, error) {
	contract, err := bindTableTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TableTestFilterer{contract: contract}, nil
}

// bindTableTest binds a generic wrapper to an already deployed contract.
func bindTableTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TableTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TableTest *TableTestRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TableTest.Contract.TableTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TableTest *TableTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _TableTest.Contract.TableTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TableTest *TableTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _TableTest.Contract.TableTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TableTest *TableTestCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TableTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TableTest *TableTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _TableTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TableTest *TableTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _TableTest.Contract.contract.Transact(opts, method, params...)
}

// Select is a free data retrieval call binding the contract method 0xfcd7e3c1.
//
// Solidity: function select(string name) constant returns(string[], int256[], string[])
func (_TableTest *TableTestCaller) Select(opts *bind.CallOpts, name string) ([]string, []*big.Int, []string, error) {
	var (
		ret0 = new([]string)
		ret1 = new([]*big.Int)
		ret2 = new([]string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _TableTest.contract.Call(opts, out, "select", name)
	return *ret0, *ret1, *ret2, err
}

// Select is a free data retrieval call binding the contract method 0xfcd7e3c1.
//
// Solidity: function select(string name) constant returns(string[], int256[], string[])
func (_TableTest *TableTestSession) Select(name string) ([]string, []*big.Int, []string, error) {
	return _TableTest.Contract.Select(&_TableTest.CallOpts, name)
}

// Select is a free data retrieval call binding the contract method 0xfcd7e3c1.
//
// Solidity: function select(string name) constant returns(string[], int256[], string[])
func (_TableTest *TableTestCallerSession) Select(name string) ([]string, []*big.Int, []string, error) {
	return _TableTest.Contract.Select(&_TableTest.CallOpts, name)
}

// Insert is a paid mutator transaction binding the contract method 0xebf3b24f.
//
// Solidity: function insert(string name, int256 item_id, string item_name) returns(int256)
func (_TableTest *TableTestTransactor) Insert(opts *bind.TransactOpts, name string, item_id *big.Int, item_name string) (*types.Transaction, *types.Receipt, error) {
	return _TableTest.contract.Transact(opts, "insert", name, item_id, item_name)
}

func (_TableTest *TableTestTransactor) AsyncInsert(handler func(*types.Receipt, error), opts *bind.TransactOpts, name string, item_id *big.Int, item_name string) (*types.Transaction, error) {
	return _TableTest.contract.AsyncTransact(opts, handler, "insert", name, item_id, item_name)
}

// Insert is a paid mutator transaction binding the contract method 0xebf3b24f.
//
// Solidity: function insert(string name, int256 item_id, string item_name) returns(int256)
func (_TableTest *TableTestSession) Insert(name string, item_id *big.Int, item_name string) (*types.Transaction, *types.Receipt, error) {
	return _TableTest.Contract.Insert(&_TableTest.TransactOpts, name, item_id, item_name)
}

func (_TableTest *TableTestSession) AsyncInsert(handler func(*types.Receipt, error), name string, item_id *big.Int, item_name string) (*types.Transaction, error) {
	return _TableTest.Contract.AsyncInsert(handler, &_TableTest.TransactOpts, name, item_id, item_name)
}

// Insert is a paid mutator transaction binding the contract method 0xebf3b24f.
//
// Solidity: function insert(string name, int256 item_id, string item_name) returns(int256)
func (_TableTest *TableTestTransactorSession) Insert(name string, item_id *big.Int, item_name string) (*types.Transaction, *types.Receipt, error) {
	return _TableTest.Contract.Insert(&_TableTest.TransactOpts, name, item_id, item_name)
}

func (_TableTest *TableTestTransactorSession) AsyncInsert(handler func(*types.Receipt, error), name string, item_id *big.Int, item_name string) (*types.Transaction, error) {
	return _TableTest.Contract.AsyncInsert(handler, &_TableTest.TransactOpts, name, item_id, item_name)
}

// Remove is a paid mutator transaction binding the contract method 0xc4f41ab3.
//
// Solidity: function remove(string name, int256 item_id) returns(int256)
func (_TableTest *TableTestTransactor) Remove(opts *bind.TransactOpts, name string, item_id *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _TableTest.contract.Transact(opts, "remove", name, item_id)
}

func (_TableTest *TableTestTransactor) AsyncRemove(handler func(*types.Receipt, error), opts *bind.TransactOpts, name string, item_id *big.Int) (*types.Transaction, error) {
	return _TableTest.contract.AsyncTransact(opts, handler, "remove", name, item_id)
}

// Remove is a paid mutator transaction binding the contract method 0xc4f41ab3.
//
// Solidity: function remove(string name, int256 item_id) returns(int256)
func (_TableTest *TableTestSession) Remove(name string, item_id *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _TableTest.Contract.Remove(&_TableTest.TransactOpts, name, item_id)
}

func (_TableTest *TableTestSession) AsyncRemove(handler func(*types.Receipt, error), name string, item_id *big.Int) (*types.Transaction, error) {
	return _TableTest.Contract.AsyncRemove(handler, &_TableTest.TransactOpts, name, item_id)
}

// Remove is a paid mutator transaction binding the contract method 0xc4f41ab3.
//
// Solidity: function remove(string name, int256 item_id) returns(int256)
func (_TableTest *TableTestTransactorSession) Remove(name string, item_id *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _TableTest.Contract.Remove(&_TableTest.TransactOpts, name, item_id)
}

func (_TableTest *TableTestTransactorSession) AsyncRemove(handler func(*types.Receipt, error), name string, item_id *big.Int) (*types.Transaction, error) {
	return _TableTest.Contract.AsyncRemove(handler, &_TableTest.TransactOpts, name, item_id)
}

// Update is a paid mutator transaction binding the contract method 0x487a5a10.
//
// Solidity: function update(string name, int256 item_id, string item_name) returns(int256)
func (_TableTest *TableTestTransactor) Update(opts *bind.TransactOpts, name string, item_id *big.Int, item_name string) (*types.Transaction, *types.Receipt, error) {
	return _TableTest.contract.Transact(opts, "update", name, item_id, item_name)
}

func (_TableTest *TableTestTransactor) AsyncUpdate(handler func(*types.Receipt, error), opts *bind.TransactOpts, name string, item_id *big.Int, item_name string) (*types.Transaction, error) {
	return _TableTest.contract.AsyncTransact(opts, handler, "update", name, item_id, item_name)
}

// Update is a paid mutator transaction binding the contract method 0x487a5a10.
//
// Solidity: function update(string name, int256 item_id, string item_name) returns(int256)
func (_TableTest *TableTestSession) Update(name string, item_id *big.Int, item_name string) (*types.Transaction, *types.Receipt, error) {
	return _TableTest.Contract.Update(&_TableTest.TransactOpts, name, item_id, item_name)
}

func (_TableTest *TableTestSession) AsyncUpdate(handler func(*types.Receipt, error), name string, item_id *big.Int, item_name string) (*types.Transaction, error) {
	return _TableTest.Contract.AsyncUpdate(handler, &_TableTest.TransactOpts, name, item_id, item_name)
}

// Update is a paid mutator transaction binding the contract method 0x487a5a10.
//
// Solidity: function update(string name, int256 item_id, string item_name) returns(int256)
func (_TableTest *TableTestTransactorSession) Update(name string, item_id *big.Int, item_name string) (*types.Transaction, *types.Receipt, error) {
	return _TableTest.Contract.Update(&_TableTest.TransactOpts, name, item_id, item_name)
}

func (_TableTest *TableTestTransactorSession) AsyncUpdate(handler func(*types.Receipt, error), name string, item_id *big.Int, item_name string) (*types.Transaction, error) {
	return _TableTest.Contract.AsyncUpdate(handler, &_TableTest.TransactOpts, name, item_id, item_name)
}

// TableTestCreateResultIterator is returned from FilterCreateResult and is used to iterate over the raw logs and unpacked data for CreateResult events raised by the TableTest contract.
type TableTestCreateResultIterator struct {
	Event *TableTestCreateResult // Event containing the contract specifics and raw log

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
func (it *TableTestCreateResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TableTestCreateResult)
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
		it.Event = new(TableTestCreateResult)
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
func (it *TableTestCreateResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TableTestCreateResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TableTestCreateResult represents a CreateResult event raised by the TableTest contract.
type TableTestCreateResult struct {
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCreateResult is a free log retrieval operation binding the contract event 0xb5636cd912a73dcdb5b570dbe331dfa3e6435c93e029e642def2c8e40dacf210.
//
// Solidity: event CreateResult(int256 count)
func (_TableTest *TableTestFilterer) FilterCreateResult(opts *bind.FilterOpts) (*TableTestCreateResultIterator, error) {

	logs, sub, err := _TableTest.contract.FilterLogs(opts, "CreateResult")
	if err != nil {
		return nil, err
	}
	return &TableTestCreateResultIterator{contract: _TableTest.contract, event: "CreateResult", logs: logs, sub: sub}, nil
}

// WatchCreateResult is a free log subscription operation binding the contract event 0xb5636cd912a73dcdb5b570dbe331dfa3e6435c93e029e642def2c8e40dacf210.
//
// Solidity: event CreateResult(int256 count)
func (_TableTest *TableTestFilterer) WatchCreateResult(opts *bind.WatchOpts, sink chan<- *TableTestCreateResult) (event.Subscription, error) {

	logs, sub, err := _TableTest.contract.WatchLogs(opts, "CreateResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TableTestCreateResult)
				if err := _TableTest.contract.UnpackLog(event, "CreateResult", log); err != nil {
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

// ParseCreateResult is a log parse operation binding the contract event 0xb5636cd912a73dcdb5b570dbe331dfa3e6435c93e029e642def2c8e40dacf210.
//
// Solidity: event CreateResult(int256 count)
func (_TableTest *TableTestFilterer) ParseCreateResult(log types.Log) (*TableTestCreateResult, error) {
	event := new(TableTestCreateResult)
	if err := _TableTest.contract.UnpackLog(event, "CreateResult", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TableTestInsertResultIterator is returned from FilterInsertResult and is used to iterate over the raw logs and unpacked data for InsertResult events raised by the TableTest contract.
type TableTestInsertResultIterator struct {
	Event *TableTestInsertResult // Event containing the contract specifics and raw log

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
func (it *TableTestInsertResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TableTestInsertResult)
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
		it.Event = new(TableTestInsertResult)
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
func (it *TableTestInsertResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TableTestInsertResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TableTestInsertResult represents a InsertResult event raised by the TableTest contract.
type TableTestInsertResult struct {
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInsertResult is a free log retrieval operation binding the contract event 0xc57b01fa77f41df77eaab79a0e2623fab2e7ae3e9530d9b1cab225ad65f2b7ce.
//
// Solidity: event InsertResult(int256 count)
func (_TableTest *TableTestFilterer) FilterInsertResult(opts *bind.FilterOpts) (*TableTestInsertResultIterator, error) {

	logs, sub, err := _TableTest.contract.FilterLogs(opts, "InsertResult")
	if err != nil {
		return nil, err
	}
	return &TableTestInsertResultIterator{contract: _TableTest.contract, event: "InsertResult", logs: logs, sub: sub}, nil
}

// WatchInsertResult is a free log subscription operation binding the contract event 0xc57b01fa77f41df77eaab79a0e2623fab2e7ae3e9530d9b1cab225ad65f2b7ce.
//
// Solidity: event InsertResult(int256 count)
func (_TableTest *TableTestFilterer) WatchInsertResult(opts *bind.WatchOpts, sink chan<- *TableTestInsertResult) (event.Subscription, error) {

	logs, sub, err := _TableTest.contract.WatchLogs(opts, "InsertResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TableTestInsertResult)
				if err := _TableTest.contract.UnpackLog(event, "InsertResult", log); err != nil {
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

// ParseInsertResult is a log parse operation binding the contract event 0xc57b01fa77f41df77eaab79a0e2623fab2e7ae3e9530d9b1cab225ad65f2b7ce.
//
// Solidity: event InsertResult(int256 count)
func (_TableTest *TableTestFilterer) ParseInsertResult(log types.Log) (*TableTestInsertResult, error) {
	event := new(TableTestInsertResult)
	if err := _TableTest.contract.UnpackLog(event, "InsertResult", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TableTestRemoveResultIterator is returned from FilterRemoveResult and is used to iterate over the raw logs and unpacked data for RemoveResult events raised by the TableTest contract.
type TableTestRemoveResultIterator struct {
	Event *TableTestRemoveResult // Event containing the contract specifics and raw log

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
func (it *TableTestRemoveResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TableTestRemoveResult)
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
		it.Event = new(TableTestRemoveResult)
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
func (it *TableTestRemoveResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TableTestRemoveResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TableTestRemoveResult represents a RemoveResult event raised by the TableTest contract.
type TableTestRemoveResult struct {
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRemoveResult is a free log retrieval operation binding the contract event 0x4b930e280fe29620bdff00c88155d46d6d82a39f45dd5c3ea114dc3157358112.
//
// Solidity: event RemoveResult(int256 count)
func (_TableTest *TableTestFilterer) FilterRemoveResult(opts *bind.FilterOpts) (*TableTestRemoveResultIterator, error) {

	logs, sub, err := _TableTest.contract.FilterLogs(opts, "RemoveResult")
	if err != nil {
		return nil, err
	}
	return &TableTestRemoveResultIterator{contract: _TableTest.contract, event: "RemoveResult", logs: logs, sub: sub}, nil
}

// WatchRemoveResult is a free log subscription operation binding the contract event 0x4b930e280fe29620bdff00c88155d46d6d82a39f45dd5c3ea114dc3157358112.
//
// Solidity: event RemoveResult(int256 count)
func (_TableTest *TableTestFilterer) WatchRemoveResult(opts *bind.WatchOpts, sink chan<- *TableTestRemoveResult) (event.Subscription, error) {

	logs, sub, err := _TableTest.contract.WatchLogs(opts, "RemoveResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TableTestRemoveResult)
				if err := _TableTest.contract.UnpackLog(event, "RemoveResult", log); err != nil {
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

// ParseRemoveResult is a log parse operation binding the contract event 0x4b930e280fe29620bdff00c88155d46d6d82a39f45dd5c3ea114dc3157358112.
//
// Solidity: event RemoveResult(int256 count)
func (_TableTest *TableTestFilterer) ParseRemoveResult(log types.Log) (*TableTestRemoveResult, error) {
	event := new(TableTestRemoveResult)
	if err := _TableTest.contract.UnpackLog(event, "RemoveResult", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TableTestUpdateResultIterator is returned from FilterUpdateResult and is used to iterate over the raw logs and unpacked data for UpdateResult events raised by the TableTest contract.
type TableTestUpdateResultIterator struct {
	Event *TableTestUpdateResult // Event containing the contract specifics and raw log

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
func (it *TableTestUpdateResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TableTestUpdateResult)
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
		it.Event = new(TableTestUpdateResult)
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
func (it *TableTestUpdateResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TableTestUpdateResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TableTestUpdateResult represents a UpdateResult event raised by the TableTest contract.
type TableTestUpdateResult struct {
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdateResult is a free log retrieval operation binding the contract event 0x8e5890af40fc24a059396aca2f83d6ce41fcef086876548fa4fb8ec27e9d292a.
//
// Solidity: event UpdateResult(int256 count)
func (_TableTest *TableTestFilterer) FilterUpdateResult(opts *bind.FilterOpts) (*TableTestUpdateResultIterator, error) {

	logs, sub, err := _TableTest.contract.FilterLogs(opts, "UpdateResult")
	if err != nil {
		return nil, err
	}
	return &TableTestUpdateResultIterator{contract: _TableTest.contract, event: "UpdateResult", logs: logs, sub: sub}, nil
}

// WatchUpdateResult is a free log subscription operation binding the contract event 0x8e5890af40fc24a059396aca2f83d6ce41fcef086876548fa4fb8ec27e9d292a.
//
// Solidity: event UpdateResult(int256 count)
func (_TableTest *TableTestFilterer) WatchUpdateResult(opts *bind.WatchOpts, sink chan<- *TableTestUpdateResult) (event.Subscription, error) {

	logs, sub, err := _TableTest.contract.WatchLogs(opts, "UpdateResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TableTestUpdateResult)
				if err := _TableTest.contract.UnpackLog(event, "UpdateResult", log); err != nil {
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

// ParseUpdateResult is a log parse operation binding the contract event 0x8e5890af40fc24a059396aca2f83d6ce41fcef086876548fa4fb8ec27e9d292a.
//
// Solidity: event UpdateResult(int256 count)
func (_TableTest *TableTestFilterer) ParseUpdateResult(log types.Log) (*TableTestUpdateResult, error) {
	event := new(TableTestUpdateResult)
	if err := _TableTest.contract.UnpackLog(event, "UpdateResult", log); err != nil {
		return nil, err
	}
	return event, nil
}
