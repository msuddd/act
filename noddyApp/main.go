package main

import "noddyApp/chateroo"

func main() {
	chateroo1 := chateroo.NewChateroo(1, "Best thing since sliced bread")
	chateroo2 := chateroo.NewChateroo(2, "A little dubious")
	chateroo1.Talk()
	chateroo2.TalkTo(chateroo1, "Good day sir!")
	chateroo1.Talk()
	chateroo2.TalkTo(chateroo1, "I hear the ballet in Prague is excellent this season!")
	chateroo1.Talk()

	chateroo1.TalkTo(chateroo2, "Hello Georgy")
	chateroo2.Talk()

	chateroo1.Summarize()
}
