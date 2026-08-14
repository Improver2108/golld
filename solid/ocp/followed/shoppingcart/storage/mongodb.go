package storage

import (
	"fmt"

	"github.com/improver2108/golld/solid/ocp/followed/shoppingcart"
)

type MongoDb struct {
}

func InitMongoDB() *MongoDb {
	return &MongoDb{}
}

func (m *MongoDb) Save(cart *shoppingcart.ShoppingCart) {
	fmt.Println("Saving shopping cart to MongoDB...")
}
