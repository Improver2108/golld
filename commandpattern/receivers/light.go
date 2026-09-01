package receivers

import "fmt"

type light struct{}

func NewLight() *light {
	return &light{}
}

func (r *light) On() {
	fmt.Println("Light is on")
}

func (r *light) Off() {
	fmt.Println("Light is off")
}
