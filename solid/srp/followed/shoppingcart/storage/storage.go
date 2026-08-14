package storage

import (
	"fmt"

	"github.com/improver2108/golld/solid/srp/followed/shoppingcart"
)

type ShoppingCartStorage struct {
	cart shoppingcart.ShoppingCart
}

func Init(cart *shoppingcart.ShoppingCart) *ShoppingCartStorage {
	return &ShoppingCartStorage{cart: *cart}
}

func (s *ShoppingCartStorage) SaveToDatabase() {
	fmt.Println("Saving shopping cart to database...")
}
