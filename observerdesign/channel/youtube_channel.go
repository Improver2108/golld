package channel

import (
	"fmt"

	"github.com/improver2108/golld/observerdesign/domain"
)

type youtubeChannel struct {
	subscribers map[domain.ISubscriber]struct{}
	latestVideo string
	name        string
}

func NewYouTubeChannel(name string) *youtubeChannel {
	return &youtubeChannel{name: name, subscribers: make(map[domain.ISubscriber]struct{})}
}

func (c *youtubeChannel) Subscribe(subscriber domain.ISubscriber) {
	c.subscribers[subscriber] = struct{}{}
}

func (c *youtubeChannel) Unsubscribe(subscriber domain.ISubscriber) {
	delete(c.subscribers, subscriber)
}

func (c *youtubeChannel) notifySubscribers() {
	for subscriber := range c.subscribers {
		subscriber.Update()
	}
}

func (c *youtubeChannel) GetName() string {
	return c.name
}

func (c *youtubeChannel) UploadVideo(video string) {
	c.latestVideo = video
	fmt.Printf("%s uploaded a video : %s\n", c.name, video)
	c.notifySubscribers()
}

func (c *youtubeChannel) GetVideoData() string {
	return c.latestVideo
}
