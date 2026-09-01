func countBinarySubstrings(s string) int {
    curr := 1
    prev := 0
    answer := 0
    for i:=1;i<len(s);i++{
        if s[i] == s[i-1] {
            curr++
        } else {
            answer += min(curr, prev)
            prev = curr
            curr = 1
        }
    }

    return answer + min(curr,prev)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}