package notification

import "fmt"

type NotificationStrategy interface {
	SendNotification(content string)
}

type notificationEngine struct {
	strategies []NotificationStrategy
	observable Observable
}

func NewNotificationEngine(observable Observable) *notificationEngine {
	return &notificationEngine{
		observable: observable,
	}
}

func (n *notificationEngine) AddNotificationStrategy(strategy NotificationStrategy) {
	n.strategies = append(n.strategies, strategy)
}

func (n *notificationEngine) Update() {
	if len(n.strategies) == 0 {
		fmt.Println("Please add some strategies to send notification")
		return
	}
	content := n.observable.GetNotificationContent()
	for _, strategy := range n.strategies {
		strategy.SendNotification(content)
	}
}
