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
	defer file.Close()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open the CSV file")
	}

	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read the CSV file")
	}

	return lines, nil
}

func main() {
	var filePath string
	var logLevel string
	var output string
	var flagFilePath string
	var FlagLogLevel string
	var FlagOutput string
	var failedConnections int
	IPMap := make(map[string]int)

	pflag.StringVarP(&flagFilePath, "file", "f", "", "path to log file")
	pflag.StringVarP(&FlagLogLevel, "level", "l", "", "log level")
	pflag.StringVarP(&FlagOutput, "output", "o", "", "output type")
	pflag.Parse()

	envFilePath, ok := os.LookupEnv("LOG_ANALYZER_FILE")
	if ok && flagFilePath == "" {
		filePath = envFilePath
	} else if !ok && flagFilePath == "" {
		fmt.Println("file path is not set")
		return
	} else {
		filePath = flagFilePath
	}

	envLogLevel, ok := os.LookupEnv("LOG_ANALYZER_LEVEL")
	if ok && FlagLogLevel == "" {
		logLevel = envLogLevel
	} else if !ok && FlagLogLevel == "" {
		logLevel = "info"
	} else {
		logLevel = FlagLogLevel
	}

	envOutput, ok := os.LookupEnv("LOG_ANALYZER_OUTPUT")
	if ok && FlagOutput == "" {
		output = envOutput
	} else if !ok && FlagOutput == "" {
		output = "console"
	} else {
		output = FlagOutput
	}

	fmt.Printf("Reading file: %s with log level: %s. Output: %s\n", filePath, logLevel, output)
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

		outputFile, err := os.Create(output)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
		defer outputFile.Close()

		var outputLine strings.Builder
		outputLine.WriteString(fmt.Sprintf("Failed connections: %d\n", failedConnections))
		_, err = outputFile.WriteString(outputLine.String())
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}

		outputLine.Reset()
		outputLine.WriteString(fmt.Sprintf("IPs: %v\n", IPMap))
		_, err = outputFile.WriteString(outputLine.String())
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
	}
}
