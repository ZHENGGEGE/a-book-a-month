package sysinit

import (
  "github.com/astaxie/beego"
  _ "github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
  _ "github.com/astaxie/beego/orm"
  _ "github.com/go-sql-driver/mysql"
)

func dbinit(alias string)  {
  dbAlias:=alias // default
  if "w"==alias||"default"==alias||len(alias)<=0 {
    dbAlias = alias
    alias = "w"
  }

  //数据库名称
  dbName:=beego.AppConfig.String("db_"+alias+"_database")
  //数据库用户名
  dbUser:=beego.AppConfig.String("db_"+alias+"_username")
  //数据库密码
  dbPwd:=beego.AppConfig.String("db_"+alias+"_password")
  //数据库ip
  dbHost:=beego.AppConfig.String("db_"+alias+"_host")
  //数据库端口
  dbPort:=beego.AppConfig.String("db_"+alias+"_port")
  //root:135246@tcp(127.0.0.1:3306)/book?charset=utf8
  orm.RegisterDataBase(dbAlias,"mysql",dbUser+dbPwd+":"+"@tcp("+dbHost+":"+dbPort+")"+dbName+"charset=utf8",30)

  isDev:= "dev"==beego.AppConfig.String("runmode")

  if "w"==alias {
    orm.RunSyncdb("default",false,isDev)
  }
}
