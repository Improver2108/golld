package main

import (
	"fmt"

	"github.com/improver2108/golld/singletondesign/simple/singleton"
)

func main() {
	s1 := singleton.GetInstance()
	s2 := singleton.GetInstance()
	fmt.Println(s1 == s2)
}
