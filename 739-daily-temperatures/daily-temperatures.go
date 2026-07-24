type tempPair struct {
    temp int
    idx int
}

func dailyTemperatures(temperatures []int) []int {
    stack := make([]tempPair, 0)
    answer := make([]int, len(temperatures))

    for i, t := range temperatures {
        for len(stack) > 0 && stack[len(stack)-1].temp < t {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            answer[top.idx] = i-top.idx
        }
        if len(stack) == 0 || stack[len(stack)-1].temp >= t {
            stack = append(stack, tempPair{
                temp: t,
                idx: i,
            })
        }
    }

    return answer
}

/*
73,74,75,71,69,72,76,73
[(75, 2), (72, 5) () ]
if curr > top {
    pop
}

[]

*/