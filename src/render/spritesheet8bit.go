package render

import "github.com/jeppe9821/go-pacman/src/image8bit"

const (
	SPRITE_WIDTH  uint8 = 8
	SPRITE_HEIGHT uint8 = 8
)

type Spritesheet8Bit struct {
	pixels []byte
	width  uint8
	height uint8
}

func LoadFromFile(filePath string) Spritesheet8Bit {
	var img = image8bit.LoadImage(filePath)
	return Spritesheet8Bit{
		width:  uint8(img.Width),
		height: uint8(img.Height),
		pixels: img.Pixels,
	}
}

func (spritesheet *Spritesheet8Bit) GetPixel(x uint8, y uint8) Pixel {
	var pixelIndex uint = (uint(x) + uint(y)*uint(spritesheet.width)) * RGBA

	return Pixel{
		r: spritesheet.pixels[pixelIndex+0],
		g: spritesheet.pixels[pixelIndex+1],
		b: spritesheet.pixels[pixelIndex+2],
		a: spritesheet.pixels[pixelIndex+3],
	}
}
