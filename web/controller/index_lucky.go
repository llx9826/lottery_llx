package controller

func (c *IndexController) GetLucky() string {
	c.Ctx.Header("Content-Type", "text/html")
	return ""
}
