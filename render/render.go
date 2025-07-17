package render

import "github.com/jeppe9821/go-pacman/utils"

type Renderable struct {
	pixels []byte
	width  int
	height int
}

type Pixel struct {
	r byte
	g byte
	b byte
	a byte
}

// If pixels is nil it will generate an empty pixels array
func Create(width int, height int, pixels []byte) Renderable {
	if pixels == nil {
		pixels = make([]byte, width*height*4)
	}

	return Renderable{
		width:  width,
		height: height,
		pixels: pixels,
	}
}

func LoadFromFile(filePath string) Renderable {
	var img = utils.LoadImage(filePath)
	return Renderable{
		width:  img.Width,
		height: img.Height,
		pixels: img.Pixels,
	}
}

func (renderable *Renderable) GetPixelsRaw() []byte {
	return renderable.pixels
}

func (renderable *Renderable) GetPixel(x int, y int) Pixel {
	pixelIndex := (x + (y * renderable.width)) << 2

	return Pixel{
		r: renderable.pixels[pixelIndex+0],
		g: renderable.pixels[pixelIndex+1],
		b: renderable.pixels[pixelIndex+2],
		a: renderable.pixels[pixelIndex+3],
	}
}

func (renderable *Renderable) DrawPixel(x int, y int, pixel Pixel) {
	colorIndex := (x + (y * renderable.width)) << 2
	renderable.pixels[colorIndex+0] = pixel.r
	renderable.pixels[colorIndex+1] = pixel.g
	renderable.pixels[colorIndex+2] = pixel.b
	renderable.pixels[colorIndex+3] = pixel.a
}

func (renderable *Renderable) DrawPixels(xp int, yp int, w int, h int, pixels []Pixel) {
	for y := yp; y < h; y++ {
		for x := xp; x < w; x++ {
			renderable.DrawPixel(x, y, pixels[(x+(y*renderable.width))])
		}
	}
}

func (renderer *Renderable) Draw(xp int, yp int, otherRenderable Renderable) {
	for y := 0; y < 16; y++ {
		for x := 0; x < 16; x++ {
			var pixel = otherRenderable.GetPixel(x, y)
			renderer.DrawPixel(x+xp, y+yp, pixel)
		}
	}
}

func (renderer *Renderable) DrawPart(xp int, yp int, xp2 int, yp2 int, w int, h int, xFlip bool, yFlip bool, otherRenderable Renderable) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			var pixel Pixel

			if xFlip {
				if yFlip {
					pixel = otherRenderable.GetPixel(((w-1)-x)+xp2, ((h-1)-y)+yp2)
				} else {
					pixel = otherRenderable.GetPixel(((w-1)-x)+xp2, y+yp2)
				}
			} else {
				if yFlip {
					pixel = otherRenderable.GetPixel(x+xp2, ((h-1)-y)+yp2)
				} else {
					pixel = otherRenderable.GetPixel(x+xp2, y+yp2)
				}
			}

			renderer.DrawPixel(x+xp, y+yp, pixel)
		}
	}
}
