func combine(n int, k int) [][]int {
    answer := make([][]int, 0)

    traverse(1, n, k, &answer, []int{})

    return answer
}

func traverse(idx, n,k int, answer *[][]int, take []int) {
    if len(take) == k {
        combin := make([]int, k)
        copy(combin, take)
        (*answer) = append(*answer, combin)
        return
    }

    for i:=idx;i<=n;i++{
        take = append(take, i)
        traverse(i+1, n, k, answer, take)
        take = take[:len(take)-1]
    }
}
