package notification

type Observer interface {
	Update()
}

type Observable interface {
	AddObserver(obs Observer)
	RemoveObserver(obs Observer)
	GetNotificationContent() string
	SetNotification(notification Notification)
}

type Notification interface {
	GetContent() string
}
