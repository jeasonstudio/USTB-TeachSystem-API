package controllers

import (
	"USTB-TeachSystem-API/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type GetCetScoreController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all cetScore
// @Success 200 {object} models.cetScore
// @router / [get]
func (cetScore *GetCetScoreController) GetAll() {
	userName := cetScore.GetString("username")
	password := cetScore.GetString("password")

	cetScore.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	cetScore.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	cetScore.Ctx.ResponseWriter.Header().Set("content-type", "application/json")             //返回数据格式是json

	res := models.GetCETScoreFromLogin(userName, password)
	cetScore.Data["json"] = res
	cetScore.ServeJSON()
}
