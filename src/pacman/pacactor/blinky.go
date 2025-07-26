package pacactor

import (
	"github.com/jeppe9821/go-pacman/src/pacman/pacmap"
	"github.com/jeppe9821/go-pacman/src/render"
)

type BlinkyChaseBehaviour struct {
}

func (BlinkyChaseBehaviour) Update(actor *Actor, tileMap pacmap.TileMap, tick int) {

}

func (BlinkyChaseBehaviour) Draw(actor Actor, screen render.Screen, spritesheet render.Spritesheet8Bit) {
	screen.Draw(actor.Position.X, actor.Position.Y, 0, 29, false, false, spritesheet)
}
