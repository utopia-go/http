package validator

import "github.com/utopia-go/http/http"

type BooleanValidator struct {
	loose bool
}

func NewBooleanValidator(loose bool) (validator BooleanValidator) {
	validator.loose = loose

	return
}

func (bv BooleanValidator) GetDescription() string {
	return "Value must be a valid boolean"
}

func (bv BooleanValidator) IsArray() bool {
	return false
}

func (bv BooleanValidator) GetType() string {
	return http.TypeBoolean
}

func (bv BooleanValidator) IsValid(value interface{}) bool {
	switch v := value.(type) {
	case bool:
		return true
	case string:
		if bv.loose {
			return v == "true" || v == "false" || v == "1" || v == "0"
		}
	case int:
	case float64:
		return int(v) == 1 || v == 0
	}

	return false
}
