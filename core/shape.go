package core

import "math"

type ShapeInter interface {
	ObjectBound() Bounds3
	// ray should be in world space, shape responsible to translate to object space if needed
	// testAlphaTexture tests for textures that 'cut away' parts of the shape surface
	Intersect(ray Ray, testAlphaTexture bool) (bool, float64, SurfaceInteraction)
	IntersectP(ray Ray, testAlphaTexture bool) bool
	Area() float64
}

func WorldBound(s ShapeData, si ShapeInter) Bounds3 {
	return s.ObjectToWorld.ApplyB(si.ObjectBound())
}

func intersectP(si ShapeInter, ray Ray, testAlphaTexture bool) bool {
	b, _, _ := si.Intersect(ray, testAlphaTexture)
	return b
}

type ShapeData struct {
	Desc                                         string
	ObjectToWorld, WorldToObject                 *Transform
	ReverseOrientation, TransformSwapsHandedness bool
}

//precompute if we need to swap the normal dir on ray intersection
func NewShapeData(objToWorld, worldToObj *Transform, reverseOrientation bool,
	desc string) ShapeData {
	return ShapeData{desc, objToWorld, worldToObj, reverseOrientation,
		objToWorld.SwapsHandedness()}
}

type Sphere struct {
	shape                                          ShapeData
	radius, zMin, zMax, thetaMin, thetaMax, phiMax float64
}

func NewSphere(objectToWorld, worldToObject *Transform, reverseOrientation bool,
	radius, zMin, zMax, phiMax float64) Sphere {
	return Sphere{
		NewShapeData(objectToWorld, worldToObject, reverseOrientation, "Sphere"),
		radius,
		Clamp(math.Max(zMin, zMax), -radius, radius),
		Clamp(math.Max(zMin, zMax), -radius, radius),
		math.Acos(Clamp(zMin/radius, -1, 1)),
		math.Asin(Clamp(zMax/radius, -1, 1)),
		phiMax}
}

func (self Sphere) ObjectBound() Bounds3 {
	r := self.radius
	return NewBounds3(Point3{-r, -r, self.zMin}, Point3{r, r, self.zMax})
}
func (self Sphere) Intersect(r Ray, testAlphaTexture bool) (bool, float64, SurfaceInteraction) {
	var phi float64
	var pHit Point3

	// Transform ray to object space
	ray := self.shape.WorldToObject.ApplyR(r)

	// Compute quadratic sphere coefficients
	dx, dy, dz := ray.Dir.X, ray.Dir.Y, ray.Dir.Z
	ox, oy, oz := ray.Orig.X, ray.Orig.Y, ray.Orig.Z
	radius := self.radius

	a := dx*dx + dy*dy + dz*dz
	b := 2 * (dx*ox + dy*oy + dz*oz)
	c := ox*ox + oy*oy + oz*oz - radius*radius

	// Solve for t values
	var t0, t1 float64
	ret, t0, t1 := Quadratic(a, b, c)
	if !ret || t0 > ray.tMax || t1 <= 0 {
		return false, 0, SurfaceInteraction{}
	}

	tShapeHit := t0
	if tShapeHit <= 0 {
		tShapeHit = t1
		if tShapeHit >= ray.tMax {
			return false, 0, SurfaceInteraction{}
		}
	}

	// Compute sphere hit position and phi
	getPosAndPhi := func(tsh float64) (Point3, float64) {
		p := ray.GetPointForT(tsh)
		// TODO: refine sphere intersection point
		if p.X == 0 && p.Y == 0 {
			p.Z = 1e-5 * radius
		}
		thisPhi := math.Atan2(p.Y, p.X)
		if phi < 0 {
			phi += 2 * math.Pi
		}
		return p, thisPhi
	}

	pHit, phi = getPosAndPhi(tShapeHit)

	// Test sphere intersection against clipping parameters, careful that z range
	// doesnt lie inside of sphere
	if (self.zMin > -radius && pHit.Z < self.zMin) ||
		(self.zMax < radius && pHit.Z > self.zMax) ||
		phi > self.phiMax {

		if tShapeHit == t1 || t1 > ray.tMax {
			return false, 0, SurfaceInteraction{}
		}
		tShapeHit = t1
		pHit, phi = getPosAndPhi(tShapeHit)
		if (self.zMin > -radius && pHit.Z < self.zMin) ||
			(self.zMax < radius && pHit.Z > self.zMax) ||
			phi > self.phiMax {
			return false, 0, SurfaceInteraction{}
		}
	}

	// Find parametric representation of sphere hit
	u := phi / self.phiMax
	theta := math.Acos(Clamp(pHit.Z/radius, -1, 1))
	v := (theta - self.thetaMin) / (self.thetaMax - self.thetaMin)
	zRadius := math.Sqrt(pHit.X*pHit.X + pHit.Y*pHit.Y)
	invZRadius := 1 / zRadius
	cosPhi := pHit.X * invZRadius
	sinPhi := pHit.Y * invZRadius
	dpdu := Vec3{-self.phiMax * pHit.Y, self.phiMax * pHit.X, 0}
	dpdv := Vec3{pHit.Z * cosPhi, pHit.Z * sinPhi, -radius * math.Sin(theta)}.Multiply(self.thetaMax - self.thetaMin)
	//Weingarten eqns for dndu and dndv
	d2Pduu := Vec3{pHit.X, pHit.Y, 0}.Multiply(-self.phiMax * self.phiMax)
	d2Pduv := Vec3{-sinPhi, cosPhi, 0}.Multiply((self.thetaMax - self.thetaMin) * pHit.Z * self.phiMax)
	d2Pdvv := Vec3{pHit.X, pHit.Y, pHit.Z}.Multiply((self.thetaMax - self.thetaMin) * -(self.thetaMax - self.thetaMin))

	E := DotV3(dpdu, dpdu)
	F := DotV3(dpdu, dpdv)
	G := DotV3(dpdv, dpdv)
	N := CrossV3(dpdu, dpdv).Normalize()
	e := DotV3(N, d2Pduu)
	f := DotV3(N, d2Pduv)
	g := DotV3(N, d2Pdvv)

	invEGF2 := 1 / (E*G - F*F)
	dndu := NormalFromVec3(dpdu.Multiply((f*F - e*G) * invEGF2).Add(dpdv.Multiply((e*F - f*E) * invEGF2)))
	dndv := NormalFromVec3(dpdu.Multiply((g*F - f*G) * invEGF2).Add(dpdv.Multiply((f*F - g*E) * invEGF2)))

	// TODO: Compute error bounds for sphere intersection

	// Initialize surface interaction for parametric information
	si := NewSurfaceInteraction(pHit, Vec3{}, ray.Dir.Inverse(), ray.Time, Point2{u, v}, dpdu, dpdv, dndu, dndv, &self.shape)
	si = self.shape.ObjectToWorld.ApplySI(si)

	// update thit for quadratic intersection
	return true, tShapeHit, si
}
func (self Sphere) IntersectP(ray Ray, testAlphaTexture bool) bool {
	return intersectP(self, ray, testAlphaTexture)
}
func (self Sphere) Area() float64 {
	return self.phiMax * self.radius * (self.zMax - self.zMin)
}
