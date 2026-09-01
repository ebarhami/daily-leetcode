func countBinarySubstrings(s string) int {
    counter := make([]int, 0)

    prev := rune(s[0])
    s += "x"
    count := 0
    for _, ch := range s {
        if ch == prev  {
            count++
        } else {
            counter = append(counter, count)
            prev = ch 
            count = 1
        }
    }

    fmt.Println(counter)
    answer := 0
    for i:=1;i<len(counter);i++{
        answer += min(counter[i], counter[i-1])
    }

    return answer
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}