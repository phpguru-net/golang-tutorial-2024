package price

import (
	"reflect"
	"testing"
)

type ExpectedInputPrices struct {
	Prices   []float64
	TaxRates []int
}

func TestGetPrices(t *testing.T) {
	tc := struct {
		name     string
		input    TaxIncludedPriceJob
		expected ExpectedInputPrices
	}{
		name:     "tc1",
		input:    *NewTaxIncludedPriceJob(&[]float64{10, 20, 30}, &[]int{}, &[][]float64{}),
		expected: ExpectedInputPrices{Prices: []float64{10, 20, 30}},
	}

	t.Run(tc.name, func(t *testing.T) {
		if !reflect.DeepEqual(tc.input.GetPrices(), &tc.expected.Prices) {
			t.Errorf("%v not equal %v", *tc.input.GetPrices(), tc.expected.Prices)
		}
	})
}

func TestGetTaxRates(t *testing.T) {
	tc := struct {
		name     string
		input    TaxIncludedPriceJob
		expected ExpectedInputPrices
	}{
		name:     "tc2",
		input:    *NewTaxIncludedPriceJob(&[]float64{}, &[]int{0, 10, 20}, &[][]float64{}),
		expected: ExpectedInputPrices{TaxRates: []int{0, 10, 20}},
	}

	t.Run(tc.name, func(t *testing.T) {
		if !reflect.DeepEqual(tc.input.GetTaxRates(), &tc.expected.TaxRates) {
			t.Errorf("%v not equal %v", *tc.input.GetTaxRates(), tc.expected.TaxRates)
		}
	})
}

func TestProcess(t *testing.T) {
	t.Run("Test Process", func(t *testing.T) {
		NewTaxIncludedPriceJob(&[]float64{}, &[]int{0, 10, 20}, &[][]float64{}).Process()
	})
}
