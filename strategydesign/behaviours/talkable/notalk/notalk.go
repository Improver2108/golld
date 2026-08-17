package notalk

import "fmt"

type NoTalk struct{}

func Init() *NoTalk {
	return &NoTalk{}
}

func (t *NoTalk) Talk() {
	fmt.Println("Not Talking")
}
