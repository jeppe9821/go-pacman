package pacman

import (
	"math"

	"github.com/jeppe9821/go-pacman/render"
)

type PacmanMap struct {
	player      *Player
	spritesheet render.Renderable
	grid        [][]int
}

func (pacmanMap PacmanMap) Update(tick int) {
	pacmanMap.player.Update(tick)
}

func (pacmanMap PacmanMap) Draw(screen render.Renderable) {
	var yGameOffset = 24

	for y := 0; y < len(pacmanMap.grid); y++ {
		for x := 0; x < len(pacmanMap.grid[y]); x++ {
			if pacmanMap.grid[y][x] == 0 {
				continue
			}

			tile := pacmanMap.grid[y][x] * 8

			screen.DrawPart(x*8, (y*8)+yGameOffset, tile%256, 8*int(math.Floor(float64(tile)/float64(256.0))), 8, 8, false, false, pacmanMap.spritesheet)
		}
	}

	pacmanMap.player.Draw(screen, pacmanMap.spritesheet)
}
