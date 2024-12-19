var w string

func exist(board [][]byte, word string) bool {
	w = word
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			if board[y][x] == word[0] && search(board, x, y, 0) {
				return true
			}
		}
	}
	return false
}

func search(board [][]byte, x, y, count int) bool {
	if count == len(w) {
		return true
	}
	if x < 0 || x == len(board[0]) || y < 0 || y == len(board) ||
		board[y][x] == '#' || board[y][x] != w[count] {
		return false
	}
	tmp := board[y][x]
	board[y][x] = '#'
	ans := search(board, x+1, y, count+1) ||
		search(board, x, y+1, count+1) ||
		search(board, x-1, y, count+1) ||
		search(board, x, y-1, count+1)
	board[y][x] = tmp
	return ans
}
