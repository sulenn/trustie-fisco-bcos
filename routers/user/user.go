package user

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sulenn/trustie-fisco-bcos/contract"
	"github.com/sulenn/trustie-fisco-bcos/contract/opensource"
	"github.com/sulenn/trustie-fisco-bcos/modules/structs"
	api "github.com/sulenn/trustie-fisco-bcos/modules/structs"

	"gitea.com/macaron/macaron"
)

// // 用户注册
// func Register(ctx *macaron.Context, opt api.CreateUserOption) {
// 	var user = &api.User{opt.Username, opt.Orgname}
// 	ctx.JSON(200, user)
// }

// select user all balance
func SelectUserAllBalance(ctx *macaron.Context, username string, logger *log.Logger) {
	if username == "" {
		ctx.JSON(http.StatusOK, api.StringEmpty)
		return
	}
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}
	config := &configs[0]

	client, err := client.Dial(config)
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	// load the contract
	contractAddress := common.HexToAddress(contract.ContractAddress)
	instance, err := opensource.NewOpenSource(contractAddress, client)
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	openSourceSession := &opensource.OpenSourceSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	usernames, tokenNames, balances, err := openSourceSession.SelectUserAllBalance(username) // call select API
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}
	var userBalancelist api.UserBalanceList
	for i := 0; i < len(usernames); i++ {
		userBalancelist = append(userBalancelist, &api.UserBalance{User: api.User{Username: usernames[i], TokenName: tokenNames[i]}, Balance: balances[i].Uint64()})
		fmt.Printf("username: %v, token_name: %v, balance: %v \n", usernames[i], tokenNames[i], balances[i])
	}
	ctx.JSON(http.StatusOK, &structs.ResponseUserBalanceList{
		UserBalanceList: userBalancelist,
		Response: structs.Response{
			Status:  0,
			Message: "query success!",
		},
	})
}

// select user balance
func SelectUserBalance(ctx *macaron.Context, username string, tokenName string, logger *log.Logger) {
	if username == "" || tokenName == "" {
		ctx.JSON(http.StatusOK, api.StringEmpty)
		return
	}
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}
	config := &configs[0]

	client, err := client.Dial(config)
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	// load the contract
	contractAddress := common.HexToAddress(contract.ContractAddress)
	instance, err := opensource.NewOpenSource(contractAddress, client)
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	openSourceSession := &opensource.OpenSourceSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	usernames, tokenNames, balances, err := openSourceSession.SelectUserBalance(username, tokenName) // call select API
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}
	if len(usernames) == 0 {
		ctx.JSON(http.StatusOK, api.ResUserNotExisted)
		return
	}
	var userBalance = api.UserBalance{
		User: api.User{
			Username:  usernames[0],
			TokenName: tokenNames[0],
		},
		Balance: balances[0].Uint64(),
	}
	fmt.Printf("username: %v, token_name: %v, balance: %v \n", usernames[0], tokenNames[0], balances[0])

	ctx.JSON(http.StatusOK, &structs.ResponseUserBalance{
		UserBalance: userBalance,
		Response: structs.Response{
			Status:  0,
			Message: "query success!",
		},
	})
}

func AddUserAmount(ctx *macaron.Context, username string, tokenName string, amount uint64, logger *log.Logger) {
	if username == "" || tokenName == "" {
		ctx.JSON(http.StatusOK, api.StringEmpty)
		return
	}
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}
	config := &configs[0]

	client, err := client.Dial(config)
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	// load the contract
	contractAddress := common.HexToAddress(contract.ContractAddress)
	instance, err := opensource.NewOpenSource(contractAddress, client)
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	openSourceSession := &opensource.OpenSourceSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
	tx, receipt, err := openSourceSession.AddUserBalance(username, tokenName, big.NewInt(int64(amount)))

	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}
	logger.Printf("tx sent: %s\n", tx.Hash().Hex())
	code, err := parseOutput(opensource.OpenSourceABI, "addUserBalance", receipt)
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}
	if code.Int64() > 0 {
		logger.Printf("inserted lines: %v\n", code)
		ctx.JSON(http.StatusOK, structs.ResUserAddAmountSucc)
		return
	} else if code.Int64() == contract.RepoUnexisted {
		logger.Printf("error code: %v, message: %v\n", code, structs.ResRepoNotExisted)
		ctx.JSON(http.StatusOK, structs.ResRepoNotExisted)
		return
	} else if code.Int64() == contract.RepoCurBalanceUpOverFlow {
		logger.Printf("error code: %v, message: %v\n", code, structs.ResRepoCurBalanceUpOverFlow)
		ctx.JSON(http.StatusOK, structs.ResRepoCurBalanceUpOverFlow)
		return
	}
	ctx.JSON(http.StatusOK, structs.ResUserAddAmountUnsucc)
}

func MinusUserAmount(ctx *macaron.Context, username string, tokenName string, amount uint64, logger *log.Logger) {
	if username == "" || tokenName == "" {
		ctx.JSON(http.StatusOK, api.StringEmpty)
		return
	}
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}
	config := &configs[0]

	client, err := client.Dial(config)
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	// load the contract
	contractAddress := common.HexToAddress(contract.ContractAddress)
	instance, err := opensource.NewOpenSource(contractAddress, client)
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	openSourceSession := &opensource.OpenSourceSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
	tx, receipt, err := openSourceSession.MinusUserBalance(username, tokenName, big.NewInt(int64(amount)))

	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}
	logger.Printf("tx sent: %s\n", tx.Hash().Hex())
	code, err := parseOutput(opensource.OpenSourceABI, "minusUserBalance", receipt)
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}
	if code.Int64() > 0 {
		logger.Printf("inserted lines: %v\n", code)
		ctx.JSON(http.StatusOK, structs.ResUserMinusAmountSucc)
		return
	} else if code.Int64() == contract.RepoUnexisted {
		logger.Printf("error code: %v, message: %v\n", code, structs.ResRepoNotExisted)
		ctx.JSON(http.StatusOK, structs.ResRepoNotExisted)
		return
	} else if code.Int64() == contract.UserBblanceNotEnough {
		logger.Printf("error code: %v, message: %v\n", code, structs.ResUserBalanceNotEnough)
		ctx.JSON(http.StatusOK, structs.ResUserBalanceNotEnough)
		return
	}
	ctx.JSON(http.StatusOK, structs.ResUserMinusAmountUnsucc)
}

func TransferAmount(ctx *macaron.Context, payer string, payee string, tokenName string, amount uint64, logger *log.Logger) {
	if payer == "" || payee == "" || tokenName == "" {
		ctx.JSON(http.StatusOK, api.StringEmpty)
		return
	}
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}
	config := &configs[0]

	client, err := client.Dial(config)
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	// load the contract
	contractAddress := common.HexToAddress(contract.ContractAddress)
	instance, err := opensource.NewOpenSource(contractAddress, client)
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	openSourceSession := &opensource.OpenSourceSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
	tx, receipt, err := openSourceSession.TransferUserBalance(payer, payee, tokenName, big.NewInt(int64(amount)))

	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}
	logger.Printf("tx sent: %s\n", tx.Hash().Hex())
	code, err := parseOutput(opensource.OpenSourceABI, "transferUserBalance", receipt)
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}
	if code.Int64() > 0 {
		logger.Printf("inserted lines: %v\n", code)
		ctx.JSON(http.StatusOK, structs.ResUserTransferAmountSucc)
		return
	} else if code.Int64() == contract.RepoUnexisted {
		logger.Printf("error code: %v, message: %v\n", code, structs.ResRepoNotExisted)
		ctx.JSON(http.StatusOK, structs.ResRepoNotExisted)
		return
	} else if code.Int64() == contract.UserBblanceNotEnough {
		logger.Printf("error code: %v, message: %v\n", code, structs.ResUserBalanceNotEnough)
		ctx.JSON(http.StatusOK, structs.ResUserBalanceNotEnough)
		return
	}
	ctx.JSON(http.StatusOK, structs.ResUserTransferAmountUnsucc)
}

func parseOutput(abiStr, name string, receipt *types.Receipt) (*big.Int, error) {
	parsed, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		fmt.Printf("parse ABI failed, err: %v", err)
	}
	var ret *big.Int
	b, err := hex.DecodeString(receipt.Output[2:])
	if err != nil {
		return nil, fmt.Errorf("decode receipt.Output[2:] failed, err: %v", err)
	}
	err = parsed.Unpack(&ret, name, b)
	if err != nil {
		return nil, fmt.Errorf("unpack %v failed, err: %v", name, err)
	}
	return ret, nil
}
