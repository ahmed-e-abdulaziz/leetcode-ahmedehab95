func lastStoneWeight(stones []int) int {
    h := &IntHeap{}
    for _, s := range stones {
        heap.Push(h, s)
    }

    for len(*h) > 1 {
        first := heap.Pop(h).(int)
        second := heap.Pop(h).(int)

        if first == second {
            continue
        }
        heap.Push(h, first - second)
    }

    if len(*h) == 0 {
        return 0
    }
    
    return (*h)[0]
}

// An IntHeap is a max-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}