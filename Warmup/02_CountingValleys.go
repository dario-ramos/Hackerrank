package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {
	var currentLevel = 0
	var insideValley bool = false
	var valleyCount int32 = 0
	var i int32
	for i = 0; i < n; i++ {
		if s[i] == 'D' {
			currentLevel--
		} else if s[i] == 'U' {
			currentLevel++
		}
		if !insideValley && currentLevel < 0 {
			insideValley = true
		}
		if insideValley && currentLevel >= 0 {
			insideValley = false
			valleyCount++
		}
	}
	return valleyCount
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	s := readLine(reader)

	result := countingValleys(n, s)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
