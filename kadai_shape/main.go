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
	return fmt.Sprintf(
		"[バリデーションエラー] %s：%s（Value：%.2f）\n [詳細] Field=%s, Value=%.2f\n",
		e.Field,
		e.Message,
		e.Value,
		e.Field,
		e.Value,
	)
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
				Message: "幅は正の数を指定してください。", Field: "Rectangle.Width", Value: shape.Width,
			}
		}
		if shape.Height <= 0 {
			return &ValidationError{
				Message: "高さは正の数を指定してください。", Field: "Rectangle.Height", Value: shape.Height,
			}
		}
	case Circle:
		if shape.Radius <= 0 {
			return &ValidationError{
				Message: "半径は正の数を指定してください。", Field: "Circle.Radius", Value: shape.Radius,
			}
		}
	default:
		return errors.New("対応していない図形です")
	}
	return nil
}

func printShapeInfo(s Shape) {
	fmt.Printf("図形：%T｜面積：%.2f ｜周長：%.2f\n", s, s.Area(), s.Perimeter())
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 10.0, Height: 5.0},
		Circle{Radius: 3.0},
		Rectangle{Width: -2.0, Height: 5.0}, // 不正（幅が不正）
		Circle{Radius: -1.0},                // 不正（半径が不正）
		Rectangle{Width: 1.0, Height: 0.0},  // 不正（高さが不正）
	}

	var calcShapes []Shape
	fmt.Println("--- 図形のバリデーション ---")
	for i, s := range shapes {
		if err := validateShape(s); err != nil {
			if valErr, ok := errors.AsType[*ValidationError](err); ok {
				fmt.Printf("#%d：%T => NG：%v", i+1, s, valErr)
			} else {
				fmt.Println("不明なエラーです")
			}
			continue
		}

		fmt.Printf("#%d：%T => バリデーションOK\n", i+1, s)
		calcShapes = append(calcShapes, s)
	}

	fmt.Println("--- 検証済み図形の一括計算 ---")
	for _, s := range calcShapes {
		printShapeInfo(s)
	}
}
