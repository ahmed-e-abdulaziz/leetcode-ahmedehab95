func minCostConnectPoints(points [][]int) int {
	h := &NodeHeap{{0, 0}}
	seen := make([]bool, len(points))
	cost, seenCount := 0, 0
	for len(points) > seenCount {
	    prevNode := heap.Pop(h).(*Node)
		if seen[prevNode.i] {
			continue
		}
		cost += prevNode.dist
		prevPoint := points[prevNode.i]
        seen[prevNode.i] = true
        seenCount++
        for i := 0; i < len(points); i++ {
			if !seen[i] {
				heap.Push(h, &Node{i, calcDist(prevPoint[0], prevPoint[1], points[i][0], points[i][1])})
			}
		}
	}
	return cost
}

func calcDist(xi, yi, xj, yj int) int {
	return int(math.Abs(float64(xi-xj))) + int(math.Abs(float64(yi-yj)))
}

type Node struct {
	i    int // index
	dist int // manhattan distance |xi - xj| + |yi - yj|
}

type NodeHeap []*Node

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Node))
}

func (h *NodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
