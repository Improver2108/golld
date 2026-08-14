package printer

import (
	"fmt"

	"github.com/improver2108/golld/solid/ocp/followed/shoppingcart"
)

type ShoppingCartPrinter struct {
	cart shoppingcart.ShoppingCart
}

func Init(cart *shoppingcart.ShoppingCart) *ShoppingCartPrinter {
	return &ShoppingCartPrinter{cart: *cart}
}

func (s *ShoppingCartPrinter) PrintInvoice() {
	fmt.Println("Shopping cart invoice:")
	for _, p := range s.cart.GetProducts() {
		fmt.Println(p.Name, " - Rs ", p.Price)
	}
	fmt.Println("Total: Rs ", s.cart.CalculateTotal())
}
