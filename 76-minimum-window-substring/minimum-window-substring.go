func minWindow(s string, t string) string {
    // required[c] = how many times c is needed
    required := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        required[t[i]]++
    }

    // window[c] = how many times c currently exists in the window
    window := make(map[byte]int)

    // Number of distinct characters whose required frequency is satisfied.
    satisfied := 0
    need := len(required)

    left := 0
    bestLeft := 0
    bestRight := len(s) + 1

    for right := 0; right < len(s); right++ {
        c := s[right]

        if required[c] > 0 {
            window[c]++

            // We just reached the required frequency.
            if window[c] == required[c] {
                satisfied++
            }
        }

        // Window is valid. Try to make it smaller.
        for satisfied == need {
            if right-left+1 < bestRight-bestLeft {
                bestLeft = left
                bestRight = right + 1
            }

            c := s[left]

            if required[c] > 0 {
                window[c]--

                // Removing c made this requirement unsatisfied.
                if window[c] < required[c] {
                    satisfied--
                }
            }

            left++
        }
    }

    if bestRight == len(s)+1 {
        return ""
    }

    return s[bestLeft:bestRight]
}
