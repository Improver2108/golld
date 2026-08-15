package bankclient

import "fmt"

type Account interface {
	Withdraw(amount float32) error
	Deposit(amount float32) error
}

type BankClient struct {
	accounts []Account
}

func Init() *BankClient {
	return &BankClient{}
}

func (b *BankClient) ProcessTransactions() {
	for _, acc := range b.accounts {
		acc.Deposit(1000)
		if err := acc.Withdraw(500); err != nil {
			error := fmt.Errorf("Error:%w", err)
			fmt.Println(error)
		}
	}
}
