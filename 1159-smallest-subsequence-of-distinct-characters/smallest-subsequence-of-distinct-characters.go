func smallestSubsequence(s string) string {
	left := [26]int{}
	for _, ch := range s {
		left[ch-'a']++
	}
	stack := []byte{}
	inStack := [26]bool{}
	for i := range s {
		ch := s[i]
		if !inStack[ch-'a'] {
			for len(stack) > 0 && ch < stack[len(stack)-1] {
				last := stack[len(stack)-1] - 'a'
				if left[last] == 0 {
					break
				}
				stack = stack[:len(stack)-1]
				inStack[last] = false
			}
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}
		left[ch-'a']--
	}
	return string(stack)
}


// func smallestSubsequence(s string) string {
//     lexiOrder := make([]int, 0)
//     exist := make(map[rune]bool)

//     for _, ch := range s {
//         if _, exist := exist[ch]; !exist{
//             lexiOrder = append(lexiOrder, int(ch - 'a'))
//             exist[ch] = true
//         }
//     }


// }

/*
cbacdcbc

a, b, c, 

a 2
b 1, 6
c 0, 3, 5, 7
d 4
*/