package singleton

type simpleSingleton struct {
	val int
}

var instance *simpleSingleton

func init() {
	instance = &simpleSingleton{}
}

func GetInstance() *simpleSingleton {
	return instance
}
