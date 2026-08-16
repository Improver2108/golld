package main

import (
	"github.com/improver2108/golld/solid/dip/followed/persistant/mongodb"
	"github.com/improver2108/golld/solid/dip/followed/persistant/sql"
	"github.com/improver2108/golld/solid/dip/followed/user/service"
)

func main() {
	mySql := sql.Init()
	mongo := mongodb.Init()

	service1 := service.Init(mySql)
	service1.StoreUser("Raina")

	service2 := service.Init(mongo)
	service2.StoreUser("Anita")
}
