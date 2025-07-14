package main

import (
	"fmt"
	"go-problem-solving/algorithms/warmup"
)

func main() {
	var a, b uint32
	fmt.Scanf("%v\n%v", &a, &b)
	result := warmup.SolveMeFirst(a, b)
	fmt.Print(result)
}
