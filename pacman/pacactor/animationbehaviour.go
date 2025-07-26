package pacactor

import (
	"github.com/jeppe9821/go-pacman/pacman/pacmath"
	"github.com/jeppe9821/go-pacman/render"
)

type AnimationBehaviour struct {
	CurrentSpriteIndex uint8
	startSpriteIndex   uint8
	endSpriteIndex     uint8
	updateEveryTick    int
}

func CreateNewAnimationBehaviour(startSpriteIndex uint8, endSpriteIndex uint8, updateEveryTick int) *AnimationBehaviour {
	return &AnimationBehaviour{
		CurrentSpriteIndex: 0,
		startSpriteIndex:   startSpriteIndex,
		endSpriteIndex:     endSpriteIndex,
		updateEveryTick:    updateEveryTick,
	}
}

func (ab *AnimationBehaviour) Update(actor *Actor, tick int) {
	if tick%ab.updateEveryTick == 0 {
		ab.CurrentSpriteIndex = ab.startSpriteIndex + (ab.CurrentSpriteIndex+1)%ab.endSpriteIndex
	}
}

func (ab AnimationBehaviour) Draw(xOffset int,
	yOffset int,
	xTileStart uint8,
	yTileStart uint8,
	xFlip bool,
	yFlip bool,
	screen render.Screen,
	spritesheet render.Spritesheet8Bit,
	position pacmath.Position2D) {

	var xSecondTileOffset int = 8

	if xFlip {
		xSecondTileOffset = -8
	}

	var ySecondTileOffset int = 8

	if yFlip {
		ySecondTileOffset = -8
	}

	if xSecondTileOffset < 0 {
		xOffset += 8
	}

	if ySecondTileOffset < 0 {
		yOffset += 8
	}

	var xp uint = uint(int(position.X) + xOffset)
	var yp uint = uint(int(position.Y) + yOffset)

	var xt uint8 = xTileStart + (ab.CurrentSpriteIndex * 2)
	var yt uint8 = yTileStart

	screen.Draw(xp, yp, xt, yt, xFlip, yFlip, spritesheet)
	screen.Draw(xp, uint(int(yp)+ySecondTileOffset), xt, yt+1, xFlip, yFlip, spritesheet)
	screen.Draw(uint(int(xp)+xSecondTileOffset), yp, xt+1, yt, xFlip, yFlip, spritesheet)
	screen.Draw(uint(int(xp)+xSecondTileOffset), uint(int(yp)+ySecondTileOffset), xt+1, yt+1, xFlip, yFlip, spritesheet)
}
