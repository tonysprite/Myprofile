package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.easyhtml.cn"
	c.Data["Email"] = "catchyouhand@qq.com"
	c.TplName = "index.tpl"
}

type DetailController struct {
	beego.Controller
}

func (c *DetailController) Get() {
	c.Data["Website"] = "easyhtml.cn"
	c.Data["Email"] = "catchyouhand@qq.com"
	c.TplName = "detail.tpl"
}
