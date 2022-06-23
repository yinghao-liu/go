package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"os"

	_ "image/png"

	_ "golang.org/x/image/bmp"
)

func main() {

}

// 图片读
func imageRead() {
	f, err := os.Open("white2.bmp")
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

	rgba := image.NewAlpha(image.Rect(0, 0, 1, 3))
	fmt.Printf("%#v\n", rgba)
}
