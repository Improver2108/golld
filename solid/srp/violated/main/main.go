package main

import (
	"github.com/improver2108/golld/solid/srp/violated/product"
	"github.com/improver2108/golld/solid/srp/violated/shoppingcart"
)

func main() {
	cart := shoppingcart.NewShoppingCart()

	cart.AddProduct(product.NewProduct("Laptop", 50000.23))
	cart.AddProduct(product.NewProduct("Mouse", 2000))

	cart.PrintInvoice()
	cart.SaveToDatabase()
}
