package controllers

import (
	"USTB-TeachSystem-API/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type GetLibBaseController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all libBase
// @Success 200 {object} models.libBase
// @router / [get]
func (libBase *GetLibBaseController) GetAll() {
	userName := libBase.GetString("username")
	password := libBase.GetString("password")

	libBase.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	libBase.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	libBase.Ctx.ResponseWriter.Header().Set("content-type", "application/json")             //返回数据格式是json

	res := models.LibLogin(userName, password, userName)
	libBase.Data["json"] = res
	libBase.ServeJSON()
}
