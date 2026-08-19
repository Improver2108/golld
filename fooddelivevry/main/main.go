package main

import (
	"fmt"

	"github.com/improver2108/golld/fooddelivevry/app"
	"github.com/improver2108/golld/fooddelivevry/payment"
	"github.com/improver2108/golld/fooddelivevry/user"
)

func main() {
	app := app.NewApp()
	user := user.Init("delhi", "yash")
	fmt.Println("User:", user.GetName(), "is active.")

	restaurantLists := app.SearchRestaurants("delhi")

	if len(restaurantLists) == 0 {
		fmt.Println("No restaurant found!!")
		return
	}
	fmt.Println("Found Restaurant:")
	for _, rest := range restaurantLists {
		fmt.Printf(" - %s\n", rest.GetName())
	}

	app.SelectRestaurant(user, restaurantLists[0])
	fmt.Println("Selected restaurant :", restaurantLists[0].GetName())

	app.AddToCart(user, "P1")
	app.AddToCart(user, "P2")

	app.PrintUserCart(user)

	order := app.CheckoutInstant(user, "Delivery", payment.NewUPIPayment("870036121"))

	app.PayForOrder(order)
}
