func numOfStrings(patterns []string, word string) int {
    answer := 0
    for _, pattern := range patterns {
        if isExist(pattern, word) {
            fmt.Println(pattern)
            answer++
        }
    }

    return answer
}

func isExist(pattern, word string) bool {
    n := len(pattern)

    for i:=0;i<=len(word)-n;i++{
        s := word[i:i+n]
        if s == pattern {
            return true
        }
    }

    return false
}