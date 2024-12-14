func evalRPN(tokens []string) int {
    s:=make(stack, 0)
    for _, t := range tokens {
        if isOp(t) {
            var x,y int
            s, y = s.Pop()
            s, x = s.Pop()
            res:=compute(t, x, y)
            s=s.Push(res)
        } else {
            num, _ := strconv.Atoi(t)
            s=s.Push(num)
        }
    }
    _, res:=s.Pop()
    return res
}

func isOp(c string) bool {
    return c == "*" || c == "/" || c == "+" || c == "-"
}

func compute(op string, x,y int) int {
    switch op {
        case "*":
            return x*y
        case "/":
            return x/y
        case "+":
            return x+y
        case "-":
            return x-y
    }
    return x+y
}

type stack []int

func (s stack) Push(v int) stack {
    return append(s, v)
}

func (s stack) Pop() (stack, int) {
    l := len(s)
    return  s[:l-1], s[l-1]
}
