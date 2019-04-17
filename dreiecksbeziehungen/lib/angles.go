package lib

import "math"

func ToDegrees(r float64) float64 {
	return (r * 180) / math.Pi
}

func ToRadians(d float64) float64 {
	return d * (math.Pi / 180)
}
