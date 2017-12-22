package geometry

import (
	amath "Anvil/math"
	"math"
)

// All bounding boxes are AABB (Axis aligned bounding boxes)
// pMin and pMax are opposing corners of the BB
// invariant: pMin < pMax for all components
type Bounds3 struct {
	pMin, pMax amath.Point3
}

// returns the point for one of the 8 corners of the BB
func (b Bounds3) Corner(i int) amath.Point3 {
	var pX, pY, pZ float64
	if i&1 != 0 {
		pX = b.pMax.X
	} else {
		pX = b.pMin.X
	}

	if i&2 != 0 {
		pY = b.pMax.Y
	} else {
		pY = b.pMin.Y
	}

	if i&4 != 0 {
		pZ = b.pMax.Z
	} else {
		pZ = b.pMin.Z
	}

	return amath.Point3{X: pX, Y: pY, Z: pZ}
}

func (b Bounds3) Contains(p amath.Point3) bool {
	return (p.X >= b.pMin.X && p.X <= b.pMax.X &&
		p.Y >= b.pMin.Y && p.Y <= b.pMax.Y &&
		p.Z >= b.pMin.Z && p.Z <= b.pMax.Z)
}

func (b Bounds3) Diagonal() amath.Vec3 {
	return b.pMax.SubtractP(b.pMin)
}

func (b Bounds3) SurfaceArea() float64 {
	d := b.Diagonal()
	return 2 * (d.X*d.Y + d.X*d.Z + d.Y*d.Z)
}

func (b Bounds3) Volume() float64 {
	d := b.Diagonal()
	return d.X * d.Y * d.Z
}

// returns index of the longest of the three axis
func (b Bounds3) MaxExtent() int {
	d := b.Diagonal()
	if d.X > d.Y && d.X > d.Z {
		return 0
	} else if d.Y > d.Z {
		return 1
	} else {
		return 2
	}
}

// lerp by the corners of the box
func (b Bounds3) Lerp(t amath.Point3) amath.Point3 {
	x := amath.Lerp(t.X, b.pMin.X, b.pMax.X)
	y := amath.Lerp(t.Y, b.pMin.Y, b.pMax.Y)
	z := amath.Lerp(t.Z, b.pMin.Z, b.pMax.Z)
	return amath.Point3{X: x, Y: y, Z: z}
}

// returns pos of point relative to box coords (0,0,0) is at pMin, (1,1,1) is at pMax
func (b Bounds3) Offset(p amath.Point3) amath.Vec3 {
	o := p.SubtractP(b.pMin)
	if b.pMax.X > b.pMin.X {
		o.X /= b.pMax.X - b.pMin.X
	}
	if b.pMax.Y > b.pMin.Y {
		o.Y /= b.pMax.Y - b.pMin.Y
	}
	if b.pMax.Z > b.pMin.Z {
		o.Z /= b.pMax.Z - b.pMin.Z
	}
	return o
}

func (b Bounds3) BoundingSphere() (amath.Point3, float64) {
	center := (b.pMin.AddP(b.pMax)).Divide(2)
	var radius float64
	if b.Contains(center) {
		radius = amath.DistanceP3(center, b.pMax)
	}
	return center, radius
}

func NewEmptyBounds3() Bounds3 {
	ret := Bounds3{}
	min := float64(math.MinInt32)
	max := float64(math.MaxInt32)
	// give bounds invalid
	ret.pMin = amath.Point3{X: max, Y: max, Z: max}
	ret.pMax = amath.Point3{X: min, Y: min, Z: min}
	return ret
}

func NewSinglePBounds3(p amath.Point3) Bounds3 {
	return Bounds3{p, p}
}

func NewBounds3(p0, p1 amath.Point3) Bounds3 {
	ret := Bounds3{}
	ret.pMin = amath.Point3{X: math.Min(p0.X, p1.X), Y: math.Min(p0.Y, p1.Y), Z: math.Min(p0.Z, p1.Z)}
	ret.pMax = amath.Point3{X: math.Max(p0.X, p1.X), Y: math.Max(p0.Y, p1.Y), Z: math.Max(p0.Z, p1.Z)}
	return ret
}

// returns new bounding box containing b and p
func UnionB3P(b Bounds3, p amath.Point3) Bounds3 {
	p1 := amath.Point3{X: math.Min(b.pMin.X, p.X), Y: math.Min(b.pMin.Y, p.Y), Z: math.Min(b.pMin.Z, p.Z)}
	p2 := amath.Point3{X: math.Max(b.pMax.X, p.X), Y: math.Max(b.pMax.Y, p.Y), Z: math.Max(b.pMax.Z, p.Z)}
	return Bounds3{p1, p2}
}

// returns new bounding box containing b0 and b1
func UnionB3B3(b0, b1 Bounds3) Bounds3 {
	p1 := amath.Point3{X: math.Min(b0.pMin.X, b1.pMin.X), Y: math.Min(b0.pMin.Y, b1.pMin.Y), Z: math.Min(b0.pMin.Z, b1.pMin.Z)}
	p2 := amath.Point3{X: math.Max(b0.pMax.X, b1.pMax.X), Y: math.Max(b0.pMax.Y, b1.pMax.Y), Z: math.Max(b0.pMax.Z, b1.pMax.Z)}
	return Bounds3{p1, p2}
}

func OverlapsB3(b1, b2 Bounds3) bool {
	x := (b1.pMax.X >= b2.pMin.X) && (b1.pMin.X <= b2.pMax.X)
	y := (b1.pMax.Y >= b2.pMin.Y) && (b1.pMin.Y <= b2.pMax.Y)
	z := (b1.pMax.Z >= b2.pMin.Z) && (b1.pMin.Z <= b2.pMax.Z)
	return (x && y && z)
}

// returns new bounds padded by t
func ExpandB3(b Bounds3, t float64) Bounds3 {
	delta := amath.Vec3{X: t, Y: t, Z: t}
	return Bounds3{b.pMin.SubtractV(delta), b.pMax.AddV(delta)}
}

type Bounds2 struct {
	pMin, pMax amath.Point2
}

// returns the point for one of the 8 corners of the BB
func (b Bounds2) Corner(i int) amath.Point2 {
	var pX, pY float64
	if i&1 != 0 {
		pX = b.pMax.X
	} else {
		pX = b.pMin.X
	}

	if i&2 != 0 {
		pY = b.pMax.Y
	} else {
		pY = b.pMin.Y
	}

	return amath.Point2{X: pX, Y: pY}
}

func (b Bounds2) Contains(p amath.Point2) bool {
	return (p.X >= b.pMin.X && p.X <= b.pMax.X &&
		p.Y >= b.pMin.Y && p.Y <= b.pMax.Y)
}

func (b Bounds2) Diagonal() amath.Vec2 {
	return b.pMax.SubtractP(b.pMin)
}

func (b Bounds2) SurfaceArea() float64 {
	d := b.Diagonal()
	return d.X * d.Y
}

// returns index of the longest of the three axis
func (b Bounds2) MaxExtent() int {
	d := b.Diagonal()
	if d.X > d.Y {
		return 0
	} else {
		return 1
	}
}

// lerp by the corners of the box
func (b Bounds2) Lerp(t amath.Point2) amath.Point2 {
	x := amath.Lerp(t.X, b.pMin.X, b.pMax.X)
	y := amath.Lerp(t.Y, b.pMin.Y, b.pMax.Y)
	return amath.Point2{X: x, Y: y}
}

// returns pos of point relative to box coords (0,0,0) is at pMin, (1,1,1) is at pMax
func (b Bounds2) Offset(p amath.Point2) amath.Vec2 {
	o := p.SubtractP(b.pMin)
	if b.pMax.X > b.pMin.X {
		o.X /= b.pMax.X - b.pMin.X
	}
	if b.pMax.Y > b.pMin.Y {
		o.Y /= b.pMax.Y - b.pMin.Y
	}
	return o
}

func (b Bounds2) BoundingCircle() (amath.Point2, float64) {
	center := (b.pMin.AddP(b.pMax)).Divide(2)
	var radius float64
	if b.Contains(center) {
		radius = amath.DistanceP2(center, b.pMax)
	}
	return center, radius
}

func NewEmptyBounds2() Bounds2 {
	ret := Bounds2{}
	min := float64(math.MinInt32)
	max := float64(math.MaxInt32)
	// give bounds invalid
	ret.pMin = amath.Point2{X: max, Y: max}
	ret.pMax = amath.Point2{X: min, Y: min}
	return ret
}

func NewSinglePBounds2(p amath.Point2) Bounds2 {
	return Bounds2{p, p}
}

func NewBounds2(p0, p1 amath.Point2) Bounds2 {
	ret := Bounds2{}
	ret.pMin = amath.Point2{X: math.Min(p0.X, p1.X), Y: math.Min(p0.Y, p1.Y)}
	ret.pMax = amath.Point2{X: math.Max(p0.X, p1.X), Y: math.Max(p0.Y, p1.Y)}
	return ret
}

// returns new bounding box containing b and p
func UnionB2P(b Bounds2, p amath.Point2) Bounds2 {
	p1 := amath.Point2{X: math.Min(b.pMin.X, p.X), Y: math.Min(b.pMin.Y, p.Y)}
	p2 := amath.Point2{X: math.Max(b.pMax.X, p.X), Y: math.Max(b.pMax.Y, p.Y)}
	return Bounds2{p1, p2}
}

// returns new bounding box containing b0 and b1
func UnionB2B2(b0, b1 Bounds2) Bounds2 {
	p1 := amath.Point2{X: math.Min(b0.pMin.X, b1.pMin.X), Y: math.Min(b0.pMin.Y, b1.pMin.Y)}
	p2 := amath.Point2{X: math.Max(b0.pMax.X, b1.pMax.X), Y: math.Max(b0.pMax.Y, b1.pMax.Y)}
	return Bounds2{p1, p2}
}

func OverlapsB2(b1, b2 Bounds2) bool {
	x := (b1.pMax.X >= b2.pMin.X) && (b1.pMin.X <= b2.pMax.X)
	y := (b1.pMax.Y >= b2.pMin.Y) && (b1.pMin.Y <= b2.pMax.Y)
	return (x && y)
}

// returns new bounds padded by t
func ExpandB2(b Bounds2, t float64) Bounds2 {
	delta := amath.Vec2{X: t, Y: t}
	return Bounds2{b.pMin.SubtractV(delta), b.pMax.AddV(delta)}
}
