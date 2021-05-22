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

type Tiles []*Tile

// Ul returns the top left longitude and latitude of a tile.
func (t *Tile) Ul() (*LongLat, error) {
	Z2 := math.Pow(2, float64(t.z))
	longDeg := float64(t.x)/Z2*360.0 - 180.0
	latRadian := math.Atan(math.Sinh(math.Pi * (1 - 2*float64(t.y)/Z2)))
	latDeg := latRadian * mercantile.R2D

	return NewLongLat(longDeg, latDeg)
}

// BoundsLatLong returns the bounding box of the Tile in LatLong
func (t *Tile) BoundsLatLong() (*LongLatBbox, error) {
	Z2 := math.Pow(2, float64(t.z))
	tlLongDeg := float64(t.x)/Z2*360.0 - 180.0
	tlLatRad := math.Atan(math.Sinh(math.Pi * (1 - 2*float64(t.y)/Z2)))
	tlLatDeg := tlLatRad * mercantile.R2D

	brLongDeg := float64(t.x+1)/Z2*360.0 - 180
	brLatRad := math.Atan(math.Sinh(math.Pi * (1 - 2*float64(t.y+1)/Z2)))
	brLatDeg := brLatRad * mercantile.R2D

	return NewLongLatBbox(tlLongDeg, brLatDeg, brLongDeg, tlLatDeg)
}

// IsValid checks whether the given tile coordinates are valid.
func (t *Tile) IsValid() bool {
	if t.x <= uint(math.Pow(2, float64(t.z))-1) {
		if t.y <= uint(math.Pow(2, float64(t.z)-1)) {
			return true
		}
	}
	return false
}

// Neighbours will return the neighbouring tiles
// The neighbors function makes no guarantees regarding neighbor tile
// ordering. The neighbors function returns up to eight neighboring tiles, where
// tiles will be omitted when they are not valid e.g. Tile(-1, -1, z).
func (t *Tile) Neighbours() (Tiles, error) {
	tiles := Tiles{}
	_, max := tileRange(t.z)

	for i := range [3]int{-1, 0, 1} {
		for j := range [3]int{-1, 0, 1} {
			if i == 0 && j == 0 {
				continue
			} else if int(t.x)+i < 0 || int(t.y)+j < 0 {
				continue
			} else if int(t.x)+i > int(max) || int(t.y)+1 > int(max) {
				continue
			}
			neighbourTile, err := NewTile(uint(int(t.x)+i), uint(int(t.y)+j), t.z)
			if err != nil {
				return nil, err
			}
			if neighbourTile.IsValid() {
				tiles = append(tiles, neighbourTile)
			}
		}
	}
	return tiles, nil
}

// Get the web mercator bounding box of a tile.
func (t *Tile) XYBounds() (*WebMercatorBbox, error) {
	tileSize := mercantile.CE / math.Pow(2, float64(t.z))
	left := float64(t.x)*tileSize - mercantile.CE/2
	right := left + tileSize

	top := (mercantile.CE / 2) - float64(t.y)*tileSize
	bottom := top - tileSize

	return NewWebMercatorBbox(left, bottom, right, top)
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
