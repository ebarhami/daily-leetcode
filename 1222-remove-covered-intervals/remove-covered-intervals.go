func removeCoveredIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i,j int) bool {
        return (intervals[i][0] < intervals[j][0]) || 
        (intervals[i][0] == intervals[j][0] && (intervals[i][1] > intervals[j][1]))
    })

    maxEnd := -1
    answer := 0
    for _, interval := range intervals {
        if interval[1] > maxEnd {
            answer++
            maxEnd = interval[1]
        }
    }

    return answer
}

/*

if 


*/