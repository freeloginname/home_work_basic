package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc    string
		counter int
	}{
		{
			desc:    "Run counter",
			counter: 10,
		},
		{
			desc:    "Run counter 0",
			counter: 0,
		},

		{
			desc:    "Run counter 42",
			counter: 42,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			counter := CounterStarter(tC.counter)
			require.Equal(t, tC.counter, counter)

		})
	}
}
