package lib

import (
	"math"
)

type Vector struct {
	X float64
	Y float64
}

func (v *Vector) Add(o *Vector) {
	v.X += o.X
	v.Y += o.Y
}

func (v *Vector) GetLength() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

func (v *Vector) GetAngle(o *Vector) float64 {
	return math.Acos(v.DotProduct(o) / (v.GetLength() * o.GetLength()))
}

func (v *Vector) GetAngleWithAxis() float64 {
	if v.X == 0 && v.Y > 0 {
		return math.Pi / 2
	} else if v.X == 0 && v.Y < 0 {
		return math.Pi * 1.5
	}
	a := math.Atan(v.Y / v.X)
	if v.Y >= 0 && v.X < 0 {
		a += math.Pi
	} else if v.Y < 0 && v.X > 0 {
		a += math.Pi * 2
	}
	return a
}

func (v *Vector) Negate() {
	v.X = -v.X
	v.Y = -v.Y
}

func (v *Vector) DotProduct(o *Vector) float64 {
	return v.X*o.X + v.Y*o.Y
}

func (v *Vector) Rotate(a float64) {
	ca := math.Cos(a)
	sa := math.Sin(a)
	v.X, v.Y = ca*v.X-sa*v.Y, sa*v.X+ca*v.Y
}

func NegateVector(v *Vector) *Vector {
	return &Vector{X: -v.X, Y: -v.Y}
}

func AddVector(v *Vector, o *Vector) *Vector {
	return &Vector{X: v.X + o.X, Y: v.Y + o.Y}
}
