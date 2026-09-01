package receivers

import "fmt"

type fan struct{}

func NewFan() *fan {
	return &fan{}
}

func (r *fan) On() {
	fmt.Println("Fan is on")
}

func (r *fan) Off() {
	fmt.Println("Fan is off")
}
