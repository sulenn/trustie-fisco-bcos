package channel

import (
	"gitea.com/macaron/macaron"
	api "github.com/sulenn/trustie-fisco-bcos/modules/structs"
)

// 在通道中实例化链码
func InstantiateChaincode(ctx *macaron.Context, opt api.InstantiateChaincodeOption) {
	var instantiateChaincodeOption = &api.InstantiateChaincodeOption{
		CommonOption:     api.CommonOption{ChannelName: opt.ChannelName, Orgname: opt.Orgname},
		Peers:            opt.Peers,
		ChaincodeName:    opt.ChaincodeName,
		ChaincodeVersion: opt.ChaincodeVersion,
		ChaincodeType:    opt.ChaincodeType,
	}
	ctx.JSON(200, instantiateChaincodeOption)
}

// 调用链码，完成交易
func InvokeChaincode(ctx *macaron.Context, opt api.InvokeChaincodeOption) {
	var invokeChaincodeOption = &api.InvokeChaincodeOption{
		CommonOption:  api.CommonOption{ChannelName: opt.ChannelName, Orgname: opt.Orgname},
		Peers:         opt.Peers,
		ChaincodeName: opt.ChaincodeName,
		FunctionName:  opt.FunctionName,
		Args:          opt.Args,
	}
	ctx.JSON(200, invokeChaincodeOption)
}

// 查询链码中数据
func QueryChaincode(ctx *macaron.Context) {
	var queryChaincodeOption = &api.QueryChaincodeOption{
		CommonOption:  api.CommonOption{ChannelName: ctx.Params(":channelName"), Orgname: ctx.Query("orgnization_name")},
		ChaincodeName: ctx.Params(":chaincodeName"),
		Peer:          ctx.Query("peer"),
		FunctionName:  ctx.Query("function_name"),
		Args:          ctx.Query("args"),
	}
	ctx.JSON(200, queryChaincodeOption)
}
