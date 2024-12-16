func spiralOrder(matrix [][]int) []int {
	var x, y int
    resPart := []int{}
	res := []int{}

	for true {
		resPart, x, y = right(matrix, x, y)
		res = append(res, resPart...)
		resPart, x, y = down(matrix, x, y)
		res = append(res, resPart...)
		resPart, x, y = left(matrix, x, y)
		res = append(res, resPart...)
		resPart, x, y = up(matrix, x, y)
		res = append(res, resPart...)
        if len(resPart) == 0 {
			break
		}
	}
	return res
}

func right(matrix [][]int, x, y int) ([]int, int, int) {
	res := []int{}
	for x < len(matrix[0]) && matrix[y][x] != -200 {
		res = append(res, matrix[y][x])
		matrix[y][x] = -200
		x++
	}
	return res, x - 1, y + 1
}

func down(matrix [][]int, x, y int) ([]int, int, int) {
	res := []int{}
	for y < len(matrix) && matrix[y][x] != -200 {
		res = append(res, matrix[y][x])
		matrix[y][x] = -200
		y++
	}
	return res, x - 1, y - 1
}

func left(matrix [][]int, x, y int) ([]int, int, int) {
	res := []int{}
	for x >= 0 && matrix[y][x] != -200 {
		res = append(res, matrix[y][x])
		matrix[y][x] = -200
		x--
	}
	return res, x + 1, y - 1
}

func up(matrix [][]int, x, y int) ([]int, int, int) {
	res := []int{}
	for y >= 0 && matrix[y][x] != -200 {
		res = append(res, matrix[y][x])
		matrix[y][x] = -200
		y--
	}
	return res, x + 1, y + 1
}
