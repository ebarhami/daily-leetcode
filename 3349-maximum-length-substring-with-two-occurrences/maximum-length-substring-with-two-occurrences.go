func maximumLengthSubstring(s string) int {
    counter := make(map[byte]int)

    answer := 0
    left := 0
    for i, _ := range s {
        counter[s[i]]++
        for counter[s[i]] > 2 {
            counter[s[left]]--
            left++
        }

        if i-left+1 > answer {
            answer = i-left+1
        }
    }

    return answer
}