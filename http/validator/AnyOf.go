package validator

import "github.com/utopia-go/http/http"

type AnyOf struct {
	failedRule http.Validator
	validators []http.Validator
	xType      string
}

func NewAnyOf(validators []http.Validator, xType string) (validator AnyOf) {
	if xType == "" {
		xType = http.TypeMixed
	}

	return
}

func (bv *AnyOf) GetDescription() string {
	if bv.failedRule != nil {
		return bv.failedRule.GetDescription()
	}

	return bv.validators[0].GetDescription()
}

func (bv *AnyOf) IsArray() bool {
	return true
}

func (bv *AnyOf) GetType() string {
	return bv.xType
}

func (bv *AnyOf) IsValid(value interface{}) bool {
	for _, validator := range bv.validators {
		valid := validator.IsValid(value)
		bv.failedRule = validator
		if valid {
			return true
		}
	}

	return false
}
