func backspaceCompare(s string, t string) bool {
    i, j := len(s)-1, len(t)-1
    for i >= 0 && j >= 0 {
        i = getIndexAfterBackspace(s, i)
        j = getIndexAfterBackspace(t, j)

        fmt.Printf("%d %d\n", i, j)

        if i>=0 && j >= 0 && s[i] != t[j] {
            return false
        }

        i--
        j--
    }
    i = getIndexAfterBackspace(s, i)
    j = getIndexAfterBackspace(t, j)

    return i == j
}

func getIndexAfterBackspace(s string, i int) int {
    counter := 0
    for i >= 0 && s[i] == '#' {
        counter = 1
        i--
        for i>=0 && counter > 0 {
            if s[i] == '#' {
                counter++
            } else {
                counter--
            }
            i--
        }
    }
    return i
}

/*
i,j = n-1, n-1
abc## <- i
ab# <- j

cntBackspace = 0
i, j = 0


*/
