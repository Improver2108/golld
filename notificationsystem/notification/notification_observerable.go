package notification

type notificationObservable struct {
	observers           map[Observer]struct{}
	currentNotification Notification
}

func NewNotificationObservable() *notificationObservable {
	return &notificationObservable{
		observers: make(map[Observer]struct{}),
	}
}

func (n *notificationObservable) AddObserver(obs Observer) {
	n.observers[obs] = struct{}{}
}

func (n *notificationObservable) RemoveObserver(obs Observer) {
	delete(n.observers, obs)
}

func (n *notificationObservable) notifyObservers() {
	for obs := range n.observers {
		obs.Update()
	}
}

func (n *notificationObservable) GetNotificationContent() string {
	return n.currentNotification.GetContent()
}

func (n *notificationObservable) SetNotification(notification Notification) {
	n.currentNotification = notification
	n.notifyObservers()
}
