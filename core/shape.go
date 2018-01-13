package core

type Shape struct {
	Desc                                         string
	pos                                          Vec3
	ReverseOrientation, TransformSwapsHandedness bool
}

func (s Shape) GetPos() Vec3 {
	return s.pos
}
