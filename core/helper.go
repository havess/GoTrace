package core

import "math"

func MachineEpsilon() float64 {
	return math.Nextafter(1, 2) - 1
}

func Gamma(n float64) float64 {
	return (n * MachineEpsilon()) / (1 - n*MachineEpsilon())
}

func Lerp(t, p0, p1 float64) float64 {
	return (1-t)*p0 + t*p1
}

func Radians(deg float64) float64 {
	return (math.Pi / 180) * deg
}

func Degress(rad float64) float64 {
	return (180 / math.Pi) * rad
}

func Xor(a, b bool) bool {
	return (a && !b) || (!a && b)
}
