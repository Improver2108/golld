package nowalk

import "fmt"

type NoWalk struct{}

func Init() *NoWalk {
	return &NoWalk{}
}

func (w *NoWalk) Walk() {
	fmt.Println("Cannot walk. Sorry mate!")
}
