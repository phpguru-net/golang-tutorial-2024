package leapyear

import (
	"fmt"

	"phpguru.net/go-apps/appbase"
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
func (l *LeapYear) Run(args ...any) {
	fmt.Println("Leap year app runs!")
}

func NewLeapApp() appbase.App {
	return &LeapYear{
		appbase.AppInformation{
			Name: "Check the input year is leap year or not",
		},
	}
}
