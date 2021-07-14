package controllers

type AboutMeController struct {
	BaseController
}

func (c *AboutMeController) Get() {
	c.TplName = "aboutme.html"
}
