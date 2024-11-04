func generate(numRows int) [][]int {
    switch numRows {
        case 1:
          return [][]int{[]int{1}}
        case 2:
          return [][]int{[]int{1}, []int{1,1}}
        case 3:
          return [][]int{[]int{1}, []int{1,1}, []int{1,2,1}}
    }
    res := generate(numRows-1)
    currentRes := []int{1}
    for i := 1; i < numRows-1; i++ {
        currentRes = append(currentRes, res[numRows-2][i-1] + res[numRows-2][i])
    }
    currentRes = append(currentRes, 1)
    return append(res, currentRes)
}