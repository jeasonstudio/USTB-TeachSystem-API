package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:GetLibBaseController"] = append(beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:GetLibBaseController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:GetCetScoreController"] = append(beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:GetCetScoreController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:GetCXScoreController"] = append(beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:GetCXScoreController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:GetClassScoreController"] = append(beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:GetClassScoreController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:GetCourseController"] = append(beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:GetCourseController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:GetExamTimeController"] = append(beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:GetExamTimeController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:UserController"] = append(beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:UserController"] = append(beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:UserController"] = append(beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:UserController"] = append(beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"put"},
			Params:           nil})

	beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:UserController"] = append(beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params:           nil})

	beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:UserController"] = append(beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/login`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:UserController"] = append(beego.GlobalControllerRouter["USTB-TeachSystem-API/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Logout",
			Router:           `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

}
