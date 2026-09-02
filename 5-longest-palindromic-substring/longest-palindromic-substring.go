func longestPalindrome(s string) string {
    l, r := 0, 0
    answer := 0

    for i:=0;i<len(s);i++{
        odd := expand(s, i, i) 
        even := expand(s, i, i+1)
        if answer >= odd && answer >= even {
            continue
        }
        if even > odd {
            r = i + (even/2)
            l = i - ((even/2)-1)
            answer = even
        } else {
            l = i - (odd/2)
            r = i + (odd/2)
            answer = odd
        }
    }

    return s[l:r+1]
}


/*

xabbad
i = 2
length = 4
l = 1, r = 4

*/


func expand(s string, l, r int) int {
    if r >= len(s) || l < 0 {
        return 0
    }
    if l == r {
        return 1 + expand(s, l-1, r+1)
    }
    if s[l] == s[r] {
        return 2 + expand(s, l-1, r+1)
    }

    return 0
}