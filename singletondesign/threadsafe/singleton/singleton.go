package singleton

import (
	"fmt"
	"sync"
)

type simpleSingleton struct {
	val int
}

var (
	instance *simpleSingleton
	once     sync.Once
)

func GetInstance() *simpleSingleton {
	once.Do(func() {
		fmt.Println("Singleton constructor called. New struct created.")
		instance = &simpleSingleton{}
	})
	return instance
}
