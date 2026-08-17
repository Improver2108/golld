package factory

type Burger interface {
	Prepare()
}

type GarlicBread interface {
	Prepare()
}

type BurgerFactory interface {
	CreateBurger(burgerType MealType) (Burger, error)
	CreateGarlicBread(breadType MealType) (GarlicBread, error)
}
