package compound_interest

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"

	"phpguru.net/go-apps/appbase"
	"phpguru.net/go-apps/help"
	"phpguru.net/go-apps/util"
)

type CompoundInterestApp struct {
	appbase.AppInformation
}

// GetAppName implements appbase.App.
// Subtle: this method shadows the method (AppInformation).GetAppName of CompoundInterestApp.AppInformation.
func (c *CompoundInterestApp) GetAppName() string {
	return c.Name
}

// Run implements appbase.App.
// A = P * (1+r)^n
func (c *CompoundInterestApp) Run(payload ...any) {
	compoundInterestCmd := flag.NewFlagSet("cp", flag.ExitOnError)
	pv := compoundInterestCmd.Float64("pv", 0, "Enter your PV(Initial Investment). Eg: 1000 USD")
	r := compoundInterestCmd.Float64("r", 0, "Enter your r(Interest rate per perod). Eg: 5% -> 0.05")
	n := compoundInterestCmd.Int("n", 0, "Enter your n(number of periods). Eg: 1 year => 1")

	compoundInterestCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", compoundInterestCmd.Name())
		fmt.Fprintf(os.Stderr, "  -pv=float64 : %s\n", "Principle or Initial Investment (required). Eg: 1000 USD => 1000")
		fmt.Fprintf(os.Stderr, "  -r=float64 : %s\n", "Interest rate (required). Eg: 5% => 0.05")
		fmt.Fprintf(os.Stderr, "  -n=int : %s\n", "Number of periods (required). Eg: 1 year => 1")
	}
	args, ok := help.ParseArgs(payload...)
	if !ok {
		return
	}
	// Parse the flags
	if err := compoundInterestCmd.Parse(args); err != nil {
		fmt.Println("Error parsing flags:", err)
		return
	}
	// validate data
	if *pv <= 0 || *r <= 0 || *n <= 0 {
		compoundInterestCmd.Usage()
		os.Exit(1)
	}
	futureValue := calculateFutureValue(*pv, *r, *n)
	output := strings.Repeat("-", 50) + "\n"
	output += fmt.Sprintf("Investment: %v", *pv) + "\n"
	output += fmt.Sprintf("Interest: %v%v", *r*100, "%") + "\n"
	output += fmt.Sprintf("Periods: %v%v", *n, "%") + "\n"
	output += fmt.Sprintf("Future value: %v", futureValue) + "\n"
	output += strings.Repeat("-", 50) + "\n"
	fmt.Print(output)
}

func calculateFutureValue(pv float64, r float64, n int) float64 {
	// A = P * (1+r)^n
	return util.RoundUp(pv*math.Pow(1+r, float64(n)), 2)
}

func NewCompoundInterestApp() appbase.App {
	return &CompoundInterestApp{
		appbase.AppInformation{
			Name: "calculate future value of your investment",
		},
	}
}
