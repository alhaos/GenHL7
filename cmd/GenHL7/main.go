package main

import (
	"fmt"
	"github.com/alhaos/GenHL7/HL7"
)

func main() {

	d := HL7.FromFile(`C:\tmp\exampleHl7.hl7`)

	fmt.Println(d.String())

}
