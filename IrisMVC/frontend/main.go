package main

import (
	"app/common"
	"app/frontend/middlerware"
	"app/frontend/web/controllers"
	"app/repositories"
	"app/services"
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"time"
)

func main() {
	//创建iris实例
	app := iris.New()
	//设置错误模式
	app.Logger().SetLevel("debug")
	//注册模板
	app.RegisterView(iris.HTML("./frontend/web/views", ".html").
		Layout("shared/layout.html").Reload(true))
	//设置模板目录
	app.HandleDir("/public", "./frontend/web/public")
	//出现异常跳转指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "页面出错"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	//连接mysql
	db, err := common.NewMysqlConn()
	if err != nil {

	}
	//上下文
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sess := sessions.New(sessions.Config{Cookie: "helloworld", Expires: 60 * time.Minute})
	userRepository := repositories.NewUserRepository("user", db)
	userService := services.NewUserService(userRepository)
	user := mvc.New(app.Party("/user"))
	user.Register(userService, ctx, sess.Start)
	user.Handle(new(controllers.UserController))

	product := repositories.NewProductManager("product", db)
	order := repositories.NewOrderManagerRepository("iorder", db)
	orderService := services.NewOrderService(order)
	productService := services.NewProductService(product)
	productPro := app.Party("/product")
	productPro.Use(middlerware.AuthConProduct)
	pro := mvc.New(productPro)
	pro.Register(productService, orderService, sess.Start)
	pro.Handle(new(controllers.ProductController))

	app.Run(
		iris.Addr("localhost:8900"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
