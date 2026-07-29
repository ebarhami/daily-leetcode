func maxAmount(initialCurrency string, pairs1 [][]string, rates1 []float64, pairs2 [][]string, rates2 []float64) float64 {
    graph1, graph2 := make(map[string]map[string]float64), make(map[string]map[string]float64)

    for i, pair := range pairs1 {
        x, y := pair[0], pair[1]

        if _, ok := graph1[x]; !ok {
            graph1[x] = make(map[string]float64)
        }
        if _, ok := graph1[y]; !ok {
            graph1[y] = make(map[string]float64)
        }
        graph1[x][y] = rates1[i]
        graph1[y][x] = 1/rates1[i]
    }

    for i, pair := range pairs2 {
        x, y := pair[0], pair[1]

        if _, ok := graph2[x]; !ok {
            graph2[x] = make(map[string]float64)
        }
        if _, ok := graph2[y]; !ok {
            graph2[y] = make(map[string]float64)
        }
        graph2[x][y] = rates2[i]
        graph2[y][x] = 1/rates2[i]
    }

    getAllRate := func(graph map[string]map[string]float64, ref string) map[string]float64 {
        rateFromRef := make(map[string]float64)
        rateFromRef[ref] = 1.0

        var dfs func(float64, string)

        dfs = func(currRate float64, ref string) {
            for next, rate := range graph[ref] {
                newRate := currRate * rate
                if newRate > rateFromRef[next] {
                    rateFromRef[next] = newRate
                    dfs(newRate, next)
                }
            }
        }

        dfs(1.0, ref)

        return rateFromRef
    }

    day1Rate := getAllRate(graph1, initialCurrency)

    answer := 1.0
    for currency1, rate1 := range day1Rate {
        day2Rate := getAllRate(graph2, currency1)
        for currency2, rate2 := range day2Rate {
            if currency2 == initialCurrency && rate1 * rate2 > answer {
                answer = rate1 * rate2
            }
        }
    }
    
    return answer
}