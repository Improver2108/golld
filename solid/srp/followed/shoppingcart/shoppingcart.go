package shoppingcart

import "github.com/improver2108/golld/solid/srp/followed/product"

type ShoppingCart struct {
	products []product.Product
}

func Init() *ShoppingCart {
	return &ShoppingCart{}
}

func (s *ShoppingCart) AddProduct(p *product.Product) {
	s.products = append(s.products, *p)
}

func (s *ShoppingCart) GetProducts() []product.Product {
	return s.products
}

func (s *ShoppingCart) CalculateTotal() float32 {
	var total float32
	for _, p := range s.products {
		total += p.Price
	}
	return total
}
