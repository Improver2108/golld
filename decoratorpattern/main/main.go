package main

import (
	"fmt"

	"github.com/improver2108/golld/decoratorpattern/character"
)

func main() {
	var mario character.Character
	mario = character.NewMario()
	fmt.Println("Basic character:", mario.GetAbilities())

	mario = character.NewHeightUp(mario)
	fmt.Println("After Height up:", mario.GetAbilities())

	mario = character.NewGunPowerUp(mario)
	fmt.Println("After Gun power:", mario.GetAbilities())

	mario = character.NewStarPowerUp(mario)
	fmt.Println("After star power:", mario.GetAbilities())

}
