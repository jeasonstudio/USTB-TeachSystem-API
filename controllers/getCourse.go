package controllers

import (
	"USTB-TeachSystem-API/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type GetScoreController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (course *GetScoreController) GetAll() {
	userName := course.GetString("username")
	password := course.GetString("password")
	semestre := course.GetString("semestre")

	res := models.GetCourseFromLogin(userName, password, semestre)
	course.Data["json"] = res
	course.ServeJSON()
}
