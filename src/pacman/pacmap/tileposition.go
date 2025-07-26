package pacmap

const (
	ONE_TILE   = 3
	TILE_EMPTY = 0
	TILE_FOOD  = 45
)

type TilePosition struct {
	X uint8
	Y uint8
}

func (tp TilePosition) GetDistanceTo(other TilePosition) uint8 {
	var xDistance uint8
	var yDistance uint8

	if other.X > tp.X {
		xDistance = other.X - tp.X
	} else {
		xDistance = tp.X - other.X
	}

	if other.Y > tp.Y {
		yDistance = other.Y - tp.Y
	} else {
		yDistance = tp.Y - other.Y
	}

	return xDistance + yDistance
}
