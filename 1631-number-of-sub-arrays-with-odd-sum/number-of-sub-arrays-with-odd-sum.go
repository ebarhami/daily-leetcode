func numOfSubarrays(arr []int) int {
    even, odd := 1, 0
    prefix := make([]int, len(arr))
    mod := 1000000007
    
    for i:=0;i<len(arr);i++{
        prefix[i] = arr[i]
        if i > 0 {
            prefix[i] += prefix[i-1]
        }
    }

    for _, num := range prefix {
        if num % 2 == 0{
            even++
        } else {
            odd++
        }
    }

    return (odd * even) % mod
}

/*
3 
all odd : C(3,3) + C(3,1) = 1 + 3 = 4
there is even number ? for each of above, pair with all even groups possibility

e.g 4 even numbers: 2,4,6,8
evenSum = (4,1) + (4,2) + (4,3) + (4,4)
(3,3)* (evenSum) +

(3,1)* evenSum


1.          1 
1 1.        2
1 2 1       4
1 3 3 1     8
1 4 6 4 1   16

1.              0 
1 1.            1
1 2 1           2
1 3 3 1         4
1 4 6 4 1       8
1 5 10 10 5 1   16

1,2,3,4,5,6,7
odd = 4
even = 3
2



*/