func maximumPopulation(logs [][]int) int {
    max, n := 0, len(logs)
    answer := 0

    for i:=0;i<n;i++{
        year := logs[i][0]
        count := 0
        for j:=0;j<n;j++{
            b, d := logs[j][0], logs[j][1]
            if b <= year && d > year {
                count++
            }
        }
        if count > max || (count == max && year < answer) {
            max = count
            answer = year
        }
    }
    return answer
}