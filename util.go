package currencymath

import (
	"fmt"
	"strconv"
	"strings"
)

type valuer interface {
	value() float64
	baseValue() int64
}

type mather interface {
	valuer
	add(mather) mather
	sub(mather) mather
	mult(mather) mather
	div(mather) mather
}

func intToFloat(decimals int, val int64) float64 {
	str := strconv.FormatInt(val, 10)
	for len(str) < decimals+1 {
		str = "0" + str
	}
	fStr := fmt.Sprintf("%s.%s", str[0:len(str)-decimals], str[len(str)-decimals:])
	fVal, _ := strconv.ParseFloat(fStr, 64)
	return fVal
}

func floatToInt(decimals int, val valuer) int64 {
	str := strconv.FormatFloat(val.value(), 'f', decimals, 64)
	iStr := strings.Replace(str, ".", "", 1)
	iVal, _ := strconv.ParseInt(iStr, 10, 64)
	return iVal
}

func floatToFloat(decimals int, val float64) float64 {
	str := strconv.FormatFloat(val, 'f', decimals, 64)
	fVal, _ := strconv.ParseFloat(str, 64)
	return fVal
}

func add(vals ...mather) mather {
	if len(vals) == 0 {
		return zero(0)
	}

	var ret = vals[0]

	if len(vals) > 1 {
		vals = vals[1:]

		for _, v := range vals {
			ret = ret.add(v)
		}
	}

	return ret
}
