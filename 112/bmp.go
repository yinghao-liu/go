package main

import (
	"fmt"
	"image"
	"log"
	"os"

	"github.com/sergeymakinen/go-bmp"
)

// 写bmp
func saveBmp(path string, img image.Image) {
	imgfile, _ := os.Create(fmt.Sprintf(path))
	defer imgfile.Close()
	err := bmp.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}
}

// 读bmp
func readBmp(path string) (img image.Image, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	img, formatName, err := image.Decode(f)
	if err != nil {
		return
	}
	fmt.Printf("formatName is %s\n", formatName)
	fmt.Printf("Bounds is %v\n", img.Bounds())
	fmt.Printf("ColorModel is %#v\n", img.ColorModel())

	return
}
