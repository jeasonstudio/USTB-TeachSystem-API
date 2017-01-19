package controllers

import (
	"USTB-TeachSystem-API/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type GetExamTimeController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all examTime
// @Success 200 {object} models.examTime
// @router / [get]
func (examTime *GetExamTimeController) GetAll() {
	userName := examTime.GetString("username")
	password := examTime.GetString("password")
	semestre := examTime.GetString("semestre")

	examTime.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	examTime.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	examTime.Ctx.ResponseWriter.Header().Set("content-type", "application/json")             //返回数据格式是json

	res := models.GetExamTimeFromLogin(userName, password, semestre)
	examTime.Data["json"] = res
	examTime.ServeJSON()
}
