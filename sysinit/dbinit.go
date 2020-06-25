package sysinit

import (
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
  _ "github.com/go-sql-driver/mysql"
)

func dbinit(aliases ...string)  {
  //如果是开发模式，则显示命令信息
  isDev:= "dev"==beego.AppConfig.String("runmode")

  if len(aliases)>0{
    for _,alias:=range aliases{
      registerDatabase(alias)
      //主库自动建表
      if "w"==alias {
        orm.RunSyncdb("default",false,isDev)
      }
    }
  }else{
    registerDatabase("w")
    orm.RunSyncdb("default",false,isDev)
  }


}

func registerDatabase(alias string)  {
  if len(alias)<=0 {
    return
  }
  //连接名称
  dbAlias:=alias // default
  if "w"==alias||"default"==alias {
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
  //root:edd%L0+h3tjj@tcp(127.0.0.1:3306)/book?charset=utf8
  orm.RegisterDataBase(dbAlias,"mysql",dbUser+dbPwd+":"+"@tcp("+dbHost+":"+dbPort+")"+dbName+"charset=utf8",30)
}
