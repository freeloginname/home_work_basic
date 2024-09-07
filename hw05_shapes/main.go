package main

import (
	"errors"
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (circle Circle) Area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

type Rectangle struct {
	wide  float64
	tight float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.wide * rectangle.tight
}

type Triangle struct {
	base float64
	high float64
}

func (triangle Triangle) Area() float64 {
	return 0.5 * triangle.base * triangle.high
}

func calculateArea(shape Shape) (float64, error) {
	var area float64
	var err error
	_, shape_implementation := interface{}(shape).(Shape)
	if !shape_implementation {
		err = errors.New("Ошибка: переданный объект не является фигурой.")
	}
	area = shape.Area()
	return area, err
}

type fakeFigure struct {
	dimension float64
}

func main() {
	var err error
	circle := Circle{radius: 5}
	circleArea, err := calculateArea(circle)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Printf("Круг: радиус %f \n", circle.radius)
	fmt.Printf("Площадь: %f \n\n", circleArea)

	rectangle := Rectangle{
		wide:  10,
		tight: 5,
	}
	rectangleArea, err := calculateArea(rectangle)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Printf("Прямоугольник: ширина %f , высота %f \n", rectangle.wide, rectangle.tight)
	fmt.Printf("Площадь: %f \n\n", rectangleArea)

	triangle := Triangle{
		base: 8,
		high: 6,
	}
	triangleArea, err := calculateArea(triangle)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Printf("Треугольник: основание %f , высота %f \n", triangle.base, triangle.high)
	fmt.Printf("Площадь: %f \n\n", triangleArea)

	// fake := fakeFigure{dimension: 1}
	// fakeArea, err := calculateArea(fake)
	// if err != nil {
	// 	fmt.Println("err")
	// }
	// fmt.Printf("Площадь фигуры без метода ее подсчета: %f \n\n", fakeArea)

}
