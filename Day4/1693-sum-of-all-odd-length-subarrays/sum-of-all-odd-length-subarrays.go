func sumOddLengthSubarrays(arr []int) int {
    answer := 0
    n := len(arr)

    for i:=0;i<n;i++{
        tot := 0
        for j:=i;j<n;j++{
            tot += arr[j]
            if (j-i+1) % 2 == 1 {
                answer += tot
            }
        }
    }

    return answer
}

/*

4 5 6 7 6 5 4

*/