func lexGreaterPermutation(s string, target string) string {
	cnt := make([]int, 26)
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
		cnt[target[i]-'a']--
	}

	// Try from right to left
	t := []byte(target)
	for i := len(s) - 1; i >= 0; i-- {
		b := t[i] - 'a'
		cnt[b]++ // Reversal of consumption
		// Check if the prefix can fully match
		if min(cnt) < 0 {
			continue
		}
		// Find the smallest available character larger than b.
		for j := b + 1; j < 26; j++ {
			if cnt[j] > 0 {
				cnt[j]--
				t[i] = byte('a' + j)
				return string(t[:i+1]) + getMinString(cnt)
			}
		}
	}

	return ""
}

func min(arr []int) int {
	m := arr[0]
	for _, v := range arr {
		if v < m {
			m = v
		}
	}
	return m
}

// Get the lexicographically smallest string (in ascending order)
func getMinString(cnt []int) string {
	var res []byte
	for i := 0; i < 26; i++ {
		res = append(res, bytes.Repeat([]byte{byte('a' + i)}, cnt[i])...)
	}
	return string(res)
}