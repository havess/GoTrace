package math

type Point struct {
    Pos Vec3
}

func (p Point) GetPos() Vec3 {
    return p.Pos
}
