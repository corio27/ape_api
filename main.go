package main

import (
<<<<<<< HEAD
	_ "github.com/corio27/pae_api/routers"
=======
	_ "pae_api/routers"
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
<<<<<<< HEAD
	orm.RegisterDataBase("default", "mysql", "root:esguerra@tcp(127.0.0.1:3306)/pma_pae")
=======
	orm.RegisterDataBase("default", "mysql", "docencia:D0c3Nc142017@tcp(10.20.0.88:3306)/pma_pae")
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

