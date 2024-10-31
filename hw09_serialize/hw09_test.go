package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHW09(t *testing.T) {
	testCases := []struct {
		desc        string
		book        Book
		expectation Book
	}{
		{
			desc: "Marshal and unmarshal",
			book: Book{
				ID:     1,
				Title:  "title",
				Author: "author",
				Year:   1,
				Size:   1,
				Rate:   1.0,
			},
			expectation: Book{
				ID:     1,
				Title:  "title",
				Author: "author",
				Year:   1,
				Size:   1,
				Rate:   1.0,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			results, err := tC.book.Marshaller()
			require.NoError(t, err)
			var newBook Book
			err = newBook.Unmarshaller(results)
			require.NoError(t, err)
			require.Equal(t, tC.expectation, newBook)
			if tC.expectation != tC.book {
				t.Errorf("expected %v, got %v", tC.expectation, tC.book)
			}
		})
	}
}

func TestHW09Slices(t *testing.T) {
	testCases := []struct {
		desc        string
		books       []Book
		expectation []Book
	}{
		{
			desc: "Slice Marshal and unmarshal",
			books: []Book{{
				ID:     1,
				Title:  "title",
				Author: "author",
				Year:   1,
				Size:   1,
				Rate:   1.0,
			},
				{
					ID:     2,
					Title:  "title2",
					Author: "author2",
					Year:   2,
					Size:   2,
					Rate:   2.0,
				}},
			expectation: []Book{{
				ID:     1,
				Title:  "title",
				Author: "author",
				Year:   1,
				Size:   1,
				Rate:   1.0,
			},
				{
					ID:     2,
					Title:  "title2",
					Author: "author2",
					Year:   2,
					Size:   2,
					Rate:   2.0,
				}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			results, err := SliceMarshaller(tC.books)
			require.NoError(t, err)
			result2, err := SliceUnmarshaller(results)
			require.NoError(t, err)
			require.Equal(t, tC.expectation, result2)
		})
	}
}
