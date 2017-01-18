package controllers

import (
	"USTB-TeachSystem-API/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type GetCourseController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Course
// @Success 200 {object} models.Course
// @router / [get]
func (course *GetCourseController) GetAll() {
	userName := course.GetString("username")
	password := course.GetString("password")
	semestre := course.GetString("semestre")
	course.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	course.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	course.Ctx.ResponseWriter.Header().Set("content-type", "application/json")             //返回数据格式是json

	res := models.GetCourseFromLogin(userName, password, semestre)
	course.Data["json"] = res
	course.ServeJSON()
}
