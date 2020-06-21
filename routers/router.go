package routers

import (
	"suzheng/a-book-a-month/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
