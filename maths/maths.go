package maths

import "math"

// Round is a light wrapper around math.Round. It returns rounding of the given
// float to the nearest unit provided.
func Round(x float64, unit float64) float64 {
	return math.Round(x/unit) * unit
}
