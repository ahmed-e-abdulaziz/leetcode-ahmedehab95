func rotate(matrix [][]int) {
    // Step 1: Transpose the matrix
    for i := 0; i < len(matrix); i++ {
        for j := i+1; j < len(matrix[0]); j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }

    // Step 2: Reverse each row of the transposed matrix
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[0])/2; j++ {
            matrix[i][j], matrix[i][len(matrix[0])-j-1] = matrix[i][len(matrix[0])-j-1], matrix[i][j]
        }
    }
}