package storage

import (
	"fmt"

	"github.com/improver2108/golld/solid/ocp/violated/shoppingcart"
)

type ShoppingCartStorage struct {
	cart shoppingcart.ShoppingCart
}

func Init(cart *shoppingcart.ShoppingCart) *ShoppingCartStorage {
	return &ShoppingCartStorage{cart: *cart}
}

func (s *ShoppingCartStorage) SaveToSQLDatabase() {
	fmt.Println("Saving shopping cart to database...")
}

func (s *ShoppingCartStorage) SaveToMongoDatabase() {
	fmt.Println("Saving shopping cart to database...")
}

func (s *ShoppingCartStorage) SaveToFile() {
	fmt.Println("Saving shopping cart to database...")
}
