package dp

import (
	"fmt"
	"math"
)

func longestCommonSubSequence(seq1, seq2 []int) int {
	rows := len(seq1) + 1
	cols := len(seq2) + 1
	matrix := initializeLcsMatrix(rows, cols)
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			in := matrix[i-1][j]
			del := matrix[i][j-1]
			mm := matrix[i-1][j-1]
			if seq1[i-1] == seq2[j-1] {
				mm += 1
			}
			matrix[i][j] = max(in, del, mm)
		}
	}
	fmt.Println(matrix)
	return matrix[rows-1][cols-1]
}

func max(nums ...int) int {
	max := math.MinInt
	for i := 0; i < len(nums); i++ {
		if max < nums[i]{
			max = nums[i]
		}
	}
	return max
}

func initializeLcsMatrix(rows, cols int) [][]int {
	mat := make([][]int, rows)
	for i := 0; i < rows; i++ {
		mat[i] = make([]int, cols)
	}
	return mat
}
