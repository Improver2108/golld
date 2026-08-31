package notification

type timestampDecorator struct {
	notification Notification
}

func NewTimestampDecorator(n Notification) *timestampDecorator {
	return &timestampDecorator{
		notification: n,
	}
}

func (n *timestampDecorator) GetContent() string {
	return "[2025-04-13 14:22:00] " + n.notification.GetContent()
}

type signatureDecorator struct {
	notification Notification
	signature    string
}

func NewSignatureDecorator(n Notification, sign string) *signatureDecorator {
	return &signatureDecorator{
		notification: n,
		signature:    sign,
	}
}

func (n *signatureDecorator) GetContent() string {
	return n.notification.GetContent() + "\n-- " + n.signature + "\n\n"
}
