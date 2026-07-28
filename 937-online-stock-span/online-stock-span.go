type Item struct {
    val int
    span int
}

type StockSpanner struct {
    stack []Item
}


func Constructor() StockSpanner {
    return StockSpanner{
        stack: make([]Item, 0),
    }
}


func (this *StockSpanner) Next(price int) int {
    currSpan := 1
    for len(this.stack) > 0 && this.stack[len(this.stack)-1].val <= price {
        currSpan += this.stack[len(this.stack)-1].span
        this.stack = this.stack[:len(this.stack)-1]
    }
    this.stack = append(this.stack, Item{
        val: price,
        span: currSpan,
    })

    return currSpan
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */