package controllers

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris/v12"
)

type IBController interface {
	InitController(app *iris.Application, controllerName string)
	RegisterValue() string
	RegisterController()
}
type BController struct {
	Ctx            iris.Context
	Db             *gorm.DB
	ControllerName string
}

func (b *BController) InitController(app *iris.Application, controllerName string) {

}
