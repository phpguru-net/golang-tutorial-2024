package main

import (
	"fmt"

	"phpguru.net/go-apps/lunar"
)

func main() {
	// Test solar to lunar conversion

	day, month, year, leap := lunar.Solar2Lunar(20, 6, 2024, 7)
	fmt.Printf("Solar to Lunar: %02d-%02d-%d, Leap: %v\n", day, month, year, leap)

	day2, month2, year2, leap2 := lunar.Solar2Lunar(1, 5, 1988, 7)
	fmt.Printf("Solar to Lunar: %02d-%02d-%d, Leap: %v\n", day2, month2, year2, leap2)

	// Test lunar to solar conversion
	// sDay, sMonth, sYear := lunar.Lunar2Solar(1, 5, 1988, false, 7)
	// fmt.Printf("Lunar to Solar: %02d-%02d-%d\n", sDay, sMonth, sYear)
	// cli.ReadArgs()
}
