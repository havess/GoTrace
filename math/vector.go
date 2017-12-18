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
