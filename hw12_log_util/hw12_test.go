package main

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc        string
		logFile     string
		logLevel    string
		outputFile  string
		expectftion string
	}{
		{
			desc:        "Read info",
			logFile:     "log.txt",
			logLevel:    "info",
			outputFile:  "console",
			expectftion: "Failed connections: 0\nIPs: map[127.0.0.1:2]\n",
		},
		{
			desc:        "Read error",
			logFile:     "log.txt",
			logLevel:    "error",
			outputFile:  "console",
			expectftion: "Failed connections: 3\nIPs: map[127.0.0.1:1 8.8.8.8:2]\n",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			out, _ := exec.Command("go", "run", "main.go", "-f", tC.logFile, "-l", tC.logLevel, "-o", tC.outputFile).Output()
			require.Equal(t, tC.expectftion, string(out))
		})
	}
}
