package core

import "math"

type Normal3 struct {
	X, Y, Z float64
}

func (n *Normal3) ToVec3() Vec3 {
	return Vec3{n.X, n.Y, n.Z}
}

func (n *Normal3) MagnitudeSq() float64 {
	return n.X*n.X + n.Y*n.Y + n.Z*n.Z
}

func (n *Normal3) Magnitude() float64 {
	return math.Sqrt(n.MagnitudeSq())
}

func (n *Normal3) HasNaN() bool {
	return math.IsNaN(n.X) || math.IsNaN(n.Y) || math.IsNaN(n.Z)
}

func (n *Normal3) Inverse() Normal3 {
	return Normal3{-n.X, -n.Y, -n.Z}
}

func (n *Normal3) Add(v2 Normal3) Normal3 {
	return Normal3{n.X + v2.X, n.Y + v2.Y, n.Z + v2.Z}
}

func (n *Normal3) Subtract(v2 Normal3) Normal3 {
	inv := v2.Inverse()
	return n.Add(inv)
}

func (n *Normal3) Multiply(f float64) Normal3 {
	return Normal3{f * n.X, f * n.Y, f * n.Z}
}

func (n *Normal3) Divide(f float64) Normal3 {
	return Normal3{n.X / f, n.Y / f, n.Z / f}
}

func (n *Normal3) Normalize() Normal3 {
	return n.Divide(n.Magnitude())
}

func (n *Normal3) MinComponent() float64 {
	return math.Min(n.X, math.Min(n.Y, n.Z))
}

func (n *Normal3) MaxComponent() float64 {
	return math.Max(n.X, math.Max(n.Y, n.Z))
}

func DotN3(m, n *Normal3) float64 {
	return m.X*n.X + m.Y*n.Y + m.Z*n.Z
}

func AbsDotN3(m, n *Normal3) float64 {
	return math.Abs(DotN3(m, n))
}

func CrossN3(m, n *Normal3) Normal3 {
	return Normal3{m.Y*n.Z - m.Z*n.Y, m.Z*n.X - m.X*n.Z, m.X*n.Y - m.Y*n.X}
}

func FaceForward(n *Normal3, v Vec3) Normal3 {
	ret := *n
	if DotV3(n.ToVec3(), v) < 0.0 {
		ret = ret.Inverse()
	}
	return ret
}
