package shoppingcart

import (
	"fmt"

	"github.com/improver2108/golld/solid/srp/violated/product"
)

type shoppingCart struct {
	products []product.Product
}

func NewShoppingCart() *shoppingCart {
	return &shoppingCart{}
}

func (s *shoppingCart) AddProduct(p *product.Product) {
	s.products = append(s.products, *p)
}

func (s *shoppingCart) GetProducts() []product.Product {
	return s.products
}

func (s *shoppingCart) calculateTotal() float32 {
	var total float32
	for _, p := range s.products {
		total += p.Price
	}
	return total
}

func (s *shoppingCart) PrintInvoice() {
	fmt.Println("Shopping cart invoice:")
	for _, p := range s.products {
		fmt.Println(p.Name, " - Rs ", p.Price)
	}
	fmt.Println("Total: Rs ", s.calculateTotal())
}

func (s *shoppingCart) SaveToDatabase() {
	fmt.Println("Saving shopping cart to database...")
}
