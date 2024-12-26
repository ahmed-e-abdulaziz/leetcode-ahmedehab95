func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) == 0 || k <= 0 || k > len(nums) {
		return nil
	}
    d := deque{[]int{}}
	res := []int{}

    for i := 0; i < len(nums); i++ {
        if !d.isEmpty() && i-k+1 > d.getFirst() {
            d.popFirst()
        }
        if !d.isEmpty() && nums[d.getFirst()] < nums[i] {
            d.empty()
        }
        for !d.isEmpty() && nums[d.getLast()] < nums[i] {
                d.popLast()
        }
        d.push(i)
        if i >= k-1 {
            res = append(res, nums[d.getFirst()])
        }
	}
	return res
}

type deque struct {
	indexes []int
}

func (d *deque) push(i int) {
	d.indexes = append(d.indexes, i)
}

func (d *deque) getFirst() int {
	return d.indexes[0]
}

func (d *deque) popFirst() {
	d.indexes = d.indexes[1:]
}

func (d *deque) getLast() int {
	return d.indexes[len(d.indexes)-1]
}

func (d *deque) popLast() {
	d.indexes = d.indexes[:len(d.indexes)-1]
}

func (d *deque) isEmpty() bool {
	return 0 == len(d.indexes)
}

func (d *deque) empty() {
    d.indexes = []int{}
}
