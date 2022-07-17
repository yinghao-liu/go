package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	data := []byte{0x00, 0x00}
	//var buffer bytes.Buffer

	// dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	// base64.StdEncoding.Encode(dst, data)
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)

}
