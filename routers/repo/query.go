package repo

import (
	"encoding/json"
	"fmt"
	"log"
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

// 查询 push 数据
func SelectPushInfo(ctx *macaron.Context, pushID string, logger *log.Logger) {
	if pushID == "" {
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

	_, pushNumber, repoID, reponame, ownername, username, branch, commitShas, time, err := openSourceSession.SelectPushInfo(pushID) // call Insert API
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	fmt.Printf("pushID: %v, pushNumber: %v, repoID: %v, reponame: %v, ownername: %v, username: %v, branch: %v, commitShas: %v, commitShas: %v   \n",
		pushID, pushNumber, repoID, reponame, ownername, username, branch, commitShas, time)

	ctx.JSON(http.StatusOK, &structs.ResponsePush{
		Response: structs.Response{
			Status:  0,
			Message: "query success!",
		},
		UploadPushOption: structs.UploadPushOption{
			PushID:     pushID,
			PushNumber: pushNumber.Uint64(),
			RepoID:     repoID,
			Reponame:   reponame,
			Ownername:  ownername,
			Username:   username,
			Branch:     branch,
			CommitShas: strings.Split(commitShas, "-"),
			Time:       time,
		},
	})
}

// 查询 pull request 最新数据
func SelectPullRequestLatestInfo(ctx *macaron.Context, pullRequestID string, logger *log.Logger) {
	if pullRequestID == "" {
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

	_, strJSON, err := openSourceSession.SelectPullRequestInfo(pullRequestID) // call Insert API
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	fmt.Printf("pullRequestID: %v   \n", strJSON)

	var uploadPullRequestOption api.UploadPullReuqestOption
	err = json.Unmarshal([]byte(strJSON), &uploadPullRequestOption)
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	ctx.JSON(http.StatusOK, &structs.ResponsePullRequest{
		Response: structs.Response{
			Status:  0,
			Message: "query success!",
		},
		UploadPullReuqestOption: uploadPullRequestOption,
	})
}

// 查询 pull request 所有数据
func SelectPullRequestAllInfo(ctx *macaron.Context, pullRequestID string, logger *log.Logger) {
	if pullRequestID == "" {
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

	_, strJSONArr, err := openSourceSession.SelectPullRequestAllInfo(pullRequestID) // call Insert API
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	fmt.Printf("pullRequestID: %v   \n", strJSONArr)

	var uploadPullRequestOption api.UploadPullReuqestOption
	var uploadPullReuqestOptionArr api.UploadPullReuqestOptionArr
	for i := 0; i < len(strJSONArr); i++ {
		err = json.Unmarshal([]byte(strJSONArr[i]), &uploadPullRequestOption)
		if err != nil {
			ctx.JSON(http.StatusOK, api.UnknownErr(err))
			return
		}
		uploadPullReuqestOptionArr = append(uploadPullReuqestOptionArr, uploadPullRequestOption)
	}
	ctx.JSON(http.StatusOK, &structs.ResponsePullRequestArr{
		Response: structs.Response{
			Status:  0,
			Message: "query success!",
		},
		UploadPullReuqestOptionArr: uploadPullReuqestOptionArr,
	})
}
