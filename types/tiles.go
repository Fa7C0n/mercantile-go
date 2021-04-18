package types

import (
	"math"

	"github.com/Fa7C0n/mercantile-go"
	mercantileerrors "github.com/Fa7C0n/mercantile-go/errors"
	log "github.com/sirupsen/logrus"
)

// Tile is an XYZ web mercator tile with x and y being the indices of the tile and
// z being the zoom level.
type Tile struct {
	x uint
	y uint
	z uint
}

// TopLeft returns the top left longitude and latitude of a tile.
func (t *Tile) TopLeft() (*LongLat, error) {
	Z2 := math.Pow(2, float64(t.z))
	longDeg := float64(t.x)/Z2*360.0 - 180.0
	latRadian := math.Atan(math.Sinh(math.Pi * (1 - 2*float64(t.y)/Z2)))
	latDeg := latRadian * mercantile.R2D

	return NewLongLat(longDeg, latDeg)
}

func (t *Tile) Bounds() (*LongLatBbox, error) {
	Z2 := math.Pow(2, float64(t.z))
	tlLongDeg := float64(t.x)/Z2*360.0 - 180.0
	tlLatRad := math.Atan(math.Sinh(math.Pi * (1 - 2*float64(t.y)/Z2)))
	tlLatDeg := tlLatRad * mercantile.R2D

	brLongDeg := float64(t.x+1)/Z2*360.0 - 180
	brLatRad := math.Atan(math.Sinh(math.Pi * (1 - 2*float64(t.y+1)/Z2)))
	brLatDeg := brLatRad * mercantile.R2D

	return NewLongLatBbox(tlLongDeg, brLatDeg, brLongDeg, tlLatDeg)
}

// NewTile takes XYZ tile index and returns a Pointer to Tile and an error.
func NewTile(x uint, y uint, z uint) (*Tile, error) {
	low, high := tileRange(z)

	if (x <= low && x >= high) || (y <= low && y >= high) {
		errMsg := "mercantile requires tile x and y to be within the range (0, 2 ^ zoom)"
		log.Error(errMsg)
		return nil, mercantileerrors.NewTileIndexOutOfRange(errMsg)
	}
	return &Tile{
		x: x,
		y: y,
		z: z,
	}, nil
}

func tileRange(zoom uint) (uint, uint) {
	var min uint = 0
	var max uint = 2 ^ zoom - 1

	return min, max
}
