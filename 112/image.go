package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"os"
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
	saveBmp(path, plt)
}

// 16色
func bitmap16(path string) {
	plt := image.NewPaletted(image.Rect(0, 0, 4, 4), color16)
	wh := plt.Rect.Dx() * plt.Rect.Dy()
	for i := 0; i < wh; i++ {
		plt.Pix[i] = (uint8)(i % 16)
	}
	fmt.Printf("plt is %#v\n", plt)
	saveBmp(path, plt)
}

// RGBA
func truecolor(path string) {
	rgb := image.NewRGBA(image.Rect(0, 0, 32, 32))
	// for i := 0; i < 16; i++ {
	// 	for j := 0; j < 16; j++ {
	// 		rgb.Set(i, j, color16[i])
	// 	}
	// }
	if err := drawRect(image.Rect(0, 0, 10, 10), rgb, color16[2]); nil != err {
		fmt.Printf("err is %s\n", err.Error())
	}
	saveBmp(path, rgb)
}

func drawRect(rect image.Rectangle, dst draw.Image, col color.Color) (err error) {
	if !dst.Bounds().Overlaps(rect) {
		return errors.New("non Overlaps")
	}
	for x := rect.Min.X; x <= rect.Max.X; x++ {
		fmt.Printf("x:%d, y:%d, color:%#v\n", x, rect.Min.Y, col)
		dst.Set(x, rect.Min.Y, col)
	}
	for y := rect.Min.Y; y <= rect.Max.Y; y++ {
		dst.Set(rect.Min.X, y, col)
	}
	for x := rect.Min.X; x <= rect.Max.X; x++ {
		dst.Set(x, rect.Max.Y, col)
	}
	for y := rect.Min.Y; y <= rect.Max.Y; y++ {
		dst.Set(rect.Max.X, y, col)
	}
	return
}
