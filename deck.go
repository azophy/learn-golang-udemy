package main

import (
  "os"
  "fmt"
  "time"
  "strings"
  "io/ioutil"
  "math/rand"
)

type deck []string

func newDeck() deck {
  cards := deck{}
  cardSuits := []string{
    "Spades",
    "Diamonds",
    "Hearts",
    "Club",
  }
  cardValues := []string{
    "Ace",
    "One",
    "Two",
  }

  for _, suit := range cardSuits {
    for _, value := range cardValues {
      cards = append(cards, suit+" of "+value)
    }
  }

  return cards
}

func newDeckFromFile(filename string) deck {
  content, err := ioutil.ReadFile(filename)
  if err != nil {
    // option #1: print error & return new deck
    // option #2: print error & quit program
    fmt.Println("error:", err)
    os.Exit(1)
  }

  s := strings.Split(string(content), ",")
  return deck(s)
}

func (d deck) print() {
  for i, card := range d {
    fmt.Println(i, card)
  }
}

func deal(d deck, handSize int) (deck, deck) {
  return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
  return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
  return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) shuffle() {
  source := rand.NewSource(time.Now().UnixNano())
  r := rand.New(source)

  for i := range d {
    newPosition := r.Intn(len(d)-1)

    d[i], d[newPosition] = d[newPosition], d[i] //swap
  }
}
