func bestClosingTime(customers string) int {
    length := len(customers)
    y, n := make([]int, length+1), make([]int, length+1)
    for i:=0;i<length;i++{
        if i > 0 {
            n[i] = n[i-1]   
        }
        if customers[i] == 'N' {
            n[i]++
        }
    }
    n[length] = n[length-1]

    for i:=length-1;i>=0;i--{
        if i != length-1 {
            y[i] = y[i+1]
        }
        if customers[i] == 'Y' {
            y[i]++
        }
    }
    y[length] = 0

    min, ans := 100005, -1
    for i:=0;i<=length;i++{
        fmt.Printf("%d : Y %d N %d\n", i, y[i], n[i])
        x := y[i]
        if i > 0 {x += n[i-1]}
        if x < min {
            min = x
            ans = i
        }
    }

    return ans
}