package math

import "math"
// -------- Vec 3 ---------
type Vec3 struct {
    X, Y, Z float64
}

func (v Vec3) MagnitudeSq() float64 {
    return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vec3) Magnitude() float64 {
    return math.Sqrt(v.MagnitudeSq())
}

func (v Vec3) HasNaN() bool {
    return math.IsNaN(v.X) || math.IsNaN(v.Y) || math.IsNaN(v.Z)
}

func (v Vec3) Inverse() Vec3 {
    return Vec3{-v.X, -v.Y, -v.Z}
}

func (v Vec3) Add(v2* Vec3) Vec3 {
    return Vec3{v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
}

func (v Vec3) Subtract(v2* Vec3) Vec3 {
    inv := v2.Inverse()
    return v.Add(&inv)
}

func (v Vec3) Multiply(f float64) Vec3 {
    return Vec3{f*v.X, f*v.Y, f*v.Z}
}

func (v Vec3) Dot(v2* Vec3) float64 {
    return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (v Vec3) AbsDot(v2* Vec3) float64 {
    return math.Abs(v.Dot(v2))
}

func (v Vec3) Cross(w* Vec3) Vec3 {
    return Vec3 {v.Y*w.Z - v.Z*w.Y, v.Z*w.X - v.X*w.Z, v.X*w.Y - v.Y*w.X}
}

func (v Vec3) Normalize() Vec3 {
    return v.Multiply(float64(1.0/v.Magnitude()))
}

func (v Vec3) MinComponent() float64 {
    return math.Min(v.X, math.Min(v.Y, v.Z))
}

func (v Vec3) MaxComponent() float64 {
    return math.Max(v.X, math.Max(v.Y, v.Z))
}

// -------- Vec 2 ---------

type Vec2 struct {
    X, Y float64
}

func (v Vec2) MagnitudeSq() float64 {
    return v.X*v.X + v.Y*v.Y
}

func (v Vec2) Magnitude() float64 {
    return math.Sqrt(v.MagnitudeSq())
}

func (v Vec2) HasNaN() bool {
    return math.IsNaN(v.X) || math.IsNaN(v.Y)
}

func (v Vec2) Inverse() Vec2 {
    return Vec2{-v.X, -v.Y}
}

func (v Vec2) Add(v2* Vec2) Vec2 {
    return Vec2{v.X + v2.X, v.Y + v2.Y}
}

func (v Vec2) Subtract(v2* Vec2) Vec2 {
    inv := v2.Inverse()
    return v.Add(&inv)
}

func (v Vec2) Multiply(f float64) Vec2 {
    return Vec2{f*v.X, f*v.Y}
}

func (v Vec2) Dot(v2* Vec2) float64 {
    return v.X*v2.X + v.Y*v2.Y
}

func (v Vec2) AbsDot(v2* Vec2) float64 {
    return math.Abs(v.Dot(v2))
}

func (v Vec2) Normalize() Vec2 {
    return v.Multiply(float64(1.0/v.Magnitude()))
}

// ---------- Public helper functions ----------

func GetBasisVec3X() Vec3 {
    return Vec3{1,0,0}
}

func GetBasisVec3Y() Vec3 {
    return Vec3{0,1,0}
}

func GetBasisVec3Z() Vec3 {
    return Vec3{0,0,1}
}

func GetBasisVec2X() Vec2 {
    return Vec2{1,0}
}

func GetBasisVec2Y() Vec2 {
    return Vec2{0,1}
}

func MinCompsVec(v Vec3, w Vec3) Vec3 {
    return Vec3{math.Min(v.X, w.X), math.Min(v.Y, w.Y), math.Min(v.Z, w.Z)}
}

func MaxCompsVec(v Vec3, w Vec3) Vec3 {
    return Vec3{math.Max(v.X, w.X), math.Max(v.Y, w.Y), math.Max(v.Z, w.Z)}
}

func MakeCoordSystem(v* Vec3, w* Vec3, u* Vec3) {
    if math.Abs(v.X) > math.Abs(v.Y) {
        sqrt := math.Sqrt(v.X * v.X + v.Z * v.Z)
        *w = Vec3{-v.Z, 0, v.X}.Multiply(1.0/sqrt)
    } else {
        sqrt := math.Sqrt(v.Y * v.Y + v.Z * v.Z)
        *w = Vec3{0, v.Z, -v.Y}.Multiply(1.0/sqrt)
    }
    *u = v.Cross(w)
}
