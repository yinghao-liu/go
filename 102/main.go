// json.go
package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func parserTypeswitches() {

}

func main() {
	// pa := &Address{"private", "Aartselaar", "Belgium"}
	// wa := &Address{"work", "Boom", "Belgium"}
	// vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// // fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:

	// // JSON format:
	// js, _ := json.Marshal(vc)
	// fmt.Printf("JSON format: %s", js)

	// var out bytes.Buffer
	// json.Indent(&out, js, "", "\t")
	// fmt.Println(out.String())
	// out.WriteTo(os.Stdout)

	// goodJSON := `true`
	// badJSON := `{"example":2:]}}`

	// fmt.Println("\n", json.Valid([]byte(goodJSON)), json.Valid([]byte(badJSON)))

	// fmt.Println(io.EOF.Error())

	// // using an encoder:
	// file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	// defer file.Close()
	// enc := json.NewEncoder(file)
	// err := enc.Encode(vc)
	// if err != nil {
	// 	log.Println("Error in encoding json")
	// }

	jsonData := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)

	var v interface{}
	json.Unmarshal(jsonData, &v)
	data := v.(map[string]interface{})

	for k, v := range data {
		switch v := v.(type) {
		case string:
			fmt.Println(k, v, "(string)")
		case float64:
			fmt.Println(k, v, "(float64)")
		case []interface{}:
			fmt.Println(k, "(array):")
			for i, u := range v {
				fmt.Println("    ", i, u)
			}
		default:
			fmt.Println(k, v, "(unknown)")
		}
	}
}
