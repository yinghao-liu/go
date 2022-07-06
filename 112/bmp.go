package main

import (
	"fmt"
	"image"
	"log"
	"os"

	"github.com/sergeymakinen/go-bmp"
)

func saveBmp(path string, img image.Image) {
	imgfile, _ := os.Create(fmt.Sprintf(path))
	defer imgfile.Close()
	err := bmp.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}
}
