package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

// type Date struct {
// 	A string
// 	// Month string
// 	// Day   string
// }

// func main() {
// 	var date Date
// 	date.Year = "1900"
// 	date.Month = "01"
// 	date.Day = "01"
// 	data, e := yaml.Marshal(&date)
// 	if nil == e {
// 		fmt.Printf("%s", string(data))
// 	}
// }

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type Date2 struct {
	Year  string
	Month string
	Day   string
	// B struct {
	// 	RenamedC int   `yaml:"c"`
	// 	D        []int `yaml:",flow"`
	// }
}

func main() {
	date := Date2{}
	date.Year = "Easy!"
	date.Month = "192.167.4"
	date.Day = "Easy!"

	// date := Date{}
	// date.A = "1900"

	// err := yaml.Unmarshal([]byte(data), &t)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&date)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	// d, err = yaml.Marshal(&date)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// fmt.Printf("--- t dump:\n%s\n\n", string(d))

	// m := make(map[interface{}]interface{})

	// err = yaml.Unmarshal([]byte(data), &m)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// fmt.Printf("--- m:\n%v\n\n", m)

	// d, err = yaml.Marshal(&m)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// fmt.Printf("--- m dump:\n%s\n\n", string(d))
}
