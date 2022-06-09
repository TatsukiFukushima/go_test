package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	cards := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}

	rand.Seed(time.Now().UnixNano())

	cards_len := len(cards)
	cards_rand := []int{}
	for i := 0; i < cards_len; i++ {
		num := rand.Intn(len(cards))
		cards_rand = append(cards_rand, cards[num])
		cards = append(cards[:num], cards[num+1:]...)
	}
	fmt.Printf("cards_rand: %v\n", cards_rand)
}
