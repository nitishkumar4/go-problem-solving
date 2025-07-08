package arrayDs

func ReverseArray(a []int32) []int32 {
	// Write your code here
	l := len(a)
	for i := 0; i <= int(l/2)-1; i++ {
		a[i], a[l-i-1] = a[l-i-1], a[i]
	}
	return a
}
