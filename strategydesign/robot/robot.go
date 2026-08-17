package robot

type WalkableRobot interface {
	Walk()
}

type TalkableRobot interface {
	Talk()
}

type Robot struct {
	walkBehaviour WalkableRobot
	talkBehaviour TalkableRobot
}

func Init(w WalkableRobot, t TalkableRobot) *Robot {
	return &Robot{walkBehaviour: w, talkBehaviour: t}
}

func (r *Robot) Walk() {
	r.walkBehaviour.Walk()
}

func (r *Robot) Talk() {
	r.talkBehaviour.Talk()
}
