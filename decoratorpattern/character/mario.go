package character

type mario struct {
}

func NewMario() *mario {
	return &mario{}
}

func (c *mario) GetAbilities() string {
	return "mario"
}
