package types

import validator "github.com/Fa7C0n/mercantile-go/types/validators"

type BBox struct {
	Left   float64
	Bottom float64
	Right  float64
	Top    float64
}

// LongLatBbox is a geographic bounding box in decimel degrees of CRS EPSG:4326.
type LongLatBbox struct {
	BBox
}

// NewLongLatBbox takes the left, bottom, right, top and returns LongLatBbox, error.
func NewLongLatBbox(left float64, bottom float64, right float64, top float64) (*LongLatBbox, error) {
	err := validator.ValidateLongs(left, right)
	if err != nil {
		return nil, err
	}
	err = validator.ValidateLats(bottom, top)
	if err != nil {
		return nil, err
	}
	return &LongLatBbox{
		BBox{
			Left:   left,
			Bottom: bottom,
			Right:  right,
			Top:    top,
		},
	}, nil
}

// WebMercatorBbox is a webMercator bounding box in of CRS EPSG:3857.
type WebMercatorBbox struct {
	BBox
}

// NewWebMercatorBbox takes left, bottom, right, top and return WebMercatorBbox, error.
func NewWebMercatorBbox(left float64, bottom float64, right float64, top float64) (*WebMercatorBbox, error) {
	err := validator.ValidateEastings(left, right)
	if err != nil {
		return nil, err
	}

	err = validator.ValidateNorthings(bottom, top)
	if err != nil {
		return nil, err
	}

	return &WebMercatorBbox{
		BBox{
			Left:   left,
			Bottom: bottom,
			Right:  right,
			Top:    top,
		},
	}, nil
}
