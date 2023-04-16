package main

import (
	"fmt"
	"lottery_llx/bootstrap"
	"lottery_llx/web/middleware"
	"lottery_llx/web/routes"
)

var port = 8080

func newApp() *bootstrap.Bootstrapper {
	// 初始化应用
	app := bootstrap.New("Go抽奖系统", "llx")
	app.Bootstrap()
	app.Configure(middleware.Configure, routes.Configure)

	return app
}
func main() {
	// 服务器集群的时候才需要区分这项设置
	// 比如：根据服务器的IP、名称、端口号等，或者运行的参数
	app := newApp()
	app.Listen(fmt.Sprintf(":%d", port))
}
