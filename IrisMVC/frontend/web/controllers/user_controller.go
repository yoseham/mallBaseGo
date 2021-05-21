package controllers

import (
	"app/datamodels"
	"app/services"
	"app/tool"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"strconv"
)

type UserController struct {
	Ctx         iris.Context
	UserService services.IUserService
	Session     *sessions.Session
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
	//ozzo-validataion 表单验证
	user := &datamodels.User{UserName: userName, NickName: nickName, HashPassword: password}
	_, err := c.UserService.AddUser(user)
	if err != nil {
		c.Ctx.Redirect("/user/error")
	} else {
		c.Ctx.Redirect("/user/login")
	}
	return
}

func (c *UserController) GetLogin() mvc.View {
	return mvc.View{
		Name: "user/login.html",
	}
}

func (c *UserController) PostLogin() mvc.Response {
	var (
		userName = c.Ctx.FormValue("userName")
		password = c.Ctx.FormValue("password")
	)
	user, isOk := c.UserService.IsPwdSuccess(userName, password)
	if !isOk {
		fmt.Println("登录失败")
		return mvc.Response{
			Path: "/user/login",
		}
	} else {
		tool.GlobalCookie(c.Ctx, "uid", strconv.FormatInt(user.ID, 10))
		c.Session.Set("userID", strconv.FormatInt(user.ID, 10))
		return mvc.Response{
			Path: "/product/",
		}
	}
}
