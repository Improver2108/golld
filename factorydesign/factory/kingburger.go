package factory

import (
	"errors"

	"github.com/improver2108/golld/factorydesign/burger"
	"github.com/improver2108/golld/factorydesign/garlicbread"
)

type KingBurger struct{}

func NewKingBurger() BurgerFactory {
	return &KingBurger{}
}

func (b *KingBurger) CreateBurger(burgerType MealType) (Burger, error) {
	switch burgerType {
	case Basic:
		return burger.NewBasicWheatBurger(), nil
	case Standard:
		return burger.NewStandardWheatBurger(), nil
	default:
		return nil, errors.New("Invalid burger type")
	}
}

func (b *KingBurger) CreateGarlicBread(breadType MealType) (GarlicBread, error) {
	switch breadType {
	case Basic:
		return garlicbread.NewBasicWheatGarlicBread(), nil
	case Standard:
		return burger.NewStandardWheatBurger(), nil
	default:
		return nil, errors.New("Invalid burger type")
	}
}
