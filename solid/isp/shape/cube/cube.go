package cube

type Cube struct {
	side float64
}

func Init(s float64) *Cube {
	return &Cube{side: s}
}

func (c *Cube) Area() float64 {
	return 6 * c.side * c.side
}

func (c *Cube) Volume() float64 {
	return c.side * c.side * c.side
}
