func smallestNumber(n int, t int) int {
    for i:=0;i<=10;i++{
        if getProductDigits(n+i) % t == 0 {
            return n+i
        }
    }
    return n
} 

func getProductDigits(n int) int {
    num := 1
    for n > 0 {
        num *= n%10
        n /= 10
    }

    return num
}