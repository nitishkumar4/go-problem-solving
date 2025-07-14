package main

import (
	"bufio"
	"fmt"
	"go-problem-solving/algorithms/warmup"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	// create a reader object
	reader := bufio.NewReader(os.Stdin)

	// read the size of the array
	fmt.Print("Enter the size of the array: ")
	size, _ := reader.ReadString('\n')
	intSize, _ := strconv.ParseInt(strings.TrimSpace(size), 10, 32)
	fmt.Println("Type:", reflect.TypeOf(intSize))

	// read the array of ints
	fmt.Print("Enter the space separated array of integers: ")
	array, _ := reader.ReadString('\n')
	arrayOfStrings := strings.Split(strings.TrimSpace(array), " ")
	var arrayOfInts []int32
	for _, element := range arrayOfStrings {
		intElement, _ := strconv.ParseInt(element, 10, 32)
		arrayOfInts = append(arrayOfInts, int32(intElement))
	}
	fmt.Println("Type: ", reflect.TypeOf(arrayOfInts))

	result := warmup.SimpleArraySum(arrayOfInts)
	fmt.Print(result)
}
