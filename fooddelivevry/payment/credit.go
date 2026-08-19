package payment

import "fmt"

type creditPayment struct {
	cardNumber string
}

func NewCreditPayment(card string) *creditPayment {
	return &creditPayment{cardNumber: card}
}

func (p *creditPayment) Pay(amount float32) {
	fmt.Println("Paid ₹", amount, " using Credit Card (", p.cardNumber, ")")
}
