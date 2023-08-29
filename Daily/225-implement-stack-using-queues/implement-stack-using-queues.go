type MyStack struct {
    qIn  []int
    qOut []int
}


func Constructor() MyStack {
    return MyStack{
        qIn: make([]int, 0),
        qOut: make([]int, 0),
    }
}


func (this *MyStack) Push(x int)  {
    for len(this.qIn) > 0 {
        fr := this.qIn[0]
        this.qIn = this.qIn[1:]

        this.qOut = append(this.qOut, fr)
    }

    this.qIn = append(this.qIn, x)
    this.qIn, this.qOut = this.qOut, this.qIn
}


func (this *MyStack) Pop() int {
    if len(this.qOut) == 0 {
        return 0
    }
    fr := this.qOut[0]
    this.qOut = this.qOut[1:]
    if len(this.qOut) == 0 {
        this.qIn, this.qOut = this.qOut, this.qIn
    }
    return fr
}


func (this *MyStack) Top() int {
    if len(this.qOut) == 0 {
        return 0
    }
    return this.qOut[0]
}


func (this *MyStack) Empty() bool {
    return len(this.qOut) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();

q2 = [4,3,1]
q1 = []
 
 */