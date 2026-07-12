func arrayRankTransform(arr []int) []int {
    temp := make([]int, len(arr))
    copy(temp, arr) 
    sort.Ints(temp)

    valToRank := make(map[int]int)
    rank := 1
    for _, val := range temp {
        if _, exist := valToRank[val]; !exist{
            valToRank[val] = rank
            rank++
        }
    }
    for i:=0;i<len(arr);i++{
        arr[i] = valToRank[arr[i]]
    }

    return arr
}