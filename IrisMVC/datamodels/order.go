package datamodels

type Order struct {
	ID          int64 `sql:"ID" gorm:"ID"`
	UserID      int64 `sql:"UserID" gorm:"UserID"`
	ProductID   int64 `sql:"ProductID" gorm:"ProductID"`
	OrderStatus int64 `sql:"OrderStatus" gorm:"OrderStatus"`
}

const (
	OrderWait    = iota
	OrderSuccess //1
	OrderFailed  //2
)
