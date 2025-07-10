package arrays

func ReverseArray(a []int32) []int32 {
	// Write your code here
	for i, l := 0, len(a); i <= int(l/2)-1; i++ {
		a[i], a[l-i-1] = a[l-i-1], a[i]
	}
	return a
}
