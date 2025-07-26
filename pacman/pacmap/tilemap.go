package pacmap

import (
	"math"

	"github.com/jeppe9821/go-pacman/pacman/pacmath"
	"github.com/jeppe9821/go-pacman/render"
)

type TileMap struct {
	Grid [31][28]uint8
}

func (tm TileMap) GetCurrentTile(tilePosition pacmath.TilePosition) uint8 {
	return tm.Grid[tilePosition.Y][tilePosition.X]
}

func (tm TileMap) TryGetNextWalkableTile(currentTilePosition pacmath.TilePosition, direction pacmath.Vector2) (bool, pacmath.TilePosition) {
	var nextTilePosition pacmath.TilePosition = currentTilePosition

	nextTilePosition.X = uint8(int(currentTilePosition.X) + direction.X)
	nextTilePosition.Y = uint8(int(currentTilePosition.Y) + direction.Y)

	if tm.Grid[nextTilePosition.Y][nextTilePosition.X] == 45 {
		return true, nextTilePosition
	}

	return false, currentTilePosition
}

func (tm TileMap) Draw(screen render.Screen, spritesheet render.Spritesheet8Bit) {
	var yGameOffset uint = 24
	var x, y uint

	for y = 0; y < uint(len(tm.Grid)); y++ {
		for x = 0; x < uint(len(tm.Grid[y])); x++ {
			if tm.Grid[y][x] == 0 {
				continue
			}

			var tile uint8 = tm.Grid[y][x]

			screen.Draw(
				x<<pacmath.ONE_TILE,
				(y<<pacmath.ONE_TILE)+yGameOffset,
				tile%25,
				uint8(math.Floor(float64(tile)/float64(25.0))),
				false,
				false,
				spritesheet)
		}
	}
}
