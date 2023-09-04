var grpMp map[byte][]string
var cacheMp map[string]int

func matchVal(s, st string) (int, bool) {
    if len(s) < len(st) {
        return 0, true
    }
    for j:=0;j<len(st);j++ {
        if s[j] != st[j] {
            return 0, true
        }
    }
    return len(st), false
}

func calc(s string) int {
    if val, ok := cacheMp[s]; ok {
        return val
    }

    if s == "" {
        cacheMp[s] = 0
        return 0
    }

    strArr, ok := grpMp[s[0]]
    if !ok {
        return 1 + calc(s[1:], )
    }

    ans := -1
    for _, str := range strArr {
        lenS, err := matchVal(s, str)
        if err {
            continue
        }
        temp := calc(s[lenS:])
        if ans == -1 {
            ans = temp
        }
        if temp < ans {
            ans = temp
        }
    }

    ans2 := 1 + calc(s[1:])
    if ans == -1 {
        cacheMp[s] = ans2
        return ans2
    }

    if ans < ans2 {
        cacheMp[s] = ans
        return ans
    }
    cacheMp[s] = ans2
    return ans2
}

func minExtraChar(s string, dict []string) int {
    cacheMp = make(map[string]int, 0)

    grpMp = make(map[byte][]string, 0)
    for _, str := range dict {
        if _, ok := grpMp[str[0]]; !ok {
            grpMp[str[0]] = make([]string, 0, len(dict))
        }
        grpMp[str[0]] = append(grpMp[str[0]], str)
    }

    return calc(s)
}