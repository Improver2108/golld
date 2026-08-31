package notification

type simpleNotification struct {
	text string
}

func NewSimpleNotification(msg string) *simpleNotification {
	return &simpleNotification{
		text: msg,
	}
}

func (n *simpleNotification) GetContent() string {
	return n.text
}
