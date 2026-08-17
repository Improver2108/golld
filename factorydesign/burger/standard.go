package burger

import "fmt"

type StandardBurger struct{}

func NewStandardBurger() *StandardBurger {
	return &StandardBurger{}
}

func (b *StandardBurger) Prepare() {
	fmt.Println("Preparing Standard Burger with bun, patty, cheese, and lettuce!")
}

type StandardWheatBurger struct{}

func NewStandardWheatBurger() *StandardWheatBurger {
	return &StandardWheatBurger{}
}

func (b *StandardWheatBurger) Prepare() {
	fmt.Println("Preparing Standard wheat Burger with bun, patty, cheese, and lettuce!")
}
