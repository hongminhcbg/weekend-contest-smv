package main

import "fmt"

func format2DArray(in [][]int) [][]int {
	row := make(map[int]struct{})
	col := make(map[int]struct{})

	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in[0]); j++ {
			if in[i][j] == 1 {
				row[i] = struct{}{}
				col[j] = struct{}{}
			}
		}
	}

	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in[0]); j++ {
			if _, ok := row[i]; ok {
				in[i][j] = 1
				continue
			}

			if _, ok := col[j]; ok {
				in[i][j] = 1
				continue
			}
		}
	}

	return in
}

func main() {
	fmt.Println(format2DArray([][]int{{1,2,3}, {3,3,3}, {0,0,1}}))
	fmt.Println(format2DArray([][]int{{1,1,1}, {3,3,3}, {0,0,1}}))
	fmt.Println(format2DArray([][]int{{2,2,2}, {2,2,2}, {2,2,2}}))
	fmt.Println(format2DArray([][]int{{1,1,1}, {1,1,1}, {1,1,1}}))
}
