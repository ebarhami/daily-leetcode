func shortestBeautifulSubstring(s string, k int) string {
	if strings.Count(s, "1") < k {
		return ""
	}
	ans := s
	cnt, left := 0, 0
	for right := 0; right < len(s); right++ {
		cnt += int(s[right] - '0')
		for cnt > k || s[left] == '0' {
			cnt -= int(s[left] - '0')
			left++
		}
		if cnt == k {
			t := s[left : right+1]
			if len(t) < len(ans) || len(t) == len(ans) && t < ans {
				ans = t
			}
		}
	}
	return ans
}