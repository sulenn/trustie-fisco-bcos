package structs

import "fmt"

type Response struct {
	Status  uint16 `json:"status"`
	Message string `json:"message"`
}

var (
	ResRepoSucc                   = &Response{Status: 0, Message: "create repo successfully!"}                               // create repo successfully!
	ResRepoExisted                = &Response{Status: 1, Message: "repo existed!"}                                           // repo existed!
	ResRepoNotExisted             = &Response{Status: 5, Message: "repo not existed!"}                                       // repo not existed!
	ResRepoBalanceOverFlow        = &Response{Status: 2, Message: "username balance overflow when create repo!"}             // username balance overflow when create repo!
	ResRepoUnsucc                 = &Response{Status: 3, Message: "create repo unsuccessfully!"}                             // create repo unsuccessfully!
	ResRepoCurBalanceUpOverFlow   = &Response{Status: 4, Message: "repo current balance up overflow when user adds amount!"} // username balance overflow when create repo!
	ResRepoCurBalanceDownOverFlow = &Response{Status: 6, Message: "repo current balance less than 0 when user adds amount!"}

	ResCommitUploadSucc               = &Response{Status: 0, Message: "commit upload successfully!"}
	ResCommitUploadUnsucc             = &Response{Status: 7, Message: "commit upload unsuccessfully!"}
	ResPushUploadSucc                 = &Response{Status: 0, Message: "push upload successfully!"}
	ResPushUploadUnsucc               = &Response{Status: 8, Message: "push upload unsuccessfully!"}
	ResPullRequestUploadSucc          = &Response{Status: 0, Message: "pull request upload successfully!"}
	ResPullRequestUploadUnsucc        = &Response{Status: 9, Message: "pull request upload unsuccessfully!"}
	ResPullRequestCommentUploadSucc   = &Response{Status: 0, Message: "pull request comment upload successfully!"}
	ResPullRequestCommentUploadUnsucc = &Response{Status: 10, Message: "pull request comment upload unsuccessfully!"}

	ResUserNotExisted           = &Response{Status: 100, Message: "user not exist!"}
	ResUserAddAmountSucc        = &Response{Status: 0, Message: "user adds amount successfully!"}
	ResUserMinusAmountSucc      = &Response{Status: 0, Message: "user minus amount successfully!"}
	ResUserTransferAmountSucc   = &Response{Status: 0, Message: "user transfers amount successfully!"}
	ResUserLockAmountSucc       = &Response{Status: 0, Message: "user locks amount successfully!"}
	ResUserUnlockAmountSucc     = &Response{Status: 0, Message: "user unlocks amount successfully!"}
	ResUserAddAmountUnsucc      = &Response{Status: 101, Message: "user adds amount unsuccessfully!"}
	ResUserMinusAmountUnsucc    = &Response{Status: 102, Message: "user minus amount unsuccessfully!"}
	ResUserLockAmountUnsucc     = &Response{Status: 103, Message: "user locks amount unsuccessfully!"}
	ResUserUnlockAmountUnsucc   = &Response{Status: 104, Message: "user unlocks amount unsuccessfully!"}
	ResUserTransferAmountUnsucc = &Response{Status: 105, Message: "user transfers amount unsuccessfully!"}
	ResUserBalanceNotEnough     = &Response{Status: 106, Message: "user's Balance not enough!"}

	TypeTransferErr = &Response{Status: 10, Message: "type transfer error!"} // type transfer error!
	StringEmpty     = &Response{Status: 11, Message: "string empty!"}        // string empty!
)

func UnknownErr(err interface{}) *Response {
	return &Response{Status: 20, Message: fmt.Sprintf("unknown error, err:%v\n", err)}
}
