package HL7

type SampleField struct {
	Value string
}

func NewSampleField(v string) *SampleField {
	return &SampleField{Value: v}
}

func (f *SampleField) GetValue(s ...int) any {
	return f.Value
}

func (f *SampleField) String() string {
	return f.Value
}