func dailyTemperatures(temps []int) []int {
	res := make([]int, len(temps))
	s := make(Stack, 0)
	for i, t := range temps {
		if len(s) > 0 {
			prev := s.Peak()
			for len(s) > 0 && temps[prev] < t {
				res[prev] = i - prev
				s.Pop()
				if len(s) > 0 {
					prev = s.Peak()
				}
			}
		}
		s.Push(i)
	}
	return res
}

// Stack implementation
type Stack []int

func (s *Stack) Push(c int) {
	*s = append(*s, c)
}

func (s *Stack) Pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *Stack) Peak() int {
	return (*s)[len(*s)-1]
}