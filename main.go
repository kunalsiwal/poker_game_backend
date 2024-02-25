package main

import "fmt"

func main() {
	nd := newDeck()
	nd.shuffle()
	fmt.Println(nd.drawCards(5))
}
