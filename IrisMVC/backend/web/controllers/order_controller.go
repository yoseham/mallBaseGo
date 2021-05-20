package controllers

import (
	"app/common"
	"app/datamodels"
	"app/services"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"strconv"
)

type OrderController struct {
	Ctx          iris.Context
	OrderService services.IOrderService
}

func (o *OrderController) GetAll() mvc.View {
	orderArray, err := o.OrderService.GetAllOrderInfo()
	fmt.Println(orderArray)
	if err != nil {
		o.Ctx.Application().Logger().Debug(err)
	}
	sended := 0
	nosend := 0
	for k, _ := range orderArray {
		if orderArray[k]["OrderStatus"] == "1" {
			sended += 1
		} else {
			nosend += 1
		}
	}
	return mvc.View{
		Name: "order/view.html",
		Data: iris.Map{
			"order":  orderArray,
			"sended": sended,
			"nosend": nosend,
		},
	}
}

func (o *OrderController) GetAdd() mvc.View {
	return mvc.View{
		Name: "order/add.html",
	}
}

func (o *OrderController) PostAdd() {
	order := &datamodels.Order{}
	o.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{TagName: "sql"})
	if err := dec.Decode(o.Ctx.Request().Form, order); err != nil {
		fmt.Println(err)
		o.Ctx.Application().Logger().Debug(err)
	}
	_, err := o.OrderService.InsertOrder(order)
	if err != nil {
		fmt.Println(err)
		o.Ctx.Application().Logger().Debug(err)
	}
	o.Ctx.Redirect("/order/all")
}

func (o *OrderController) GetDelete() {
	idString := o.Ctx.URLParam("id")
	id, err := strconv.ParseInt(idString, 10, 16)
	if err != nil {
		o.Ctx.Application().Logger().Debug(err)
	}
	isOk, err := o.OrderService.DeleteOrderByID(id)
	if err != nil {
		o.Ctx.Application().Logger().Debug(err)
	}
	if isOk {
		o.Ctx.Application().Logger().Debug("删除订单成功，id=" + idString)
	} else {
		o.Ctx.Application().Logger().Debug("删除订单失败，id=" + idString)
	}
	o.Ctx.Redirect("/order/all")
}

func (o *OrderController) PostUpdate() {
	order := &datamodels.Order{}
	o.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{TagName: "sql"})
	if err := dec.Decode(o.Ctx.Request().Form, order); err != nil {
		o.Ctx.Application().Logger().Debug(err)
	}
	err := o.OrderService.UpdateOrder(order)
	if err != nil {
		o.Ctx.Application().Logger().Debug(err)
	}
	o.Ctx.Redirect("/order/all")
}

func (o *OrderController) GetManager() mvc.View {
	idString := o.Ctx.URLParam("id")
	id, err := strconv.ParseInt(idString, 10, 16)
	if err != nil {
		o.Ctx.Application().Logger().Debug(err)
	}
	order, err := o.OrderService.GetOrderByID(id)
	if err != nil {
		o.Ctx.Application().Logger().Debug(err)
	}
	return mvc.View{
		Name: "order/manager.html",
		Data: iris.Map{
			"order": order,
		},
	}
}
