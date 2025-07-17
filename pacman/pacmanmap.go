package pacman

import "github.com/jeppe9821/go-pacman/render"

var SpriteMap = map[string]SpritePos{
	"corner_R_TOP_outer":           {x: 8 * 0, y: 8 * 0},
	"corner_L_TOP_outer":           {x: 8 * 1, y: 8 * 0},
	"straight_R_SIDE_outer":        {x: 8 * 2, y: 8 * 0},
	"straight_L_SIDE_outer":        {x: 8 * 3, y: 8 * 0},
	"corner_R_BOTTOM_outer":        {x: 8 * 4, y: 8 * 0},
	"corner_L_BOTTOM_outer":        {x: 8 * 5, y: 8 * 0},
	"straight_corner_L_TOP":        {x: 8 * 6, y: 8 * 0},
	"straight_corner_L_BOTTOM":     {x: 8 * 7, y: 8 * 0},
	"straight_corner_R_TOP":        {x: 8 * 8, y: 8 * 0},
	"straight_corner_R_BOTTOM":     {x: 8 * 9, y: 8 * 0},
	"straight_TOP_outer":           {x: 8 * 10, y: 8 * 0},
	"straight_TOP_outer2":          {x: 8 * 11, y: 8 * 0},
	"straight_BOTTOM_outer":        {x: 8 * 12, y: 8 * 0},
	"straight_BOTTOM_outer2":       {x: 8 * 13, y: 8 * 0},
	"straight_inner":               {x: 8 * 14, y: 8 * 0},
	"straight_inner2":              {x: 8 * 15, y: 8 * 0},
	"unknown5":                     {x: 8 * 16, y: 8 * 0},
	"unknown6":                     {x: 8 * 17, y: 8 * 0},
	"unknown7":                     {x: 8 * 18, y: 8 * 0},
	"unknown8":                     {x: 8 * 19, y: 8 * 0},
	"straight_TOP_inner":           {x: 8 * 20, y: 8 * 0},
	"straight_TOP_inner2":          {x: 8 * 21, y: 8 * 0},
	"corner_R_TOP_inner_SHORT":     {x: 8 * 22, y: 8 * 0},
	"corner_L_TOP_inner_SHORT":     {x: 8 * 23, y: 8 * 0},
	"straight_SIDE_R_inner":        {x: 8 * 24, y: 8 * 0},
	"straight_SIDE_L_inner":        {x: 8 * 25, y: 8 * 0},
	"corner_R_BOTTOM_inner_SHORT":  {x: 8 * 26, y: 8 * 0},
	"corner_L_BOTTOM_inner_SHORT":  {x: 8 * 27, y: 8 * 0},
	"corner_R_TOP_ghost":           {x: 8 * 28, y: 8 * 0},
	"corner_L_TOP_ghost":           {x: 8 * 29, y: 8 * 0},
	"corner_R_BOTTOM_ghost":        {x: 8 * 30, y: 8 * 0},
	"door_L_ghost":                 {x: 8 * 0, y: 8 * 1},
	"door_R_ghost":                 {x: 8 * 1, y: 8 * 1},
	"corner_L_TOP_inner_LONG":      {x: 8 * 2, y: 8 * 1},
	"corner_R_TOP_inner_LONG":      {x: 8 * 3, y: 8 * 1},
	"corner_L_BOTTOM_inner_LONG":   {x: 8 * 4, y: 8 * 1},
	"corner_R_BOTTOM_inner_LONG":   {x: 8 * 5, y: 8 * 1},
	"corner_R_TOP_inner_SHORT2":    {x: 8 * 6, y: 8 * 1},
	"corner_L_TOP_inner_SHORT2":    {x: 8 * 7, y: 8 * 1},
	"corner_R_BOTTOM_inner_SHORT2": {x: 8 * 8, y: 8 * 1},
	"corner_L_BOTTOM_inner_SHORT2": {x: 8 * 9, y: 8 * 1},
	"straight_corner_outer_L":      {x: 8 * 10, y: 8 * 1},
	"straight_corner_outer_R":      {x: 8 * 11, y: 8 * 1},
	"straight_weird_inbetween":     {x: 8 * 12, y: 8 * 1},
}

type SpritePos struct {
	x int
	y int
}

type PacmanMap struct {
	spritesheet render.Renderable
}

func (pacman PacmanMap) DrawPart(screen render.Renderable, x int, y int, id string) {
	var yGameOffset = 24 / 8

	screen.DrawPart(x*8, (y+yGameOffset)*8, SpriteMap[id].x, SpriteMap[id].y, 8, 8, false, false, pacman.spritesheet)
}

func (pacman PacmanMap) DrawRect(screen render.Renderable, xStart int, yStart int, w int, h int) {
	//corners
	pacman.DrawPart(screen, xStart, yStart, "corner_L_TOP_inner_SHORT2")
	pacman.DrawPart(screen, xStart+w, yStart, "corner_R_TOP_inner_SHORT2")
	pacman.DrawPart(screen, xStart, yStart+h, "corner_L_BOTTOM_inner_SHORT2")
	pacman.DrawPart(screen, xStart+w, yStart+h, "corner_R_BOTTOM_inner_SHORT2")

	for i := xStart; i < xStart+w-1; i++ {
		pacman.DrawPart(screen, i+1, yStart, "straight_inner")
		pacman.DrawPart(screen, i+1, yStart+h, "straight_TOP_inner")
	}
	for i := yStart; i < yStart+h-1; i++ {
		pacman.DrawPart(screen, xStart, i+1, "straight_SIDE_L_inner")
		pacman.DrawPart(screen, xStart+w, i+1, "straight_SIDE_R_inner")
	}
}

func (pacman PacmanMap) DrawWeirdPenis(screen render.Renderable) {
	pacman.DrawPart(screen, 7, 6, "corner_L_TOP_inner_SHORT2")
	pacman.DrawPart(screen, 8, 6, "corner_R_TOP_inner_SHORT2")

	pacman.DrawPart(screen, 7, 7, "straight_SIDE_L_inner")
	pacman.DrawPart(screen, 7, 8, "straight_SIDE_L_inner")
	pacman.DrawPart(screen, 7, 9, "straight_SIDE_L_inner")
	pacman.DrawPart(screen, 7, 10, "straight_SIDE_L_inner")
	pacman.DrawPart(screen, 7, 11, "straight_SIDE_L_inner")
	pacman.DrawPart(screen, 7, 12, "straight_SIDE_L_inner")

	pacman.DrawPart(screen, 7, 13, "corner_L_BOTTOM_inner_SHORT2")
	pacman.DrawPart(screen, 8, 13, "corner_R_BOTTOM_inner_SHORT2")

	pacman.DrawPart(screen, 8, 10, "corner_L_TOP_inner_SHORT")
	pacman.DrawPart(screen, 8, 11, "straight_SIDE_R_inner")
	pacman.DrawPart(screen, 8, 12, "straight_SIDE_R_inner")

	pacman.DrawPart(screen, 8, 8, "straight_SIDE_R_inner")
	pacman.DrawPart(screen, 8, 7, "straight_SIDE_R_inner")

}

func (pacman PacmanMap) DrawInner(screen render.Renderable) {
	//L
	pacman.DrawRect(screen, 2, 2, 3, 2)
	pacman.DrawRect(screen, 7, 2, 4, 2)
	pacman.DrawRect(screen, 2, 6, 3, 1)

	//R
	pacman.DrawRect(screen, 16, 2, 4, 2)
	pacman.DrawRect(screen, 22, 2, 3, 2)
	pacman.DrawRect(screen, 22, 6, 3, 1)

	pacman.DrawRect(screen, 7, 15, 1, 4)
	pacman.DrawRect(screen, 7, 21, 4, 1)
	pacman.DrawRect(screen, 16, 21, 4, 1)

	pacman.DrawWeirdPenis(screen)
}

func (pacman PacmanMap) DrawOuter(screen render.Renderable) {
	/** START OF OUTER DRAW **/
	/////////////////////// TOP ///////////////////////
	pacman.DrawPart(screen, 0, 0, "corner_L_TOP_outer")
	pacman.DrawPart(screen, 1, 0, "straight_TOP_outer")
	pacman.DrawPart(screen, 2, 0, "straight_TOP_outer")
	pacman.DrawPart(screen, 3, 0, "straight_TOP_outer")
	pacman.DrawPart(screen, 4, 0, "straight_TOP_outer")
	pacman.DrawPart(screen, 5, 0, "straight_TOP_outer")
	pacman.DrawPart(screen, 6, 0, "straight_TOP_outer")
	pacman.DrawPart(screen, 7, 0, "straight_TOP_outer")
	pacman.DrawPart(screen, 8, 0, "straight_TOP_outer")
	pacman.DrawPart(screen, 9, 0, "straight_TOP_outer")
	pacman.DrawPart(screen, 10, 0, "straight_TOP_outer")
	pacman.DrawPart(screen, 11, 0, "straight_TOP_outer")
	pacman.DrawPart(screen, 12, 0, "straight_TOP_outer")
	pacman.DrawPart(screen, 13, 0, "straight_corner_outer_R")
	pacman.DrawPart(screen, 14, 0, "straight_weird_inbetween")
	pacman.DrawPart(screen, 15, 0, "straight_corner_outer_L")
	pacman.DrawPart(screen, 16, 0, "straight_TOP_outer")
	pacman.DrawPart(screen, 17, 0, "straight_TOP_outer")
	pacman.DrawPart(screen, 18, 0, "straight_TOP_outer")
	pacman.DrawPart(screen, 19, 0, "straight_TOP_outer")
	pacman.DrawPart(screen, 20, 0, "straight_TOP_outer")
	pacman.DrawPart(screen, 21, 0, "straight_TOP_outer")
	pacman.DrawPart(screen, 22, 0, "straight_TOP_outer")
	pacman.DrawPart(screen, 23, 0, "straight_TOP_outer")
	pacman.DrawPart(screen, 24, 0, "straight_TOP_outer")
	pacman.DrawPart(screen, 25, 0, "straight_TOP_outer")
	pacman.DrawPart(screen, 26, 0, "straight_TOP_outer")
	pacman.DrawPart(screen, 27, 0, "corner_R_TOP_outer")

	///////////////////// SIDE_L_PART1 /////////////////////
	pacman.DrawPart(screen, 0, 1, "straight_L_SIDE_outer")
	pacman.DrawPart(screen, 0, 2, "straight_L_SIDE_outer")
	pacman.DrawPart(screen, 0, 3, "straight_L_SIDE_outer")
	pacman.DrawPart(screen, 0, 4, "straight_L_SIDE_outer")
	pacman.DrawPart(screen, 0, 5, "straight_L_SIDE_outer")
	pacman.DrawPart(screen, 0, 6, "straight_L_SIDE_outer")
	pacman.DrawPart(screen, 0, 7, "straight_L_SIDE_outer")
	pacman.DrawPart(screen, 0, 8, "straight_L_SIDE_outer")

	pacman.DrawPart(screen, 0, 9, "corner_L_BOTTOM_outer")
	pacman.DrawPart(screen, 1, 9, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 2, 9, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 3, 9, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 4, 9, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 5, 9, "corner_R_TOP_inner_SHORT2")

	pacman.DrawPart(screen, 5, 10, "straight_L_SIDE_outer")
	pacman.DrawPart(screen, 5, 11, "straight_L_SIDE_outer")
	pacman.DrawPart(screen, 5, 12, "straight_L_SIDE_outer")

	pacman.DrawPart(screen, 5, 13, "corner_R_BOTTOM_inner_SHORT2")
	pacman.DrawPart(screen, 4, 13, "straight_TOP_outer")
	pacman.DrawPart(screen, 3, 13, "straight_TOP_outer")
	pacman.DrawPart(screen, 2, 13, "straight_TOP_outer")
	pacman.DrawPart(screen, 1, 13, "straight_TOP_outer")
	pacman.DrawPart(screen, 0, 13, "straight_TOP_outer")

	///////////////////// SIDE_L_PART2 /////////////////////
	pacman.DrawPart(screen, 0, 15, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 1, 15, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 2, 15, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 3, 15, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 4, 15, "straight_BOTTOM_outer")

	pacman.DrawPart(screen, 5, 15, "corner_R_TOP_inner_SHORT2")
	pacman.DrawPart(screen, 5, 16, "straight_L_SIDE_outer")
	pacman.DrawPart(screen, 5, 17, "straight_L_SIDE_outer")
	pacman.DrawPart(screen, 5, 18, "straight_L_SIDE_outer")
	pacman.DrawPart(screen, 5, 19, "corner_R_BOTTOM_inner_SHORT2")

	pacman.DrawPart(screen, 4, 19, "straight_TOP_outer")
	pacman.DrawPart(screen, 3, 19, "straight_TOP_outer")
	pacman.DrawPart(screen, 2, 19, "straight_TOP_outer")
	pacman.DrawPart(screen, 1, 19, "straight_TOP_outer")
	pacman.DrawPart(screen, 0, 19, "corner_L_TOP_outer")

	pacman.DrawPart(screen, 0, 20, "straight_L_SIDE_outer")
	pacman.DrawPart(screen, 0, 21, "straight_L_SIDE_outer")
	pacman.DrawPart(screen, 0, 22, "straight_L_SIDE_outer")
	pacman.DrawPart(screen, 0, 23, "straight_L_SIDE_outer")

	//weird nip thing
	pacman.DrawPart(screen, 0, 24, "straight_corner_L_BOTTOM")
	pacman.DrawPart(screen, 1, 24, "straight_inner")
	pacman.DrawPart(screen, 2, 24, "corner_R_TOP_inner_SHORT2")
	pacman.DrawPart(screen, 2, 25, "corner_R_BOTTOM_inner_SHORT2")
	pacman.DrawPart(screen, 1, 25, "straight_TOP_inner")
	pacman.DrawPart(screen, 0, 25, "straight_corner_R_BOTTOM")

	pacman.DrawPart(screen, 0, 26, "straight_L_SIDE_outer")
	pacman.DrawPart(screen, 0, 27, "straight_L_SIDE_outer")
	pacman.DrawPart(screen, 0, 28, "straight_L_SIDE_outer")
	pacman.DrawPart(screen, 0, 29, "straight_L_SIDE_outer")
	pacman.DrawPart(screen, 0, 30, "corner_L_BOTTOM_outer")

	///////////////////// FOOTER /////////////////////
	pacman.DrawPart(screen, 1, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 2, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 3, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 4, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 5, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 6, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 7, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 8, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 9, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 10, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 11, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 12, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 13, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 14, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 15, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 16, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 17, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 18, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 19, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 20, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 21, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 22, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 23, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 24, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 25, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 26, 30, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 27, 30, "corner_R_BOTTOM_outer")

	///////////////////// SIDE_R_PART1 /////////////////////
	pacman.DrawPart(screen, 27, 1, "straight_R_SIDE_outer")
	pacman.DrawPart(screen, 27, 2, "straight_R_SIDE_outer")
	pacman.DrawPart(screen, 27, 3, "straight_R_SIDE_outer")
	pacman.DrawPart(screen, 27, 4, "straight_R_SIDE_outer")
	pacman.DrawPart(screen, 27, 5, "straight_R_SIDE_outer")
	pacman.DrawPart(screen, 27, 6, "straight_R_SIDE_outer")
	pacman.DrawPart(screen, 27, 7, "straight_R_SIDE_outer")
	pacman.DrawPart(screen, 27, 8, "straight_R_SIDE_outer")

	pacman.DrawPart(screen, 27, 9, "corner_R_BOTTOM_outer")
	pacman.DrawPart(screen, 26, 9, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 25, 9, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 24, 9, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 23, 9, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 22, 9, "corner_L_TOP_inner_SHORT2")

	pacman.DrawPart(screen, 22, 10, "straight_R_SIDE_outer")
	pacman.DrawPart(screen, 22, 11, "straight_R_SIDE_outer")
	pacman.DrawPart(screen, 22, 12, "straight_R_SIDE_outer")

	pacman.DrawPart(screen, 22, 13, "corner_L_BOTTOM_inner_SHORT2")
	pacman.DrawPart(screen, 23, 13, "straight_TOP_outer")
	pacman.DrawPart(screen, 24, 13, "straight_TOP_outer")
	pacman.DrawPart(screen, 25, 13, "straight_TOP_outer")
	pacman.DrawPart(screen, 26, 13, "straight_TOP_outer")
	pacman.DrawPart(screen, 27, 13, "straight_TOP_outer")

	///////////////////// SIDE_R_PART2 /////////////////////
	pacman.DrawPart(screen, 27, 15, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 26, 15, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 25, 15, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 24, 15, "straight_BOTTOM_outer")
	pacman.DrawPart(screen, 23, 15, "straight_BOTTOM_outer")

	pacman.DrawPart(screen, 22, 15, "corner_L_TOP_inner_SHORT2")
	pacman.DrawPart(screen, 22, 16, "straight_R_SIDE_outer")
	pacman.DrawPart(screen, 22, 17, "straight_R_SIDE_outer")
	pacman.DrawPart(screen, 22, 18, "straight_R_SIDE_outer")
	pacman.DrawPart(screen, 22, 19, "corner_L_BOTTOM_inner_SHORT2")

	pacman.DrawPart(screen, 23, 19, "straight_TOP_outer")
	pacman.DrawPart(screen, 24, 19, "straight_TOP_outer")
	pacman.DrawPart(screen, 25, 19, "straight_TOP_outer")
	pacman.DrawPart(screen, 26, 19, "straight_TOP_outer")
	pacman.DrawPart(screen, 27, 19, "corner_R_TOP_outer")

	pacman.DrawPart(screen, 27, 20, "straight_R_SIDE_outer")
	pacman.DrawPart(screen, 27, 21, "straight_R_SIDE_outer")
	pacman.DrawPart(screen, 27, 22, "straight_R_SIDE_outer")
	pacman.DrawPart(screen, 27, 23, "straight_R_SIDE_outer")

	pacman.DrawPart(screen, 27, 24, "straight_corner_L_TOP")
	pacman.DrawPart(screen, 26, 24, "straight_inner")
	pacman.DrawPart(screen, 25, 24, "corner_L_TOP_inner_SHORT2")
	pacman.DrawPart(screen, 25, 25, "corner_L_BOTTOM_inner_SHORT2")
	pacman.DrawPart(screen, 26, 25, "straight_TOP_inner")
	pacman.DrawPart(screen, 27, 25, "straight_corner_R_TOP")

	pacman.DrawPart(screen, 27, 26, "straight_R_SIDE_outer")
	pacman.DrawPart(screen, 27, 27, "straight_R_SIDE_outer")
	pacman.DrawPart(screen, 27, 28, "straight_R_SIDE_outer")
	pacman.DrawPart(screen, 27, 29, "straight_R_SIDE_outer")
	pacman.DrawPart(screen, 27, 30, "corner_R_BOTTOM_outer")
	/** END OF OUTER DRAW **/
}
