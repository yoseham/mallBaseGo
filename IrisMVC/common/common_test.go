package common

import (
	"app/datamodels"
	"reflect"
	"testing"
)

func TestDataToStructByTagSql(t *testing.T) {
	data := map[string]string{
		"ID":           "1",
		"ProductName":  "imoocTest",
		"ProductNum":   "123",
		"ProductUrl":   "http://url",
		"ProductImage": "http://image",
	}
	product := datamodels.Product{}
	DataToStructByTagSql(data, &product)

	want := datamodels.Product{1, "imoocTest", 123, "http://image", "http://url"}
	if !reflect.DeepEqual(want, product) {
		t.Errorf("want %v get %v", want, product)
	}
}
