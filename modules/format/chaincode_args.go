package format

import (
	"encoding/json"
	"errors"
)

type ChaincodeArgs struct {
	Func string   `json:"Func"` // 调用链码时指定的函数名
	Args []string `json:"Args"` // 调用链码函数时传入的参数
}

func ArgsJson(function string, args ...string) (*string, error) { //链码参数转字符串
	if len(function) == 0 {
		return nil, errors.New("function name is not nil")
	}
	chaincodeArgs := &ChaincodeArgs{
		Func: function,
		Args: args,
	}
	chaincodeArgsToJsonBytes, err := json.Marshal(chaincodeArgs)
	if err != nil {
		return nil, err
	}
	chaincodeArgsToJsonString := string(chaincodeArgsToJsonBytes)
	return &chaincodeArgsToJsonString, nil
}
