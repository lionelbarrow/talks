package main

import "braintree"

func main() {
  myCard := braintree.CreditCard{"314159265359", "5359", 1, 17}
  myCard.Charge(100)
}
