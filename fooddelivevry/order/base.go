package order

import (
	"errors"

	"github.com/improver2108/golld/fooddelivevry/menuitem"
	"github.com/improver2108/golld/fooddelivevry/restaurant"
	"github.com/improver2108/golld/fooddelivevry/user"
)

type Payment interface {
	Pay(amount float32)
}

type baseOrder struct {
	id         int
	restaurant *restaurant.Restaurant
	items      []*menuitem.MenuItem
	payment    Payment
	user       *user.User
	total      float32
}

func newBaseOrder(r *restaurant.Restaurant, i []*menuitem.MenuItem, p Payment, u *user.User) *baseOrder {
	var total float32
	for _, item := range i {
		total += item.GetPrice()
	}
	return &baseOrder{restaurant: r, items: i, payment: p, user: u, total: total}
}

func (o *baseOrder) GetOrderId() int {
	return o.id
}

func (o *baseOrder) GetUser() *user.User {
	return o.user
}

func (o *baseOrder) ProcessPayment() error {
	if o.payment == nil {
		return errors.New("Please choose a payment mode first")
	}
	o.payment.Pay(o.total)
	return nil
}
