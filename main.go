package main

import (
	"github.com/astaxie/beego"
	_ "github.com/iamMarkchu/rose/helpers/log"
	_ "github.com/iamMarkchu/rose/models"
	_ "github.com/iamMarkchu/rose/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
