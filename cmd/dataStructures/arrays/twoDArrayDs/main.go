package main

import (
	"bufio"
	"fmt"
	"go-problem-solving/dataStructures/arrays"
	"os"
	"strconv"
	"strings"
)

func main() {
	// create the reader object
	reader := bufio.NewReader(os.Stdin)

	var twoDArray [][]int32
	var oneDArray []int32
	for i := 0; i < 6; i++ {
		row, err := reader.ReadString('\n')
		checkError(err)
		rowString := strings.Split(strings.TrimSpace(row), " ")
		for index := range rowString {
			element, err := strconv.ParseInt(rowString[index], 10, 32)
			checkError(err)
			oneDArray = append(oneDArray, int32(element))
		}
		twoDArray = append(twoDArray, oneDArray)
	}

	result := arrays.HourglassSum(twoDArray)
	fmt.Println(result)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
