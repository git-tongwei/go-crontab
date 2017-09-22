package main

import (
	"html/template"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/go-crontab/jobs"
	"github.com/go-crontab/models"
	_ "github.com/go-crontab/routers"
)

const (
	VERSION = "1.0.1"
)

func init() {
	//初始化数据模型
	models.Init()
	jobs.InitJobs()
}

func main() {

	// 设置默认404页面
	beego.ErrorHandler("404", func(rw http.ResponseWriter, r *http.Request) {
		t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/error/404.html")
		data := make(map[string]interface{})
		data["content"] = "page not found"
		t.Execute(rw, data)
	})

	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.Run()
}
