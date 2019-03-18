package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	var maxHourglassSum int32 = -70 // Something lower than -9 * 7
	// Hourglasses exist only for the inner matrix elements (i.e. non border)
	// So, iterate over those and calculate the hourglass sum for each
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			var currentHSum = arr[i-1][j-1] + arr[i-1][j] + arr[i-1][j+1] +
				arr[i][j] +
				arr[i+1][j-1] + arr[i+1][j] + arr[i+1][j+1]
			if currentHSum > maxHourglassSum {
				maxHourglassSum = currentHSum
			}
		}

	}
	return maxHourglassSum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := hourglassSum(arr)

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
