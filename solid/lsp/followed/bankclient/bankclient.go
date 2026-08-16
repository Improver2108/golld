package bankclient

type DepositOnlyAccount interface {
	Deposit(amount float32) error
}

type WithdrawableAccount interface {
	DepositOnlyAccount
	Withdraw(amount float32) error
}

type BankClient struct {
	depositOnlyAccounts  []DepositOnlyAccount
	withdrawableAccounts []WithdrawableAccount
}

func Init(depositOnlyAccounts []DepositOnlyAccount, withdrawableAccounts []WithdrawableAccount) *BankClient {
	return &BankClient{depositOnlyAccounts, withdrawableAccounts}
}

func (b *BankClient) ProcessTransactions() {
	for _, acc := range b.depositOnlyAccounts {
		acc.Deposit(5000)
	}
	for _, acc := range b.withdrawableAccounts {
		acc.Deposit(1000)
		acc.Withdraw(500)
	}
}
