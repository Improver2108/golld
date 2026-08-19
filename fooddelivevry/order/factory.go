package order

import (
	"github.com/improver2108/golld/fooddelivevry/cart"
	"github.com/improver2108/golld/fooddelivevry/menuitem"
	"github.com/improver2108/golld/fooddelivevry/restaurant"
	"github.com/improver2108/golld/fooddelivevry/user"
)

type instantOrderFactory struct{}

func NewInstantOrderFactory() *instantOrderFactory {
	return &instantOrderFactory{}
}

func (f *instantOrderFactory) CreateOrder(user *user.User, cart *cart.Cart, rest *restaurant.Restaurant, menuItems []*menuitem.MenuItem, payment Payment, orderType string) Order {
	base := newBaseOrder(rest, menuItems, payment, user)
	var order Order
	switch orderType {
	case "delivery":
		order = NewDeliveryOrder(base, user.GetLocation())
	default:
		order = NewPickupOrder(base, rest.GetLocation())
	}
	return order
}
