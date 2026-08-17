package singleton

import "fmt"

type simpleSingleton struct {
	val int
}

var instance *simpleSingleton

func GetInstance() *simpleSingleton {
	if instance == nil {
		fmt.Println("Singleton constructor called. New struct created.")
		instance = &simpleSingleton{}
	}
	return instance
}
