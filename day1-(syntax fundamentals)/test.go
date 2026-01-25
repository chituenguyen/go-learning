package main

func maximalSquare(matrix [][]byte) int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m := len(matrix)
	n := len(matrix[0])

	// Build prefix sum array
	countOne := make([][]int, m)
	for i := 0; i < m; i++ {
		countOne[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			a, b, c := 0, 0, 0

			if i != 0 {
				a = countOne[i-1][j]
			}
			if j != 0 {
				b = countOne[i][j-1]
			}
			if i != 0 && j != 0 {
				c = countOne[i-1][j-1]
			}

			// Build prefix sum
			val := 0
			if matrix[i][j] == '1' {
				val = 1
			}
			countOne[i][j] = a + b - c + val
		}
	}

	maxSide := 0

	// Try each position as top-left corner
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			maxPossibleSide := min(m-i, n-j)

			c := 0
			if i != 0 && j != 0 {
				c = countOne[i-1][j-1]
			}

			// Try square sizes from largest to smallest
			for side := maxPossibleSide; side >= 1; side-- {
				lowerRightRow := i + side - 1
				lowerRightCol := j + side - 1

				totalOne := countOne[lowerRightRow][lowerRightCol]

				a := 0
				if i != 0 {
					a = countOne[i-1][lowerRightCol]
				}

				b := 0
				if j != 0 {
					b = countOne[lowerRightRow][j-1]
				}

				// Count ones in the square
				oneInSquare := totalOne - a - b + c

				// If all cells are 1, update max
				if oneInSquare == side*side {
					if side > maxSide {
						maxSide = side
					}
					break
				}
			}
		}
	}

	// Return area (side squared)
	return maxSide * maxSide
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
