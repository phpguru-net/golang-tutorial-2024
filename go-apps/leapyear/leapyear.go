package leapyear

import (
	"flag"
	"fmt"
	"os"

	"phpguru.net/go-apps/appbase"
	"phpguru.net/go-apps/help"
)

type LeapYear struct {
	appbase.AppInformation
}

// GetAppName implements appbase.App.
// Subtle: this method shadows the method (AppInformation).GetAppName of LeapYear.AppInformation.
func (l *LeapYear) GetAppName() string {
	return l.Name
}

// Run implements appbase.App.
func (l *LeapYear) Run(payload ...any) {
	leapYearCmd := flag.NewFlagSet("leapyear", flag.ExitOnError)
	leapYearNumber := leapYearCmd.Int("year", 0, "Enter the year to check for leap year")
	leapYearCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", leapYearCmd.Name())
		fmt.Fprintf(os.Stderr, "  -year=int : %s\n", "Specify the year to check (required)")
	}
	args, ok := help.ParseArgs(payload...)
	if !ok {
		return
	}
	// Parse the flags
	if err := leapYearCmd.Parse(args); err != nil {
		fmt.Println("Error parsing flags:", err)
		return
	}
	if *leapYearNumber == 0 {
		fmt.Println("The 'year' flag is required.")
		leapYearCmd.Usage()
		os.Exit(1)
	}
	fmt.Printf("The year [%v] %v Leap Year and has [%v] days\n", *leapYearNumber, "is", 366)
}

func NewLeapApp() appbase.App {
	return &LeapYear{
		appbase.AppInformation{
			Name: "Check the input year is leap year or not",
		},
	}
}
