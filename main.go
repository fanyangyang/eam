package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/fanyangyang/eam/models"
	_ "github.com/fanyangyang/eam/models"
	_ "github.com/fanyangyang/eam/routers"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "./data/test.db")
	orm.RunSyncdb("default", false, true)
}

func main() {
	models.NewOrm()
	beego.Run()
}
