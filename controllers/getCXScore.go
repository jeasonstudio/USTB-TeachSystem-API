package controllers

import (
	"USTB-TeachSystem-API/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type GetCXScoreController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all cxScore
// @Success 200 {object} models.cxScore
// @router / [get]
func (cxScore *GetCXScoreController) GetAll() {
	userName := cxScore.GetString("username")
	password := cxScore.GetString("password")

	cxScore.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	cxScore.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	cxScore.Ctx.ResponseWriter.Header().Set("content-type", "application/json")             //返回数据格式是json

	res := models.GetCXScoreFromLogin(userName, password)
	cxScore.Data["json"] = res
	cxScore.ServeJSON()
}
