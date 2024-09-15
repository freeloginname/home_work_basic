//go:build hw06_test

package main

import (
	"errors"
	"testing"

	"github.com/freeloginname/home_work_basic/hw06_testing/hw03"
	"github.com/freeloginname/home_work_basic/hw06_testing/hw04"
	"github.com/freeloginname/home_work_basic/hw06_testing/hw05"

	"github.com/stretchr/testify/require"
	// "github.com/stretchr/testify/require"
)

func TestHW03(t *testing.T) {
	testCases := []struct {
		desc  string
		size  int
		board string
		err   error
	}{
		{
			desc:  "Default one",
			size:  8,
			board: "# # # # \n # # # #\n# # # # \n # # # #\n# # # # \n # # # #\n# # # # \n # # # #\n",
			err:   nil,
		},
		{
			desc:  "1 size",
			size:  1,
			board: "#\n",
			err:   nil,
		},
		{
			desc:  "2 size",
			size:  2,
			board: "# \n #\n",
			err:   nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := hw03.HW03(tC.size)
			/**
			* можно ли сделать универсальную проверку средствами require
			* для ситуаций когда ошидаешь ошибку и когда ее не должно быть?
			**/
			require.NoError(t, err)
			require.Equal(t, tC.board, got)
		})
	}

	testCases = []struct {
		desc  string
		size  int
		board string
		err   error
	}{
		{
			desc:  "Zero size",
			size:  0,
			board: "",
			err:   errors.New("invalid board size"),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := hw03.HW03(tC.size)
			require.Error(t, err)
			require.Equal(t, tC.board, got)
		})
	}
}

func TestHW04(t *testing.T) {
	testCases := []struct {
		desc      string
		compareBy hw04.ComparisonType
		// firstBook hw04.Book
		// secondBook hw04.Book
		firstBookId      int
		firstBookTitle   string
		firstBookAuthor  string
		firstBookYear    int
		firstBookSize    int
		firstBookRate    float32
		SecondBookId     int
		SecondBookTitle  string
		SecondBookAuthor string
		SecondBookYear   int
		SecondBookSize   int
		SecondBookRate   float32
		expectation      bool
	}{
		{
			/**
			* как можно определить значения для Book на этом этапе,
			* если они задаются через вызов методов?
			 **/
			desc:             "Year comparison",
			compareBy:        hw04.Year,
			firstBookId:      1,
			firstBookTitle:   "aaa",
			firstBookAuthor:  "aaaaa",
			firstBookYear:    1990,
			firstBookSize:    10,
			firstBookRate:    2.0,
			SecondBookId:     2,
			SecondBookTitle:  "bb",
			SecondBookAuthor: "bbb",
			SecondBookYear:   1991,
			SecondBookSize:   20,
			SecondBookRate:   5.0,
			expectation:      false,
		},
		{
			desc:             "Size comparison",
			compareBy:        hw04.Size,
			firstBookId:      1,
			firstBookTitle:   "aaa",
			firstBookAuthor:  "aaaaa",
			firstBookYear:    1990,
			firstBookSize:    10,
			firstBookRate:    2.0,
			SecondBookId:     2,
			SecondBookTitle:  "bb",
			SecondBookAuthor: "bbb",
			SecondBookYear:   1991,
			SecondBookSize:   20,
			SecondBookRate:   5.0,
			expectation:      false,
		},
		{
			desc:             "Rate comparison",
			compareBy:        hw04.Rate,
			firstBookId:      1,
			firstBookTitle:   "aaa",
			firstBookAuthor:  "aaaaa",
			firstBookYear:    1990,
			firstBookSize:    10,
			firstBookRate:    2.0,
			SecondBookId:     2,
			SecondBookTitle:  "bb",
			SecondBookAuthor: "bbb",
			SecondBookYear:   1991,
			SecondBookSize:   20,
			SecondBookRate:   5.0,
			expectation:      false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			bookComparer := hw04.NewBookComparer(tC.compareBy)
			var firstBook hw04.Book
			var secondBook hw04.Book
			firstBook.SetBookID(tC.firstBookId)
			firstBook.SetBookTitle(tC.firstBookTitle)
			firstBook.SetBookAuthor(tC.firstBookAuthor)
			firstBook.SetBookYear(tC.firstBookYear)
			firstBook.SetBookSize(tC.firstBookSize)
			firstBook.SetBookRate(tC.firstBookRate)
			secondBook.SetBookID(tC.SecondBookId)
			secondBook.SetBookTitle(tC.SecondBookTitle)
			secondBook.SetBookAuthor(tC.SecondBookAuthor)
			secondBook.SetBookYear(tC.SecondBookYear)
			secondBook.SetBookSize(tC.SecondBookSize)
			secondBook.SetBookRate(tC.SecondBookRate)
			got := bookComparer.CompareBooks(firstBook, secondBook)
			require.Equal(t, tC.expectation, got)
		})
	}
}

func TestHW05(t *testing.T) {
	testCases := []struct {
		desc        string
		radius      float64
		width       float64
		tight       float64
		base        float64
		high        float64
		shape_type  hw05.ShapeType
		expectation float64
		err         error
	}{
		{
			desc:        "Circle",
			radius:      5,
			shape_type:  hw05.CircleType,
			expectation: 78.53981633974483,
			err:         nil,
		},
		{
			desc:        "Rectangle",
			width:       10,
			tight:       5,
			shape_type:  hw05.RectangleType,
			expectation: 50,
			err:         nil,
		},
		{
			desc:        "Triangle",
			base:        8,
			high:        6,
			shape_type:  hw05.TriangleType,
			expectation: 24,
			err:         nil,
		},
		{
			desc:        "Zero",
			base:        0,
			high:        6,
			shape_type:  hw05.TriangleType,
			expectation: 0,
			err:         nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var shape hw05.Shape
			switch tC.shape_type {
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
