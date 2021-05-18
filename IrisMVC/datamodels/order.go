package datamodels

type Order struct {
	ID          int64 `sql:"ID""`
	UserID      int64 `sql:"UserID""`
	ProductID   int64 `sql:"ProductID"`
	OrderStatus int64 `sql:"OrderStatus"`
}

const (
	OrderWait    = iota
	OrderSuccess //1
	OrderFailed  //2
)
