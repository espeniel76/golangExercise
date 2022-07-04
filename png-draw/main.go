package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {

	go EduceImage("pokeball.png", "#B8F500")
}

func EduceImage(imgSrc string, hexValue string) {
	imgSource, _ := os.Open(imgSrc)
	imgLayer, _ := png.Decode(imgSource)
	defer imgSource.Close()

	b := imgLayer.Bounds()
	imgResult := image.NewRGBA(b)

	m := image.NewRGBA(b)
	colorRgba, _ := parseHexColor(hexValue)
	draw.Draw(m, m.Bounds(), &image.Uniform{colorRgba}, image.Point{}, draw.Src)

	draw.Draw(imgResult, b, m, image.Point{}, draw.Src)
	draw.Draw(imgResult, b, imgLayer, image.Point{}, draw.Over)

	third, _ := os.Create("imgResult.jpg")
	jpeg.Encode(third, imgResult, &jpeg.Options{90})
	defer third.Close()
}

func parseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return
}
