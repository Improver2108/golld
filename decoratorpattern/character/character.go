package character

type Character interface {
	GetAbilities() string
}

type characterDecorator struct {
	character Character
}

func newCharacter(c Character) *characterDecorator {
	return &characterDecorator{character: c}
}

type heightUp struct {
	*characterDecorator
}

func NewHeightUp(c Character) *heightUp {
	return &heightUp{
		characterDecorator: newCharacter(c),
	}
}

func (c *heightUp) GetAbilities() string {
	return c.character.GetAbilities() + " with HeightUp"
}

type gunPowerUp struct {
	*characterDecorator
}

func NewGunPowerUp(c Character) *gunPowerUp {
	return &gunPowerUp{
		characterDecorator: newCharacter(c),
	}
}

func (c *gunPowerUp) GetAbilities() string {
	return c.character.GetAbilities() + " with Gun"
}

type starPowerUp struct {
	*characterDecorator
}

func NewStarPowerUp(c Character) *starPowerUp {
	return &starPowerUp{
		characterDecorator: newCharacter(c),
	}
}

func (c *starPowerUp) GetAbilities() string {
	return c.character.GetAbilities() + " with Star"
}
