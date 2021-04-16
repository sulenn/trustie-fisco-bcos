package format

import (
	"testing"
)

func Test_ArgsJson(T *testing.T) {
	chaincodeArgsToJsonString, err := ArgsJson("transfer", "nudt1", "sulenn", "qiubing", "10")
	if err != nil {
		T.Error(err)
	}
	if *chaincodeArgsToJsonString != "{\"Func\":\"transfer\",\"Args\":[\"nudt1\",\"sulenn\",\"qiubing\",\"10\"]}" {
		T.Error(*chaincodeArgsToJsonString)
	}
}
