package shape

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
}

func (rec Rectangle) Perimeter() float64 {
	return (rec.Width + rec.Height) * 2
}

func (rec Rectangle) Area() float64 {
	return rec.Width * rec.Height
}

func (cir Circle) Area() float64 {
	return math.Pi * cir.Radius * cir.Radius
}

func (tri Triangle) Area() float64 {
	return 0.5 * (tri.Base * tri.Height)
}
