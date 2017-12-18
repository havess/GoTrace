package lights

import (
    "Anvil/math"
)

type Light struct {
    pos math.Vec3
    power float64
}

func (l Light) getPos() math.Vec3 {
    return l.pos
}

func (l Light) getPow() float64 {
    return l.power
}
