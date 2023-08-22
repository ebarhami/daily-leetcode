func convertToTitle(columnNumber int) string {
    answer := ""
    
    for columnNumber > 0 {
        columnNumber--

        answer = fmt.Sprintf("%s%s", string(rune(int('A') + columnNumber % 26)), answer)

        columnNumber = columnNumber/26
    }
    return answer
}

func getOrderInAlphabet(x int) int {
    a := x % 26
    if a == 0 {a = 26}
    return a
}