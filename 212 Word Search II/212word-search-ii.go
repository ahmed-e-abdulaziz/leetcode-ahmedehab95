var directions = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func findWords(board [][]byte, words []string) []string {
	trie := new(TrieNode)
	for _, word := range words {
		if len(board)*len(board[0]) >= len(word) {
			trie.insert(word)
		}
	}
	res := make([]string, 0, len(words))
	for y, row := range board {
		for x, ch := range row {
            if node := trie.Children[ch-'a']; node != nil {
                find(board, x, y, &res, trie)
            }
			
		}
	}
	return res
}

func find(board [][]byte, x, y int, res *[]string, node *TrieNode) {
	char := board[y][x]
	node = node.Children[char-'a']
	if node == nil {
		return
	}
	if node.Word != "" {
		*res = append(*res, node.Word)
		node.Word = ""
	}

	board[y][x] = '#'
	for _, dir := range directions {
		y, x := y+dir[0], x+dir[1]
		if y >= 0 && y < len(board) && x >= 0 && x < len(board[0]) && board[y][x] != '#' {
			find(board, x, y, res, node)
		}
	}
	board[y][x] = char
}

type TrieNode struct {
	Children [26]*TrieNode
	Word     string
}

func (tn *TrieNode) insert(word string) {
	for _, ch := range word {
		idx := ch - 'a'
		if tn.Children[idx] == nil {
			tn.Children[idx] = new(TrieNode)
		}
		tn = tn.Children[idx]
	}
	tn.Word = word
}
