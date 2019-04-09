package avro

import (
	"encoding/json"

	"github.com/valyala/fastjson"
)

// ArraySchema -
type ArraySchema struct {
	Type  Type   `json:"type"`
	Items Schema `json:"items"`
}

// TypeName -
func (t *ArraySchema) TypeName() Type {
	return TypeArray
}

// MarshalJSON -
func (t *ArraySchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(t)
}

func translateValue2ArraySchema(value *fastjson.Value) (Schema, error) {

	return nil, nil
}

// NewArraySchema -
func NewArraySchema(items Schema) *ArraySchema {
	return &ArraySchema{
		Type:  TypeArray,
		Items: items,
	}
}
