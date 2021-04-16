package structs

type InstantiateChaincodeOption struct {
	CommonOption
	Peers              []string  `json:"peers"`
	ChaincodeName      string    `json:"chaincode_name"`
	ChaincodeVersion   string    `json:"chaincode_version"`
	ChaincodeType      string    `json:"chaincode_type"`   // 链码所用语言
}

type InvokeChaincodeOption struct {
	CommonOption
	Peers              []string  `json:"peers"`
	ChaincodeName      string    `json:"chaincode_name"`
	FunctionName       string    `json:"function_name"`   // 调用链码中的方法名
	Args               []string  `json:"args"`            // 调用链码时传入的参数
}

type QueryChaincodeOption struct {
	CommonOption
	ChaincodeName      string    `json:"chaincode_name"`
	Peer               string    `json:"peer"`
	FunctionName       string    `json:"function_name"`   // 调用链码中的方法名
	Args               string    `json:"args"`            // 调用链码时传入的参数
}