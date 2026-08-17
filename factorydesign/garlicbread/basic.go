package garlicbread

import "fmt"

type BasicGarlicBread struct{}

func NewBasicGarlicBread() *BasicGarlicBread {
	return &BasicGarlicBread{}
}

func (b *BasicGarlicBread) Prepare() {
	fmt.Println("Preparing Basic Garlic bread with bun, patty, and ketchup!")
}

type BasicWheatGarlicBread struct{}

func NewBasicWheatGarlicBread() *BasicWheatGarlicBread {
	return &BasicWheatGarlicBread{}
}

func (b *BasicWheatGarlicBread) Prepare() {
	fmt.Println("Preparing Basic Wheat Garlic bread with bun, wheat, patty, and ketchup!")
}
