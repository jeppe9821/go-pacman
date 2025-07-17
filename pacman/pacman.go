package pacman

import "github.com/jeppe9821/go-pacman/render"

type Pacman struct {
	pacmanMap PacmanMap
}

func Create() Pacman {
	return Pacman{
		pacmanMap: PacmanMap{
			spritesheet: render.LoadFromFile("assets/spritesheet.png"),
		},
	}
}

func (pacman Pacman) Draw(screen render.Renderable) {
	pacman.pacmanMap.DrawOuter(screen)
	pacman.pacmanMap.DrawInner(screen)
}
