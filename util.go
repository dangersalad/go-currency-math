package currencymath

import (
	"fmt"
	"github.com/pkg/errors"
	"math"
	"strconv"
	"strings"
)

type roundingOpt string

const (
	// RoundingOptFloor causes all float values to ignore extra
	// digits, so a fiat value of 0.009 would be 0.00
	RoundingOptFloor roundingOpt = "floor"
	// RoundingOptCeil causes all float values to be rounded up, so a
	// fiat value of 0.001 would be 0.01
	RoundingOptCeil roundingOpt = "ceil"
	// RoundingOptRound causes all float values to be rounded
	// naturally, so a fiat value of 0.001 would be 0.00 and 0.005
	// would be 0.01
	RoundingOptRound roundingOpt = "round"
)

var roundingOpts = []roundingOpt{
	RoundingOptFloor,
	RoundingOptCeil,
	RoundingOptRound,
}

var packageLevelRounding = roundingOpt("")

// SetRounding sets the way to handle rounding floats. This can only
// be set once and will effect all package level operations from then
// on. The default is RoundingOptRound
func SetRounding(r roundingOpt) error {
	if packageLevelRounding != "" {
		return errors.New("currencymath rounding policy already set")
	}
	for _, o := range roundingOpts {
		if roundingOpt(r) == o {
			packageLevelRounding = roundingOpt(r)
			return nil
		}
	}
	// if nothing matched, panic, there is no reason to report this
	// error
	return errors.Errorf("invalid currencymath rounding policy %s", r)
}

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

func floatToInt(r roundingOpt, decimals int, val float64) int64 {
	str := getFloatString(r, decimals, val)
	iStr := strings.Replace(str, ".", "", 1)
	iVal, _ := strconv.ParseInt(iStr, 10, 64)
	return iVal
}

func floatToFloat(r roundingOpt, decimals int, val float64) float64 {
	str := getFloatString(r, decimals, val)
	fVal, _ := strconv.ParseFloat(str, 64)
	return fVal
}

func getFloatString(r roundingOpt, decimals int, val float64) string {
	// get next digit
	str := strconv.FormatFloat(val, 'f', decimals+1, 64)
	strLen := len(str)
	endStr := fmt.Sprintf("%s.%s", str[strLen-2:strLen-1], str[strLen-1:])
	endVal, _ := strconv.ParseFloat(endStr, 64)
	var endInt int
	switch r {
	case RoundingOptFloor:
		endInt = int(math.Floor(endVal))
	case RoundingOptCeil:
		endInt = int(math.Ceil(endVal))
	default: // RoundingOptRound
		endInt = int(math.Round(endVal))
	}
	return fmt.Sprintf("%s%s", str[:len(str)-2], strconv.Itoa(endInt))

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
