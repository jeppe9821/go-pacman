package pacactor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jeppe9821/go-pacman/src/pacman/pacmap"
	"github.com/jeppe9821/go-pacman/src/pacman/pacmath"
	"github.com/jeppe9821/go-pacman/src/render"
)

type PlayerBehaviour struct {
	AnimationBehaviour *AnimationBehaviour
	oldDirectionVector pacmath.Vector2
}

func (pb *PlayerBehaviour) Tick(actor *Actor, tileMap pacmap.TileMap) bool {
	couldFindNext, nextTile := tileMap.TryGetNextWalkableTile(actor.TilePosition, actor.DirectionVector)
	result := true

	if !couldFindNext {
		couldFindNext, nextTile = tileMap.TryGetNextWalkableTile(actor.TilePosition, pb.oldDirectionVector)
		result = false
	}

	if couldFindNext {
		actor.TilePosition.X = nextTile.X
		actor.TilePosition.Y = nextTile.Y
	}

	return result
}

func (pb *PlayerBehaviour) Update(actor *Actor, tileMap pacmap.TileMap, tick int) {
	pb.AnimationBehaviour.Update(actor, tick)

	if tick == 0 {
		pb.oldDirectionVector = actor.DirectionVector
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		actor.DirectionVector = pacmath.Vector2{X: -1, Y: 0}
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		actor.DirectionVector = pacmath.Vector2{X: 1, Y: 0}
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		actor.DirectionVector = pacmath.Vector2{X: 0, Y: -1}
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		actor.DirectionVector = pacmath.Vector2{X: 0, Y: 1}
	}

	actor.Position.X = uint(int(actor.Position.X) - int(pacmath.Sign(int(actor.Position.X)-8*int(actor.TilePosition.X))))
	actor.Position.Y = uint(int(actor.Position.Y) - int(pacmath.Sign(int(actor.Position.Y)-8*int(actor.TilePosition.Y))))

	if actor.Position.Y == 8*uint(actor.TilePosition.Y) && actor.Position.X == 8*uint(actor.TilePosition.X) {
		if !pb.Tick(actor, tileMap) {
			actor.DirectionVector = pb.oldDirectionVector
		} else {
			pb.oldDirectionVector = actor.DirectionVector
		}
	}
}

func (pb PlayerBehaviour) Draw(actor Actor, screen render.Screen, spritesheet render.Spritesheet8Bit) {
	var xTileStart uint8 = uint8(pacmath.Abs(pb.oldDirectionVector.Y) * 4)
	var xFlip bool = pb.oldDirectionVector.X == -1
	var yFlip bool = pb.oldDirectionVector.Y == -1

	pb.AnimationBehaviour.Draw(-4, 20, xTileStart, 2, xFlip, yFlip, screen, spritesheet, actor.Position)
}
