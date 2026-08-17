package normalwalk

import "fmt"

type NormalWalk struct{}

func Init() *NormalWalk {
	return &NormalWalk{}
}

func (w *NormalWalk) Walk() {
	fmt.Println("Normally walk")
}
