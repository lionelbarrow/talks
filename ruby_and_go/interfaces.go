package numbers

type Number interface {
  Value() float64
}

type Rational struct {
  top int
  bottom int
}

func (this Rational) Value() float64 {
  return float64(this.top) / float64(this.bottom)
}

func (this Rational) Add(other Number) Number {
  return this.Valu() + other.Value()

type Irrational struct {
  approximateValue float64
}


