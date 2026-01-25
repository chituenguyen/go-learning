package main

import (
	"fmt"
	"sort"
)

func main() {
	// Test maximalSquare
	// matrix := [][]byte{
	// 	{'1', '0', '1', '0', '0'},
	// 	{'1', '0', '1', '1', '1'},
	// 	{'1', '1', '1', '1', '1'},
	// 	{'1', '0', '0', '1', '0'},
	// }
	// fmt.Println(maximalSquare(matrix))

	// Test combinationSum
	arr := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum(arr, target))
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	if candidates[0] > target || len(candidates) == 0 {
		return [][]int{}
	}

	var result = [][]int{}
	var path = []int{}
	var dfs func(start_index int, remain int)
	dfs = func(start_index int, remain int) {
		if remain == 0 {
			var tmp = make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
		}

		for i := start_index; i < len(candidates); i++ {
			if i > start_index && candidates[i] == candidates[i-1] {
				continue
			}
			if candidates[i] > remain {
				break
			}

			path = append(path, candidates[i])
			dfs(i+1, remain-candidates[i])

			path = path[:len(path)-1]
		}
	}

	dfs(0, target)

	return result
}
