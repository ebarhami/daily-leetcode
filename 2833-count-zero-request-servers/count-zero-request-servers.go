type q struct {
    query int
    idx int
}
func countServers(n int, logs [][]int, x int, queries []int) []int {
    answer := make([]int, len(queries))
    queryIdx := make([]q, 0)
    for i, query := range queries {
        queryIdx = append(queryIdx, q{
            query: query,
            idx: i,
        })
    }

    sort.Slice(queryIdx, func(i, j int) bool {
        return queryIdx[i].query < queryIdx[j].query
    })
    sort.Slice(logs, func(i, j int) bool {
        return logs[i][1] < logs[j][1]
    })

    freq := make(map[int]int)

	l, r := 0, 0
	active := 0

	for _, q := range queryIdx {
		// Expand right side:
		// include logs with timestamp <= query.
		for r < len(logs) && logs[r][1] <= q.query {
			server := logs[r][0]

			if freq[server] == 0 {
				active++
			}

			freq[server]++
			r++
		}

		// Shrink left side:
		// valid window is [query-x, query].
		for l < r && logs[l][1] < q.query-x {
			server := logs[l][0]

			freq[server]--

			if freq[server] == 0 {
				active--
				delete(freq, server)
			}

			l++
		}

		answer[q.idx] = n - active
	}

	return answer
}

/*
[2,1], [3,1] ,[1,2], [2,4]
*/