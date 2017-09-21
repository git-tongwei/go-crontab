package controllers

import (
	"time"

	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/go-crontab/libs"
	"github.com/go-crontab/models"
)

type ServerController struct {
	BaseController
}

func (this *ServerController) List() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}
	result, count := models.GetTaskServerList(page, this.pageSize)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		if v.Type == 0 {
			row["type"] = "密码"
		} else {
			row["type"] = "密钥"
		}
		row["server_name"] = v.ServerName
		row["server_ip"] = v.ServerIp
		row["detail"] = v.Detail
		row["port"] = v.Port
		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		list[k] = row
	}
	this.Data["pageTitle"] = "服务器列表"
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("ServerController.List"), true).ToString()
	this.display()
}
func (this *ServerController) Add() {
	if this.isPost() {
		server := new(models.TaskServer)
		server.ServerName = strings.TrimSpace(this.GetString("server_name"))
		server.ServerAccount = strings.TrimSpace(this.GetString("server_account"))
		server.ServerIp = strings.TrimSpace(this.GetString("server_ip"))
		server.Port, _ = strconv.Atoi(this.GetString("port"))
		server.Type, _ = strconv.Atoi(this.GetString("type"))
		server.PrivateKeySrc = strings.TrimSpace(this.GetString("private_key_src"))
		server.PublicKeySrc = strings.TrimSpace(this.GetString("public_key_src"))
		server.Password = strings.TrimSpace(this.GetString("password"))
		server.Detail = strings.TrimSpace(this.GetString("detail"))
		server.CreateTime = time.Now().Unix()
		server.UpdateTime = time.Now().Unix()
		server.Status = 0
		_, err := models.AddTaskServer(server)
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "添加服务器"
	this.display()
}
func (this *ServerController) Edit() {
	id, _ := this.GetInt("id")
	server, err := models.GetTaskServerById(id)
	if err != nil {
		this.showMsg(err.Error())
	}

	if this.isPost() {
		server.ServerName = strings.TrimSpace(this.GetString("server_name"))
		server.ServerAccount = strings.TrimSpace(this.GetString("server_account"))
		server.ServerIp = strings.TrimSpace(this.GetString("server_ip"))
		server.Port, _ = strconv.Atoi(this.GetString("port"))
		server.Type, _ = strconv.Atoi(this.GetString("type"))
		server.Id, _ = strconv.Atoi(this.GetString("id"))
		server.PrivateKeySrc = strings.TrimSpace(this.GetString("private_key_src"))
		server.PublicKeySrc = strings.TrimSpace(this.GetString("public_key_src"))
		server.Password = strings.TrimSpace(this.GetString("password"))
		server.Detail = strings.TrimSpace(this.GetString("detail"))
		server.UpdateTime = time.Now().Unix()
		server.Status = 0
		err := server.Update()
		if err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	this.Data["pageTitle"] = "编辑服务器"
	this.Data["server"] = server
	this.display()
}
