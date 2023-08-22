func convertToTitle(columnNumber int) string {
    answer := ""
    
    for columnNumber > 0 {
        val := getOrderInAlphabet(columnNumber)
        answer = fmt.Sprintf("%s%s", string(rune(int('A') + val - 1)), answer)
        residu := 0
        if columnNumber % 26 == 0 {residu = 1}
        columnNumber = columnNumber/26 - residu
    }
    return answer
}

func getOrderInAlphabet(x int) int {
    a := x % 26
    if a == 0 {a = 26}
    return a
}