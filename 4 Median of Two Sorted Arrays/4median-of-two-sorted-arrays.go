func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	h := &IntHeap{}
	heap.Init(h)
	var i, j, count int

	for count <= (len(nums1)+len(nums2))/2 {
		if j >= len(nums2) || (i < len(nums1) && nums1[i] < nums2[j]) {
			heap.Push(h, nums1[i])
			i++
		} else {
			heap.Push(h, nums2[j])
			j++
		}
		count++
	}
	if (len(nums1)+len(nums2))%2 == 0 {
		return float64(heap.Pop(h).(int) + heap.Pop(h).(int)) / 2.0
	}
	return float64(heap.Pop(h).(int))
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
