package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	const toSearch = "a"
	// How many a's appear in a single prefix?
	var perPrefix = strings.Count(s, toSearch)
	// How many complete prefixes appear in n?
	var fullPrefixesCount = n / int64(len(s))
	// How many a's in the remaining characters?
	var remainingChars = n % int64(len(s))
	var remainingCount = strings.Count(s[0:remainingChars], toSearch)
	return fullPrefixesCount*int64(perPrefix) + int64(remainingCount)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

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
