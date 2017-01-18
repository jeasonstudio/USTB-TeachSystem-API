package main

import (
	_ "USTB-TeachSystem-API/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
	}
	beego.Alert("USTB TeachSystem API Beego Started")
	beego.Run()
}
