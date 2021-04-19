package validator

import (
	mercantileerrors "github.com/Fa7C0n/mercantile-go/errors"
	log "github.com/sirupsen/logrus"
)

func ValidateLongs(longs ...float64) error {
	for long := range longs {
		if long < -180.0 || long > 180.0 {
			errmsg := "longitude should be within the range of -180 to 180"
			log.Error(errmsg)
			return mercantileerrors.NewCoordinateOutOfRangeError(errmsg)
		}
	}

	return nil
}

func ValidateLats(lats ...float64) error {
	for lat := range lats {
		if lat < -90.0 || lat > 90.0 {
			errmsg := "latitude should be within the range of -90 to 90"
			log.Errorf(errmsg)
			return mercantileerrors.NewCoordinateOutOfRangeError(errmsg)
		}
	}

	return nil
}

// TODO: add proper validation for mercator limits.
func ValidateEastings(eastings ...float64) error {
	return nil
}

// TODO: add proper validation for mercator limits.
func ValidateNorthings(northings ...float64) error {
	return nil
}
