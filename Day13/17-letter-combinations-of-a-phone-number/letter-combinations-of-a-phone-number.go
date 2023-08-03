func letterCombinations(digits string) []string {
    phone := make(map[rune][]string)
    phone['2'] = []string{"a", "b", "c"}
    phone['3'] = []string{"d", "e", "f"}
    phone['4'] = []string{"g", "h", "i"}
    phone['5'] = []string{"j", "k", "l"}
    phone['6'] = []string{"m", "n", "o"}
    phone['7'] = []string{"p", "q", "r", "s"}
    phone['8'] = []string{"t", "u", "v"}
    phone['9'] = []string{"w", "x", "y", "z"}

    answer := make([]string, 0)
    newAnswer := make([]string, 0)
    for _, digit := range digits {
        d := rune(digit)
        if len(answer) == 0 {
            for _, letter := range phone[d] {
                answer = append(answer, letter)
            }
        } else {
            newAnswer = make([]string, 0)
            for _, str := range answer {
                for _, letter := range phone[d] {
                    newAnswer = append(newAnswer, str+letter)
                }
            }
            answer = newAnswer
        }
    }
    return answer
}