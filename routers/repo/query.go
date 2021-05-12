package repo

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

var ()

// 查询 commit 数据
func SelectCommitInfo(ctx *macaron.Context, commitHash string, logger *log.Logger) {
	if commitHash == "" {
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

	_, repoID, author, email, time, message, commitDiff, err := openSourceSession.SelectCommitInfo(commitHash) // call Insert API
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	fmt.Printf("commitHash: %v, repoID: %v, author: %v, email: %v, curStimeupply: %v, message: %v, commitDiff: %v  \n", commitHash, repoID, author, email, time, message, commitDiff)

	ctx.JSON(http.StatusOK, &structs.ResponseCommit{
		Response: structs.Response{
			Status:  0,
			Message: "query success!",
		},
		UploadCommitOption: structs.UploadCommitOption{
			CommitHash: commitHash,
			RepoID:     repoID,
			Author:     author,
			Email:      email,
			Time:       time,
			Content:    message,
			CommitDiff: commitDiff,
		},
	})
}
