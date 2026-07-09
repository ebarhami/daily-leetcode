func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
    par := make(map[int]int)

    for i:=0;i<n;i++{
        if i == 0 {
            par[nums[i]] = i
        } else {
            prev, curr := nums[i-1], nums[i]
            if curr - prev <= maxDiff {
                par[curr] = par[prev]
            } else {
                par[curr] = i
            }
        }
    }

    answer := make([]bool, 0)
    for _, query := range queries {
        u, v := query[0], query[1]
        if par[nums[u]] == par[nums[v]] {
            answer = append(answer, true)
        } else {
            answer = append(answer, false)
        }
    }

    return answer
}