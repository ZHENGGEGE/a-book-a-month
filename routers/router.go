package routers

import (
  "github.com/astaxie/beego"
  "suzheng/a-book-a-month/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
