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
 * Complete the 'permutationEquation' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY p as parameter.
 */
func indexOf(element int32, data []int32) int32 {
	for k, v := range data {
		if element == v {
			return int32(k)
		}
	}
	return -1 //not found.
}

func permutationEquation(p []int32) []int32 {
	// Write your code here
	result := []int32{}
	var i int32
	for i = 1; i <= int32(len(p)); i++ {
		result = append(result, indexOf(indexOf(i, p)+1, p)+1)
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	pTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var p []int32

	for i := 0; i < int(n); i++ {
		pItemTemp, err := strconv.ParseInt(pTemp[i], 10, 64)
		checkError(err)
		pItem := int32(pItemTemp)
		p = append(p, pItem)
	}

	result := permutationEquation(p)

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
