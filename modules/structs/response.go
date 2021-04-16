package structs

type Response struct {
	Status  uint16 `json:"status"`
	Message string `json:"message"`
}

var (
	ResRepoSucc            = &Response{Status: 0, Message: "create repo successfully!"}                   // create repo successfully!
	ResRepoExisted         = &Response{Status: 1, Message: "repo existed!"}                               // repo existed!
	ResRepoBalanceOverFlow = &Response{Status: 1, Message: "username balance overflow when create repo!"} // username balance overflow when create repo!
	ResRepoUnsucc          = &Response{Status: 2, Message: "create repo unsuccessfully!"}                 // create repo unsuccessfully!

	TypeTransferErr = &Response{Status: 10, Message: "type transfer error!"} // type transfer error!
	StringEmpty     = &Response{Status: 10, Message: "string empty!"}        // string empty!
)
