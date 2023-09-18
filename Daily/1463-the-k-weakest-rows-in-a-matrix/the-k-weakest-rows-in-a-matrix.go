func kWeakestRows(mat [][]int, k int) []int {
    scoreByIdx := make([][]int, len(mat))
    for i:=0;i<len(mat);i++{
        score := count(mat[i])
        scoreByIdx[i] = []int{i, score}
    }

    sort.SliceStable(scoreByIdx, func(i,j int) bool {
        return scoreByIdx[i][1] < scoreByIdx[j][1]
    })

    answer := make([]int, k)
    for i:=0;i<k;i++{
        answer[i] = scoreByIdx[i][0]
    }
    
    return answer
}

func count(arr []int) int {
    l, r := 0, len(arr)-1
    size := -1

    for l<=r {
        mid := (l+r)/2

        if arr[mid] == 1 {
            size = mid
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return size + 1
}

/*
l, r = 3, 4



*/