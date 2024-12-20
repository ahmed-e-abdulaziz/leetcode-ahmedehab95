type MinStack struct {
    Vals  []int
    Mins  []int
}


func Constructor() MinStack {
    return MinStack{}
}


func (ms *MinStack) Push(val int)  {
    ms.Vals = append(ms.Vals, val)
    if len(ms.Mins) == 0 || val <= ms.GetMin() {
        ms.Mins = append(ms.Mins, val)
    } else {
        ms.Mins = append(ms.Mins, ms.GetMin())
    }
}


func (ms *MinStack) Pop()  {
    ms.Vals = ms.Vals[:len(ms.Vals)-1]
    ms.Mins = ms.Mins[:len(ms.Mins)-1]
}


func (ms *MinStack) Top() int {
    return ms.Vals[len(ms.Vals)-1]
}


func (ms *MinStack) GetMin() int {
    return ms.Mins[len(ms.Mins)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */