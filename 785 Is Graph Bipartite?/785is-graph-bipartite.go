func isBipartite(graph [][]int) bool {
    color:=make([]int, len(graph))
	for idx, _ := range graph {
        if color[idx]==0 && !canBeColored(graph, color, 1, idx) {
            return false
        }
    }
    return true
}

func canBeColored(graph [][]int, colors []int, wantedColor, node int) bool {
    if colors[node] != 0 {
        return colors[node] == wantedColor
    }
    colors[node] = wantedColor
	for _, neighbour := range graph[node] {
        if !canBeColored(graph, colors, -wantedColor, neighbour) {
			return false
		}
	}

    return true
}