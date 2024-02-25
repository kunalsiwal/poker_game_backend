package main

import (
	"testing"
)

func compareSlices(a, b deck) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestDeckShuffle(t *testing.T) {

  nd:=newDeck()
  nd2:=newDeck()
  nd=nd.shuffle()
  nd2=nd2.shuffle()
	if compareSlices(nd, nd2){
		t.Errorf("Shuffling is not correct here is nd1,nd2 and nd respectively,%v,%v", nd, nd2)
	}
}

