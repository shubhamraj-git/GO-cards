package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

//create new type of deck

type deck []string //slice of string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, i := range cardSuits {
		for _, j := range cardValues {
			cards = append(cards, i+" of "+j)
		}
	}
	return cards
}

func deal(d deck, handsize int) (deck, deck) {
	card_new := d[:handsize]
	card_left := d[handsize:]
	return card_new, card_left

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// func newDeckFromFile(filename string) deck {
// 	bs, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		// if you have error, give a new decks of cards
// 		// log the error and entirely quit
// 		fmt.Println("Error:", err)
// 		os.Exit(1)
// 	}
// 	return deck(strings.Split(string(bs), ","))
// }

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
