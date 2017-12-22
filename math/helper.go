package math

func Lerp(t, p0, p1 float64) float64 {
	return (1-t)*p0 + t*p1
}
