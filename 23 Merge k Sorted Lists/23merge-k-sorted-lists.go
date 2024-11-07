/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    pq := NewMinHeap()
    heap.Init(pq)
    for _, v := range lists {
        for v != nil {
            heap.Push(pq, v.Val)
            v = v.Next
        }
    }
    
    sol := new(ListNode)
    head := sol

    for pq.Len() > 0 {
        head.Next = NewListNode(heap.Pop(pq).(int))
        head = head.Next
    }

    return sol.Next
}

func NewListNode(val int) *ListNode {
    return &ListNode{
        Val: val,
        Next: nil,
    }
}

type MinHeap struct {
    items []int
}

func NewMinHeap() *MinHeap{
    return &MinHeap{
        items: make([]int, 0),
    }
}

func (h MinHeap) Len() int { return len(h.items) }
func (h MinHeap) Less(i, j int) bool { return h.items[i] < h.items[j] }
func (h MinHeap) Swap(i, j int) { h.items[i], h.items[j] = h.items[j], h.items[i] }
func (h *MinHeap) Push(x interface{}) { h.items = append(h.items, x.(int)) }
func (h *MinHeap) Pop() interface{} {
    min := h.items[h.Len()-1]
    h.items = h.items[:h.Len()-1]
    return min
}