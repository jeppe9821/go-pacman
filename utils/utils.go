package utils

import (
	"image"
	"image/color"
	_ "image/png"
	"os"
)

type Image struct {
	Pixels []byte
	Width  uint16
	Height uint16
}

func LoadImage(path string) Image {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	pixels := make([]byte, width*height<<2)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := color.RGBAModel.Convert(img.At(x, y)).(color.RGBA)
			i := (y*width + x) << 2
			pixels[i+0] = c.R
			pixels[i+1] = c.G
			pixels[i+2] = c.B
			pixels[i+3] = c.A
		}
	}

	return Image{
		Pixels: pixels,
		Width:  uint16(width),
		Height: uint16(height),
	}
}
