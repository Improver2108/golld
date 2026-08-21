package main

import (
	"github.com/improver2108/golld/observerdesign/channel"
	"github.com/improver2108/golld/observerdesign/subscriber"
)

func main() {
	channel := channel.NewYouTubeChannel("shader-cache")

	subs1 := subscriber.NewYouTubeSubscriber("yash", channel)
	subs2 := subscriber.NewYouTubeSubscriber("mithu", channel)

	channel.Subscribe(subs1)
	channel.Subscribe(subs2)

	channel.UploadVideo("Spider man walkthrough part 1")
}
