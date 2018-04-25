package core

import (
	"Anvil/media"
)

type Interaction struct {
	p    Point3
	time float64
	/*
	   For interactions that lie along a ray, store negative ray direction in 'wo'
	   (omega not) which is the notation we use for outgoing direction when computing
	   lighting at points.
	*/
	pError, wo      Vec3
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

func NewSurfaceInteraction(
	p Point3,
	pError, wo Vec3,
	time float64,
	med media.MediumInterface,
	uv Point2,
	dpdu, dpdv Vec3,
	dndu, dndv Normal3,
	shape *Shape) SurfaceInteraction {
	normal := NormalFromVec3(CrossV3(dpdu, dpdv).Normalize())
	interaction := NewInteraction(p, normal, pError, wo, time, nil)
	shading := struct {
		n, dndu, dndv Normal3
		dpdu, dpdv    Vec3
	}{normal, dndu, dndv, dpdu, dpdv}

	if shape != nil && Xor(shape.ReverseOrientation, shape.TransformSwapsHandedness) {
		normal = normal.Multiply(-1)
		shading.n = shading.n.Multiply(-1)
	}
	return SurfaceInteraction{interaction, uv, dpdu, dpdv, dndu, dndv, shape, shading}
}

func (si SurfaceInteraction) SetShadingGeometry(dpdus, dpdvs Vec3,
	dndus, dndvs Normal3,
	orientationIsAuthorative bool) {
	si.shading.n = NormalFromVec3(CrossV3(dpdus, dpdvs)).Normalize()
	if si.shape != nil && Xor(si.shape.ReverseOrientation, si.shape.TransformSwapsHandedness) {
		si.shading.n = si.shading.n.Multiply(-1)
	}
	if orientationIsAuthorative {
		si.inter.n = FaceForward(&si.inter.n, si.shading.n.ToVec3())
	} else {
		si.shading.n = FaceForward(&si.shading.n, si.inter.n.ToVec3())
	}
	si.shading.dpdu = dpdus
	si.shading.dpdv = dpdvs
	si.shading.dndu = dndus
	si.shading.dndv = dndvs
}
