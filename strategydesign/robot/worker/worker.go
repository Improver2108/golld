package worker

import (
	"fmt"

	"github.com/improver2108/golld/strategydesign/robot"
)

type WorkerRobot struct {
	*robot.Robot
}

func Init(w robot.WalkableRobot, t robot.TalkableRobot) *WorkerRobot {
	return &WorkerRobot{
		Robot: robot.Init(w, t),
	}
}

func (r *WorkerRobot) Projection() {
	fmt.Println("Displaying friendly Worker features...")
}
