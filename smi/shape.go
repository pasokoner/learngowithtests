package smi

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {

	return 2 * (r.Height + r.Width)
}

func (rectangle Rectangle) Area() float64 {

	return rectangle.Height * rectangle.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {

	return c.Radius * c.Radius * math.Pi
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {

	return 0.5 * (t.Base * t.Height)
}
