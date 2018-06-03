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

func Clamp(v, min, max float64) (ret float64) {
	if min > max {
		return 0
	}
	if v < min {
		ret = min
	} else if v > max {
		ret = max
	}
	return ret
}

func Quadratic(a, b, c float64) (bool, float64, float64) {
	disc := b*b - 4*a*c
	if disc < 0 {
		return false, 0, 0
	}
	root := math.Sqrt(disc)
	q := -0.5 * (b + root)
	if b < 0 {
		q = -0.5 * (b - root)
	}
	t0, t1 := q/a, c/q
	if t1 < t0 {
		t0, t1 = t1, t0
	}
	return true, t0, t1
}
