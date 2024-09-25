package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHW08(t *testing.T) {
	testCases := []struct {
		desc        string
		data        []int
		target      int
		expectation int
	}{
		{
			desc:        "middle",
			data:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:      5,
			expectation: 4,
		},
		{
			desc:        "first",
			data:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:      1,
			expectation: 0,
		},
		{
			desc:        "last",
			data:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:      10,
			expectation: 9,
		},
		{
			desc:        "not found",
			data:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:      0,
			expectation: -1,
		},
		{
			desc:        "empty",
			data:        []int{},
			target:      0,
			expectation: -1,
		},
		{
			desc:        "empty target",
			data:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:      0,
			expectation: -1,
		},
		{
			desc:        "empty data",
			data:        []int{},
			target:      0,
			expectation: -1,
		},
		{
			desc:        "target not found 2",
			data:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:      11,
			expectation: -1,
		},
		{
			desc:        "data not sorted",
			data:        []int{1, 2, 11, 4, 5, 6, 7, 8, 9, 10},
			target:      9,
			expectation: 8,
		},
		{
			desc:        "data not sorted 1",
			data:        []int{1, 2, 11, 4, 5, 6, 7, 8, 9, 10},
			target:      2,
			expectation: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			results := BinarySearch(tC.data, tC.target)
			require.Equal(t, tC.expectation, results)
		})
	}
}
