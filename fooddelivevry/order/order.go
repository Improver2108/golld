package order

import (
	"github.com/improver2108/golld/fooddelivevry/cart"
	"github.com/improver2108/golld/fooddelivevry/menuitem"
	"github.com/improver2108/golld/fooddelivevry/restaurant"
	"github.com/improver2108/golld/fooddelivevry/user"
)

type Order interface {
	GetOrderId() int
	GetAddress() string
	GetType() string
	ProcessPayment() error
	GetUser() *user.User
}

type OrderFactory interface {
	CreateOrder(user *user.User, cart *cart.Cart, rest *restaurant.Restaurant, menuItems []*menuitem.MenuItem, payment Payment, orderType string) Order
}
