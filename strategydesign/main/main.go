package main

import (
	"github.com/improver2108/golld/strategydesign/behaviours/talkable/normaltalk"
	"github.com/improver2108/golld/strategydesign/behaviours/talkable/notalk"
	"github.com/improver2108/golld/strategydesign/behaviours/walk/normalwalk"
	"github.com/improver2108/golld/strategydesign/behaviours/walk/nowalk"
	"github.com/improver2108/golld/strategydesign/robot/companion"
	"github.com/improver2108/golld/strategydesign/robot/worker"
)

func main() {
	robot1 := companion.Init(normalwalk.Init(), normaltalk.Init())
	robot1.Walk()
	robot1.Talk()
	robot1.Projection()

	robot2 := worker.Init(nowalk.Init(), notalk.Init())
	robot2.Walk()
	robot2.Talk()
	robot2.Projection()
}
