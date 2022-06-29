package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func sort_matrix(mat [][]int, row_to_sort int) [][]int {
	sorted_row := mat[row_to_sort]

	li := make([][]int, len(sorted_row))
	for k, v := range sorted_row {
		kV := []int{v, k}
		li[k] = kV
	}
	sort.Slice(li, func(i, j int) bool {
		if li[i][0] == li[j][0] {
			return li[i][1] < li[j][1]
		} else {
			return li[i][0] < li[j][0]
		}
	})

	sort_index := make([]int, len(li))
	sort_lst := make([]int, len(li))
	for i, val := range li {
		sort_lst[i] = val[0]
		sort_index[i] = val[1]
	}

	new_mt_flip := make([][]int, len(mat))
	for idx, vl := range mat {
		if idx == row_to_sort {
			new_mt_flip[idx] = sort_lst
			continue
		}

		s_sub := make([]int, len(sort_index))
		for s_i, s_v := range sort_index {
			s_sub[s_i] = vl[s_v]
		}
		new_mt_flip[idx] = s_sub
	}

	return new_mt_flip
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		in.ReadString('\n')

		// get rows and cols
		var n, m int
		fmt.Fscan(in, &n, &m)

		// create and fill matrix
		var nextNum int
		matrix := make([][]int, n)
		for mi := range matrix {
			for col := 0; col < m; col++ {
				fmt.Fscan(in, &nextNum)
				matrix[mi] = append(matrix[mi], nextNum)
			}
		}

		// create flip matrix
		mt_flip := make([][]int, m)
		for j := range mt_flip {
			mt_flip[j] = make([]int, n)
		}

		// fill flip matrix
		for nextRow := 0; nextRow < m; nextRow++ {
			for nextCol := 0; nextCol < n; nextCol++ {
				mt_flip[nextRow][nextCol] = matrix[nextCol][nextRow]
			}
		}

		// clicks
		var clicks int
		fmt.Fscan(in, &clicks)

		// clicks on row and sort flip matrix
		for cli := 0; cli < clicks; cli++ {
			var nextClick int
			fmt.Fscan(in, &nextClick)
			mt_flip = sort_matrix(mt_flip, nextClick-1)

		}

		// print clicked matrix
		for fli := 0; fli < n; fli++ {
			for flj := range mt_flip {
				fmt.Fprintf(out, "%v ", mt_flip[flj][fli])
			}
			fmt.Fprint(out, "\n")
		}
	}
}
