package http

type Validator interface {
	GetDescription() string
	IsArray() bool
	IsValid(value interface{}) bool
	GetType() string
}

const TypeBoolean = "boolean"
const TypeInteger = "integer"
const TypeFloat = "double"
const TypeString = "string"
const TypeArray = "array"
const TypeObject = "object"
const TypeMixed = "mixed"
