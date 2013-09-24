package braintree

type CreditCard struct {
  number string
  Last4 string
  ExpirationMonth int
  ExpirationYear int
}

func (this CreditCard) Charge(amount int) err {
  return this.runTransaction(amount)
}

func (this CreditCard) runTransaction(amount int) {
  // hidden
}
