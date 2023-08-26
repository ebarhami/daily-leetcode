func findLongestChain(pairs [][]int) int {
    sort.Slice(pairs, func(i,j int)bool {
        return pairs[i][1] < pairs[j][1]
    })

    total := 0
    curr := -2000
    for i:=0;i<len(pairs);i++{
        if pairs[i][0] > curr {
            total++
            curr = pairs[i][1]
        }
    }

    return total
}