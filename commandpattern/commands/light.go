package commands

type light interface {
	On()
	Off()
}

type lightCommand struct {
	light light
}

func NewLightCommand(l light) *lightCommand {
	return &lightCommand{light: l}
}

func (c *lightCommand) Execute() {
	c.light.On()
}

func (c *lightCommand) Undo() {
	c.light.Off()
}
