package main

import (
	"app/backend/web/controllers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	//创建iris实例
	app := iris.New()
	//设置错误模式
	app.Logger().SetLevel("debug")
	//注册模板
	app.RegisterView(iris.HTML("./backend/web/views",".html").
		Layout("shared/layout.html").Reload(true))
	//设置模板目录
	app.HandleDir("/assets","./backend/web/assets")
	//出现异常跳转指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message","页面出错"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	//注册控制器
	mvc.New(app.Party("/hello")).Handle(new(controllers.MovieController))


	//启动服务
	app.Run(
		iris.Addr("localhost:8899"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		)
}
