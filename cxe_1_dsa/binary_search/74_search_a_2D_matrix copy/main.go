package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows, cols := len(matrix), len(matrix[0])
	l, r := 0, rows*cols-1

	for l <= r {
		mid := l + (r-l)/2
		curr := matrix[mid/cols][mid%cols]
		if curr == target {
			return true
		}
		if curr > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	target := 3
	result := searchMatrix(matrix, target)

	if result {
		println("Target found in the matrix.")
	} else {
		println("Target not found in the matrix.")
	}
}