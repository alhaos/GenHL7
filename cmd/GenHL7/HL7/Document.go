package HL7

import (
	"bufio"
	"io"
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

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	reader := bufio.NewReader(f)

	for i := 0; ; i++ {
		line, _, errReadLine := reader.ReadLine()
		if errReadLine == io.EOF {
			break
		}
		if errReadLine != nil {

			log.Fatalln(errReadLine)
		}

		if i == 0 {
			FieldSeparator = string(line[3])
			InFieldSeparator = string(line[4])
		} else {
			sLine := string(line)
			if sLine == "" {
				continue
			}

			splits := strings.Split(sLine, FieldSeparator)

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
