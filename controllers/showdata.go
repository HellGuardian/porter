package controllers

import (
	"github.com/astaxie/beego"
)

type ShowDataController struct {
	beego.Controller
}

func (l * ShowDataController) Get() {
	l.TplName = "showdata.html"
}