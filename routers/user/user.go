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
