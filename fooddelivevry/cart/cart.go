package cart

import (
	"github.com/improver2108/golld/fooddelivevry/menuitem"
	"github.com/improver2108/golld/fooddelivevry/restaurant"
)

type Cart struct {
	restaurant *restaurant.Restaurant
	items      []*menuitem.MenuItem
}

func Init() *Cart {
	return &Cart{}
}

func (c *Cart) AddItem(item *menuitem.MenuItem) {
	c.items = append(c.items, item)
}

func (c *Cart) GetItems() []*menuitem.MenuItem {
	return c.items
}

func (c *Cart) GetTotalCost() float32 {
	var sum float32
	for _, item := range c.items {
		sum += item.GetPrice()
	}
	return sum
}

func (c *Cart) IsEmpty() bool {
	return c.restaurant == nil || len(c.items) == 0
}

func (c *Cart) Clear() {
	c.items = nil
	c.restaurant = nil
}

func (c *Cart) GetRestaurant() *restaurant.Restaurant {
	return c.restaurant
}

func (c *Cart) SetRestaurant(r *restaurant.Restaurant) {
	c.restaurant = r
}
