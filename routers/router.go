package routers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"github.com/go-crontab/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/login", &controllers.MainController{}, "*:Login")
	beego.Router("/logout", &controllers.MainController{}, "*:Logout")
	beego.Router("/profile", &controllers.MainController{}, "*:Profile")
	beego.Router("/gettime", &controllers.MainController{}, "*:GetTime")
	beego.Handler("/captcha/*.png", captcha.Server(240, 80))
	beego.Router("/help", &controllers.HelpController{}, "*:Index")
	beego.AutoRouter(&controllers.ServerController{})
	beego.AutoRouter(&controllers.GroupController{})
	beego.AutoRouter(&controllers.TaskController{})
}
