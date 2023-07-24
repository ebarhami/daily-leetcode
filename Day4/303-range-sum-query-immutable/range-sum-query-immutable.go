type NumArray struct {
    arr []int
}


func Constructor(nums []int) NumArray {
    for i:=1;i<len(nums);i++{
        nums[i] = nums[i]+nums[i-1]
    }
    return NumArray{
        arr: nums,
    }
}


func (this *NumArray) SumRange(left int, right int) int {
    if left == 0 {
        return this.arr[right]
    }
    return this.arr[right] - this.arr[left-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */