// @APIVersion 1.0.0
// @Title USTB-TeachSystem-API
// @Description Here is an APIVersion for USTB-TeachSystem, just for fun.
// @Contact me@jeasonstudio.cn
// @TermsOfServiceUrl https://github.com/jeasonstudio
// @License MIT
// @LicenseUrl https://github.com/jeasonstudio/USTB-TeachSystem-API/blob/master/LICENSE
package routers

import (
	"USTB-TeachSystem-API/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		// beego.NSNamespace("/user.ustbsu",
		// 	beego.NSInclude(
		// 		&controllers.UserController{},
		// 	),
		// ),
		beego.NSNamespace("/course.ustbsu",
			beego.NSInclude(
				&controllers.GetCourseController{},
			),
		),
		beego.NSNamespace("/cxScore.ustbsu",
			beego.NSInclude(
				&controllers.GetCXScoreController{},
			),
		),
		beego.NSNamespace("/classScore.ustbsu",
			beego.NSInclude(
				&controllers.GetClassScoreController{},
			),
		),
		beego.NSNamespace("/examTime.ustbsu",
			beego.NSInclude(
				&controllers.GetExamTimeController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
