
func orangesRotting(grid [][]int) int {
    maxY := len(grid)
    maxX := len(grid[0])
	isVisited := make([][]bool, maxY)
	for i := 0; i < maxY; i++ {
		isVisited[i] = make([]bool, maxX)
	}
	fresh := 0
	q := [][2]int{}
	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			if grid[i][j] == 1 {
				fresh++
			}
			if grid[i][j] == 2 {
				q = append(q, [2]int{i, j})
			}
		}
	}
    if fresh == 0 {
        return 0
    }
	directions := [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	minutes := -1
	for len(q) > 0 {
		minutes++
		qLen := len(q)
		for i := 0; i < qLen; i++ {
			curRot := q[0]
			q = q[1:]
            for _, off := range directions {
                y:=curRot[0]+off[0]
                x:=curRot[1]+off[1]
                if y > -1 && y < maxY && x > -1 && x < maxX && grid[y][x] == 1 {
                    fresh--
                    grid[y][x]=2
                    q = append(q, [2]int{y, x})
                }
            }
		}
	}
    if fresh == 0 {
        return minutes
    }
	return -1
}