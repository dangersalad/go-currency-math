package currencymath

// Bitcoin is the representation of a bitcoin value
type Bitcoin float64

const (
	decimalsBTC = 8
)

// BitcoinFromSatoshi returns a Bitcoin from a satoshi value, for
// converting from an API that uses integers
func BitcoinFromSatoshi(s int64) Bitcoin {
	return Bitcoin(intToFloat(decimalsBTC, s))
}

// Satoshi returns the satoshi value for the bitcoin, for talking to
// APIs that use integers.
func (b Bitcoin) Satoshi() int64 {
	return b.baseValue()
}

// Add adds another bitcoin value to this one and returns a new number
func (b Bitcoin) Add(v Bitcoin) Bitcoin {
	return b.add(v).(Bitcoin)
}

// Sub subrtacts another bitcoin value from this one and returns a new number
func (b Bitcoin) Sub(v Bitcoin) Bitcoin {
	return b.sub(v).(Bitcoin)
}

// Mult multiplies another bitcoin value with this one and returns a new number
func (b Bitcoin) Mult(v Bitcoin) Bitcoin {
	return b.mult(v).(Bitcoin)
}

// Div divides this bitcoin value by the provided value and returns a new number
func (b Bitcoin) Div(v Bitcoin) Bitcoin {
	return b.div(v).(Bitcoin)
}

func (b Bitcoin) value() float64 {
	return floatToFloat(decimalsBTC, float64(b))
}

func (b Bitcoin) baseValue() int64 {
	return floatToInt(decimalsBTC, b)
}

func (b Bitcoin) add(v mather) mather {
	return BitcoinFromSatoshi(b.baseValue() + v.baseValue())
}

func (b Bitcoin) sub(v mather) mather {
	return BitcoinFromSatoshi(b.baseValue() - v.baseValue())
}

func (b Bitcoin) mult(v mather) mather {
	return Bitcoin(floatToFloat(decimalsBTC, b.value()*v.value()))
}

func (b Bitcoin) div(v mather) mather {
	return Bitcoin(floatToFloat(decimalsBTC, b.value()/v.value()))
}

// BitcoinToFiat converts a bitcoin value to a fiat value based on the
// supplied rate
func BitcoinToFiat(b Bitcoin, ex Fiat) Fiat {
	return ex.mult(b).(Fiat)
}
