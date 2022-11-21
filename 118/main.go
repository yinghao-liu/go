package main

import (
	"encoding/xml"
	"fmt"
)

type Date struct {
	Year  string
	Month string
	Day   string
}

func main() {
	var date Date
	date.Year = "1900"
	date.Month = "01"
	date.Day = "01"
	data, e := xml.Marshal(date)
	if nil == e {
		fmt.Printf(string(data))
	}
}
