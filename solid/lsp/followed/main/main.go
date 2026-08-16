package main

import (
	"github.com/improver2108/golld/solid/lsp/followed/account/current"
	"github.com/improver2108/golld/solid/lsp/followed/account/fixed"
	"github.com/improver2108/golld/solid/lsp/followed/account/saving"
	"github.com/improver2108/golld/solid/lsp/followed/bankclient"
)

func main() {
	depositOnlyAccounts := []bankclient.DepositOnlyAccount{}
	depositOnlyAccounts = append(depositOnlyAccounts, fixed.Init())

	withdrawableAccounts := []bankclient.WithdrawableAccount{}
	withdrawableAccounts = append(withdrawableAccounts, current.Init(), saving.Init())

	client := bankclient.Init(depositOnlyAccounts, withdrawableAccounts)
	client.ProcessTransactions()
}
