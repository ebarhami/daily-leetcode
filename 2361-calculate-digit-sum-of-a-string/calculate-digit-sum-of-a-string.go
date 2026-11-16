func digitSum(s string, k int) string {
    for len(s) > k {
        s = doRound(s, k)
    }

    return s
}

func doRound(s string, k int) string {
    if len(s) <= k {
        return s
    }
    answer := ""
    for i:=0;i<len(s);i+=k{
        newS := s[i:min(i+k, len(s))]
        answer += sum(newS)
    }
    return answer
}

func sum(s string) string {
    if len(s) == 1 {
        return s
    }
    answer := 0
    for _, ch := range s {
        c := string(ch)
        num, _ := strconv.Atoi(c)
        answer += num
    }

    newString := strconv.Itoa(answer)
    return newString
}

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}