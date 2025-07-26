package render

const (
	RGBA = 4
)

type Screen struct {
	pixels []byte
	width  uint
	height uint
}

type Pixel struct {
	r byte
	g byte
	b byte
	a byte
}

// If pixels is nil it will generate an empty pixels array
func CreateScreen(width uint, height uint, pixels []byte) Screen {
	if pixels == nil {
		pixels = make([]byte, width*height*RGBA)
	}

	return Screen{
		width:  uint(width),
		height: uint(height),
		pixels: pixels,
	}
}

func (renderable *Screen) GetPixelsRaw() []byte {
	return renderable.pixels
}

func (renderable *Screen) DrawPixel(xp uint, yp uint, pixel Pixel) {
	colorIndex := (xp + (yp * renderable.width)) * 4
	renderable.pixels[colorIndex+0] = pixel.r
	renderable.pixels[colorIndex+1] = pixel.g
	renderable.pixels[colorIndex+2] = pixel.b
	renderable.pixels[colorIndex+3] = pixel.a
}

func (screen *Screen) Draw(xp uint,
	yp uint,
	xt uint8,
	yt uint8,
	xFlip bool,
	yFlip bool,
	spritesheet Spritesheet8Bit) {

	type Sprite8Bit struct {
		x     uint8
		y     uint8
		xFlip bool
		yFlip bool
	}

	sprite := Sprite8Bit{
		x:     (xt * SPRITE_WIDTH) + (1 * xt),
		y:     (yt * SPRITE_HEIGHT) + (1 * yt),
		xFlip: xFlip,
		yFlip: yFlip,
	}

	var x, y uint8

	for y = 0; y < SPRITE_HEIGHT; y++ {
		for x = 0; x < SPRITE_WIDTH; x++ {

			var pixel Pixel
			var xSpritesheet, ySpritsheet uint8
			if sprite.xFlip {
				if sprite.yFlip {
					xSpritesheet = ((SPRITE_WIDTH - 1) - x) + sprite.x
					ySpritsheet = ((SPRITE_HEIGHT - 1) - y) + sprite.y
				} else {
					xSpritesheet = ((SPRITE_WIDTH - 1) - x) + sprite.x
					ySpritsheet = y + sprite.y
				}
			} else {
				if sprite.yFlip {
					xSpritesheet = x + sprite.x
					ySpritsheet = ((SPRITE_HEIGHT - 1) - y) + sprite.y
				} else {
					xSpritesheet = x + sprite.x
					ySpritsheet = y + sprite.y
				}
			}

			pixel = spritesheet.GetPixel(xSpritesheet, ySpritsheet)
			screen.DrawPixel(uint(x)+xp, uint(y)+yp, pixel)
		}
	}
}
