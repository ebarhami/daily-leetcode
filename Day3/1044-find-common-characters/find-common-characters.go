func commonChars(words []string) []string {
    base := getCount(words[0])

    for i:=1;i<len(words);i++{
        comp := getCount(words[i])
        for key, val := range base {
            base[key] = min(val, comp[key])
        }
    }
    answer := make([]string, 0)
    for key, val := range base {
        if val == 0 {continue}
        for i:=0;i<val;i++{
            answer = append(answer, string(key))
        }
    }

    return answer
}

func min(a, b int) int {
    if a < b {return a}
    return b
} 

func getCount(s string) map[byte]int {
    count := make(map[byte]int)

    for j:=0;j<len(s);j++{
        count[s[j]]++
    }
    return count
}