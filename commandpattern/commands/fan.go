package commands

type fan interface {
	On()
	Off()
}

type fanCommand struct {
	fan fan
}

func NewFanCommand(f fan) *fanCommand {
	return &fanCommand{fan: f}
}

func (c *fanCommand) Execute() {
	c.fan.On()
}

func (c *fanCommand) Undo() {
	c.fan.Off()
}
