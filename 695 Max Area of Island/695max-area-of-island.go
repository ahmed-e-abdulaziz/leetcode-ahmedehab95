func maxAreaOfIsland(grid [][]int) int {
    max := 0
    for y, row := range grid {
        for x, cell := range row {
            if cell == 1 {
                max = getMax(max, dfs(grid, y, x))
            }
        }
    }
    return max
}

func dfs(grid [][]int, row, col int) int {
    if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) || grid[row][col] == 0 {
        return 0
    }
    grid[row][col] = 0
    count := 1
    count+=dfs(grid, row, col+1)
    count+=dfs(grid, row+1, col)
    count+=dfs(grid, row, col-1)
    count+=dfs(grid, row-1, col)
    return count
}

func getMax(x,y int) int {
    if x > y {
        return x
    }
    return y
}