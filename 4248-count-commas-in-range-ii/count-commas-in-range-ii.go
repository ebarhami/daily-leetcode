func countCommas(n int64) int64 {
    ceiling := int64(1000)
    answer := int64(0)
    multiplier := int64(1)
    
    for ceiling <= n {
        nextCeil := ceiling * 1000
        if n < nextCeil {
            answer += (n-ceiling+1) * multiplier
        } else {
            answer += (nextCeil - ceiling) * multiplier
        }

        multiplier++
        ceiling = nextCeil
    }

    return answer
}

/*
1000 - 999.999 = 998.999 + 1 = 999000
1.000.000 - 999.999.999 = 998.999.999 + 1 = 999.000.000 * 2

*/