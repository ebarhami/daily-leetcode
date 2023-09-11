func groupThePeople(groupSizes []int) [][]int {
    groups := make(map[int][][]int)

    n := len(groupSizes)
    
    for i:=0;i<n;i++{
        size := groupSizes[i]
        if _, ok := groups[size]; !ok {
            groups[size] = make([][]int, 1)
        }
        j := 0
        for len(groups[size][j]) == size {
            j++
            if j == len(groups[size]) {
                groups[size] = append(groups[size], []int{})
            }
        }
        groups[size][j] = append(groups[size][j], i)
    }

    answer := make([][]int, 0)

    for _, arr := range groups {
        answer = append(answer, arr...)
    }

    return answer
}