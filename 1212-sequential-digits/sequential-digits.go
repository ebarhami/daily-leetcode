func sequentialDigits(low int, high int) []int {
    s := "123456789"
    answer := make([]int, 0)
    n := len(s)
    
    for i:=0;i<n;i++{
        for j:=i+1;j<=n;j++{
            num := toInt(s[i:j])
            if num < low || num > high {
                continue
            }
            answer = append(answer, num)
        }
    }

    sort.Ints(answer)
    return answer
}


func toInt(s string) int {
    num := 0
    for i:=0;i<len(s);i++{
        num *= 10
        num += int(s[i] - '0')
    }

    return num
}