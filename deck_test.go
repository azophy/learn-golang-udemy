package main

import (
  "os"
  "testing"
)

/* notice that the func name is start with Uppercase, unlike our previous funcs
which started with lowercase
*/
func TestNewDeck(t *testing.T) {
  d := newDeck()

  if len(d) != 12 {
    // %v is example of a formatted string, similar to Sprintf in C language
    t.Errorf("Expected deck length of 12, but got %v", len(d))
  }

  if d[0] != "Ace of Spades" {
    t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
  }

  if d[len(d)-1] != "Three of Clubs" {
    t.Errorf("Expected first card of Three of Clubs, but got %v", d[len(d)-1])
  }
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
  os.Remove("_decktesting")

  originalDeck := newDeck()
  originalDeck.saveToFile("_decktesting")

  loadedDeck := newDeckFromFile("_decktesting")

  if len(originalDeck) != len(loadedDeck) {
    t.Errorf("Expected the same deck size between original & loaded deck, but got size %v for loadedDeck vs %v for originalDeck", len(loadedDeck), len(originalDeck))
  }

  os.Remove("_decktesting")
}
