package companion

import (
	"fmt"

	"github.com/improver2108/golld/strategydesign/robot"
)

type CompanionRobot struct {
	*robot.Robot
}

func Init(w robot.WalkableRobot, t robot.TalkableRobot) *CompanionRobot {
	return &CompanionRobot{
		Robot: robot.Init(w, t),
	}
}

func (r *CompanionRobot) Projection() {
	fmt.Println("Displaying friendly companion features...")
}
