package factory

import (
	"errors"

	"github.com/improver2108/golld/factorydesign/burger"
	"github.com/improver2108/golld/factorydesign/garlicbread"
)

type SinghBurger struct{}

func NewSinghBurger() BurgerFactory {
	return &SinghBurger{}
}

func (b *SinghBurger) CreateBurger(burgerType MealType) (Burger, error) {
	switch burgerType {
	case Basic:
		return burger.NewBasicBurger(), nil
	case Standard:
		return burger.NewStandardBurger(), nil
	case Premium:
		return burger.NewPremiumBurger(), nil
	default:
		return nil, errors.New("Invalid burger type")
	}
}

func (b *SinghBurger) CreateGarlicBread(breadType MealType) (GarlicBread, error) {
	switch breadType {
	case Basic:
		return garlicbread.NewBasicGarlicBread(), nil
	case Standard:
		return garlicbread.NewStandardGarlicBread(), nil
	default:
		return nil, errors.New("Invalid burger type")
	}
}
