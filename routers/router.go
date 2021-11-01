package routers

import (
	"Myprofile/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/detail", &controllers.DetailController{})
}
