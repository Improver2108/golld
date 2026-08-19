package user

import "github.com/improver2108/golld/fooddelivevry/cart"

type User struct {
	id       int
	location string
	name     string
	cart     *cart.Cart
}

func Init(location, name string) *User {
	return &User{location: location, name: name, cart: cart.Init()}
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}
func (u *User) GetLocation() string {
	return u.location
}

func (u *User) SetLocation(location string) {
	u.location = location
}

func (u *User) GetCart() *cart.Cart {
	return u.cart
}
