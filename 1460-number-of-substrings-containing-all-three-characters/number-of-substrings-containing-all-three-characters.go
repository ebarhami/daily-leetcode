func numberOfSubstrings(s string) int {
    n := len(s)
    freq:=make([][]int, 3)
    for i:=0;i<3;i++{
        freq[i] = make([]int, n)
    }

    for i, ch := range s {
        if i > 0 {
            freq[0][i] += freq[0][i-1]
            freq[1][i] += freq[1][i-1]
            freq[2][i] += freq[2][i-1]
        }
        if ch == 'a' {
            freq[0][i]++
        }
        if ch == 'b' {
            freq[1][i]++
        }
        if ch == 'c' {
            freq[2][i]++
        }
    }

    l, r := 0,0
    answer := 0

    for l <= r && r < len(s) {
        if satisfied(freq, s, l, r) {
            answer+=(n-r)
            l++
        } else {
            r++
        }
    }

    return answer
}

func satisfied(freq [][]int, s string, l, r int) bool {
    a := freq[0][r] - freq[0][l]
    if s[l] == 'a' {
        a++
    }

    b := freq[1][r] - freq[1][l]
    if s[l] == 'b' {
        b++
    }

    c := freq[2][r] - freq[2][l]
    if s[l] == 'c' {
        c++
    }

    return a > 0 && b > 0 && c > 0
}

/*
abcabc
abc

*/