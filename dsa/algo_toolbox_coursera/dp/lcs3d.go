package dp

func lcs3d(seq1, seq2, seq3 []int) int {
	rows := len(seq1) + 1
	cols := len(seq2) + 1
	depth := len(seq3) + 1
	matrix := initialize3dMatrix(rows, cols, depth)

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			for k := 1; k < depth; k++ {
				mm := matrix[i-1][j-1][k-1]
				if seq1[i-1] == seq2[j-1] && seq2[j-1] == seq3[k-1] {
					mm = mm + 1
				}
				matrix[i][j][k] = max(mm, matrix[i-1][j][k], matrix[i][j-1][k], matrix[i][j][k-1])
			}
		}
	}
	return matrix[rows-1][cols-1][depth-1]
}

func initialize3dMatrix(rows int, cols int, depth int) [][][]int {
	matrix := make([][][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([][]int, cols)
		for j := range matrix[i] {
			matrix[i][j] = make([]int, depth)
		}
	}
	return matrix
}
