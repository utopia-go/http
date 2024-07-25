package validator

import "github.com/utopia-go/http/http"

type AllOf struct {
	failedRule http.Validator
	validators []http.Validator
	xType      string
}

func NewDefaultAllOf() *AllOf {
	return NewAllOf(make([]http.Validator, 0), http.TypeMixed)
}

func NewAllOf(validators []http.Validator, xType string) *AllOf {
	return &AllOf{validators: validators, xType: xType}
}

func (bv *AllOf) GetDescription() string {
	if bv.failedRule != nil {
		return bv.failedRule.GetDescription()
	}

	return bv.validators[0].GetDescription()
}

func (bv *AllOf) IsArray() bool {
	return true
}

func (bv *AllOf) GetType() string {
	return bv.xType
}

func (bv *AllOf) IsValid(value interface{}) bool {
	for _, validator := range bv.validators {
		valid := validator.IsValid(value)

		if !valid {
			bv.failedRule = validator
			return false
		}
	}

	return true
}
