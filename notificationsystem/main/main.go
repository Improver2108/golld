package main

import "github.com/improver2108/golld/notificationsystem/notification"

func main() {
	notificationObservable := notification.NewNotificationObservable()
	notificationService := notification.GetNotificationServiceInstance(notificationObservable)

	logger := notification.NewLogger(notificationObservable)
	notificationObservable.AddObserver(logger)

	notificationEngine := notification.NewNotificationEngine(notificationObservable)
	notificationEngine.AddNotificationStrategy(notification.NewEmail("random.person@gmail.com"))
	notificationEngine.AddNotificationStrategy(notification.NewSMS("+91 9876543210"))
	notificationObservable.AddObserver(notificationEngine)

	var noti notification.Notification
	noti = notification.NewSimpleNotification("Your order has been shipped")
	noti = notification.NewTimestampDecorator(noti)
	noti = notification.NewSignatureDecorator(noti, "Customer Care")

	notificationService.SendNotification(noti)
}
