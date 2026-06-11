func assignEdgeWeights(e [][]int) int {
    n := 0
    for k := range e {
        n = max(n, e[k][0], e[k][1])
    }
    al := make([][]int, n + 1) // adjacency list
    for k := range e {
        al[e[k][0]] = append(al[e[k][0]], e[k][1])
        al[e[k][1]] = append(al[e[k][1]], e[k][0])
    }
    md := 0
    q := []int{1}
    was := make([]bool, n + 1) // visited nodes
    was[1] = true
    for len(q) > 0 { // start bfs
        nq := []int{}
        for k := range q {
            was[q[k]] = true
            for i := range al[q[k]] {
                if !was[al[q[k]][i]] {
                    nq = append(nq, al[q[k]][i])
                }
            }
        }
        q = nq
        if len(q) > 0 { md ++ } // increase depth if queue still contains nodes
    }
    return pow_mod(md-1, 1000000007)
}

func pow_mod(exp, mod int) int { // number^2 % mod
    if exp == 0 { return 1 }
    res := 2
	for exp > 1 {
		res = (2 * res) % mod
		exp --
	}
	return res
}