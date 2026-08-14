package storage

import (
	"fmt"

	"github.com/improver2108/golld/solid/ocp/followed/shoppingcart"
)

type File struct {
}

func InitFile() *File {
	return &File{}
}

func (f *File) Save(cart *shoppingcart.ShoppingCart) {
	fmt.Println("Saving shopping cart to a file...")
}
