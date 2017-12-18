package shapes

import (
    "Anvil/math"
)

type Shape struct {
    Desc string
    pos math.Vec3
}

func (s Shape) GetPos() math.Vec3 {
    return s.pos
}
