package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

//create new type dick as slice of strings
type deck []string

// create new function that belongs to deck type
// print each individual card
// receiver function -> any variable of type deck has access to this method
// convention -> abbreviated name of type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// create a new hand
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Hearts", "Clubs", "Spades", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	// replace _ cause we don't care about i and j
	for _, suit := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suit)
		}
	}
	return cards
}

// deal some cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err, "Failed to load deck, generating my own!")
		return newDeck()
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	//create seed and random number gen
	seed := rand.NewSource(time.Now().UnixNano())
	ran := rand.New(seed)
	for i := range d {
		pos := ran.Intn(len(d) - 1)
		//by default go uses same seed
		//one liner swap, take what is at new pos assign to i, take what is i assign to new pos
		d[i], d[pos] = d[pos], d[i]
	}
}
