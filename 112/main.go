package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"

	_ "image/png"
	"os"

	_ "golang.org/x/image/bmp"
)

func main() {
	f, err := os.Open("white.bmp")
	if err != nil {
		panic(err)
	}
	img, formatName, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("formatName is %s\n", formatName)
	fmt.Printf("Bounds is %v\n", img.Bounds())
	fmt.Printf("ColorModel is %#v\n", img.ColorModel())
}
