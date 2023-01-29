package HL7

import (
	"log"
	"os"
	"strings"
)

type Document struct {
	Segments []*Segment
}

var FieldSeparator string
var InFieldSeparator string

func NewDocument() *Document {
	return new(Document)
}

func FromFile(filename string) *Document {

	d := NewDocument()

	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	FieldSeparator = string(bytes[3])
	InFieldSeparator = string(bytes[4])

	lines := strings.Split(string(bytes), "\r")

	for _, line := range lines {

		if line == "" {
			continue
		}

		splits := strings.Split(line, FieldSeparator)

		s := NewSegment(splits[0])

		for _, v := range splits[1:] {
			if strings.Contains(v, InFieldSeparator) {
				s.Fields = append(s.Fields, NewCompositeField(
					strings.Split(v, InFieldSeparator),
				))
			} else {
				s.Fields = append(s.Fields, NewSampleField(v))
			}
		}
		d.Segments = append(d.Segments, s)
	}
	return d
}

func (d *Document) String() string {
	sb := strings.Builder{}
	for i, s := range d.Segments {
		if i == 0 {
			sb.WriteString(s.String())
		} else {
			sb.WriteString("\n" + s.String())
		}
	}
	return sb.String()
}
