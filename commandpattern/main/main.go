package main

import (
	"fmt"

	"github.com/improver2108/golld/commandpattern/commands"
	"github.com/improver2108/golld/commandpattern/receivers"
	"github.com/improver2108/golld/commandpattern/remote"
)

func main() {
	livingRoomLight := receivers.NewLight()
	ceilingFan := receivers.NewFan()

	remote := remote.NewRemoteController()

	remote.SetCommand(0, commands.NewLightCommand(livingRoomLight))
	remote.SetCommand(1, commands.NewFanCommand(ceilingFan))

	fmt.Println("Toggling light button 0 ----")
	remote.PressButton(0)
	remote.PressButton(0)

	fmt.Println("Toggling fan button 1 ----")
	remote.PressButton(1)
	remote.PressButton(1)
}
