package routers

import (
	"github.com/astaxie/beego"
	"shark/controllers"
)

func init () {
	beego.Router("/", &controllers.MainController{})
}
