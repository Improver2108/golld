package rectangle

type Rectangle struct {
	length float64
	width  float64
}

func Init(l, w float64) *Rectangle {
	return &Rectangle{l, w}
}

func (r *Rectangle) Area() float64 {
	return r.length * r.width
}
