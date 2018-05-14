package currencymath

import (
	"math/rand"
	"testing"
	"time"
)

func TestRounding(t *testing.T) {
	downVal := 0.123456781
	upVal := 0.123456785
	testRound(t, RoundingOptCeil, downVal, 0.12345679)
	testRound(t, RoundingOptCeil, upVal, 0.12345679)
	testRound(t, RoundingOptFloor, downVal, 0.12345678)
	testRound(t, RoundingOptFloor, upVal, 0.12345678)
	testRound(t, RoundingOptRound, downVal, 0.12345678)
	testRound(t, RoundingOptRound, upVal, 0.12345679)
}

func testRound(t *testing.T, r roundingOpt, v, e float64) {
	v1 := floatToFloat(r, 8, v)
	if v1 != e {
		t.Logf("%s %.8f != %.8f, from %.9f", r, v1, e, v1)
		t.Fail()
	}
}

func TestSettingRound(t *testing.T) {
	err := SetRounding(RoundingOptCeil)
	if err != nil {
		t.Logf("setting rounding: %+v", err)
		t.FailNow()
	}
	if packageLevelRounding != RoundingOptCeil {
		t.Log("rounding not set correctly")
		t.FailNow()
	}
}

func TestSettingRoundFail(t *testing.T) {
	packageLevelRounding = RoundingOptCeil
	err := SetRounding(RoundingOptFloor)
	if err == nil {
		t.Logf("no error setting rounding after initial setting")
		t.FailNow()
	}
	if packageLevelRounding != RoundingOptCeil {
		t.Log("rounding got set again")
		t.FailNow()
	}
}

func TestSettingRoundInvalid(t *testing.T) {
	packageLevelRounding = roundingOpt("")
	err := SetRounding("foobar")
	if err == nil {
		t.Logf("no error setting rounding to an invalid value")
		t.FailNow()
	}
	if packageLevelRounding != "" {
		t.Logf("rounding got set to an invalid value %s", packageLevelRounding)
		t.FailNow()
	}
}

func TestSatoshiConversion(t *testing.T) {
	v := BitcoinFromSatoshi(8675309)
	if v != 0.08675309 {
		t.Fail()
	}
}

func TestBitcoinAdding(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	vals := []Bitcoin{}
	totalInt := int64(0)

	for i := 0; i < 1000; i++ {
		n := rand.Float64()
		vals = append(vals, Bitcoin(n))
		totalInt += floatToInt(packageLevelRounding, decimalsBTC, n)
	}

	v := Bitcoin(0)

	for _, b := range vals {
		v = v.Add(b)
	}

	total := BitcoinFromSatoshi(totalInt)
	if v != total {
		t.Logf("total %f does not match expected value %f", v, total)
		t.Fail()
	}
}

func TestBitcoinPricing(t *testing.T) {
	v := BitcoinToFiat(Bitcoin(1), Fiat(1000))
	if v != 1000 {
		t.Fail()
	}
	v2 := FiatToBitcoin(Fiat(1), Fiat(1000))
	if v2 != 0.001 {
		t.Fail()
	}
}
