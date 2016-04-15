/*
SpotlightController
*/


package controllers

import (
	"github.com/astaxie/beego"
	"go-MyVIT/api"
)

// Operations about login
type SpotlightController struct {
	beego.Controller
}

// @Title Get
// @Description find object by objectid
// @Success 200
// @Failure 403 parameters missing
// @router / [get]
func (o *SpotlightController) Get() {
	regNo := o.Input().Get("regNo")
	psswd := o.Input().Get("psswd")
	campus := o.Ctx.Input.Param(":campus")
	var baseuri string
	if campus == "vellore" {
		baseuri = "https://academics.vit.ac.in"
	} else {
		baseuri = "https://academicscc.vit.ac.in"
	}
	if regNo != "" && psswd != "" {
		resp := api.Spotlight(regNo, psswd, baseuri)
		o.Data["json"] = resp
	}
	o.ServeJSON()
}