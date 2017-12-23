package math

import (
	"Anvil/media"
	"math"
)

type Ray struct {
	Orig       Point3
	Dir        Vec3
	tMax, Time float64
	medium     *media.Medium
}

func (r Ray) GetPointForT(t float64) Point3 {
	return r.Orig.AddV(r.Dir.Multiply(t)) // return o + t*d
}

func NewEmptyRay() Ray {
	return Ray{tMax: math.Inf(1), Time: 0.0, medium: nil}
}

func NewRay(p Point3, d Vec3, tMax, time float64, medium *media.Medium) Ray {
	return Ray{p, d, tMax, time, medium}
}

/*
   Structure to give us two auxiliary rays to some actual ray, each offset by
   one sample in the x and y directions in the view plane/film plane
   by computing the area that the three rays project onto an object being
   shaded, textures can estimate an area to average over for proper
   antialiasing
*/
type RayDifferential struct {
	R                  *Ray
	HasDifferentials   bool
	rxOrigin, ryOrigin Point3
	rxDir, ryDir       Vec3
}

func NewEmptyRayDiff() RayDifferential {
	return RayDifferential{HasDifferentials: false}
}

func NewRayDiff(p Point3, d Vec3, tMax, time float64, medium *media.Medium) RayDifferential {
	r := NewRay(p, d, tMax, time, medium)
	ret := NewEmptyRayDiff()
	ret.R = &r
	return ret
}

// update differential rays for estimated sampling of 's'
func (r *RayDifferential) ScaleRayDifferentials(s float64) {
	o := r.R.Orig
	d := r.R.Dir

	r.rxOrigin = o.AddV(r.rxOrigin.SubtractP(o).Multiply(s))
	r.ryOrigin = o.AddV(r.ryOrigin.SubtractP(o).Multiply(s))
	r.rxDir = d.Add(r.rxDir.Subtract(d).Multiply(s))
	r.ryDir = d.Add(r.ryDir.Subtract(d).Multiply(s))
}

func NewRayDifferential(r *Ray) RayDifferential {
	// set initial differentials to false since the neighbouring rays are not known yet
	return RayDifferential{R: r, HasDifferentials: false}
}
