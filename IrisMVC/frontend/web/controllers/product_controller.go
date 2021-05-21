package controllers

import (
	"app/datamodels"
	"app/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"strconv"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService services.IproductService
	OrderService   services.IOrderService
	Session        *sessions.Session
}

func (p *ProductController) GetDetail() mvc.View {
	idString := p.Ctx.URLParam("productID")
	id, _ := strconv.ParseInt(idString, 10, 16)
	product, err := p.ProductService.GetProductByID(id)
	if err != nil {
		p.Ctx.Application().Logger().Error(err)
	}
	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/view.html",
		Data: iris.Map{
			"product": product,
		},
	}
}

func (p *ProductController) GetOrder() mvc.View {
	productIDString := p.Ctx.URLParam("productID")
	userID := p.Ctx.GetCookie("uid")
	productID, err := strconv.Atoi(productIDString)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	product, err := p.ProductService.GetProductByID(int64(productID))
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	var orderID int64
	showMessage := "抢购失败！"
	if product.ProductNum > 0 {
		product.ProductNum -= 1
		err := p.ProductService.UpdateProduct(product)
		if err != nil {
			p.Ctx.Application().Logger().Debug(err)
		}
		userID, err := strconv.Atoi(userID)
		if err != nil {
			p.Ctx.Application().Logger().Debug(err)
		}
		order := &datamodels.Order{UserID: int64(userID), ProductID: int64(productID), OrderStatus: datamodels.OrderSuccess}
		orderID, err = p.OrderService.InsertOrder(order)
		if err != nil {
			p.Ctx.Application().Logger().Debug(err)
		} else {
			showMessage = "抢购成功"
		}
	}
	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/result.html",
		Data: iris.Map{
			"orderID":     orderID,
			"showMessage": showMessage,
		},
	}
}
