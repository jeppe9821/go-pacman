package pacactor

import (
	"github.com/jeppe9821/go-pacman/pacman/pacmap"
	"github.com/jeppe9821/go-pacman/pacman/pacmath"
	"github.com/jeppe9821/go-pacman/render"
)

type Sprite struct {
	Position pacmath.Vector2
	Sprite   render.Screen
}

type Actor struct {
	Position         pacmath.Position2D
	TilePosition     pacmath.TilePosition
	DirectionVector  pacmath.Vector2
	Behaviours       []Behaviour
	CurrentBehaviour Behaviour
}

func CreateActor(tilePosition pacmath.TilePosition, behaviours ...Behaviour) *Actor {
	return &Actor{
		Position: pacmath.Position2D{
			X: uint(tilePosition.X) << pacmath.ONE_TILE,
			Y: uint(tilePosition.Y) << pacmath.ONE_TILE,
		},
		TilePosition:     tilePosition,
		DirectionVector:  pacmath.Vector2{X: -1, Y: 0},
		Behaviours:       behaviours,
		CurrentBehaviour: behaviours[0],
	}
}

type Behaviour interface {
	Update(actor *Actor, tileMap pacmap.TileMap, tick int)
	Draw(actor Actor, screen render.Screen, spritesheet render.Spritesheet8Bit)
}

/*
func (p *Entity) GetCurrentTile(grid [][]int) (int, int) {
	return p.xTile / 8, p.yTile / 8
}

func (p *Entity) GetNextHorizontalTile(grid [][]int, dir int) int {
	return p.xTile + (8 * (dir*2 - 1))
}

func (p *Entity) GetNextVerticalTile(grid [][]int, dir int) int {
	return p.yTile + (8 * ((dir%2)*2 - 1))
}

func (p *Entity) GetClosestTileToTarget(grid [][]int, xp int, yp int, x2 int, y2 int) (int, bool) {
	log.Println("Try same direction")
	//figure out which directions we're allowed to walk
	allowedTilesActualTile := make([]int, 4)
	allowedTiles := make([]int, 4)
	allowedTilesDir := make([]bool, 4)
	for i := 0; i < 4; i++ {
		allowedTiles[i] = 9999

		var tile int
		if i < 2 {
			tile = p.GetNextHorizontalTile(grid, i)

			if tile < 0 {
				continue
			}

			if p.CanMoveNextHorizontalTile(grid, tile) {
				allowedTilesActualTile[i] = tile
				xx := math.Abs(float64(tile - xp)) //* math.Abs(float64(tile-xp))
				yy := math.Abs(float64(y2 - yp))   //* math.Abs(float64(y2-yp))
				allowedTiles[i] = int(xx + yy)
				allowedTilesDir[i] = true
			}
		} else {
			tile = p.GetNextVerticalTile(grid, i)

			if tile < 0 {
				continue
			}

			if p.CanMoveNextVerticalTile(grid, tile) {
				allowedTilesActualTile[i] = tile
				xx := math.Abs(float64(x2 - xp))   //* math.Abs(float64(x2-xp))
				yy := math.Abs(float64(tile - yp)) //* math.Abs(float64(tile-yp))
				allowedTiles[i] = int(xx + yy)
				allowedTilesDir[i] = false
			}
		}
	}

	biggest := 0
	howManyDoWeHave := 0
	for i := 0; i < 4; i++ {
		if allowedTiles[i] != 9999 {
			howManyDoWeHave++
		}
		if allowedTiles[i] < allowedTiles[biggest] {
			biggest = i
		}
	}

	hlastTile := -1
	vlastTile := -1
	newDir := -1

	if p.dir == 0 {
		hlastTile = p.GetNextHorizontalTile(grid, 1)
		newDir = 1
	} else if p.dir == 1 {
		hlastTile = p.GetNextHorizontalTile(grid, 0)
		newDir = 0
	} else if p.dir == 2 {
		vlastTile = p.GetNextVerticalTile(grid, 3)
		newDir = 3
	} else if p.dir == 3 {
		vlastTile = p.GetNextVerticalTile(grid, 2)
		newDir = 2
	}

	if howManyDoWeHave == 1 && (allowedTilesActualTile[biggest] == hlastTile || allowedTilesActualTile[biggest] == vlastTile) {
		//move backwards if its our only option
		log.Println("WE ARE MOVING BACKWARDS!")
		p.dir = newDir

		lastTile := -1

		if allowedTilesActualTile[biggest] == hlastTile {
			lastTile = hlastTile
		} else {
			lastTile = vlastTile
		}

		return lastTile, allowedTilesDir[biggest]
	} else if howManyDoWeHave > 1 && (allowedTilesActualTile[biggest] == hlastTile || allowedTilesActualTile[biggest] == vlastTile) {
		//choose any of the other ones if we're trying to move backwards but have other options
		log.Println("We are doing edge case scenario")
		for i := 2; ; {
			if allowedTiles[i] == 9999 || (allowedTilesActualTile[i] == hlastTile || allowedTilesActualTile[i] == vlastTile) {

				log.Printf("H_ARE THEY THE SAME? %v, %v", allowedTilesActualTile[i], hlastTile)
				log.Printf("V_ARE THEY THE SAME? %v, %v", allowedTilesActualTile[i], vlastTile)

				if allowedTilesActualTile[i] == hlastTile {
					if i == 0 {
						i = 3
					} else if i == 1 {
						i = 2
					}

					log.Printf("H_LAST_TILE_STUCK?: %v, %v, %v, %v", allowedTiles[0], allowedTiles[1], allowedTiles[2], allowedTiles[3])

					if allowedTiles[3] == 9999 || allowedTiles[2] == 9999 {
						if i == 2 {
							i = 0
						} else if i == 3 {
							i = 1
						}
					}
				} else if allowedTilesActualTile[i] == vlastTile {
					if i == 2 {
						i = 0
					} else if i == 3 {
						i = 1
					}

					log.Printf("i: %v, V_LAST_TILE_STUCK?: %v, %v, %v, %v", i, allowedTiles[0], allowedTiles[1], allowedTiles[2], allowedTiRenderable

					if allowedTiles[0] == 9999 || allowedTiles[1] == 9999 {
						if i == 0 {
							i = 2
						} else if i == 1 {
							i = 3
						}
					}
				} else {
					if i == 0 {
						i = 3
					} else if i == 1 {
						i = 2
					} else if i == 2 {
						i = 0
					} else if i == 3 {
						i = 1
					}
					continue
				}

				continue
			}

			p.dir = i
			return allowedTilesActualTile[i], allowedTilesDir[i]
		}
	}

	//Move normally
	p.dir = biggest
	return allowedTilesActualTile[biggest], allowedTilesDir[biggest]
}

func (p *Entity) CanMoveNextHorizontalTile(grid [][]int, nextTile int) bool {
	if grid[p.yTile/8][nextTile/8] == 46 || grid[p.yTile/8][nextTile/8] == 45 {
		return true
	} else {
		return false
	}
}

func (p *Entity) CanMoveNextVerticalTile(grid [][]int, nextTile int) bool {
	if grid[nextTile/8][p.xTile/8] == 46 || grid[nextTile/8][p.xTile/8] == 45 {
		return true
	} else {
		return false
	}
}
*/
