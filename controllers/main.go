package controllers

import (
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils"
	"github.com/dchest/captcha"
	"github.com/go-crontab/libs"
	"github.com/go-crontab/models"
)

type MainController struct {
	BaseController
}

//首页
func (this *MainController) Index() {
	this.Data["pageTitle"] = "系统概况"
	this.display()
}

//个人信息
func (this *MainController) Profile() {
	beego.ReadFromRequest(&this.Controller)
	user, _ := models.GetUserById(this.userId)
	if this.isPost() {
		user.Email = this.GetString("email")
		user.Update()
		password1 := this.GetString("password1")
		password2 := this.GetString("password2")
		if password1 != "" {
			if len(password1) < 6 {
				this.ajaxMsg("密码长度必须大于6位", MSG_ERR)
			} else if password2 != password1 {
				this.ajaxMsg("两次输入的密码不一致", MSG_ERR)
			} else {
				user.Salt = string(utils.RandomCreateBytes(10))
				user.Password = libs.Md5([]byte(password1 + user.Salt))
				user.Update()
			}
		}
		this.ajaxMsg("", MSG_OK)
	}
	this.Data["pageTitle"] = "资料修改"
	this.Data["user"] = user
	this.display()
}

func (this *MainController) Login() {

	if this.userId > 0 {
		this.redirect("/")
	}
	beego.ReadFromRequest(&this.Controller)
	if this.isPost() {
		flash := beego.NewFlash()
		errmsg := ""

		username := strings.TrimSpace(this.GetString("username"))
		password := strings.TrimSpace(this.GetString("password"))
		captchaValue := strings.TrimSpace(this.GetString("captcha"))
		remember := this.GetString("remember")
		captchaId := this.GetString("captchaId")
		if username != "" && password != "" && captchaValue != "" {

			//验证码校验
			if !captcha.VerifyString(captchaId, captchaValue) {
				errmsg = "验证码错误！"
				flash.Error(errmsg)
				flash.Store(&this.Controller)
				this.redirect(beego.URLFor("MainController.Login"))
			}

			user, err := models.GetUserByName(username)

			if err != nil || user.Password != libs.Md5([]byte(password+user.Salt)) {
				errmsg = "帐号或密码错误"
			} else if user.Status == -1 {
				errmsg = "该帐号已禁用"
			} else {
				user.LastIp = this.getClientIp()
				user.LastLogin = time.Now().Unix()
				models.UpdateUser(user)

				authkey := libs.Md5([]byte(this.getClientIp() + "|" + user.Password + user.Salt))
				if remember == "yes" {
					this.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 7*86400)
				} else {
					this.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 86400)
				}
				this.redirect(beego.URLFor("TaskController.List"))
			}
			flash.Error(errmsg)
			flash.Store(&this.Controller)
			this.redirect(beego.URLFor("MainController.Login"))

		}
	}

	//验证码
	d := struct {
		CaptchaId string
	}{
		captcha.NewLen(4),
	}
	this.Data["CaptchaId"] = d.CaptchaId
	this.TplName = "public/login.html"
}

func (this *MainController) Logout() {
	this.Ctx.SetCookie("auth", "")
	this.redirect(beego.URLFor("MainController.Login"))
}

// 获取系统时间
func (this *MainController) GetTime() {
	out := make(map[string]interface{})
	out["time"] = time.Now().UnixNano() / int64(time.Millisecond)
	this.jsonResult(out)
}
