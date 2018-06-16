package core

type Primitive interface {
	WorldBound() Bounds3
	Intersect(Ray) (bool, SurfaceInteraction)
	GetAreaLight() *AreaLight
	GetMaterial() *Material
	//TODO: ComputeScatteringFunctions
}

// Represents a single shape in a scene
type GeometricPrimitive struct {
	shape     ShapeInter
	material  *Material
	areaLight *AreaLight
	//TODO:  MediumInterface
}

func NewGeometricPrimitive(shape ShapeInter, material *Material, areaLight *AreaLight) GeometricPrimitive {
	return GeometricPrimitive{shape, material, areaLight}
}

func (self GeometricPrimitive) WorldBound() Bounds3 {
	return self.shape.ObjectBound()
}

func (self GeometricPrimitive) Intersect(r Ray) (bool, SurfaceInteraction) {
	b, tHit, si := self.shape.Intersect(r, false)
	if !b {
		return false, SurfaceInteraction{}
	}
	r.tMax = tHit
	si.primitive = self
	// TODO: init mediumInterface
	return true, si
}

func (self GeometricPrimitive) GetAreaLight() *AreaLight {
	return self.areaLight
}
func (self GeometricPrimitive) GetMaterial() *Material {
	return self.material
}
