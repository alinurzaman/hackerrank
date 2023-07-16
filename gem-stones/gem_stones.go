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
 * Complete the 'gemstones' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING_ARRAY arr as parameter.
 */

func gemstones(arr []string) int32 {
	var list []rune
	first := arr[0]
	for _, ch := range first {
		if !contains(list, ch) {
			list = append(list, ch)
		}
	}
	for i := 1; i < len(arr); i++ {
		tmp := list[:0]
		for _, p := range list {
			if letterContains(arr[i], p) {
				tmp = append(tmp, p)
			}
		}
		list = tmp
	}

	return int32(len(list))
}

func contains(arr []rune, key rune) bool {
	for _, v := range arr {
		if v == key {
			return true
		}
	}
	return false
}

func letterContains(letter string, key rune) bool {
	for _, v := range letter {
		if v == key {
			return true
		}
	}
	return false
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

	var arr []string

	for i := 0; i < int(n); i++ {
		arrItem := readLine(reader)
		arr = append(arr, arrItem)
	}

	result := gemstones(arr)

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
