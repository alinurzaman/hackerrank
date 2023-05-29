package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the catAndMouse function below.
func catAndMouse(x int32, y int32, z int32) string {
	var A int32 = 0
	var B int32 = 0
	if x < z {
		A = z - x
	} else if x > z {
		A = x - z
	}
	if y < z {
		B = z - y
	} else if y > z {
		B = y - z
	}
	if A < B {
		return "Cat A"
	} else if A > B {
		return "Cat B"
	} else {
		return "Mouse C"
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		xyz := strings.Split(readLine(reader), " ")

		xTemp, err := strconv.ParseInt(xyz[0], 10, 64)
		checkError(err)
		x := int32(xTemp)

		yTemp, err := strconv.ParseInt(xyz[1], 10, 64)
		checkError(err)
		y := int32(yTemp)

		zTemp, err := strconv.ParseInt(xyz[2], 10, 64)
		checkError(err)
		z := int32(zTemp)

		result := catAndMouse(x, y, z)

		fmt.Fprintf(writer, "%s\n", result)
	}

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
