package main

import "github.com/improver2108/golld/solid/dip/violated/user/service"

func main() {
	service := service.Init()
	service.StoreUserToMongo("Raina")
	service.StoreUserToSQL("Anita")
}
