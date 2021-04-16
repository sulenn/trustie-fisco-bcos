package main

import (
	"github.com/sulenn/trustie-fisco-bcos/modules/os"
	"github.com/sulenn/trustie-fisco-bcos/routers/routes"
)

func main() {
	m := routes.NewMacaron()
	os.SetListeningPort("8000") // 设置监听端口
	m.Run()                     // 如果存在环境变量 port，则读取。否则默认监听 4000
}
