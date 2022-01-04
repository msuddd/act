package chateroo

import (
	"fmt"
	pretty "github.com/inancgumus/prettyslice"
	"github.com/mitchellh/colorstring"
)

type chateroo struct {
	iteration                 int
	greetingMessage           string
	messagesRecieved          []string
	incomingConnectionsAmount int
}

func (c chateroo) Talk() {
	colorstring.Printf("[underline][bold][blue]Iteration %v: %v\n", c.iteration, c.greetingMessage)
	colorstring.Printf("[green]I have received %v messages\n\n", c.incomingConnectionsAmount)
	if len(c.messagesRecieved) > 0 {
		colorstring.Printf("[light_gray]The last greeting_message is as follows: \n [red]%v \n", c.messagesRecieved[len(c.messagesRecieved)-1])
	}
}

func (c chateroo) Summarize() {
	colorstring.Printf("[green]I have received %v messages\n\n", c.incomingConnectionsAmount)
	if len(c.messagesRecieved) > 0 {
		colorstring.Println("[light_gray]All messages follow: ")
		pretty.Show("messages", c.messagesRecieved)
	}
}

func (c chateroo) TalkTo(r *chateroo, message string) {
	r.tell(fmt.Sprintf("FROM %v: \n %v \n", c.iteration, message))
}
func (c *chateroo) tell(message string) {
	c.incomingConnectionsAmount += 1
	c.messagesRecieved = append(c.messagesRecieved, message)
}
func NewChateroo(identifier int, message string) (c *chateroo) {
	c = &chateroo{identifier, message, []string{}, 0}
	return
}
