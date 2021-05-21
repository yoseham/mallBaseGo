package datamodels

type User struct {
	ID           int64  `sql:"ID" json:"ID" form:"ID"`
	NickName     string `sql:"NickName" json:"NickName" form:"NickName"`
	UserName     string `sql:"UserName" json:"UserName" form:"UserName"`
	HashPassword string `sql:"Password" json:"Password" form:"Password"`
}
