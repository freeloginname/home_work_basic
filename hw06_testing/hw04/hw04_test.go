package hw04_test

import (
	"testing"

	"github.com/freeloginname/home_work_basic/hw06_testing/hw04"
	"github.com/stretchr/testify/require"
)

func TestHW04(t *testing.T) {
	testCases := []struct {
		desc             string
		compareBy        hw04.ComparisonType
		firstBook        hw04.Book
		secondBook       hw04.Book
		firstBookID      int
		firstBookTitle   string
		firstBookAuthor  string
		firstBookYear    int
		firstBookSize    int
		firstBookRate    float32
		SecondBookID     int
		SecondBookTitle  string
		SecondBookAuthor string
		SecondBookYear   int
		SecondBookSize   int
		SecondBookRate   float32
		expectation      bool
	}{
		{
			/**
			* как можно определить значения для Book на этом этапе,.
			* если они задаются через вызов методов?.
			 **/
			desc:             "Year comparison",
			compareBy:        hw04.Year,
			firstBookID:      1,
			firstBookTitle:   "aaa",
			firstBookAuthor:  "aaaaa",
			firstBookYear:    1990,
			firstBookSize:    10,
			firstBookRate:    2.0,
			SecondBookID:     2,
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
			firstBookID:      1,
			firstBookTitle:   "aaa",
			firstBookAuthor:  "aaaaa",
			firstBookYear:    1990,
			firstBookSize:    10,
			firstBookRate:    2.0,
			SecondBookID:     2,
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
			firstBookID:      1,
			firstBookTitle:   "aaa",
			firstBookAuthor:  "aaaaa",
			firstBookYear:    1990,
			firstBookSize:    10,
			firstBookRate:    2.0,
			SecondBookID:     2,
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
			firstBook.SetBookID(tC.firstBookID)
			firstBook.SetBookTitle(tC.firstBookTitle)
			firstBook.SetBookAuthor(tC.firstBookAuthor)
			firstBook.SetBookYear(tC.firstBookYear)
			firstBook.SetBookSize(tC.firstBookSize)
			firstBook.SetBookRate(tC.firstBookRate)
			secondBook.SetBookID(tC.SecondBookID)
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
