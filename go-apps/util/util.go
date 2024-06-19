package util

import (
	"math"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func RoundUp(number float64, decimalNumbers int) float64 {
	pow := math.Pow(10, float64(decimalNumbers))
	return math.Round(number*pow) / pow
}
