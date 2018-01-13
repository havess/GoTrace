package core

import (
	"Anvil/media"
)

type Interaction struct {
	p               Point3
	time            float64
	pError, wo      Vec3 //position error and negative ray direction
	n               Normal3
	mediumInterface *media.MediumInterface
}

func (i Interaction) IsSurfaceInteraction() bool {
	return i.n != Normal3{}
}

func NewInteraction(p Point3, n Normal3, pError, wo Vec3, time float64, med *media.MediumInterface) Interaction {
	return Interaction{p, time, pError, wo, n, med}
}

type SurfaceInteraction struct {
	inter      Interaction
	uv         Point2
	dpdu, dpdv Vec3
	dndu, dndv Normal3
	shape      *Shape
	shading    struct {
		n, dndu, dndv Normal3
		dpdu, dpdv    Vec3
	}
}

func NewSurfaceInteraction(p Point3, pError, wo Vec3, time float64, med media.MediumInterface,
	uv Point2, dpdu, dpdv Vec3, dndu, dndv Normal3, shape *Shape) SurfaceInteraction {
	normal := NormalFromVec3(CrossV3(dpdu, dpdv).Normalize())
	interaction := NewInteraction(p, normal, pError, wo, time, nil)
	s := struct {
		n, dndu, dndv Normal3
		dpdu, dpdv    Vec3
	}{normal, dndu, dndv, dpdu, dpdv}

	/**/
	if shape != nil &&
		(shape.ReverseOrientation || shape.TransformSwapsHandedness && !(shape.ReverseOrientation && shape.TransformSwapsHandedness)) {
		normal = normal.Multiply(-1)
	}
	return SurfaceInteraction{interaction, uv, dpdu, dpdv, dndu, dndv, shape, s}
}
