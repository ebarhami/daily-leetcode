func backspaceCompare(s string, t string) bool {
    si := len(s)-1
    ti := len(t)-1

    // iterate from last to first character for both strings
    for si >= 0 && ti >= 0 {

        si = clearBackspaces(s, si)
        ti = clearBackspaces(t, ti)
        
        // if current characters are not same after taking care
        // of the end backspaces, strings are not same
        if si >= 0 && ti >= 0 && s[si] != t[ti] {
            return false
        }

        si -= 1
        ti -= 1
    }

    // clear any final backspaces remaining
    si = clearBackspaces(s, si)
    ti = clearBackspaces(t, ti)

    return si == ti
}

// given a string s and a position i, decrement i to the point where
// s[i] won't be deleted/backspaced
func clearBackspaces(s string, i int) int {
    backspacesCnt := 0
    // keep track of backspaces and apply them by decrementing i
    for i >= 0 && (backspacesCnt > 0 || s[i] == '#') {
        if s[i] == '#' {
            // we'll take care of this in future loop iterations
            backspacesCnt++
        } else {
            // we're taking care of an earlier backspace in this iteration
            backspacesCnt--
        }
        i--
    }

    // we've reached a point where s[i] is neither a backspace nor affected by a backspace
    return i
}