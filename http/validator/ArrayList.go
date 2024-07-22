package validator

import (
	"fmt"
	"github.com/utopia-go/http/http"
)

type ArrayList struct {
	validator http.Validator
	length    int
}

func NewArrayList(validator http.Validator, length int) (nav ArrayList) {
	nav.validator = validator
	nav.length = length

	return
}

func (bv ArrayList) GetDescription() string {
	msg := "Value must a valid array"

	if bv.length > 0 {
		msg += fmt.Sprintf(" no longer than %d items", bv.length)
	}

	msg += fmt.Sprintf(" and: %v", bv.validator.GetDescription())

	return msg
}

func (bv ArrayList) IsArray() bool {
	return true
}

func (bv ArrayList) GetType() string {
	return bv.validator.GetType()
}

func (bv ArrayList) GetValidator() http.Validator {
	return bv.validator
}

func (bv ArrayList) IsValid(value interface{}) bool {
	switch v := value.(type) {
	case []interface{}:
		if bv.length > 0 && len(v) > bv.length {
			return false
		}

		for _, value := range v {
			if !bv.validator.IsValid(value) {
				return false
			}
		}
	default:
		return false
	}

	return true
}
