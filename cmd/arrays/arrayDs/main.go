package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// create reader object
	reader := bufio.NewReader(os.Stdin)

	// input from user
	fmt.Print("Enter the number of elements in the array")
	size, _ := reader.ReadString('\n')
	intSize, _ := strconv.ParseInt(size, 10, 32)

	fmt.Print("Enter the space seperated array elements")
	array, _ := reader.ReadString('\n')

	// parse array elements from the input string
	var arrayOfInts []int
	arrayOfStrings := strings.Split(strings.TrimSpace(array), " ")
	for i := 0; i < int(intSize); i++ {
		element, _ := strconv.ParseInt(arrayOfStrings[i], 10, 32)
		arrayOfInts = append(arrayOfInts, int(element))
	}

	// call reverseArray
	newArray := reverseArray(arrayOfInts)
	if newArray != nil {
		fmt.Println("Reversed Array: ", newArray)
	}
}
