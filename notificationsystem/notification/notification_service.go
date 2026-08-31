package notification

import (
	"sync"
)

type notificationService struct {
	notifications []Notification
	observable    Observable
}

var (
	instance *notificationService
	once     sync.Once
)

func newNotificationService(observable Observable) *notificationService {
	return &notificationService{
		observable: observable,
	}
}

func GetNotificationServiceInstance(observable Observable) *notificationService {
	once.Do(func() {
		instance = newNotificationService(observable)
	})
	return instance
}

func (n *notificationService) SendNotification(notification Notification) {
	n.notifications = append(n.notifications, notification)
	n.observable.SetNotification(notification)
}
