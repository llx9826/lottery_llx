package controller

import (
	"github.com/kataras/iris/v12"
	"lottery_llx/model"
	"lottery_llx/service"
)

type IndexController struct {
	Ctx         iris.Context
	ServiceGift service.GiftService
}

func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "welcome to 抽奖系统"
}

func (c *IndexController) GetGifts() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""
	datalist := c.ServiceGift.GetAll(false)
	list := make([]model.LtGift, 0)
	for _, data := range datalist {
		// 正常状态的才需要放进来
		if data.SysStatus == 0 {
			list = append(list, data)
		}
	}
	rs["gifts"] = list
	return rs
}
