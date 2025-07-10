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
	// create reader object
	reader := bufio.NewReader(os.Stdin)

	// input from user
	fmt.Print("Enter the number of elements in the array ")
	size, _ := reader.ReadString('\n')
	intSize, _ := strconv.ParseInt(strings.TrimSpace(size), 10, 32)

	fmt.Print("Enter the space seperated array elements ")
	array, _ := reader.ReadString('\n')

	// parse array elements from the input string
	var arrayOfInts []int32
	arrayOfStrings := strings.Split(strings.TrimSpace(array), " ")

	for i := 0; i < int(intSize); i++ {
		element, _ := strconv.ParseInt(arrayOfStrings[i], 10, 32)
		arrayOfInts = append(arrayOfInts, int32(element))
	}

	// call ReverseArray
	newArray := arrays.ReverseArray(arrayOfInts)
	if newArray != nil {
		fmt.Println("Reversed Array: ", newArray)
	}
}
