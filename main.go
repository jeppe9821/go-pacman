package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jeppe9821/go-pacman/pacman"
	"github.com/jeppe9821/go-pacman/render"
)

const WINDOW_WIDTH int = 448
const WINDOW_HEIGHT int = 564

type Game struct {
	frame  *ebiten.Image
	screen render.Renderable
	pacman pacman.Pacman
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.frame == nil {
		g.frame = ebiten.NewImage(WINDOW_WIDTH, WINDOW_HEIGHT)
	}

	g.frame.ReadPixels(g.screen.GetPixelsRaw())
	g.pacman.Draw(g.screen)
	g.frame.WritePixels(g.screen.GetPixelsRaw())
	screen.DrawImage(g.frame, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WINDOW_WIDTH >> 1,
		WINDOW_HEIGHT >> 1
}

func main() {
	game := &Game{
		screen: render.Create(WINDOW_WIDTH, WINDOW_HEIGHT, nil),
		pacman: pacman.Create(),
	}

	ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Go-Pacman")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
