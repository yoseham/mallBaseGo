package services

import (
	"app/common"
	"app/datamodels"
	"app/repositories"
	"reflect"
	"testing"
)

func TestProductService_GetAllProduct(t *testing.T) {
	db, err := common.NewMysqlConn()
	if err != nil {
		t.Error(err)
	}
	productRepository := repositories.NewProductManager("product", db)
	productService := NewProductService(productRepository)
	products, err := productService.GetAllProduct()
	if err != nil {
		t.Error(err)
	}
	product := &datamodels.Product{1, "apple", 10, "http://image", "http://url"}
	want := []*datamodels.Product{product}

	if !reflect.DeepEqual(products, want) {
		t.Errorf("got %v want %v", products, want)
	}
}
