type ProductOfNumbers struct {
    prefix []int
}

func Constructor() ProductOfNumbers {
    return ProductOfNumbers{
        prefix: []int{1},
    }
}

func (this *ProductOfNumbers) Add(num int) {
    if num == 0 {
        this.prefix = []int{1}
        return
    }

    n := len(this.prefix)
    this.prefix = append(this.prefix, this.prefix[n-1]*num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
    n := len(this.prefix)

    if k >= n {
        return 0
    }

    return this.prefix[n-1] / this.prefix[n-1-k]
}