package currencymath

type zero float64

func (zero) value() float64 {
	return 0.0
}

func (zero) baseValue() int64 {
	return 0
}

func (zero) add(v mather) mather {
	return v
}

func (zero) sub(v mather) mather {
	return v
}

func (zero) mult(v mather) mather {
	return v
}

func (zero) div(v mather) mather {
	return v
}
