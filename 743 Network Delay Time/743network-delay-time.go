func networkDelayTime(times [][]int, n int, k int) int {
	graph := map[int][]Dist{}

	for _, time := range times {
		u, v, w := time[0], time[1], time[2]
		graph[u] = append(graph[u], Dist{v, w})
	}

	h := &DistHeap{{k, 0}}
	arrivalTime := map[int]int{}

	for h.Len() > 0 {
		newDist := heap.Pop(h).(Dist)

		if _, ok := arrivalTime[newDist.Vertex]; ok {
			continue
		}

		arrivalTime[newDist.Vertex] = newDist.Weight

		for _, next := range graph[newDist.Vertex] {
			next.Weight += newDist.Weight
			heap.Push(h, next)
		}
	}

	if len(arrivalTime) < n {
		return -1
	}

	max := 0
	for _, w := range arrivalTime {
		if w > max {
			max = w
		}
	}

	return max
}

type Dist struct {
	Vertex int
	Weight int
}

type DistHeap []Dist

func (h DistHeap) Len() int           { return len(h) }
func (h DistHeap) Less(i, j int) bool { return h[i].Weight < h[j].Weight }
func (h DistHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *DistHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Dist))
}

func (h *DistHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
