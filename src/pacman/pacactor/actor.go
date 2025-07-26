package pacactor

import (
	"github.com/jeppe9821/go-pacman/src/pacman/pacmap"
	"github.com/jeppe9821/go-pacman/src/pacman/pacmath"
	"github.com/jeppe9821/go-pacman/src/render"
)

type Sprite struct {
	Position pacmath.Vector2
	Sprite   render.Screen
}

type Actor struct {
	Position         pacmath.Position2D
	TilePosition     pacmap.TilePosition
	DirectionVector  pacmath.Vector2
	Behaviours       []Behaviour
	CurrentBehaviour Behaviour
}

func CreateActor(tilePosition pacmap.TilePosition, behaviours ...Behaviour) *Actor {
	return &Actor{
		Position: pacmath.Position2D{
			X: uint(tilePosition.X) << pacmap.ONE_TILE,
			Y: uint(tilePosition.Y) << pacmap.ONE_TILE,
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
