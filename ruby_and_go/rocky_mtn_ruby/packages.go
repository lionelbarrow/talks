package braintree

func Charge(card CreditCard, amount int) bool {
	if card.valid() {
		card.charge(amount)
		return true
	}
	return false
}

type CreditCard struct {
	Last4  string
	Type   string
	number string
}

func (c CreditCard) String() string {
	return c.Type + " ending in " + c.Last4
}

func (c CreditCard) valid() bool

func (c CreditCard) charge(amount int)
