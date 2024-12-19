
func searchMatrix(matrix [][]int, target int) bool {
    length := len(matrix[0])
	l, r := 0, (len(matrix)*len(matrix[0]))-1
	var mid int
	for l <= r {
		mid = (l + r) / 2
        col, row := mid/length, mid%length
		if matrix[col][row] > target {
			r = mid - 1
		} else if matrix[col][row] < target {
			l = mid + 1
		} else {
			return true
		}
	}
	return false
}
