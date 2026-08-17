package main

import (
	"fmt"

	"github.com/improver2108/golld/factorydesign/factory"
)

func main() {
	myFactory := factory.NewSinghBurger()
	burger, err := myFactory.CreateBurger(factory.Basic)
	if err != nil {
		fmt.Println("no object created")
		return
	}
	burger.Prepare()
	garlic, err := myFactory.CreateGarlicBread(factory.Basic)
	if err != nil {
		fmt.Println("no object created")
		return
	}
	garlic.Prepare()
}
