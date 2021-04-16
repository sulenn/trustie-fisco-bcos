package chaincode

import (
	"gitea.com/macaron/macaron"
	api "github.com/sulenn/trustie-fisco-bcos/modules/structs"
)

func Install(ctx *macaron.Context, opt api.InstallChaincodeOption) {
	var installChaincodeOption = &api.InstallChaincodeOption{
		Orgname:          opt.Orgname,
		Peers:            opt.Peers,
		ChaincodeName:    opt.ChaincodeName,
		ChaincodePath:    opt.ChaincodePath,
		ChaincodeVersion: opt.ChaincodeVersion,
		ChaincodeType:    opt.ChaincodeType,
	}
	ctx.JSON(200, installChaincodeOption)
}
