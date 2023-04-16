package routes

import (
	mvc "github.com/kataras/iris/v12/mvc"
	"lottery_llx/bootstrap"
	"lottery_llx/service"
	"lottery_llx/web/controller"
)

func Configure(b *bootstrap.Bootstrapper) {
	giftService := service.NewGiftServie()

	index := mvc.New(b.Party("/"))
	index.Register(giftService)
	index.Handle(new(controller.IndexController))

}
