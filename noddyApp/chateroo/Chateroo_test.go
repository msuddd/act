package chateroo

import (
	"strings"
	"testing"
)

func TestNewChateroo(t *testing.T) {
	chateroo1 := NewChateroo(1, "Best thing since sliced bread")
	if chateroo1.iteration != 1 {
		t.Errorf("Number of iterations was not correct, expected 1 got %v", chateroo1.iteration)
	}
	if chateroo1.greetingMessage != "Best thing since sliced bread" {
		t.Errorf("Greeting message not as expected!")
	}
}

func TestChateroo_TalkTo(t *testing.T) {
	chateroo1 := NewChateroo(1, "Best thing since sliced bread")
	chateroo2 := NewChateroo(1, "Snappy dresser")
	chateroo2.TalkTo(chateroo1, "hello")

	if len(chateroo1.messagesRecieved) != 1 {
		t.Errorf("Number of iterations was not correct, expected 1 got %v", len(chateroo1.messagesRecieved))
		t.Fail()
	}
	lastMessage := chateroo1.messagesRecieved[len(chateroo1.messagesRecieved)-1]
	if !strings.Contains(lastMessage, "hello") {
		t.Errorf("Expected hello got %v", lastMessage)
	}
}
