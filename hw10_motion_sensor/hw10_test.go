package main

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "Test run",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			main()
		})
	}
}
