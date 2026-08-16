package current

import (
	"errors"
	"fmt"
)

type CurrentAccount struct {
	balance float32
}

func Init() *CurrentAccount {
	return &CurrentAccount{}
}

func (s *CurrentAccount) Deposit(amount float32) error {
	s.balance += amount
	fmt.Println("Deposited: ", amount, " in Current Account. New Balance: ", s.balance)
	return nil
}

func (s *CurrentAccount) Withdraw(amount float32) error {
	if s.balance >= amount {
		s.balance -= amount
		fmt.Println("Withdrawn: ", amount, " from Current Account. New Balance: ", s.balance)
		return nil
	}
	return errors.New("Insufficient funds in Current Account!")
}
