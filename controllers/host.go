package controllers

import (
	"github.com/astaxie/beego"
)

type HostController struct {
	beego.Controller
}

func (h *HostController) Get() {
	h.TplName = "host.html"
}