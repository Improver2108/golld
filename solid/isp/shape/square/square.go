package square

type Square struct {
	side float64
}

func Init(side float64) *Square {
	return &Square{side: side}
}

func (s *Square) Area() float64 {
	return s.side * s.side
}
