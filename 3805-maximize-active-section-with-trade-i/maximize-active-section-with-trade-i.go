func maxActiveSectionsAfterTrade(s string) int {
    currZero, prevZero := 0, 0
    total := 0

    maxZero := 0
    for i, ch := range s {
        if ch == '1' {
            if i > 0 && s[i-1] == '0' {
                prevZero = currZero
                currZero = 0
            }
            total++
        } else {
            currZero++
        }
        if prevZero > 0 && currZero > 0 {
            
            maxZero = max(maxZero, prevZero + currZero)
        }
    } 

    return total + maxZero
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}