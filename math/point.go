package math

import "math"

type Point3 struct {
	X, Y, Z float64
}

func (p Point3) HasNaN() bool {
	return math.IsNaN(p.X) || math.IsNaN(p.Y) || math.IsNaN(p.Z)
}

func (p Point3) ToVec() Vec3 {
	return Vec3{p.X, p.Y, p.Z}
}

func (p Point3) AddV(v Vec3) Point3 {
	return Point3{p.X + v.X, p.Y + v.Y, p.Z + v.Z}
}

// So that it's possible to do weighted sums of points later on
func (p Point3) AddP(p2 Point3) Point3 {
	return Point3{p.X + p2.X, p.Y + p2.Y, p.Z + p2.Z}
}

func (p Point3) SubtractV(v2 Vec3) Point3 {
	inv := v2.Inverse()
	return p.AddV(inv)
}

func (p Point3) SubtractP(p2 Point3) Vec3 {
	return Vec3{p.X - p2.X, p.Y - p2.Y, p.Z - p2.Z}
}

// Same reasoning as AddP
func (p Point3) Multiply(f float64) Point3 {
	return Point3{f * p.X, f * p.Y, f * p.Z}
}

func (p Point3) Divide(f float64) Point3 {
	return Point3{p.X / f, p.Y / f, p.Z / f}
}

type Point2 struct {
	X, Y float64
}

func (p Point2) HasNaN() bool {
	return math.IsNaN(p.X) || math.IsNaN(p.Y)
}

func (p Point2) ToVec() Vec2 {
	return Vec2{p.X, p.Y}
}

func (p Point2) AddV(v Vec2) Point2 {
	return Point2{p.X + v.X, p.Y + v.Y}
}

func (p Point2) AddP(p2 Point2) Point2 {
	return Point2{p.X + p2.X, p.Y + p2.Y}
}

func (p Point2) SubtractV(v2 Vec2) Point2 {
	inv := v2.Inverse()
	return p.AddV(inv)
}

func (p Point2) SubtractP(p2 Point2) Vec2 {
	return Vec2{p.X - p2.X, p.Y - p2.Y}
}

// Same reasoning as AddP
func (p Point2) Multiply(f float64) Point2 {
	return Point2{f * p.X, f * p.Y}
}

func (p Point2) Divide(f float64) Point2 {
	return Point2{p.X / f, p.Y / f}
}

// ------ Exported helpers --------

func DistanceP3(p1, p2 Point3) float64 {
	return p2.SubtractP(p1).Magnitude()
}

func DistanceP3Sq(p1, p2 Point3) float64 {
	return p2.SubtractP(p1).MagnitudeSq()
}

func DistanceP2(p1, p2 Point2) float64 {
	return p2.SubtractP(p1).Magnitude()
}

func DistanceP2Sq(p1, p2 Point2) float64 {
	return p2.SubtractP(p1).MagnitudeSq()
}

func LerpP3(f float64, p1, p2 Point3) Point3 {
	return p1.Multiply(1 - f).AddP(p2.Multiply(f))
}

func LerpP2(f float64, p1, p2 Point2) Point2 {
	return p1.Multiply(1 - f).AddP(p2.Multiply(f))
}

func MinP3(p1, p2 Point3) Point3 {
	return Point3{math.Min(p1.X, p2.X), math.Min(p1.Y, p2.Y), math.Min(p1.Z, p2.Z)}
}

func MinP2(p1, p2 Point2) Point2 {
	return Point2{math.Min(p1.X, p2.X), math.Min(p1.Y, p2.Y)}
}

func MaxP3(p1, p2 Point3) Point3 {
	return Point3{math.Max(p1.X, p2.X), math.Max(p1.Y, p2.Y), math.Max(p1.Z, p2.Z)}
}

func MaxP2(p1, p2 Point2) Point2 {
	return Point2{math.Max(p1.X, p2.X), math.Max(p1.Y, p2.Y)}
}

func FloorP3(p *Point3) Point3 {
	return Point3{math.Floor(p.X), math.Floor(p.Y), math.Floor(p.Z)}
}

func FloorP2(p Point2) Point2 {
	return Point2{math.Floor(p.X), math.Floor(p.Y)}
}

func CeilP3(p *Point3) Point3 {
	return Point3{math.Ceil(p.X), math.Ceil(p.Y), math.Ceil(p.Z)}
}

func CeilP2(p Point2) Point2 {
	return Point2{math.Ceil(p.X), math.Ceil(p.Y)}
}

func AbsP3(p *Point3) Point3 {
	return Point3{math.Abs(p.X), math.Abs(p.Y), math.Abs(p.Z)}
}

func AbsP2(p Point2) Point2 {
	return Point2{math.Abs(p.X), math.Abs(p.Y)}
}
