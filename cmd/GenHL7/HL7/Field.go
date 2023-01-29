package HL7

type Field interface {
	GetValue(index1 ...int) any
	String() string
}
