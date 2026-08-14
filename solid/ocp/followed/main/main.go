package main

import (
	"github.com/improver2108/golld/solid/ocp/followed/product"
	"github.com/improver2108/golld/solid/ocp/followed/shoppingcart"
	"github.com/improver2108/golld/solid/ocp/followed/shoppingcart/printer"
	"github.com/improver2108/golld/solid/ocp/followed/shoppingcart/storage"
)

type Database interface {
	Save(shoppingcart.ShoppingCart)
}

func main() {
	cart := shoppingcart.Init()

	cart.AddProduct(product.Init("Laptop", 50000))
	cart.AddProduct(product.Init("Mouse", 8000))

	printer := printer.Init(cart)
	printer.PrintInvoice()

	sqlDb := storage.InitSQL()
	file := storage.InitFile()
	mongo := storage.InitMongoDB()

	sqlDb.Save(cart)
	file.Save(cart)
	mongo.Save(cart)
}

