package conversion

import "strconv"

func StringsToFloat(strings []string) (*[]float64, error) {
	floatValues := make([]float64, len(strings))

	for lineIndex, line := range strings {
		floatVal, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, err
		}
		floatValues[lineIndex] = floatVal
	}

	return &floatValues, nil
}
