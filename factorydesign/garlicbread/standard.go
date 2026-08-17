package garlicbread

import "fmt"

type StandardGarlic struct{}

func NewStandardGarlicBread() *StandardGarlic {
	return &StandardGarlic{}
}

func (b *StandardGarlic) Prepare() {
	fmt.Println("Preparing Standard garlic bread with bun, patty, cheese, and lettuce!")
}

type StandardWheatGarlicBread struct{}

func NewStandardWheatGarlicBread() *StandardWheatGarlicBread {
	return &StandardWheatGarlicBread{}
}

func (b *StandardWheatGarlicBread) Prepare() {
	fmt.Println("Preparing Standard wheat Burger with bun, patty, cheese, and lettuce!")
}
