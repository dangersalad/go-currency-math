package currencymath

import (
	"math/rand"
	"testing"
	"time"
)

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
		totalInt += floatToInt(decimalsBTC, Bitcoin(n))
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
