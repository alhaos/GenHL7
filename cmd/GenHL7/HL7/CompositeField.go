package HL7

import "strings"

type CompositeField struct {
	Value []string
}

func NewCompositeField(values []string) *CompositeField {
	return &CompositeField{values}
}

func (f *CompositeField) GetValue(i ...int) any {
	return f.Value[i[0]]
}

func (f *CompositeField) String() string {
	return strings.Join(f.Value, InFieldSeparator)
}
