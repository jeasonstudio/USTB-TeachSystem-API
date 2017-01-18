package controllers

import (
	"USTB-TeachSystem-API/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type GetClassScoreController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all classScore
// @Success 200 {object} models.classScore
// @router / [get]
func (classScore *GetClassScoreController) GetAll() {
	userName := classScore.GetString("username")
	password := classScore.GetString("password")

	classScore.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	classScore.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	classScore.Ctx.ResponseWriter.Header().Set("content-type", "application/json")             //返回数据格式是json

	res := models.GetClassScoreFromLogin(userName, password)
	classScore.Data["json"] = res
	classScore.ServeJSON()
}
