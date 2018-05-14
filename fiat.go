package currencymath

// Fiat is the representation of a fiat value
type Fiat float64

const (
	decimalsFIAT = 2
)

// FiatFromCents returns a Fiat from a cent value, for
// converting from an API that uses integers
func FiatFromCents(s int64) Fiat {
	return Fiat(intToFloat(decimalsFIAT, s))
}

// Cents returns the satoshi value for the fiat, for talking to
// APIs that use integers.
func (b Fiat) Cents() int64 {
	return b.baseValue()
}

// Add adds another fiat value to this one and returns a new number
func (b Fiat) Add(v Fiat) Fiat {
	return b.add(v).(Fiat)
}

// Sub subrtacts another fiat value from this one and returns a new number
func (b Fiat) Sub(v Fiat) Fiat {
	return b.sub(v).(Fiat)
}

// Mult multiplies another fiat value with this one and returns a new number
func (b Fiat) Mult(v Fiat) Fiat {
	return b.mult(v).(Fiat)
}

// Div divides this fiat value by the provided value and returns a new number
func (b Fiat) Div(v Fiat) Fiat {
	return b.div(v).(Fiat)
}

func (b Fiat) value() float64 {
	return float64(b)
}

func (b Fiat) baseValue() int64 {
	return floatToInt(decimalsFIAT, b)
}

func (b Fiat) add(v mather) mather {
	return FiatFromCents(b.baseValue() + v.baseValue())
}

func (b Fiat) sub(v mather) mather {
	return FiatFromCents(b.baseValue() - v.baseValue())
}

func (b Fiat) mult(v mather) mather {
	return Fiat(floatToFloat(decimalsFIAT, b.value()*v.value()))
}

func (b Fiat) div(v mather) mather {
	return Fiat(floatToFloat(decimalsFIAT, b.value()/v.value()))
}

// FiatToBitcoin converts a fiat value to a bitcoin value based on the
// supplied rate
func FiatToBitcoin(f, ex Fiat) Bitcoin {
	return Bitcoin(f).div(ex).(Bitcoin)
}
