package repo

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sulenn/trustie-fisco-bcos/contract"
	"github.com/sulenn/trustie-fisco-bcos/contract/opensource"
	"github.com/sulenn/trustie-fisco-bcos/modules/structs"
	api "github.com/sulenn/trustie-fisco-bcos/modules/structs"

	"gitea.com/macaron/macaron"
)

var ()

// 上传 commit 数据
func UploadCommitInfo(ctx *macaron.Context, opt api.UploadCommitOption, logger *log.Logger) {
	commitHash := opt.CommitHash
	repoID := opt.RepoID
	author := opt.Author
	email := opt.Email
	time := opt.Time
	content := opt.Content
	commitDiff := opt.CommitDiff

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

	tx, receipt, err := openSourceSession.AddCommitData(commitHash, repoID, author, email, time, content, commitDiff) // call Insert API
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}
	logger.Printf("tx sent: %s\n", tx.Hash().Hex())
	// insertedLines, err := strconv.Atoi(receipt.Output[2:])
	code, err := parseOutput(opensource.OpenSourceABI, "addCommitData", receipt)
	if err != nil {
		logger.Panic("error when transfer string to int: ", err)
	}
	if code.Int64() > 0 {
		logger.Printf("inserted lines: %v\n", code)
		ctx.JSON(http.StatusOK, structs.ResCommitUploadSucc)
		return
	}
	ctx.JSON(http.StatusOK, structs.ResCommitUploadUnsucc)
}

// 上传 push 数据
func UploadPushInfo(ctx *macaron.Context, opt api.UploadPushOption, logger *log.Logger) {
	pushID := opt.PushID
	pushNumber := opt.PushNumber
	repoID := opt.RepoID
	reponame := opt.Reponame
	ownername := opt.Ownername
	username := opt.Username
	branch := opt.Branch
	commitShas := strings.Join(opt.CommitShas, "-") // 拼接字符串
	time := opt.Time

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

	tx, receipt, err := openSourceSession.AddPushData(pushID, big.NewInt(int64(pushNumber)), repoID, reponame, ownername, username, branch, commitShas, time) // call Insert API
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}
	logger.Printf("tx sent: %s\n", tx.Hash().Hex())
	// insertedLines, err := strconv.Atoi(receipt.Output[2:])
	code, err := parseOutput(opensource.OpenSourceABI, "addPushData", receipt)
	if err != nil {
		logger.Panic("error when transfer string to int: ", err)
	}
	if code.Int64() > 0 {
		logger.Printf("inserted lines: %v\n", code)
		ctx.JSON(http.StatusOK, structs.ResPushUploadSucc)
		return
	}
	ctx.JSON(http.StatusOK, structs.ResPushUploadUnsucc)
}

// 上传 pull request 数据
func UploadPullRequestInfo(ctx *macaron.Context, opt api.UploadPullReuqestOption, logger *log.Logger) {
	pullRequestID := opt.PullRequestID
	bytesJson, err := json.Marshal(opt)
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
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

	tx, receipt, err := openSourceSession.AddPullRequestData(pullRequestID, string(bytesJson)) // call Insert API
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}
	logger.Printf("tx sent: %s\n", tx.Hash().Hex())
	// insertedLines, err := strconv.Atoi(receipt.Output[2:])
	code, err := parseOutput(opensource.OpenSourceABI, "addPullRequestData", receipt)
	if err != nil {
		logger.Panic("error when transfer string to int: ", err)
	}
	if code.Int64() > 0 {
		logger.Printf("inserted lines: %v\n", code)
		ctx.JSON(http.StatusOK, structs.ResPullRequestUploadSucc)
		return
	}
	ctx.JSON(http.StatusOK, structs.ResPullRequestUploadUnsucc)
}

// 上传 pull request comment 数据
func UploadPullRequestCommentInfo(ctx *macaron.Context, opt api.UploadPullReuqestCommentOption, logger *log.Logger) {
	pullRequestCommentID := opt.PullRequestCommentID
	bytesJson, err := json.Marshal(opt)
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
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

	tx, receipt, err := openSourceSession.AddPullRequestCommentData(pullRequestCommentID, string(bytesJson)) // call Insert API
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}
	logger.Printf("tx sent: %s\n", tx.Hash().Hex())
	// insertedLines, err := strconv.Atoi(receipt.Output[2:])
	code, err := parseOutput(opensource.OpenSourceABI, "addPullRequestCommentData", receipt)
	if err != nil {
		logger.Panic("error when transfer string to int: ", err)
	}
	if code.Int64() > 0 {
		logger.Printf("inserted lines: %v\n", code)
		ctx.JSON(http.StatusOK, structs.ResPullRequestCommentUploadSucc)
		return
	}
	ctx.JSON(http.StatusOK, structs.ResPullRequestCommentUploadUnsucc)
}

// 上传 issue数据
func UploadIssueInfo(ctx *macaron.Context, opt api.UploadIssueOption, logger *log.Logger) {
	issueID := opt.IssueID
	bytesJson, err := json.Marshal(opt)
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
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

	tx, receipt, err := openSourceSession.AddIssueData(issueID, string(bytesJson)) // call Insert API
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}
	logger.Printf("tx sent: %s\n", tx.Hash().Hex())
	// insertedLines, err := strconv.Atoi(receipt.Output[2:])
	code, err := parseOutput(opensource.OpenSourceABI, "addIssueData", receipt)
	if err != nil {
		logger.Panic("error when transfer string to int: ", err)
	}
	if code.Int64() > 0 {
		logger.Printf("inserted lines: %v\n", code)
		ctx.JSON(http.StatusOK, structs.ResIssueUploadSucc)
		return
	}
	ctx.JSON(http.StatusOK, structs.ResIssueUploadUnsucc)
}

// 上传 issue comment 数据
func UploadIssueCommentInfo(ctx *macaron.Context, opt api.UploadIssueCommentOption, logger *log.Logger) {
	issueCommentID := opt.IssueCommentID
	bytesJson, err := json.Marshal(opt)
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
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

	tx, receipt, err := openSourceSession.AddIssueCommentData(issueCommentID, string(bytesJson)) // call Insert API
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}
	logger.Printf("tx sent: %s\n", tx.Hash().Hex())
	// insertedLines, err := strconv.Atoi(receipt.Output[2:])
	code, err := parseOutput(opensource.OpenSourceABI, "addIssueCommentData", receipt)
	if err != nil {
		logger.Panic("error when transfer string to int: ", err)
	}
	if code.Int64() > 0 {
		logger.Printf("inserted lines: %v\n", code)
		ctx.JSON(http.StatusOK, structs.ResIssueCommentUploadSucc)
		return
	}
	ctx.JSON(http.StatusOK, structs.ResIssueCommentUploadUnsucc)
}
