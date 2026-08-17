package burger

import "fmt"

type PremiumBurger struct{}

func NewPremiumBurger() *PremiumBurger {
	return &PremiumBurger{}
}

func (b *PremiumBurger) Prepare() {
	fmt.Println("Preparing Premium Burger with gourmet bun, premium patty, cheese, lettuce, and secret sauce!")
}
