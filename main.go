package main

import "fmt"

func main() {
  myDeck := newDeck()
  //myDeck := newDeckFromFile("coba.txt")

  //fmt.Println("Retrieved deck content")
  //fmt.Println("Original deck content")
  //myDeck.print()

  myDeck.shuffle()
  fmt.Println("Shuffle result:")
  myDeck.print()

  //fmt.Println("Dealing 5 cards")
  //myHand, myDeck := deal(myDeck, 5)
  //myHand.print()

  //fmt.Println("Remaining cards")
  //myDeck.print()

  //fmt.Println("toString")
  //fmt.Println(myDeck.toString())

  //_ = myDeck.saveToFile("coba.txt")
}
