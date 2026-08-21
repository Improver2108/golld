package subscriber

import (
	"fmt"

	"github.com/improver2108/golld/observerdesign/domain"
)

type youTubeChannel interface {
	domain.IChannel
	GetVideoData() string
}

type youTubeSubscriber struct {
	name    string
	channel youTubeChannel
}

func NewYouTubeSubscriber(name string, channel youTubeChannel) *youTubeSubscriber {
	return &youTubeSubscriber{name: name, channel: channel}
}

func (s *youTubeSubscriber) Update() {
	fmt.Printf("hey %s, %s have uploaded a video:%s\n", s.name, s.channel.GetName(), s.channel.GetVideoData())
}
