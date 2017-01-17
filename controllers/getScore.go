package controllers

import (
	"USTB-TeachSystem-API/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type getScoreController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *getScoreController) GetAll() {
	users := models.GetAllUsers()
	uid := u.GetString("uid")
	u.Data["json"] = users
	u.Data["json"] = uid
	u.ServeJSON()
}
