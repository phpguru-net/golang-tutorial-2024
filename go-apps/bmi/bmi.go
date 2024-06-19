package bmi

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"phpguru.net/go-apps/appbase"
	"phpguru.net/go-apps/help"
	"phpguru.net/go-apps/util"
)

const (
	BMI_DATA = "./bmi-categories.json"
)

type BmiApp struct {
	appbase.AppInformation
}

// GetAppName implements appbase.App.
// Subtle: this method shadows the method (AppInformation).GetAppName of BmiApp.AppInformation.
func (b *BmiApp) GetAppName() string {
	return b.Name
}

// Run implements appbase.App.
func (b *BmiApp) Run(payload ...any) {
	bmiCmd := flag.NewFlagSet("leapyear", flag.ExitOnError)
	// weight, height
	weight := bmiCmd.Float64("weight", 0, "Enter your weight (kilograms)")
	height := bmiCmd.Float64("height", 0, "Enter your weight (meters)")

	bmiCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", bmiCmd.Name())
		fmt.Fprintf(os.Stderr, "  -weight=float64 : %s\n", "Specify the weight to check (required)")
		fmt.Fprintf(os.Stderr, "  -height=float64 : %s\n", "Specify the weight to check (required)")
	}
	args, ok := help.ParseArgs(payload...)
	if !ok {
		return
	}
	// Parse the flags
	if err := bmiCmd.Parse(args); err != nil {
		fmt.Println("Error parsing flags:", err)
		os.Exit(1)
	}
	if *weight <= 0 || *height <= 0 {
		fmt.Println("Weight or height is not valid!")
		bmiCmd.Usage()
		os.Exit(1)
	}
	// calculate bmi
	categories := importData()
	var bmiRange Category
	var bmiValue float64 = util.RoundUp(*weight/(*height**height), 2)
	for _, category := range categories {
		if category.isInRange(bmiValue) {
			bmiRange = category
			break
		}
	}
	fmt.Printf("Your bmi is %v and you are %v !", bmiValue, bmiRange.Classification)
}

type Category struct {
	Classification string  `json:"classification"`
	Min            float64 `json:"min"`
	Max            float64 `json:"max"`
}

func (c *Category) Print() {
	fmt.Println(c.Classification, c.Max, c.Min)
}

func newCategory(classification string, min float64, max float64) *Category {
	return &Category{
		Classification: classification,
		Min:            min,
		Max:            max,
	}
}

func (c *Category) isInRange(bmiValue float64) bool {
	if c.Max == 0 {
		return bmiValue >= c.Min
	}
	return bmiValue >= c.Min && bmiValue <= c.Max
}

func importData() []Category {
	var categories []Category
	data, err := util.ReadFile(BMI_DATA)
	if err != nil {
		fmt.Println(err)
	}

	// jsonData := `[
	//     {"classification": "Low weight (thin)", "min": 0, "max": 18.49},
	//     {"classification": "Normal", "min": 18.5, "max": 24.99},
	//     {"classification": "Overweight", "min": 25, "max": 0},
	//     {"classification": "Pre-obesity", "min": 25, "max": 29.99},
	//     {"classification": "Obesity degree I", "min": 30, "max": 34.99},
	//     {"classification": "Obesity class II", "min": 35, "max": 39.99},
	//     {"classification": "Obesity degree III", "min": 40, "max": 0}
	// ]`

	// data = []byte(jsonData)

	e := json.Unmarshal(data, &categories)
	if e != nil {
		fmt.Println(e)
	}
	return categories
}

func NewBmiApp() appbase.App {
	return &BmiApp{
		appbase.AppInformation{
			Name: "calculate BMI (Body Mass Index) Metric",
		},
	}
}
