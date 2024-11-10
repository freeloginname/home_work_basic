package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

func ReadCsv(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open the CSV file")
	}
	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read the CSV file")
	}
	defer file.Close()
	return lines, nil
}

func ReadParams() (string, string, string, error) {
	var flagFilePath string
	var FlagLogLevel string
	var FlagOutput string
	var filePath string
	var logLevel string
	var output string

	pflag.StringVarP(&flagFilePath, "file", "f", "", "path to log file")
	pflag.StringVarP(&FlagLogLevel, "level", "l", "", "log level")
	pflag.StringVarP(&FlagOutput, "output", "o", "", "output type")
	pflag.Parse()

	envFilePath, ok := os.LookupEnv("LOG_ANALYZER_FILE")
	switch {
	case ok && flagFilePath == "":
		filePath = envFilePath
	case !ok && flagFilePath == "":
		return "", "", "", errors.New("file path is not set")
	default:
		filePath = flagFilePath
	}

	envLogLevel, ok := os.LookupEnv("LOG_ANALYZER_LEVEL")
	switch {
	case ok && FlagLogLevel == "":
		logLevel = envLogLevel
	case !ok && FlagLogLevel == "":
		logLevel = "info"
	default:
		logLevel = FlagLogLevel
	}

	envOutput, ok := os.LookupEnv("LOG_ANALYZER_OUTPUT")
	switch {
	case ok && FlagOutput == "":
		output = envOutput
	case !ok && FlagOutput == "":
		output = "console"
	default:
		output = FlagOutput
	}
	return filePath, logLevel, output, nil
}

func main() {
	var failedConnections int
	IPMap := make(map[string]int)
	filePath, logLevel, output, err := ReadParams()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	// for debug:
	// fmt.Printf("Reading file: %s with log level: %s. Output: %s\n", filePath, logLevel, output)
	logs, err := ReadCsv(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	for _, line := range logs {
		if line[0] != logLevel {
			continue
		}
		IPMap[line[2]]++
		if strings.ToLower(line[3]) == "fail" {
			failedConnections++
		}
	}

	if output == "console" {
		fmt.Printf("Failed connections: %d\n", failedConnections)
		fmt.Printf("IPs: %v\n", IPMap)
	} else {
		outputFile, errr := os.Create(output)
		if errr != nil {
			fmt.Printf("Error: %v", errr)
			return
		}
		defer outputFile.Close()

		var outputLine strings.Builder
		outputLine.WriteString(fmt.Sprintf("Failed connections: %d\n", failedConnections))
		_, errr = outputFile.WriteString(outputLine.String())
		if errr != nil {
			fmt.Printf("Error: %v", errr)
			return
		}

		outputLine.Reset()
		outputLine.WriteString(fmt.Sprintf("IPs: %v\n", IPMap))
		_, errr = outputFile.WriteString(outputLine.String())
		if errr != nil {
			fmt.Printf("Error: %v", errr)
			return
		}
	}
}
