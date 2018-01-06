package core

import (
	"fmt"
	"math"
)

// Transform represents transform matrices and their inverses
type Transform struct {
	//this is pretty memory heavy, if used naively may run into issues 8byte*16 * 2 is a lot of mem
	m, mInv Matrix4x4f
}

func (t Transform) ApplyP(p Point3) Point3 {
	x, y, z := p.X, p.Y, p.Z
	xp := t.m[0][0]*x + t.m[0][1]*y + t.m[0][2]*z + t.m[0][3]
	yp := t.m[1][0]*x + t.m[1][1]*y + t.m[1][2]*z + t.m[1][3]
	zp := t.m[2][0]*x + t.m[2][1]*y + t.m[2][2]*z + t.m[2][3]
	wp := t.m[3][0]*x + t.m[3][1]*y + t.m[3][2]*z + t.m[3][3]
	if wp == 1 {
		return Point3{xp, yp, zp}
	}
	return Point3{xp, yp, zp}.Divide(wp)
}

func (t Transform) ApplyPE(p Point3) (Point3, Vec3) {
	var err Vec3
	var ret Point3

	x, y, z := p.X, p.Y, p.Z
	// Compute transformed coordinates
	xp := t.m[0][0]*x + t.m[0][1]*y + t.m[0][2]*z + t.m[0][3]
	yp := t.m[1][0]*x + t.m[1][1]*y + t.m[1][2]*z + t.m[1][3]
	zp := t.m[2][0]*x + t.m[2][1]*y + t.m[2][2]*z + t.m[2][3]
	wp := t.m[3][0]*x + t.m[3][1]*y + t.m[3][2]*z + t.m[3][3]

	// Compute absolute error for transformed point
	xAbsSum := (math.Abs(t.m[0][0]*x) + math.Abs(t.m[0][1]*y) +
		math.Abs(t.m[0][2]*z) + math.Abs(t.m[0][3]))
	yAbsSum := (math.Abs(t.m[1][0]*x) + math.Abs(t.m[1][1]*y) +
		math.Abs(t.m[1][2]*z) + math.Abs(t.m[1][3]))
	zAbsSum := (math.Abs(t.m[2][0]*x) + math.Abs(t.m[2][1]*y) +
		math.Abs(t.m[2][2]*z) + math.Abs(t.m[2][3]))

	err = Vec3{xAbsSum, yAbsSum, zAbsSum}.Multiply(Gamma(3))
	if wp == 0 {
		fmt.Print("CANNOT DIVIDE BY 0")
		return ret, err
	}
	if wp == 1 {
		ret = Point3{xp, yp, zp}
	} else {
		ret = Point3{xp, yp, zp}.Divide(wp)
	}
	return ret, err
}

// ApplyV applies the transform to a vec, we assume homogenous behavior i.e weight of 0
func (t Transform) ApplyV(v Vec3) Vec3 {
	x, y, z := v.X, v.Y, v.Z
	return Vec3{t.m[0][0]*x + t.m[0][1]*y + t.m[0][2]*z,
		t.m[1][0]*x + t.m[1][1]*y + t.m[1][2]*z,
		t.m[2][0]*x + t.m[2][1]*y + t.m[2][2]*z}
}

/* to maintain orthogonality normals must be transformed by the inverse of the transpose of the transformation matrix
   let S be the transform for the normal and M be the transform for some point on the surface
   0 = n . t
   0 = n^T t
   to maintain orthogonality the following must be true
   0 = (n')^T t'
   0 = (Sn)^T Mt
   0 = n^T S^T Mt implies S^T M = identity mat
   this implies
   S^T = M^(-1)
   S = (M^-1)^T
*/
func (t Transform) ApplyN(n Normal3) Normal3 {
	x, y, z := n.X, n.Y, n.Z
	return Normal3{t.mInv[0][0]*x + t.mInv[1][0]*y + t.mInv[2][0]*z,
		t.mInv[0][1]*x + t.mInv[1][1]*y + t.mInv[2][1]*z,
		t.mInv[0][2]*x + t.mInv[1][2]*y + t.mInv[2][2]*z}
}

func (t Transform) ApplyR(r Ray) Ray {
	o, _ := t.ApplyPE(r.Orig)
	d := t.ApplyV(r.Dir)
	tMax := float64(0)
	return Ray{o, d, tMax, r.Time, r.medium}
}

func (t Transform) ApplyB(b Bounds3) Bounds3 {
	ret := Bounds3{pMin: t.ApplyP(Point3{b.pMin.X, b.pMin.Y, b.pMin.Z})}
	ret = UnionB3P(ret, t.ApplyP(Point3{b.pMax.X, b.pMin.Y, b.pMin.Z}))
	ret = UnionB3P(ret, t.ApplyP(Point3{b.pMin.X, b.pMax.Y, b.pMin.Z}))
	ret = UnionB3P(ret, t.ApplyP(Point3{b.pMin.X, b.pMin.Y, b.pMax.Z}))
	ret = UnionB3P(ret, t.ApplyP(Point3{b.pMin.X, b.pMax.Y, b.pMax.Z}))
	ret = UnionB3P(ret, t.ApplyP(Point3{b.pMax.X, b.pMax.Y, b.pMin.Z}))
	ret = UnionB3P(ret, t.ApplyP(Point3{b.pMax.X, b.pMin.Y, b.pMax.Z}))
	ret = UnionB3P(ret, t.ApplyP(Point3{b.pMax.X, b.pMax.Y, b.pMax.Z}))
	return ret
}

func (t Transform) Inverse() Transform {
	return Transform{t.mInv, t.m}
}

func (t Transform) Transpose() Transform {
	return Transform{t.m.Transpose(), t.mInv.Transpose()}
}

func (t Transform) IsIdentity() bool {
	m := New4x4IDMat()
	return t.m == m && t.mInv == m
}

// HasScale applies transform to 3 axis, see if they change length noticeably
func (t Transform) HasScale() bool {
	notOne := func(n float64) bool {
		return n < 0.999 || n > 1.001
	}
	x := t.ApplyV(Vec3{1, 0, 0}).MagnitudeSq()
	y := t.ApplyV(Vec3{0, 1, 0}).MagnitudeSq()
	z := t.ApplyV(Vec3{0, 0, 1}).MagnitudeSq()

	return notOne(x) || notOne(y) || notOne(z)
}

// NewTransform returns an identity transform
func NewTransform() Transform {
	m := New4x4IDMat()
	return Transform{m, m}
}

// NewTransformFromMat creates a transform from a matrix
func NewTransformFromMat(mat Matrix4x4f) (t Transform) {
	t.m = mat
	var b bool
	b, t.mInv = mat.Inverse()
	if !b {
		t.mInv = New4x4IDMat()
	}
	return t
}

// NewTransformFromArr creates a transform from an array
func NewTransformFromArr(mat [4][4]float64) Transform {
	return NewTransformFromMat(Matrix4x4f(mat))
}

func NewTransformWithInv(mat, matInv Matrix4x4f) Transform {
	return Transform{mat, matInv}
}

func IsEqualTransform(t0, t1 Transform) bool {
	return t0.m == t1.m && t0.mInv == t1.mInv
}

func Translate(delta Vec3) (t Transform) {
	m := NewMat4x4f(1, 0, 0, delta.X,
		0, 1, 0, delta.Y,
		0, 0, 1, delta.Z,
		0, 0, 0, 1)
	mInv := NewMat4x4f(1, 0, 0, -delta.X,
		0, 1, 0, -delta.Y,
		0, 0, 1, -delta.Z,
		0, 0, 0, 1)
	return Transform{m, mInv}
}

func Scale(x, y, z float64) Transform {
	m := NewMat4x4f(x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1)
	mInv := NewMat4x4f(1/x, 0, 0, 0,
		0, 1/y, 0, 0,
		0, 0, 1/z, 0,
		0, 0, 0, 1)
	return Transform{m, mInv}
}

func RotateX(theta float64) Transform {
	s := math.Sin(Radians(theta))
	c := math.Sin(Radians(theta))
	m := NewMat4x4f(1, 0, 0, 0,
		0, c, -s, 0,
		0, s, c, 0,
		0, 0, 0, 1)
	return NewTransformWithInv(m, m.Transpose())
}

func RotateY(theta float64) Transform {
	s := math.Sin(Radians(theta))
	c := math.Sin(Radians(theta))
	m := NewMat4x4f(c, 0, -s, 0,
		0, 1, 0, 0,
		s, 0, c, 0,
		0, 0, 0, 1)
	return NewTransformWithInv(m, m.Transpose())
}

func RotateZ(theta float64) Transform {
	s := math.Sin(Radians(theta))
	c := math.Sin(Radians(theta))
	m := NewMat4x4f(c, -s, 0, 0,
		s, c, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1)
	return NewTransformWithInv(m, m.Transpose())
}

func RotateFromAxis(theta float64, axis Vec3) Transform {
	a := axis.Normalize()
	s := math.Sin(Radians(theta))
	c := math.Sin(Radians(theta))
	var m Matrix4x4f
	m[0][0] = a.X*a.X + (1-a.X*a.X)*c
	m[0][1] = a.X*a.Y*(1-c) - a.Z*s
	m[0][2] = a.X*a.Z*(1-c) + a.Y*s
	m[0][3] = 0

	m[1][0] = a.Y*a.X*(1-c) + a.Z*s
	m[1][1] = a.Y*a.Y + (1-a.Y*a.Y)*c
	m[1][2] = a.Y*a.Z*(1-c) - a.X*s
	m[1][3] = 0

	m[2][0] = a.Z*a.X*(1-c) - a.Y*s
	m[2][1] = a.Z*a.Y*(1-c) + a.X*s
	m[2][2] = a.Z*a.Z + (1-a.Z*a.Z)*c
	m[2][3] = 0
	return Transform{m, m.Transpose()}
}

func LookAt(pos, look Point3, up Vec3) Transform {
	var cameraToWorld Matrix4x4f
	cameraToWorld[0][3] = pos.X
	cameraToWorld[1][3] = pos.Y
	cameraToWorld[2][3] = pos.Z
	cameraToWorld[3][3] = 1

	dir := look.SubtractP(pos).Normalize()
	left := CrossV3(up.Normalize(), dir).Normalize()
	newUp := CrossV3(dir, left)

	cameraToWorld[0][0] = left.X
	cameraToWorld[1][0] = left.Y
	cameraToWorld[2][0] = left.Z
	cameraToWorld[3][0] = 0

	cameraToWorld[0][1] = newUp.X
	cameraToWorld[1][1] = newUp.Y
	cameraToWorld[2][1] = newUp.Z
	cameraToWorld[3][1] = 0

	cameraToWorld[0][2] = dir.X
	cameraToWorld[1][2] = dir.Y
	cameraToWorld[2][2] = dir.Z
	cameraToWorld[3][2] = 0

	var inv Matrix4x4f
	_, inv = cameraToWorld.Inverse()

	return Transform{inv, cameraToWorld}
}
