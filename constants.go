package mercantile

import "math"

const (
	R2D       = 180 / math.Pi
	RE        = 6378137.0
	CE        = 2 * math.Pi * RE
	EPSILON   = 1e-14
	LLEPSILON = 1e-11
)
