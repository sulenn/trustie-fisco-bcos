package structs

import "fmt"

type Response struct {
	Status  uint16 `json:"status"`
	Message string `json:"message"`
}

var (
	ResRepoSucc            = &Response{Status: 0, Message: "create repo successfully!"}                   // create repo successfully!
	ResRepoExisted         = &Response{Status: 1, Message: "repo existed!"}                               // repo existed!
	ResRepoBalanceOverFlow = &Response{Status: 2, Message: "username balance overflow when create repo!"} // username balance overflow when create repo!
	ResRepoUnsucc          = &Response{Status: 3, Message: "create repo unsuccessfully!"}                 // create repo unsuccessfully!

	ResUserNotExisted = &Response{Status: 100, Message: "user not exist!"}

	TypeTransferErr = &Response{Status: 10, Message: "type transfer error!"} // type transfer error!
	StringEmpty     = &Response{Status: 11, Message: "string empty!"}        // string empty!
)

func UnknownErr(err interface{}) *Response {
	return &Response{Status: 20, Message: fmt.Sprintf("unknown error, err:%v\n", err)}
}
