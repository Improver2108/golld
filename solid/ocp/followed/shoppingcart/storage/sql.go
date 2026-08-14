package storage

import (
	"fmt"

	"github.com/improver2108/golld/solid/ocp/followed/shoppingcart"
)

type SQL struct {
}

func InitSQL() *SQL {
	return &SQL{}
}

func (s *SQL) Save(cart *shoppingcart.ShoppingCart) {
	fmt.Println("Saving shopping cart to SQL DB...")
}
