package fixed

import (
	"fmt"
)

type FixedAccount struct {
	balance float32
}

func Init() *FixedAccount {
	return &FixedAccount{}
}

func (s *FixedAccount) Deposit(amount float32) error {
	s.balance += amount
	fmt.Println("Deposited: ", amount, " in Fixed Account. New Balance: ", s.balance)
	return nil
}
