func fizzBuzz(n int) []string {
    ans := make([]string, n)
    for i:=1;i<=n;i++{
        s := toString(i)
        if i % 3 + i % 5 == 0 {
            s = "FizzBuzz"
        } else if i % 3 == 0 {
            s = "Fizz"
        } else if i % 5 == 0 {
            s = "Buzz"
        }
        ans[i-1] = s
    }

    return ans
}

func toString(x int) string {
    s := ""
    for x > 0 {
        s = string(x%10 + '0') + s
        x = x/10
    }
    return s
}