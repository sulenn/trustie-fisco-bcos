package routes

import (
	"gitea.com/macaron/binding"
	"gitea.com/macaron/macaron"
	api "github.com/sulenn/trustie-fisco-bcos/modules/structs"
	"github.com/sulenn/trustie-fisco-bcos/routers/repo"
	"github.com/sulenn/trustie-fisco-bcos/routers/user"
)

// NewMacaron is to create a new macaron
func NewMacaron() *macaron.Macaron {
	var m = macaron.Classic() // 经典 macaron，包括 Logger、Recovery 和 Static
	m.Use(macaron.Renderer()) // 数据渲染中间件，将数据渲染成 JSON 格式或 HTML 格式
	RegisterRoutes(m)         // 注册路由
	return m
}

// RegisterRoutes is register routes
func RegisterRoutes(m *macaron.Macaron) {
	bind := binding.Bind // 数据绑定和验证辅助模块

	m.Get("/", func(ctx *macaron.Context) {
		ctx.JSON(200, "hello world")
	})

	m.Group("/user", func() {
		// m.Post("/register", bind(api.CreateUserOption{}), user.Register) // 注册用户
		m.Get("/query/allbalance", user.SelectUserAllBalance) // 查询用户所有项目Token
		m.Get("/query/balance", user.SelectUserBalance)       // 查询用户单个项目Token
	})

	// m.Group("/channel", func() {
	// 	m.Post("/create", bind(api.CreateChannelOption{}), channel.Create) // 创建通道
	// 	m.Post("/join", bind(api.JoinChannelOption{}), channel.Join)       // Perr 节点加入通道
	// 	m.Group("/chaincode", func() {
	// 		m.Post("/instantiate", bind(api.InstantiateChaincodeOption{}), channel.InstantiateChaincode) // 实例化链码
	// 		m.Post("/invoke", bind(api.InvokeChaincodeOption{}), channel.InvokeChaincode)                // 调用链码
	// m.Get("/query/:channelName/:chaincodeName", channel.QueryChaincode) // 查询链码，注意 channelName 不能为 channel_name，无法识别，导致 api 路径出错，404 错误
	// 	})
	// })

	// m.Group("/chaincode", func() {
	// 	m.Post("/install", bind(api.InstallChaincodeOption{}), chaincode.Install) // 安装链码
	// })

	m.Group("/repos", func() {
		m.Post("/create", bind(api.CreateRepoOption{}), repo.Create) // 在区块链账本中新建一个项目
		// m.Group("/amount", func() {                                  // 和 token 值相关路由
		// 	m.Post("/transfer", bind(api.TransferTokenOption{}), repo.Transfer) // 转账
		// 	m.Post("/add", bind(api.AddTokenOption{}), repo.Add)                // 增加账户余额
		// 	m.Post("/minus", bind(api.MinusTokenOption{}), repo.Minus)          // 减少账户余额
		m.Get("/query/basic", repo.QueryBasic) // 查询账户余额
		// 	m.Get("/repo/query", repo.QueryRepoAmount)                          // 查询项目总金额
		// })
		// m.Group("/commit", func() {
		// 	m.Post("/upload", bind(api.UploadCommitOption{}), repo.Upload) // 记录 commit 数据
		// })
	})
}
