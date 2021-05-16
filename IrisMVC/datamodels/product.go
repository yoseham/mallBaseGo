package datamodels

type Product struct {
	ID           int64  `json:"ID" sql:"ID" imooc:"ID"`
	ProductName  string `json:"ProductName" sql:"ProductName" imooc:"ProductName"`
	ProductNum   int64  `json:"ProductNum" sql:"ProductNum" imooc:"ProductNum"`
	ProductImage string `json:"ProductImage" sql:"ProductImage" imooc:"ProductImage"`
	ProductUrl   string `json:"ProductUrl" sql:"ProductUrl" imooc:"ProductUrl"`
}
