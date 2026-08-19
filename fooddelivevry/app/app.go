package app

import (
	"fmt"

	"github.com/improver2108/golld/fooddelivevry/menuitem"
	"github.com/improver2108/golld/fooddelivevry/order"
	"github.com/improver2108/golld/fooddelivevry/restaurant"
	"github.com/improver2108/golld/fooddelivevry/restaurant/manager"
	"github.com/improver2108/golld/fooddelivevry/user"
)

type App struct{}

func NewApp() *App {
	initializeApp()
	return &App{}
}

func initializeApp() {
	r1 := restaurant.Init("Bikaner", "Delhi")
	r1.AddMenuItem(menuitem.Init("P1", "Chole Bhature", 120))
	r1.AddMenuItem(menuitem.Init("P2", "Samosa", 15))

	r2 := restaurant.Init("Haldiram", "Kolkata")
	r2.AddMenuItem(menuitem.Init("P1", "Raj Kachori", 80))
	r2.AddMenuItem(menuitem.Init("P2", "Pav Bhaji", 100))
	r2.AddMenuItem(menuitem.Init("P3", "Dhokla", 50))

	r3 := restaurant.Init("Saravana Bhavan", "Chennai")
	r3.AddMenuItem(menuitem.Init("P1", "Masala Dosa", 90))
	r3.AddMenuItem(menuitem.Init("P2", "Idli Vada", 60))
	r3.AddMenuItem(menuitem.Init("P3", "Filter Coffee", 30))

	restaurantManager := manager.GetInstance()
	restaurantManager.AddRestaurant(r1)
	restaurantManager.AddRestaurant(r2)
	restaurantManager.AddRestaurant(r3)
}

func (a *App) SearchRestaurants(location string) []*restaurant.Restaurant {
	return manager.GetInstance().SearchByLocation(location)
}

func (a *App) SelectRestaurant(user *user.User, rest *restaurant.Restaurant) {
	cart := user.GetCart()
	cart.SetRestaurant(rest)
}

func (a *App) AddToCart(user *user.User, itemCode string) {
	restaurant := user.GetCart().GetRestaurant()
	if restaurant == nil {
		fmt.Println("Please select a restaurant first.")
		return
	}
	for _, item := range restaurant.GetMenu() {
		if item.GetCode() == itemCode {
			user.GetCart().AddItem(item)
			break
		}
	}
}

func (a *App) CheckoutInstant(user *user.User, orderType string, payment order.Payment) order.Order {
	return a.checkout(user, orderType, payment, order.NewInstantOrderFactory())
}

func (a *App) checkout(user *user.User, orderType string, payment order.Payment, orderFactory order.OrderFactory) order.Order {
	if user.GetCart().IsEmpty() {
		return nil
	}
	cart := user.GetCart()
	orderedRestaurant := cart.GetRestaurant()
	itemOrdered := cart.GetItems()

	order := orderFactory.CreateOrder(user, cart, orderedRestaurant, itemOrdered, payment, orderType)
	return order
}

func (a *App) PayForOrder(order order.Order) {
	if err := order.ProcessPayment(); err != nil {
		fmt.Println("Payment unsuccessful")
		return
	}
	fmt.Println("Payment completed")
	order.GetUser().GetCart().Clear()
}

func (a *App) PrintUserCart(user *user.User) {
	fmt.Println("Items in cart:")
	fmt.Println("-----------------")
	for _, item := range user.GetCart().GetItems() {
		fmt.Println(item.GetCode(), ":", item.GetName(), ":", item.GetPrice())
	}
	fmt.Println("-----------------")
	fmt.Println("Grand total : Rs", user.GetCart().GetTotalCost())
}
