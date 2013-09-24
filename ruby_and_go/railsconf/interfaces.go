package animals

type Animal interface { Eat(Animal) }

type Lion struct {}

func (this Lion) Eat(prey Animal) {} // Lion implements Animal

func (this Lion) Roar() {} // Lion has other methods

func Feed(hunter, prey Animal) { hunter.Eat(prey) }
