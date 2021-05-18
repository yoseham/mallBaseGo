package main

import (
	"app/backend/web/controllers"
	"app/common"
	"app/repositories"
	"app/services"
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	//创建iris实例
	app := iris.New()
	//设置错误模式
	app.Logger().SetLevel("debug")
	//注册模板
	app.RegisterView(iris.HTML("./backend/web/views", ".html").
		Layout("shared/layout.html").Reload(true))
	//设置模板目录
	app.HandleDir("/assets", "./backend/web/assets")
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

	//注册控制器
	mvc.New(app.Party("/hello")).Handle(new(controllers.MovieController))
	productRepository := repositories.NewProductManager("product", db)
	productService := services.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, productService)
	product.Handle(new(controllers.ProductController))

	orderRepository := repositories.NewOrderManagerRepository("order", db)
	orderService := services.NewOrderService(orderRepository)
	orderParty := app.Party("/order")
	order := mvc.New(orderParty)
	order.Register(ctx, orderService)
	order.Handle(new(controllers.OrderController))
	//启动服务
	app.Run(
		iris.Addr("localhost:8899"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
