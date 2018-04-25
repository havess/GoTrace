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
func (self Sphere) Intersect(ray Ray, testAlphaTexture bool) (bool, float64, SurfaceInteraction) {

}
func (self Sphere) IntersectP(ray Ray, testAlphaTexture bool) bool {
	return intersectP(self, ray, testAlphaTexture)
}
func (self Sphere) Area() float64 {

}
