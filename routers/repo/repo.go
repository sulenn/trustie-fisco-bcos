package repo

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

var ()

// Create is to create a repository in blockchain ledger
func Create(ctx *macaron.Context, opt api.CreateRepoOption, logger *log.Logger) { // 创建项目
	tokenName := opt.TokenName
	owner := opt.Owner
	totalSupply := big.NewInt(int64(opt.TotalSupply))
	token_balance := opt.TokenBalance

	var usernameArr = []string{}
	var balanceArr = []*big.Int{}
	var curSupply int64 = 0
	for i := 0; i < len(token_balance); i++ {
		username, ok := token_balance[i][0].(string)
		if ok != true {
			ctx.JSON(http.StatusOK, structs.TypeTransferErr)
			return
		}
		usernameArr = append(usernameArr, username)
		balance, ok := token_balance[i][1].(float64)
		curSupply += int64(balance)
		if ok != true {
			ctx.JSON(http.StatusOK, structs.TypeTransferErr)
			return
		}
		balanceArr = append(balanceArr, big.NewInt(int64(balance)))
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

	tx, receipt, err := openSourceSession.CreateRepo(tokenName, owner, totalSupply, big.NewInt(curSupply), usernameArr, balanceArr) // call Insert API
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}
	logger.Printf("tx sent: %s\n", tx.Hash().Hex())
	// insertedLines, err := strconv.Atoi(receipt.Output[2:])
	code, err := parseOutput(opensource.OpenSourceABI, "createRepo", receipt)
	if err != nil {
		logger.Panic("error when transfer string to int: ", err)
	}
	if code.Int64() > 0 {
		logger.Printf("inserted lines: %v\n", code)
		ctx.JSON(http.StatusOK, structs.ResRepoSucc)
		return
	} else if code.Int64() == contract.RepoExisted {
		logger.Printf("error code: %v, message: %v\n", code, structs.ResRepoExisted)
		ctx.JSON(http.StatusOK, structs.ResRepoExisted)
		return
	} else if code.Int64() == contract.BalanceOverFlow {
		logger.Printf("error code: %v, message: %v\n", code, structs.ResRepoBalanceOverFlow)
		ctx.JSON(http.StatusOK, structs.ResRepoBalanceOverFlow)
		return
	}
	ctx.JSON(http.StatusOK, structs.ResRepoUnsucc)
}

func QueryBasic(ctx *macaron.Context, tokenName string, logger *log.Logger) {
	if tokenName == "" {
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

	_, owner, totalSupply, curSupply, err := openSourceSession.SelectRepoBasicInfo(tokenName) // call Insert API
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	fmt.Printf("tokenName: %v, owner: %v, totalSupply: %v, curSupply: %v  \n", tokenName, owner, totalSupply, curSupply)

	ctx.JSON(http.StatusOK, &structs.ReponseRepo{
		Response: structs.Response{
			Status:  0,
			Message: "query success!",
		},
		Repo: structs.Repo{
			TokenName:   tokenName,
			Owner:       owner,
			TotalSupply: totalSupply.Uint64(),
			CurSupply:   curSupply.Uint64(),
		},
	})
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

// // Transfer is to transfer amount from one person to another
// func Transfer(ctx *macaron.Context, opt api.TransferTokenOption) { // 转账
// 	tokenSymbol := opt.TokenSymbol
// 	from := opt.From
// 	to := opt.To
// 	amount := fmt.Sprintf("%v", opt.Amount)

// 	chainArgsToJsonString, err := format.ArgsJson("transfer", tokenSymbol, from, to, amount)
// 	if err != nil {
// 		ctx.Error(http.StatusInternalServerError, fmt.Sprintf("trasfer:%v", err))
// 		return
// 	}

// 	err = chaincode.InvokeOther(*chainArgsToJsonString)

// 	if err != nil {
// 		ctx.Error(http.StatusInternalServerError, fmt.Sprintf("trasfer:%v", err))
// 		return
// 	}

// 	userBalanceList := structs.UserBalanceList{}

// 	payerAmout, err := queryUserBalance(tokenSymbol, from) // 付款人余额信息
// 	if err != nil {
// 		ctx.Error(http.StatusInternalServerError, fmt.Sprintf("It is failed when query payer amount: %v", err))
// 		return
// 	}

// 	payer := &structs.UserBalance{
// 		Username:    from,
// 		TokenSymbol: tokenSymbol,
// 		Balance:     *payerAmout,
// 	}
// 	userBalanceList = append(userBalanceList, payer)

// 	payeeAmout, err := queryUserBalance(tokenSymbol, to) // 收款人余额信息
// 	if err != nil {
// 		ctx.Error(http.StatusInternalServerError, fmt.Sprintf("It is failed when query payee amount: %v", err))
// 		return
// 	}

// 	payee := &structs.UserBalance{
// 		Username:    to,
// 		TokenSymbol: tokenSymbol,
// 		Balance:     *payeeAmout,
// 	}
// 	userBalanceList = append(userBalanceList, payee)

// 	ctx.JSON(http.StatusOK, userBalanceList)
// }

// // Add is to help user add amount
// func Add(ctx *macaron.Context, opt api.AddTokenOption) { // 增加用户余额
// 	tokenSymbol := opt.TokenSymbol
// 	username := opt.Username
// 	amount := fmt.Sprintf("%v", opt.Amount)

// 	err := add(tokenSymbol, username, amount)
// 	if err != nil {
// 		ctx.Error(http.StatusInternalServerError, fmt.Sprintf("add: %v", err))
// 		return
// 	}

// 	userAmout, err := queryUserBalance(tokenSymbol, username) // 用户余额信息
// 	if err != nil {
// 		ctx.Error(http.StatusInternalServerError, fmt.Sprintf("It is failed when query  user amount: %v", err))
// 		return
// 	}

// 	user := &structs.UserBalance{
// 		Username:    username,
// 		TokenSymbol: tokenSymbol,
// 		Balance:     *userAmout,
// 	}

// 	ctx.JSON(http.StatusOK, user)
// }

// func add(reponame string, username string, amount string) error { // 用户添加余额
// 	chainArgsToJsonString, err := format.ArgsJson("add", reponame, username, amount)
// 	if err != nil {
// 		return err
// 	}

// 	err = chaincode.InvokeOther(*chainArgsToJsonString) // 调用 chaincode

// 	return err
// }

// // Minus is to help user minus amount
// func Minus(ctx *macaron.Context, opt api.MinusTokenOption) { // 减少用户余额
// 	tokenSymbol := opt.TokenSymbol
// 	username := opt.Username
// 	amount := fmt.Sprintf("%v", opt.Amount)

// 	err := minus(tokenSymbol, username, amount)

// 	if err != nil {
// 		ctx.Error(http.StatusInternalServerError, fmt.Sprintf("minus: %v", err))
// 		return
// 	}

// 	userAmout, err := queryUserBalance(tokenSymbol, username) // 用户余额信息
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, err)
// 		return
// 	}

// 	user := &structs.UserBalance{
// 		Username:    username,
// 		TokenSymbol: tokenSymbol,
// 		Balance:     *userAmout,
// 	}

// 	ctx.JSON(http.StatusOK, user)
// }

// func minus(reponame string, username string, amount string) error {
// 	chainArgsToJsonString, err := format.ArgsJson("minus", reponame, username, amount)
// 	if err != nil {
// 		return err
// 	}

// 	err = chaincode.InvokeOther(*chainArgsToJsonString) // 调用 chaincode

// 	return err
// }

// // QueryUserBalance is to query user balance by reponame and username
// func QueryUserBalance(ctx *macaron.Context) { // 查询余额
// 	reponame := ctx.Query("reponame")
// 	if reponame == "" {
// 		ctx.JSON(http.StatusPreconditionFailed, "reponame is not nil")
// 		return
// 	}
// 	username := ctx.Query("username")
// 	if username == "" {
// 		ctx.JSON(http.StatusPreconditionFailed, "username is not nil")
// 		return
// 	}

// 	amount, err := queryUserBalance(reponame, username)
// 	if err != nil {
// 		ctx.Error(http.StatusInternalServerError, fmt.Sprintf("queryUserBalance: %v", err))
// 		return
// 	}

// 	userBalance := &api.UserBalance{
// 		Username:    username,
// 		TokenSymbol: reponame,
// 		Balance:     *amount,
// 	}
// 	ctx.JSON(http.StatusOK, userBalance)
// }

// func queryUserBalance(reponame string, username string) (*uint64, error) { // 查询单次余额
// 	chainArgsToJsonString, err := format.ArgsJson("userBalance", reponame, username)
// 	if err != nil {
// 		return nil, err
// 	}

// 	balance, err := chaincode.InvokeBalance(*chainArgsToJsonString)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return balance, err
// }

// // QueryRepoAmount is to query repository amount by reponame
// func QueryRepoAmount(ctx *macaron.Context) { // 查询项目总金额
// 	reponame := ctx.Query("reponame")
// 	if reponame == "" {
// 		ctx.Error(http.StatusPreconditionFailed, "reponame is not nil")
// 		return
// 	}
// 	username := ctx.Query("username")

// 	amount, err := queryRepoAmount(reponame)
// 	if err != nil {
// 		ctx.Error(http.StatusInternalServerError, fmt.Sprintf("queryRepoAmount: %v", err))
// 		return
// 	}

// 	userBalance := &api.Repo{
// 		Owner:       username,
// 		TokenSymbol: reponame,
// 		TotalSupply: *amount,
// 	}
// 	ctx.JSON(http.StatusOK, userBalance)
// }

// func queryRepoAmount(reponame string) (*uint64, error) { // 查询项目总金额
// 	chainArgsToJsonString, err := format.ArgsJson("repoAmount", reponame)
// 	if err != nil {
// 		return nil, err
// 	}

// 	balance, err := chaincode.InvokeBalance(*chainArgsToJsonString)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return balance, err
// }

// // Upload is to record commit info to blockchain
// func Upload(ctx *macaron.Context, opt api.UploadCommitOption) { // 记录 commit 信息
// 	commitString, err := json.Marshal(opt)
// 	if err != nil {
// 		ctx.Error(http.StatusInternalServerError, fmt.Sprintf("upload: %v", err))
// 	}

// 	chainArgsToJsonString, err := format.ArgsJson("commit", string(commitString)) // 记录 commit 信息
// 	if err != nil {
// 		ctx.Error(http.StatusInternalServerError, fmt.Sprintf("upload: %v", err))
// 	}

// 	err = chaincode.InvokeOther(*chainArgsToJsonString) // 调用 chaincode
// 	if err != nil {
// 		ctx.Error(http.StatusInternalServerError, fmt.Sprintf("upload: %v", err))
// 	}

// 	err = add(opt.Reponame, opt.Username, strconv.Itoa(opt.ModificationLines)) // 为用户加钱
// 	if err != nil {
// 		ctx.Error(http.StatusInternalServerError, fmt.Sprintf("upload: %v", err))
// 	}

// 	userAmout, err := queryUserBalance(opt.Reponame, opt.Username) // 用户余额信息
// 	if err != nil {
// 		ctx.Error(http.StatusInternalServerError, fmt.Sprintf("It is failed when query  user amount: %v", err))
// 		return
// 	}

// 	user := &structs.UserBalance{
// 		Username:    opt.Username,
// 		TokenSymbol: opt.Reponame,
// 		Balance:     *userAmout,
// 	}

// 	ctx.JSON(http.StatusOK, user)

// }
