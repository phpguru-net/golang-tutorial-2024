package helpers

import (
	"encoding/json"
)

func ParseJSON(object any) ([]byte, error) {
	return json.Marshal(object)
}

func Add[T int64 | float64 | string](a, b T) T {
	return a + b
}
