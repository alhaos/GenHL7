package HL7

import "strings"

type Segment struct {
	Header string
	Fields []Field
}

func NewSegment(header string) *Segment {
	return &Segment{Header: header}
}

func (s *Segment) String() string {
	sb := strings.Builder{}
	sb.WriteString(s.Header)
	for _, f := range s.Fields {
		sb.WriteString("|" + f.String())
	}
	return sb.String()
}
