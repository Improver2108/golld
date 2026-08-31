package notification

import "fmt"

type logger struct {
	notificationObservable Observable
}

func NewLogger(observable Observable) *logger {
	return &logger{
		notificationObservable: observable,
	}
}

func (l *logger) Update() {
	fmt.Println("Logging new notification : \n", l.notificationObservable.GetNotificationContent())
}
