func countBits(n int) []int {
    answer := make([]int, n+1)

    answer[0] = 0
    for i:=1;i<=n;i++{
        answer[i] = answer[i/2]
        if i%2 == 1 {
            answer[i]++
        }
    }

    return answer
}

/*
1. 1
2. 10
3. 11
4. 100
5. 101
6. 110
7. 111
8. 1000
9. 1001
10. 1010
11. 1011
12. 1100
13. 1101
14. 1110
15. 1111




*/