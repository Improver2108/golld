package normaltalk

import "fmt"

type NormalTalk struct{}

func Init() *NormalTalk {
	return &NormalTalk{}
}

func (t *NormalTalk) Talk() {
	fmt.Println("Talking normally")
}
