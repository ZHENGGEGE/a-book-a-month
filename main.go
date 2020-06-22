package main

import (
  "github.com/astaxie/beego"
  _ "suzheng/a-book-a-month/routers"
  _ "suzheng/a-book-a-month/sysinit"
)

func main() {
	beego.Run()
}

