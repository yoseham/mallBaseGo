package controllers

import (
	"app/services"
	"github.com/kataras/iris/v12"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService services.ProductService
}

//func (p *ProductController) GetAll() mvc.View {
//	productArray, _ := p.ProductService.GetAllProduct()
//	return mvc.View{
//
//	}
//}
