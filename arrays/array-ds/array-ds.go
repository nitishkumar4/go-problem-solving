package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"reflect"
	"strconv"
)

func reverseArray(a []int32) []int32 {
	// Write your code here
	l := len(a)
	for i := 0; i <= int(l/2)-1; i++ {
		a[i], a[l-i-1] = a[l-i-1], a[i]
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number of elements in array:")
	size, _ := reader.ReadString('\n') // Read full line
	size = strings.TrimSpace(size) // remove newline and spaces
	fmt.Println(size)
	
	fmt.Print("Enter space seperated array elements:")
	input, _ := reader.ReadString('\n') // Read full line
	input = strings.TrimSpace(input) // remove newline and spaces
	fmt.Println(input)
	
	array := strings.Split(input, " ")
	fmt.Println(array)
	
	fmt.Println("Type:", reflect.TypeOf(array))
	
	var intArray []int32
	
	for i:=0; i<len(array); i++ {
		n, _ := strconv.ParseInt(array[i],10,32)
		intArray = append(intArray, int32(n))
	}
	
	fmt.Println("Integer Array", intArray)
	fmt.Println("Type:", reflect.TypeOf(intArray))
	
	fmt.Print("Reversed array:")
	fmt.Println(reverseArray(intArray))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
