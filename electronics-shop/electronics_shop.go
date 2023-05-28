package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the getMoneySpent function below.
 */
func getMoneySpent(keyboards []int, drives []int, b int) int {
	/*
	 * Write your code here.
	 */
	total := -1
	sort.Ints(keyboards)
	sort.Ints(drives)
	if keyboards[0]+drives[0] > b {
		return total
	}
	if keyboards[len(keyboards)-1]+drives[len(drives)-1] <= b {
		return keyboards[len(keyboards)-1] + drives[len(drives)-1]
	}

	for i := len(keyboards) - 1; i >= 0; i-- {
		for j := len(drives) - 1; j >= 0; j-- {
			if keyboards[i]+drives[j] <= b && keyboards[i]+drives[j] > total {
				total = keyboards[i] + drives[j]
			}
		}
	}
	return total
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	bnm := strings.Split(readLine(reader), " ")

	bTemp, err := strconv.ParseInt(bnm[0], 10, 64)
	checkError(err)
	b := int(bTemp)

	nTemp, err := strconv.ParseInt(bnm[1], 10, 64)
	checkError(err)
	n := int(nTemp)

	mTemp, err := strconv.ParseInt(bnm[2], 10, 64)
	checkError(err)
	m := int(mTemp)

	keyboardsTemp := strings.Split(readLine(reader), " ")

	var keyboards []int

	for keyboardsItr := 0; keyboardsItr < int(n); keyboardsItr++ {
		keyboardsItemTemp, err := strconv.ParseInt(keyboardsTemp[keyboardsItr], 10, 64)
		checkError(err)
		keyboardsItem := int(keyboardsItemTemp)
		keyboards = append(keyboards, keyboardsItem)
	}

	drivesTemp := strings.Split(readLine(reader), " ")

	var drives []int

	for drivesItr := 0; drivesItr < int(m); drivesItr++ {
		drivesItemTemp, err := strconv.ParseInt(drivesTemp[drivesItr], 10, 64)
		checkError(err)
		drivesItem := int(drivesItemTemp)
		drives = append(drives, drivesItem)
	}

	/*
	 * The maximum amount of money she can spend on a keyboard and USB drive, or -1 if she can't purchase both items
	 */

	moneySpent := getMoneySpent(keyboards, drives, b)

	fmt.Fprintf(writer, "%d\n", moneySpent)

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
