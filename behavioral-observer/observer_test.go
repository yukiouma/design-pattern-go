package behavioralobserver

import (
	"behavioral-observer/internal/biz"
	"testing"
)

const (
	channelA = "channelA"
	channelB = "channelB"
)

func TestObserver(t *testing.T) {
	newsPublisher := biz.NewNewsPublihser(channelA, channelB)

	userClient1 := biz.NewUserClient("user1")
	userClient2 := biz.NewUserClient("user2")

	newsPublisher.Subsrcibe(channelA, userClient1)
	newsPublisher.Subsrcibe(channelB, userClient2)

	newsPublisher.Publish(channelA, "news of channel A")
	newsPublisher.Publish(channelB, "news of channel B")

	if len(userClient1.EventHistory()) != 1 {
		t.Fatalf("expect user client 1 receive 1 news, but go %d", len(userClient1.EventHistory()))
	}

	if len(userClient2.EventHistory()) != 1 {
		t.Fatalf("expect user client 2 receive 1 news, but go %d", len(userClient1.EventHistory()))
	}

	userClient1.CleanCache()
	userClient2.CleanCache()

	newsPublisher.UnSubsrcibe(channelB, userClient2)

	newsPublisher.Publish(channelA, "news of channel A")
	newsPublisher.Publish(channelB, "news of channel B")

	if len(userClient1.EventHistory()) != 1 {
		t.Fatalf("expect user client 1 receive 3 news, but go %d", len(userClient1.EventHistory()))
	}

	if len(userClient2.EventHistory()) != 0 {
		t.Fatalf("expect user client 2 receive 1 news, but go %d", len(userClient1.EventHistory()))
	}
}
