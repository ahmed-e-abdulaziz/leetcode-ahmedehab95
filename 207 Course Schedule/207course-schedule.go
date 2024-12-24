
func canFinish(numCourses int, edges [][]int) bool {
	courses := make([][]int, numCourses)
	for _, edge := range edges {
		if courses[edge[0]] == nil {
			courses[edge[0]] = make([]int, 0)
		}
		courses[edge[0]] = append(courses[edge[0]], edge[1])
	}

	visitMatrix := make([]byte, numCourses)

	for i := 0; i < numCourses; i++ {
		if findCycle(courses, visitMatrix, i) {
			return false
		}
	}
	return true
}
func findCycle(courses [][]int, visits []byte, start int) bool {
    if visits[start] == 1 {
        return true
    }
    if visits[start] == 2 {
        return false
    }
	visits[start] = 1
	for _, next := range courses[start] {
		if findCycle(courses, visits, next) {
			return true
		}
	}
	visits[start] = 2
	return false
}
