func totalNumbers(digits []int) int {
    exist := make(map[int]bool)
    n := len(digits)
    answer := 0
    for i:=0;i<n;i++{
        for j:=0;j<n;j++{
            for k:=0;k<n;k++{
                if i == j || i == k || k == j {
                    continue
                }
                num := digits[i] * 100 + digits[j] * 10 + digits[k] 
                if ok, _ := exist[num]; !ok && num >= 100 && num % 2 == 0 {
                    answer++
                }
                exist[num] = true
            }
        }
   }

   return answer
}

/*
[0,2,2]
last = 3 -> 2 (unique even)
first = 2 -> 1 (unique total non zero)
mid = 3 -> 1 (all - 2)

*/