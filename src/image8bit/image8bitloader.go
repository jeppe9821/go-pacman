package image8bit

import (
	"image"
	"image/color"
	_ "image/png"
	"os"
)

type Image struct {
	Pixels []byte
	Width  uint8
	Height uint8
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
		Width:  uint8(width),
		Height: uint8(height),
	}
}
