package main

import (
	"github.com/improver2108/golld/solid/ocp/violated/product"
	"github.com/improver2108/golld/solid/ocp/violated/shoppingcart"
	"github.com/improver2108/golld/solid/ocp/violated/shoppingcart/printer"
	"github.com/improver2108/golld/solid/ocp/violated/shoppingcart/storage"
)

func main() {
	cart := shoppingcart.Init()

	cart.AddProduct(product.Init("Laptop", 50000))
	cart.AddProduct(product.Init("Mouse", 8000))

	printer := printer.Init(cart)
	printer.PrintInvoice()

	db := storage.Init(cart)

	db.SaveToSQLDatabase()
	db.SaveToFile()
	db.SaveToMongoDatabase()
}
