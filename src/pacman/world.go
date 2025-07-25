package pacman

import (
	"github.com/jeppe9821/go-pacman/src/pacman/pacactor"
	"github.com/jeppe9821/go-pacman/src/pacman/pacmap"
	"github.com/jeppe9821/go-pacman/src/render"
)

type World struct {
	player  *pacactor.Actor
	blinky  *pacactor.Actor
	TileMap pacmap.TileMap
}

func (w World) Update(tick int) {
	w.player.CurrentBehaviour.Update(w.player, w.TileMap, tick)
	w.blinky.CurrentBehaviour.Update(w.blinky, w.TileMap, tick)
}

func (w World) Draw(screen render.Screen, spritesheet render.Spritesheet8Bit) {
	w.TileMap.Draw(screen, spritesheet)
	w.player.CurrentBehaviour.Draw(*w.player, screen, spritesheet)
	w.blinky.CurrentBehaviour.Draw(*w.blinky, screen, spritesheet)
}
