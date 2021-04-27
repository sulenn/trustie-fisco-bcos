package routes

import (
	"gitea.com/macaron/binding"
	"gitea.com/macaron/macaron"
	api "github.com/sulenn/trustie-fisco-bcos/modules/structs"
	"github.com/sulenn/trustie-fisco-bcos/routers/fiscobcos"
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

	// 统一路由入口
	m.Group("/trustie", func() {
		m.Post("/fiscobcos", bind(api.FiscoBcos{}), fiscobcos.HandleAllRoutes)
	})
}
