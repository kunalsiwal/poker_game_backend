package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

type deck []string

func newDeck() deck {
	suites := []string{"Spade", "Club", "Heart", "Diamond"}
	value := []string{"Ace", "Duce", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "King", "Queen", "Jack"}
	result := deck{}
	for _, suite := range suites {
		for _, value := range value {
			append_val := value + " of " + suite
			result = append(result, append_val)
		}
	}
	return result
}

func (d deck) print() {
	for _, value := range d {
		fmt.Println(value, ",")
	}
}

func (d deck) drawCards(n int8) deck {
	if int8(len(d)) > n {
		return d[:n]
	}
	return d
}

func (d deck) shuffle() deck {
	for i := range d {
		index := generateRandomNumber(int64(len(d)))
		d[index], d[i] = d[i], d[index]
	}
	return d
}

func generateRandomNumber(n int64) int64 {
	now := time.Now() // Get the current time
	unixNano := now.UnixNano()
	r, err := rand.Int(rand.Reader, big.NewInt(unixNano))
	if err != nil {
		fmt.Println("error:", err)
		return 0
	}
  result:=r.Int64()
	return (result+unixNano)%n
}
