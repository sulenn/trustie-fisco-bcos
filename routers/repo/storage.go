package repo

import (
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
