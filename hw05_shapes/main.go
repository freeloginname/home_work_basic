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
	Radius float64
}

func (circle Circle) Area() float64 {
	return math.Pi * math.Pow(circle.Radius, 2)
}

type Rectangle struct {
	Wide  float64
	Tight float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Wide * rectangle.Tight
}

type Triangle struct {
	Base float64
	High float64
}

func (triangle Triangle) Area() float64 {
	return 0.5 * triangle.Base * triangle.High
}

func calculateArea(s any) (float64, error) {
	var area float64
	var err error
	_, shapeImplementation := s.(Shape)
	if !shapeImplementation {
		err = errors.New("ошибка: переданный объект не является фигурой")
		return area, err
	}
	area = s.(Shape).Area()
	return area, err
}

type fakeFigure struct {
	dimension float64
}

func main() {
	var err error
	circle := Circle{Radius: 5}
	circleArea, err := calculateArea(circle)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Printf("Круг: радиус %f \n", circle.Radius)
	fmt.Printf("Площадь: %f \n\n", circleArea)

	rectangle := Rectangle{
		Wide:  10,
		Tight: 5,
	}
	rectangleArea, err := calculateArea(rectangle)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Printf("Прямоугольник: ширина %f , высота %f \n", rectangle.Wide, rectangle.Tight)
	fmt.Printf("Площадь: %f \n\n", rectangleArea)

	triangle := Triangle{
		Base: 8,
		High: 6,
	}
	triangleArea, err := calculateArea(triangle)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Printf("Треугольник: основание %f , высота %f \n", triangle.Base, triangle.High)
	fmt.Printf("Площадь: %f \n\n", triangleArea)

	fake := fakeFigure{dimension: 1}
	fakeArea, err := calculateArea(fake)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Площадь фигуры без метода ее подсчета: %f \n\n", fakeArea)
	}
}
