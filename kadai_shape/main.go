package main

import (
	"errors"
	"fmt"
	"math"
)

type ValidationError struct {
	Message string
	Field   string
	Value   float64
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Field=%s, Value=%v", e.Message, e.Value)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func validateShape(s Shape) error {
	switch shape := s.(type) {
	case Rectangle:
		if shape.Width <= 0 {
			return &ValidationError{
				Message: "Rectangle.Width", Field: "幅", Value: shape.Width,
			}
		}
		if shape.Height <= 0 {
			return &ValidationError{
				Message: "Rectangle.Height", Field: "高さ", Value: shape.Height,
			}
		}
	case Circle:
		if shape.Radius <= 0 {
			return &ValidationError{
				Message: "Circle.Radius", Field: "半径", Value: shape.Radius,
			}
		}
	default:
		return errors.New("対応していない図形です")
	}
	return nil
}
