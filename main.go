package main

import (
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

	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.Run()
}
