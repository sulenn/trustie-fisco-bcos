package repo

import (
	"encoding/json"
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

	_, strJSON, err := openSourceSession.SelectCommitInfo(commitHash) // call Insert API
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	fmt.Printf("commitHash: %v   \n", strJSON)

	var uploadCommitOption api.UploadCommitOption
	err = json.Unmarshal([]byte(strJSON), &uploadCommitOption)
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	ctx.JSON(http.StatusOK, &structs.ResponseCommit{
		Response: structs.Response{
			Status:  0,
			Message: "query success!",
		},
		UploadCommitOption: uploadCommitOption,
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

	_, strJSON, err := openSourceSession.SelectPushInfo(pushID) // call Insert API
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	fmt.Printf("pushID: %v  \n", strJSON)

	var uploadPushOption api.UploadPushOption
	err = json.Unmarshal([]byte(strJSON), &uploadPushOption)
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	ctx.JSON(http.StatusOK, &structs.ResponsePush{
		Response: structs.Response{
			Status:  0,
			Message: "query success!",
		},
		UploadPushOption: uploadPushOption,
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

// 查询 pull request comment 最新数据
func SelectPullRequestCommentLatestInfo(ctx *macaron.Context, pullRequestCommentID string, logger *log.Logger) {
	if pullRequestCommentID == "" {
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

	_, strJSON, err := openSourceSession.SelectPullRequestCommentInfo(pullRequestCommentID) // call Insert API
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	fmt.Printf("pullRequestID: %v   \n", strJSON)

	var uploadPullRequestCommentOption api.UploadPullReuqestCommentOption
	err = json.Unmarshal([]byte(strJSON), &uploadPullRequestCommentOption)
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	ctx.JSON(http.StatusOK, &structs.ResponsePullRequestComment{
		Response: structs.Response{
			Status:  0,
			Message: "query success!",
		},
		UploadPullReuqestCommentOption: uploadPullRequestCommentOption,
	})
}

// 查询 pull request comment 所有数据
func SelectPullRequestCommentAllInfo(ctx *macaron.Context, pullRequestCommentID string, logger *log.Logger) {
	if pullRequestCommentID == "" {
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

	_, strJSONArr, err := openSourceSession.SelectPullRequestCommentAllInfo(pullRequestCommentID) // call Insert API
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	fmt.Printf("pullRequestID: %v   \n", strJSONArr)

	var uploadPullRequestCommentOption api.UploadPullReuqestCommentOption
	var uploadPullReuqestCommentOptionArr api.UploadPullReuqestCommentOptionArr
	for i := 0; i < len(strJSONArr); i++ {
		err = json.Unmarshal([]byte(strJSONArr[i]), &uploadPullRequestCommentOption)
		if err != nil {
			ctx.JSON(http.StatusOK, api.UnknownErr(err))
			return
		}
		uploadPullReuqestCommentOptionArr = append(uploadPullReuqestCommentOptionArr, uploadPullRequestCommentOption)
	}
	ctx.JSON(http.StatusOK, &structs.ResponsePullRequestCommentArr{
		Response: structs.Response{
			Status:  0,
			Message: "query success!",
		},
		UploadPullReuqestCommentOptionArr: uploadPullReuqestCommentOptionArr,
	})
}

// 查询 issue 最新数据
func SelectIssueLatestInfo(ctx *macaron.Context, issueID string, logger *log.Logger) {
	if issueID == "" {
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

	_, strJSON, err := openSourceSession.SelectIssueInfo(issueID) // call Insert API
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	fmt.Printf("issueID: %v   \n", strJSON)

	var uploadIssueOption api.UploadIssueOption
	err = json.Unmarshal([]byte(strJSON), &uploadIssueOption)
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	ctx.JSON(http.StatusOK, &structs.ResponseIssue{
		Response: structs.Response{
			Status:  0,
			Message: "query success!",
		},
		UploadIssueOption: uploadIssueOption,
	})
}

// 查询 issue 所有数据
func SelectIssueAllInfo(ctx *macaron.Context, issueID string, logger *log.Logger) {
	if issueID == "" {
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

	_, strJSONArr, err := openSourceSession.SelectIssueAllInfo(issueID) // call Insert API
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	fmt.Printf("pullRequestID: %v   \n", strJSONArr)

	var uploadIssueOption api.UploadIssueOption
	var uploadIssueOptionArr api.UploadIssueOptionArr
	for i := 0; i < len(strJSONArr); i++ {
		err = json.Unmarshal([]byte(strJSONArr[i]), &uploadIssueOption)
		if err != nil {
			ctx.JSON(http.StatusOK, api.UnknownErr(err))
			return
		}
		uploadIssueOptionArr = append(uploadIssueOptionArr, uploadIssueOption)
	}
	ctx.JSON(http.StatusOK, &structs.ResponseIssueArr{
		Response: structs.Response{
			Status:  0,
			Message: "query success!",
		},
		UploadIssueOptionArr: uploadIssueOptionArr,
	})
}

// 查询 issue Comment 最新数据
func SelectIssueCommentLatestInfo(ctx *macaron.Context, issueCommentID string, logger *log.Logger) {
	if issueCommentID == "" {
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

	_, strJSON, err := openSourceSession.SelectIssueCommentInfo(issueCommentID) // call Insert API
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	fmt.Printf("issueCommentID: %v   \n", strJSON)

	var uploadIssueCommentOption api.UploadIssueCommentOption
	err = json.Unmarshal([]byte(strJSON), &uploadIssueCommentOption)
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	ctx.JSON(http.StatusOK, &structs.ResponseIssueComment{
		Response: structs.Response{
			Status:  0,
			Message: "query success!",
		},
		UploadIssueCommentOption: uploadIssueCommentOption,
	})
}

// 查询 issue comment 所有数据
func SelectIssueCommentAllInfo(ctx *macaron.Context, issueCommentID string, logger *log.Logger) {
	if issueCommentID == "" {
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

	_, strJSONArr, err := openSourceSession.SelectIssueCommentAllInfo(issueCommentID) // call Insert API
	if err != nil {
		ctx.JSON(http.StatusOK, api.UnknownErr(err))
		return
	}

	fmt.Printf("pullRequestID: %v   \n", strJSONArr)

	var uploadIssueCommentOption api.UploadIssueCommentOption
	var uploadIssueCommentOptionArr api.UploadIssueCommentOptionArr
	for i := 0; i < len(strJSONArr); i++ {
		err = json.Unmarshal([]byte(strJSONArr[i]), &uploadIssueCommentOption)
		if err != nil {
			ctx.JSON(http.StatusOK, api.UnknownErr(err))
			return
		}
		uploadIssueCommentOptionArr = append(uploadIssueCommentOptionArr, uploadIssueCommentOption)
	}
	ctx.JSON(http.StatusOK, &structs.ResponseIssueCommentArr{
		Response: structs.Response{
			Status:  0,
			Message: "query success!",
		},
		UploadIssueCommentOptionArr: uploadIssueCommentOptionArr,
	})
}
