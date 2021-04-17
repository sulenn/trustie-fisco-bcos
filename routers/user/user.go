package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sulenn/trustie-fisco-bcos/contract"
	"github.com/sulenn/trustie-fisco-bcos/contract/opensource"
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
		logger.Panic(err)
		return
	}
	config := &configs[0]

	client, err := client.Dial(config)
	if err != nil {
		logger.Panic(err)
		return
	}

	// load the contract
	contractAddress := common.HexToAddress(contract.ContractAddress)
	instance, err := opensource.NewOpenSource(contractAddress, client)
	if err != nil {
		logger.Panic(err)
		return
	}

	openSourceSession := &opensource.OpenSourceSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	usernames, tokenNames, balances, err := openSourceSession.SelectUserAllBalance(username) // call select API
	if err != nil {
		logger.Panic(err)
		return
	}
	var userBalancelist api.UserBalanceList
	for i := 0; i < len(usernames); i++ {
		userBalancelist = append(userBalancelist, &api.UserBalance{User: api.User{Username: usernames[i], TokenName: tokenNames[i]}, Balance: balances[i].Uint64()})
		fmt.Printf("username: %v, token_name: %v, balance: %v \n", usernames[i], tokenNames[i], balances[i])
	}
	ctx.JSON(http.StatusOK, userBalancelist)
}

// select user balance
func SelectUserBalance(ctx *macaron.Context, username string, tokenName string, logger *log.Logger) {
	if username == "" || tokenName == "" {
		ctx.JSON(http.StatusOK, api.StringEmpty)
		return
	}
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		logger.Panic(err)
		return
	}
	config := &configs[0]

	client, err := client.Dial(config)
	if err != nil {
		logger.Panic(err)
		return
	}

	// load the contract
	contractAddress := common.HexToAddress(contract.ContractAddress)
	instance, err := opensource.NewOpenSource(contractAddress, client)
	if err != nil {
		logger.Panic(err)
		return
	}

	openSourceSession := &opensource.OpenSourceSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	usernames, tokenNames, balances, err := openSourceSession.SelectUserBalance(username, tokenName) // call select API
	if err != nil {
		logger.Panic(err)
		return
	}
	var userBalancelist api.UserBalanceList
	for i := 0; i < len(usernames); i++ {
		userBalancelist = append(userBalancelist, &api.UserBalance{User: api.User{Username: usernames[i], TokenName: tokenNames[i]}, Balance: balances[i].Uint64()})
		fmt.Printf("username: %v, token_name: %v, balance: %v \n", usernames[i], tokenNames[i], balances[i])
	}
	ctx.JSON(http.StatusOK, userBalancelist)
}
