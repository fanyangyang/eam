package controllers

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {
	c.Ctx.Output.Body([]byte("hello world"))
}
