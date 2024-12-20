func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	pairs := map[byte]byte{'{': '}', '(': ')', '[': ']'}
	stack := make(Stack, 0)
	for i := 0; i < len(s); i++ {
		if _, ok := pairs[s[i]]; ok {
			stack.Push(s[i])
		} else if len(stack) == 0 || pairs[stack.Pop()] != s[i] {
			return false
		}

	}
	return len(stack) == 0
}

type Stack []byte

func (s *Stack) Push(c byte) {
	*s = append(*s, c)
}

func (s *Stack) Pop() byte {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}
