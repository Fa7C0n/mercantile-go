package mercantileerrors

import "fmt"

type MercantileError struct {
	Message string
}

func (m *MercantileError) Error() string {
	return fmt.Sprintf("ERROR: %v", m.Message)
}

type InvalidLatitudeError struct {
	MercantileError
}

func NewInvalidLatitudeError(message string) *InvalidLatitudeError {
	return &InvalidLatitudeError{
		MercantileError{
			Message: message,
		},
	}
}

// coordinateOutOfRangeError is raised when the given coordinate pair doesn't conform
// to the coordinate limits of CRS
type CoordinateOutOfRangeError struct {
	MercantileError
}

// CoordinateOut
func NewCoordinateOutOfRangeError(message string) *CoordinateOutOfRangeError {
	return &CoordinateOutOfRangeError{
		MercantileError{
			Message: message,
		},
	}
}

type TileIndexOutOfRange struct {
	MercantileError
}

func NewTileIndexOutOfRange(message string) *TileIndexOutOfRange {
	return &TileIndexOutOfRange{
		MercantileError{
			Message: message,
		},
	}
}

type InvalidZoomError struct {
	MercantileError
}

type ParentTileError struct {
	MercantileError
}

type QuadKeyError struct {
	MercantileError
}

type TileArgParsingError struct {
	MercantileError
}

type TileError struct {
	MercantileError
}
