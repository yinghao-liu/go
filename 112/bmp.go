package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	"github.com/sergeymakinen/go-bmp"
)

// bitmap 2色 位深度 1
var color2 = []color.Color{
	color.RGBA{R: 0xff, G: 0x00, B: 0x0, A: 0xff},  // 红色
	color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff}, // 蓝色
}

// 16色 位深度 4
var color16 = []color.Color{
	color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}, // black 黑色
	color.RGBA{R: 0x80, G: 0x00, B: 0x00, A: 0xff}, // maroon 栗色
	color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff}, // green  绿色
	color.RGBA{R: 0x80, G: 0x80, B: 0x00, A: 0xff}, // olive 橄榄色
	color.RGBA{R: 0x00, G: 0x00, B: 0x80, A: 0xff}, // navy 藏青色
	color.RGBA{R: 0x80, G: 0x00, B: 0x80, A: 0xff}, // purple 紫色
	color.RGBA{R: 0x00, G: 0x80, B: 0x80, A: 0xff}, // teal 茶色
	color.RGBA{R: 0x80, G: 0x80, B: 0x80, A: 0xff}, // gray 灰色
	color.RGBA{R: 0xc0, G: 0xc0, B: 0xc0, A: 0xff}, // silver 银色
	color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}, // red 红色
	color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}, // lime 柠檬色
	color.RGBA{R: 0xb0, G: 0xb0, B: 0xb0, A: 0xff}, // yellow 黄色
	color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff}, // blue 蓝色
	color.RGBA{R: 0xff, G: 0x00, B: 0xff, A: 0xff}, // fuchsia 紫红
	color.RGBA{R: 0x00, G: 0xff, B: 0xff, A: 0xff}, // aqua 浅绿色
	color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}, // white 白色

}

// 位图
func bitmap(path string) {
	imgfile, _ := os.Create(fmt.Sprintf(path))
	defer imgfile.Close()

	plt := image.NewPaletted(image.Rect(0, 0, 2, 3), color2)
	wh := plt.Rect.Dx() * plt.Rect.Dy()
	for i := 0; i < wh; i++ {
		plt.Pix[i] = (uint8)(i % 2)
	}
	fmt.Printf("plt is %#v\n", plt)
	err := bmp.Encode(imgfile, plt)
	if err != nil {
		log.Fatal(err)
	}
}

// 16色
func bitmap16(path string) {
	imgfile, _ := os.Create(fmt.Sprintf(path))
	defer imgfile.Close()

	plt := image.NewPaletted(image.Rect(0, 0, 4, 4), color16)
	wh := plt.Rect.Dx() * plt.Rect.Dy()
	for i := 0; i < wh; i++ {
		plt.Pix[i] = (uint8)(i % 16)
	}
	fmt.Printf("plt is %#v\n", plt)
	err := bmp.Encode(imgfile, plt)
	if err != nil {
		log.Fatal(err)
	}
}
