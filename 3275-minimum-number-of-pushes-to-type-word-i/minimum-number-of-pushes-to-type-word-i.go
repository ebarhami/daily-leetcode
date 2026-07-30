func minimumPushes(word string) int {
	n := len(word)
	ans := 0
	for i := 0; i < n; i++ {
		ans += i/8 + 1
	}
	return ans
}