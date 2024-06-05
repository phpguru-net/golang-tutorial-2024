package helpers

import (
	"encoding/json"
)

func ParseJSON(object any) ([]byte, error) {
	return json.Marshal(object)
}
