package routers

import (
	"github.com/astaxie/beego"
	"shark/controllers"
)

func init () {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
}
