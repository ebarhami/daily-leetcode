type Item struct {
    word string
    dist int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
    neighbor := make(map[string][]string)

    wordList = append(wordList, beginWord)
    for i:=0;i<len(wordList);i++ {
        for j:=i+1;j<len(wordList);j++ {
            x := wordList[i]
            y := wordList[j]
            dist := distance(x, y)
            if dist != 1 {
                continue
            }
            if _, exist := neighbor[x]; !exist {
                neighbor[x] = make([]string, 0)
            }
            if _, exist := neighbor[y]; !exist {
                neighbor[y] = make([]string, 0)
            }
            neighbor[x] = append(neighbor[x], y)
            neighbor[y] = append(neighbor[y], x)
        }
    }

    bestDistance := make(map[string]int)
    q := make([]Item, 0)
    q = append(q, Item{
        word: beginWord,
        dist: 1,
    })
    bestDistance[beginWord] = 0
    for len(q) != 0 {
        top := q[0]
        q = q[1:]
        if top.word == endWord {
            return top.dist
        }
        currDis := top.dist + 1
        for _, neigh := range neighbor[top.word] {
            if best, exist := bestDistance[neigh]; !exist || (exist && currDis < best) {
                q = append(q, Item{
                    word: neigh,
                    dist: currDis,
                })
                bestDistance[neigh] = currDis
            }
        }
    }
    return 0
}

func distance(s1, s2 string) int {
    if len(s1) != len(s2) {
        return 999
    }

    answer := 0
    for i:=0;i<len(s1);i++{
        if s1[i] != s2[i] {
            answer++
        }
    }

    return answer
}