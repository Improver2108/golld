package singleton

import "fmt"

type ViolateSingleton struct {
	val int
}

func Init() *ViolateSingleton {
	fmt.Println("Singleton constructor called. New struct created.")
	return &ViolateSingleton{}
}
