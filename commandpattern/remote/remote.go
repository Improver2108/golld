package remote

import "fmt"

type Command interface {
	Execute()
	Undo()
}

const numButtons = 4

type remoteController struct {
	buttons       []Command
	buttonPressed []bool
}

func NewRemoteController() *remoteController {
	return &remoteController{
		buttons:       make([]Command, numButtons),
		buttonPressed: make([]bool, numButtons),
	}
}

func (r *remoteController) SetCommand(idx int, cmd Command) {
	if idx < 0 && idx >= numButtons {
		return
	}
	r.buttons[idx] = cmd
	r.buttonPressed[idx] = false
}

func (r *remoteController) PressButton(idx int) {
	if idx < 0 && idx >= numButtons && r.buttons[idx] == nil {
		fmt.Println("No command assigned to button")
		return
	}
	if r.buttonPressed[idx] {
		r.buttons[idx].Undo()
	} else {
		r.buttons[idx].Execute()
	}
	r.buttonPressed[idx] = !r.buttonPressed[idx]
}
