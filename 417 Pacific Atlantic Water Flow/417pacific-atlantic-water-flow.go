func pacificAtlantic(heights [][]int) [][]int {
	maxI := len(heights)
	maxJ := len(heights[0])
	res := [][]int{}
	pacific, atlantic := make([][]bool, maxI), make([][]bool, maxI)
	for x := 0; x < maxI; x++ {
		pacific[x] = make([]bool, maxJ)
		atlantic[x] = make([]bool, maxJ)
	}

	for i := 0; i < maxI; i++ {
        dfs(heights, i, 0, pacific)
		dfs(heights, i, maxJ-1, atlantic)
	}

	for j := 0; j < maxJ; j++ {
        dfs(heights, 0, j, pacific)
		dfs(heights, maxI-1, j, atlantic)
	}			

	for i := 0; i < maxI; i++ {
		for j := 0; j < maxJ; j++ {
			if pacific[i][j] && atlantic[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}

func dfs(graph [][]int, i, j int, visited [][]bool) {
	if visited[i][j] {
		return
	}
	visited[i][j] = true
	for _, off := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		if canVisit(i, j, off, graph) {
			dfs(graph, i+off[0], j+off[1], visited)
		}
	}
}

func canVisit(i, j int, off []int, graph [][]int) bool {
	return i+off[0] > -1 &&
		i+off[0] < len(graph) &&
		j+off[1] > -1 &&
		j+off[1] < len(graph[0]) &&
		graph[i+off[0]][j+off[1]] >= graph[i][j]
}
