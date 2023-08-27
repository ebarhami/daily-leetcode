func canCross(stones []int) bool {
    valid := make(map[int]map[int]int)

    stoneMap := make(map[int]struct{})

    for _, stone := range stones {
        stoneMap[stone] = struct{}{}
    }
    
    return traverse(stoneMap, valid, 0, stones[0], stones[len(stones)-1]) == 1
}

func traverse(stones map[int]struct{}, valid map[int]map[int]int, jump int, pos int, target int) int {
    if _, ok := valid[pos]; ok {
        if valid[pos][jump] != 0 {
            return valid[pos][jump]
        }
    } else {
        valid[pos] = make(map[int]int)
    }
    if pos == target {
        return 1
    }
    valid[pos][jump] = -1
    for i:=jump-1;i<=jump+1;i++{
        if i <= 0 {continue}
        if _, exist := stones[pos+i]; exist {
            valid[pos][jump] = max(valid[pos][jump], traverse(stones, valid, i, pos+i, target))
        }
    }

    return valid[pos][jump]
}

func max(a, b int) int {
    if a > b {return a}
    return b
}