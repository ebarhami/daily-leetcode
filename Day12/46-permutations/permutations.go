func permute(nums []int) [][]int {
    answer := make([][]int, 0)
    visited := make(map[int]bool)

    traverse(&answer, visited, nums, []int{})

    return answer
}

func traverse(answer *[][]int, visited map[int]bool, nums []int, bucket []int) {
    if len(bucket) == len(nums) {
        permutation := make([]int, len(bucket))
        copy(permutation, bucket)
        (*answer) = append(*answer, permutation)
        return
    }

    for i:=0;i<len(nums);i++{
        if visited[i] == false {
            visited[i] = true
            bucket = append(bucket, nums[i])
            traverse(answer, visited, nums, bucket)
            visited[i] = false
            bucket = bucket[:len(bucket)-1]
        }
    }
}
