package main

import (
	"github.com/improver2108/golld/solid/lsp/violated/account/current"
	"github.com/improver2108/golld/solid/lsp/violated/account/saving"
	"github.com/improver2108/golld/solid/lsp/violated/bankclient"
)

func main() {
	accounts := []bankclient.Account{}
	accounts = append(
		accounts,
		saving.Init(),
		current.Init(),
		// fixed.Init(), //will throw error. we have to solve this using LSP
	)

	client := bankclient.Init(accounts)
	client.ProcessTransactions()
}
