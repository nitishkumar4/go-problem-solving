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

// Problem link: https://www.hackerrank.com/challenges/compare-the-triplets/problem?isFullScreen=true
func main() {
	// create reader object
	reader := bufio.NewReader(os.Stdin)

	// ask user for first input
	fmt.Print("Enter space separated first triplet:")
	triplet1, _ := reader.ReadString('\n')
	triplet1Array := strings.Split(strings.TrimSpace(triplet1), " ")
	fmt.Print("Type", reflect.TypeOf(triplet1Array))
	var triplet1ArrayInt []int32
	for index := range triplet1Array {
		value, _ := strconv.ParseInt(strings.TrimSpace(triplet1Array[index]), 10, 32)
		triplet1ArrayInt = append(triplet1ArrayInt, int32(value))
	}
	fmt.Println("Type", reflect.TypeOf(triplet1ArrayInt))

	// ask user of second input
	fmt.Print("Enter space separated second triplet")
	triplet2, _ := reader.ReadString('\n')
	triplet2Array := strings.Split(strings.TrimSpace(triplet2), " ")
	fmt.Print("Type", reflect.TypeOf(triplet2Array))
	var triplet2ArrayInt []int32
	for index := range triplet2Array {
		value, _ := strconv.ParseInt(strings.TrimSpace(triplet2Array[index]), 10, 32)
		triplet2ArrayInt = append(triplet2ArrayInt, int32(value))
	}
	fmt.Println("Type", reflect.TypeOf(triplet2ArrayInt))

	// call CompareTheTriplets
	result := warmup.CompareTheTriplets(triplet1ArrayInt, triplet2ArrayInt)
	fmt.Print(result)

}
