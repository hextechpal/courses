package dp

import "math"

func editDistance(str1, str2 string) int {
	rows := len(str1) + 1
	cols := len(str2) + 1
	matrix := initializeMatrix(rows, cols)

	for i:=1; i<rows; i++{
		for j := 1; j < cols; j++ {
			in := matrix[i-1][j] + 1
			del := matrix[i][j-1] + 1
			mm := matrix[i-1][j-1]
			if str1[i-1] != str2[j-1]{
				mm += 1
			}
			matrix[i][j] = min(in, del, mm)
		}
	}
	return matrix[rows-1][cols-1]
}

func min(nums ...int) int {
	min := math.MaxInt
	for i := 0; i < len(nums); i++ {
		if min > nums[i]{
			min = nums[i]
		}
	}
	return min
}

func initializeMatrix(rows, cols int) [][]int {
	mat := make([][]int, rows)
	for i := 0; i < rows; i++ {
		mat[i] = make([]int, cols)
	}

	// Initialise first row to indexes
	for i := 0; i < cols; i++ {
		mat[0][i] = i
	}

	// Initialise first column to indexes
	for i := 0; i < rows; i++ {
		mat[i][0] = i
	}
	return mat
}
