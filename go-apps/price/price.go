package price

import (
	"fmt"

	"phpguru.net/go-apps/price/conversion"
	"phpguru.net/go-apps/price/filemanager"
)

type TaxIncludedPriceJob struct {
	prices            []float64
	taxRates          []int
	pricesIncludedTax [][]float64
}

type Result struct {
	TaxRate          int
	InputPrices      []float64
	TaxIncludedPrice map[string]string
}

func NewTaxIncludedPriceJob(prices *[]float64, taxRates *[]int, pricesIncludedTax *[][]float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		prices:            *prices,
		taxRates:          *taxRates,
		pricesIncludedTax: *pricesIncludedTax,
	}
}

func (job *TaxIncludedPriceJob) GetPrices() *[]float64 {
	return &job.prices
}

func (job *TaxIncludedPriceJob) GetTaxRates() *[]int {
	return &job.taxRates
}

func (job *TaxIncludedPriceJob) GetPricesIncludedTax() *[][]float64 {
	return &job.pricesIncludedTax
}

func (job *TaxIncludedPriceJob) ReadPricesFromFile(path string) {

	lines, err := filemanager.NewFileManager(path).ReadLines()
	if err != nil {
		fmt.Println("Can not open file " + path)
		fmt.Println(err.Error())
		return
	}

	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		fmt.Println("Can not parse values!")
		fmt.Println(err.Error())
		return
	}

	job.prices = *prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.ReadPricesFromFile("../prices.txt")
	job.taxRates = []int{10, 20, 30}

	for _, taxRate := range job.taxRates {
		pricesIncludedTax := make(map[string]string)
		for _, price := range job.prices {
			taxIncludedPrice := price*1 + float64(taxRate)/100
			pricesIncludedTax[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
		}
		result := Result{
			TaxRate:          taxRate,
			InputPrices:      job.prices,
			TaxIncludedPrice: pricesIncludedTax,
		}
		filemanager.WriteJsonFile(fmt.Sprintf("result_%v.json", taxRate), result)
	}

}
