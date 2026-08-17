package burger

import "fmt"

type BasicBurger struct{}

func NewBasicBurger() *BasicBurger {
	return &BasicBurger{}
}

func (b *BasicBurger) Prepare() {
	fmt.Println("Preparing Basic Burger with bun, patty, and ketchup!")
}

type BasicWheatBurger struct{}

func NewBasicWheatBurger() *BasicWheatBurger {
	return &BasicWheatBurger{}
}

func (b *BasicWheatBurger) Prepare() {
	fmt.Println("Preparing Basic Wheat Burger with bun, wheat, patty, and ketchup!")
}
