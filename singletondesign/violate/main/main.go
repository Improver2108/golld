package main

import (
	"fmt"

	"github.com/improver2108/golld/singletondesign/violate/singleton"
)

func main() {
	s1 := singleton.Init()
	s2 := singleton.Init()
	fmt.Println(s1 == s2)
}
