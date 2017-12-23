package math

import "math"

func Lerp(t, p0, p1 float64) float64 {
	return (1-t)*p0 + t*p1
}

func Radians(deg float64) float64 {
	return (math.Pi / 180) * deg
}

func Degress(rad float64) float64 {
	return (180 / math.Pi) * rad
}
