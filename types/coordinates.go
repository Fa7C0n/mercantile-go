package types

import (
	"math"

	"github.com/Fa7C0n/mercantile-go"
	validator "github.com/Fa7C0n/mercantile-go/types/validators"
)

// LongLat is a pair of Longitude and Latitude in decimel degrees and of CRS EPSG:4326
type LongLat struct {
	Long float64
	Lat  float64
}

func (l *LongLat) XY() (*MercatorXY, error) {
	x := mercantile.RE * (l.Long * (math.Pi / 180.0))

	var y float64
	if l.Lat <= -90.0 {
		y = math.Inf(-1)
	} else if l.Lat >= 90.0 {
		y = math.Inf(1)
	} else {
		y = mercantile.RE * math.Log(math.Tan((math.Pi*0.25)+(0.5*(l.Lat*(math.Pi/180)))))
	}
	return NewMercatorXY(x, y)
}

// NewLongLat take a pair of longitude and latitude and return a pointer to LongLat
func NewLongLat(long float64, lat float64) (*LongLat, error) {
	if err := validator.ValidateLongs(long); err != nil {
		return nil, err
	}

	if err := validator.ValidateLats(lat); err != nil {
		return nil, err
	}

	return &LongLat{
		Long: long,
		Lat:  lat,
	}, nil
}

type MercatorXY struct {
	X float64
	Y float64
}

func (m *MercatorXY) ToLongLat() (*LongLat, error) {
	long := m.X * mercantile.R2D / mercantile.RE
	lat := ((math.Pi * 0.5) - 2.0*math.Atan(math.Exp(-m.Y/mercantile.RE))) * mercantile.R2D

	return NewLongLat(long, lat)
}

func NewMercatorXY(x float64, y float64) (*MercatorXY, error) {
	if err := validator.ValidateEastings(x); err != nil {
		return nil, err
	}

	if err := validator.ValidateNorthings(y); err != nil {
		return nil, err
	}

	return &MercatorXY{
		X: x,
		Y: y,
	}, nil
}
