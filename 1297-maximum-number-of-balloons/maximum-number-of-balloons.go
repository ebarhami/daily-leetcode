func maxNumberOfBalloons(text string) int {
    freq := make(map[rune]int)

    for _, ch := range text {
        freq[ch]++
    }
    answer := freq['b']
    answer = min(answer, freq['a'])
    answer = min(answer, freq['l']/2)
    answer = min(answer, freq['o']/2)
    answer = min(answer, freq['n'])

    return answer
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}