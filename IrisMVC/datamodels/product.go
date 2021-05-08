package datamodels

type Product struct {
	ID int64 `json:"id" imooc:"id"`
	ProductName string `json:"ProductName" sql:"ProductName"`
}
