package domain

type ISubscriber interface {
	Update()
}

type IChannel interface {
	Subscribe(subscriber ISubscriber)
	Unsubscribe(subscriber ISubscriber)
	GetName() string
}
