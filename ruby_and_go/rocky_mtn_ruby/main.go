package main

import (
  "braintree"
  "fmt"
  "os"
)

func main() {
  cardNumber := os.Args[1]
  amount := os.Args[2]

  card := braintree.NewCreditCard(cardNumber)

  success := braintree.Card(card, 100)

  if success {
    fmt.Println("Charged %s $100", card)
  } else {
    fmt.Println("Failed to charge %s", card)
  }
}
