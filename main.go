package main

import (
	"github.com/astaxie/beego/orm"
	"webApp/controllers"
	"webApp/models"
	_ "webApp/routers"
	"github.com/astaxie/beego"
)

func init()  {
	models.RegisterDB()
}
func main() {
	orm.Debug=true
	orm.RunSyncdb("default",false,true)
	beego.Router("/",&controllers.MainController{})
	beego.Run()
}

