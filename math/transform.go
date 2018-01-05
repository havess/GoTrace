package math

import "math"

// Transform represents transform matrices and their inverses
type Transform struct {
	//this is pretty memory heavy, if used naively may run into issues 8byte*16 * 2 is a lot of mem
	m, mInv Matrix4x4f
}

// ApplyV applies the transform to a vec, we assume homogenous behavior i.e weight of 0
func (t Transform) ApplyV(v Vec3) Vec3 {
	x, y, z := v.X, v.Y, v.Z
	return Vec3{t.m[0][0]*x + t.m[0][1]*y + t.m[0][2]*z,
		t.m[1][0]*x + t.m[1][1]*y + t.m[1][2]*z,
		t.m[2][0]*x + t.m[2][1]*y + t.m[2][2]*z}
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
