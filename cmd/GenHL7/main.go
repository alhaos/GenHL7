package main

import (
	"fmt"
	"github.com/alhaos/GenHL7/HL7"
	"strings"
)

func main() {
	d := HL7.FromFile(`D:\repository\GenHL7\test\1674155633668.dat `)
	var account string
	sb := strings.Builder{}
	sb.WriteString("Account,TestName,TestResult")
	for _, s := range d.Segments {
		switch s.Header {
		case "OBR":
			account = s.Fields[1].String()
		case "OBX":
			sb.WriteString("\n" + account + "," + s.Fields[2].String() + "," + s.Fields[3].String())
		}
	}
	fmt.Println(sb.String())
}
