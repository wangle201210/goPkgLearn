package main

import (
	_ "github.com/go-sql-driver/mysql"

	_ "beegoTest/routers"

	"github.com/astaxie/beego"
)

func main() {
	//dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", beego.AppConfig.String("db_username"), beego.AppConfig.String("db_password"), beego.AppConfig.String("db_host"), beego.AppConfig.String("db_port"), beego.AppConfig.String("db_database"))
	//orm.Debug = true
	//err := orm.RegisterDataBase("default", "mysql", dataSource)
	//if err != nil {
	//	panic(err)
	//}
	beego.Run(":8888")
}
