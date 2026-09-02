package main

import "github.com/improver2108/golld/facadedesign/facade"

func main() {
	computer := facade.NewComputerFacade()
	computer.StartComputer()
}
