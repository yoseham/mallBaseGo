package repositories

import (
	"app/common"
	"app/datamodels"
	"reflect"
	"testing"
)

func TestProductManager_SelectAll(t *testing.T) {
	db, err := common.NewMysqlConn()
	if err != nil {
		t.Error(err)
	}
	productManager := NewProductManager("product", db)
	products, err := productManager.SelectAll()
	if err != nil {
		t.Error(err)
	}
	want := []*datamodels.Product{&datamodels.Product{1, "apple", 10, "http://image", "http://url"}}
	if !reflect.DeepEqual(products, want) {
		t.Errorf("got %v want %v", products, want)
	}
}

func TestProductManager_SelectByKey(t *testing.T) {
	db, err := common.NewMysqlConn()
	if err != nil {
		t.Error(err)
	}
	productManager := NewProductManager("product", db)
	product, err := productManager.SelectByKey(1)
	if err != nil {
		t.Error(err)
	}
	want := &datamodels.Product{1, "apple", 10, "http://image", "http://url"}
	if !reflect.DeepEqual(product, want) {
		t.Errorf("got %v want %v", product, want)
	}

}
