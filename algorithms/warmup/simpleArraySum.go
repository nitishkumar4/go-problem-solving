package warmup

func SimpleArraySum(ar []int32) int32 {
	sum := int32(0)
	for _, element := range ar {
		sum = sum + element
	}
	return sum
}
