package pacman

import (
	"github.com/jeppe9821/go-pacman/render"
)

type Player struct {
	spriteIndex int
	x           int
	y           int
	xDir        int
}

func (p *Player) Update(tick int) {
	if tick == 0 {
		p.x = 13 * 8
		p.y = 25 * 8
		p.xDir = -1
	}
	if tick%5 == 0 {
		p.spriteIndex = (p.spriteIndex + 1) % 2
	}
	if tick%30 == 0 {
		p.x += 8 * p.xDir
	}
}

func (p *Player) Draw(screen render.Renderable, spritesheet render.Renderable) {
	xFlip := p.xDir == -1
	screen.DrawPart(p.x+4, p.y+4, p.spriteIndex*13, 8*2, 13, 13, xFlip, false, spritesheet)
}
