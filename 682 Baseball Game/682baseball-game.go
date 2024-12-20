func calPoints(operations []string) int {
	s := make(Stack, 0)
	for _, op := range operations {
		switch op[0] {
		case 'C':
			s.Pop()
		case 'D':
			num := s.Peak()
			s.Put(num * 2)
		case '+':
			x := s.Peak()
			y := s.PeakBefore()
			s.Put(x + y)
		default:
			num, _ := strconv.Atoi(op)
			s.Put(num)
		}
	}
	res := 0
	for len(s) > 0 {
		res += s.Pop()
	}
	return res
}

type Stack []int

func (s *Stack) Put(num int) {
	*s = append(*s, num)
}

func (s *Stack) Pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *Stack) Peak() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) PeakBefore() int {
	return (*s)[len(*s)-2]
}