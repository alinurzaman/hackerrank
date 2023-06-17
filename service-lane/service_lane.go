package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'serviceLane' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY cases
 */

func serviceLane(width []int32, cases [][]int32) []int32 {
	// Write your code here
	var max []int32
	for _, c := range cases {
		slices := width[c[0] : c[1]+1]
		max = append(max, getMin(slices))
	}
	return max
}

func getMin(values []int32) int32 {
	var min int32 = values[0]
	for _, value := range values {
		if min > value {
			min = value
		}
	}
	return min
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	tTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	t := int32(tTemp)

	widthTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var width []int32

	for i := 0; i < int(n); i++ {
		widthItemTemp, err := strconv.ParseInt(widthTemp[i], 10, 64)
		checkError(err)
		widthItem := int32(widthItemTemp)
		width = append(width, widthItem)
	}

	var cases [][]int32
	for i := 0; i < int(t); i++ {
		casesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var casesRow []int32
		for _, casesRowItem := range casesRowTemp {
			casesItemTemp, err := strconv.ParseInt(casesRowItem, 10, 64)
			checkError(err)
			casesItem := int32(casesItemTemp)
			casesRow = append(casesRow, casesItem)
		}

		if len(casesRow) != 2 {
			panic("Bad input")
		}

		cases = append(cases, casesRow)
	}

	result := serviceLane(width, cases)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
