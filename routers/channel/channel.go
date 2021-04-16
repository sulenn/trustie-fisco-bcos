package channel

import (
	api "github.com/sulenn/trustie-fisco-bcos/modules/structs"

	"gitea.com/macaron/macaron"
)

// 创建 channel
func Create(ctx *macaron.Context, opt api.CreateChannelOption) {
	var channel = &api.CreateChannelOption{
		CommonOption:      api.CommonOption{ChannelName: opt.ChannelName, Orgname: opt.Orgname},
		ChannelConfigPath: opt.ChannelConfigPath,
	}
	ctx.JSON(200, channel)
}

// 节点加入通道
func Join(ctx *macaron.Context, opt api.JoinChannelOption) {
	var joinChannelOption = &api.JoinChannelOption{
		CommonOption: api.CommonOption{ChannelName: opt.ChannelName, Orgname: opt.Orgname},
		Peers:        opt.Peers,
	}
	ctx.JSON(200, joinChannelOption)
}
