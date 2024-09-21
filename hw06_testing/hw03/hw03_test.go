package hw03_test

import (
	"errors"
	"testing"

	"github.com/freeloginname/home_work_basic/hw06_testing/hw03"
	"github.com/stretchr/testify/require"
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
			/**
			* можно ли сделать универсальную проверку средствами require.
			* для ситуаций когда ошидаешь ошибку и когда ее не должно быть?.
			**/
			require.Equal(t, tC.err, err)
			// require.NoError(t, err)
			require.Equal(t, tC.board, got)
		})
	}
}
