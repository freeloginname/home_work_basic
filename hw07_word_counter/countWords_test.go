package countWords_test

import (
	"testing"

	countWords "github.com/freeloginname/home_work_basic/hw07_word_counter"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc        string
		inputString string
		expectation map[string]int
	}{
		{
			desc:        "simple test",
			inputString: "hello world",
			expectation: map[string]int{"hello": 1, "world": 1},
		},
		{
			desc:        "simple test with punctuation",
			inputString: "hello, world!",
			expectation: map[string]int{"hello": 1, "world": 1},
		},
		{
			desc:        "empty string",
			inputString: "",
			expectation: map[string]int{},
		},
		{
			desc:        "empty string with spaces",
			inputString: "   ",
			expectation: map[string]int{},
		},
		{
			desc:        "empty string with punctuation",
			inputString: ",.!?;",
			expectation: map[string]int{},
		},
		{
			desc:        "multiple words with punctuation",
			inputString: "hello, world! hello@%, world!  ,.!?;",
			expectation: map[string]int{"hello": 1, "world": 2, "hello@%": 1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			wordsAppearance := countWords.PublicCountWords(tC.inputString)
			require.Equal(t, tC.expectation, wordsAppearance)
		})
	}
}