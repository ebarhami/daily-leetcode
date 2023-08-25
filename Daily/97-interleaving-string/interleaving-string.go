func isInterleave(s1 string, s2 string, s3 string) bool {
    if len(s1) + len(s2) != len(s3) {
        return false
    }

    cache := make([][]int, len(s1)+1)
    for i := 0; i <= len(s1); i++ {
        cache[i] = make([]int, len(s2)+1)
        for j := 0; j <= len(s2); j++ {
            cache[i][j] = -1
        }
    }
    return isInterleaved(0, 0, 0, s1, s2, s3, cache)
}

func isInterleaved(i, j, k int, s1, s2, s3 string, cache [][]int) bool {
    if k == len(s3) {
        if i == len(s1) && j == len(s2) {
            return true
        }
        return false
    }

    if cache[i][j] != -1 {
        return cache[i][j] == 1
    }

    cache[i][j] = 0
    res := false
    if i < len(s1) && s1[i] == s3[k] {
        res = res || isInterleaved(i+1, j, k+1, s1, s2, s3, cache)
    }
    if j < len(s2) && s2[j] == s3[k] {
        res = res || isInterleaved(i, j+1, k+1, s1, s2, s3, cache)
    }
    if res {
        cache[i][j] = 1
    }
    return res
}