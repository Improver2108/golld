package saving

import (
	"errors"
	"fmt"
)

type SavingAccount struct {
	balance float32
}

func Init() *SavingAccount {
	return &SavingAccount{}
}

func (s *SavingAccount) Deposit(amount float32) error {
	s.balance += amount
	fmt.Println("Deposited: ", amount, " in Savings Account. New Balance: ", s.balance)
	return nil
}

func (s *SavingAccount) Withdraw(amount float32) error {
	if s.balance >= amount {
		s.balance -= amount
		fmt.Println("Withdrawn: ", amount, " from Savings Account. New Balance: ", s.balance)
		return nil
	}
	return errors.New("Insufficient funds in Savings Account!")
}
