package controllers

import (
	"flash-sale/datamodels"
	"flash-sale/services"
	"flash-sale/tool"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"strconv"
)

type UserController struct {
	Ctx     iris.Context
	Service services.IUserService
	Session *sessions.Session
}

func (c *UserController) GetRegister() mvc.View {
	return mvc.View{
		Name: "user/register.html",
	}
}

func (c *UserController) PostRegister() {
	var (
		nickName = c.Ctx.FormValue("nickName")
		userName = c.Ctx.FormValue("userName")
		password = c.Ctx.FormValue("password")
	)

	user := &datamodels.User{
		UserName:     userName,
		NickName:     nickName,
		HashPassword: password,
	}

	_, err := c.Service.AddUser(user)
	c.Ctx.Application().Logger().Debug(err)
	if err != nil {
		c.Ctx.Redirect("/user/error")
		return
	}
	c.Ctx.Redirect("/user/login")
	return
}

func (c *UserController) GetLogin() mvc.View {
	return mvc.View{
		Name: "user/login.html",
	}
}

func (c *UserController) PostLogin() mvc.Response {
	//获取用户提交的表单信息
	var (
		userName = c.Ctx.FormValue("userName")
		password = c.Ctx.FormValue("password")
	)
	//验证账号密码正确
	user, isOk := c.Service.Login(userName, password)
	if !isOk {
		return mvc.Response{
			Path: "/user/login",
		}
	}
	//写入用户ID到cookie中
	tool.GlobalCookie(c.Ctx, "uid", strconv.FormatInt(user.ID, 10))
	c.Session.Set("userID",strconv.FormatInt(user.ID,10))
	return mvc.Response{
		Path: "/product/",
	}

}
