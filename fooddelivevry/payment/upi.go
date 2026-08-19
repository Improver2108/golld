package payment

import "fmt"

type upiPayment struct {
	mobile string
}

func NewUPIPayment(mob string) *upiPayment {
	return &upiPayment{mobile: mob}
}

func (p *upiPayment) Pay(amount float32) {
	fmt.Println("Paid ₹", amount, " using UPI (", p.mobile, ")")
}
