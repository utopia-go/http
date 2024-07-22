package validator

import (
	"encoding/json"
	"github.com/utopia-go/http/http"
)

type Assoc struct {
	length int
}

func NewAssoc(length int) (validator Assoc) {
	if length <= 0 {
		length = 65535
	}
	validator.length = length

	return
}

func (a Assoc) GetDescription() string {
	return "Value must be a valid object"
}

func (a Assoc) IsArray() bool {
	return true
}

func (a Assoc) GetType() string {
	return http.TypeArray
}

func (a Assoc) IsValid(value interface{}) bool {
	if _, ok := value.(map[string]interface{}); ok {
		jsonString, _ := json.Marshal(value)

		if len(jsonString) > a.length {
			return false
		}

		return true
	}

	return false
}
