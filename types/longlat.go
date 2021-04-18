package types

import validator "github.com/Fa7C0n/mercantile-go/types/validators"

// LongLat is a pair of Longitude and Latitude in decimel degrees and of CRS EPSG:4326
type LongLat struct {
	Long float64
	Lat  float64
}

// NewLongLat take a pair of longitude and latitude and return a pointer to LongLat
func NewLongLat(long float64, lat float64) (*LongLat, error) {
	err := validator.ValidateLongs(long)
	if err != nil {
		return nil, err
	}
	err = validator.ValidateLats(lat)
	if err != nil {
		return nil, err
	}

	return &LongLat{
		Long: long,
		Lat:  lat,
	}, nil
}
