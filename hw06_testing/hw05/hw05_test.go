package hw05_test

import (
	"errors"
	"testing"

	"github.com/freeloginname/home_work_basic/hw06_testing/hw05"
	"github.com/stretchr/testify/require"
)

func TestHW05(t *testing.T) {
	testCases := []struct {
		desc        string
		radius      float64
		width       float64
		tight       float64
		base        float64
		high        float64
		shapeType   hw05.ShapeType
		expectation float64
		err         error
	}{
		{
			desc:        "Circle",
			radius:      5,
			shapeType:   hw05.CircleType,
			expectation: 78.53981633974483,
			err:         nil,
		},
		{
			desc:        "Rectangle",
			width:       10,
			tight:       5,
			shapeType:   hw05.RectangleType,
			expectation: 50,
			err:         nil,
		},
		{
			desc:        "Triangle",
			base:        8,
			high:        6,
			shapeType:   hw05.TriangleType,
			expectation: 24,
			err:         nil,
		},
		{
			desc:        "Zero",
			base:        0,
			high:        6,
			shapeType:   hw05.TriangleType,
			expectation: 0,
			err:         nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var shape hw05.Shape
			switch tC.shapeType {
			case hw05.CircleType:
				shape = hw05.Circle{Radius: tC.radius}
			case hw05.RectangleType:
				shape = hw05.Rectangle{Wide: tC.width, Tight: tC.tight}
			case hw05.TriangleType:
				shape = hw05.Triangle{Base: tC.base, High: tC.high}
			}
			got, err := hw05.CalculateArea(shape)
			require.NoError(t, err)
			require.Equal(t, tC.expectation, got)
		})
	}

	errTestCases := []struct {
		desc        string
		dimension   float64
		expectation float64
		err         error
	}{
		{
			desc:        "Fake figure",
			dimension:   5.0,
			expectation: 0.0,
			err:         errors.New("ошибка: переданный объект не является фигурой"),
		},
	}
	for _, tC := range errTestCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := hw05.CalculateArea(hw05.FakeFigure{Dimension: tC.dimension})
			require.Equal(t, tC.expectation, got)
			require.Equal(t, tC.err, err)
		})
	}
}
