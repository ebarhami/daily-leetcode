func countBinarySubstrings(s string) int {
    prev, curr := 0, 1
    answer := 0

    for i := 1; i < len(s); i++ {
        if s[i] == s[i-1] {
            curr++
        } else {
            answer += min(prev, curr)
            prev = curr
            curr = 1
        }
    }

    return answer + min(prev, curr)
}