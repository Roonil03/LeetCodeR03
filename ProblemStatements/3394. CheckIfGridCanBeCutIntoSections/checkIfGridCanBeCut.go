package main

import "sort"

func checkValidCuts(n int, a [][]int) bool {
	n = len(a)
	x := make([][2]int, n)
	y := make([][2]int, n)
	for i := 0; i < n; i++ {
		x[i] = [2]int{a[i][0], a[i][2]}
		y[i] = [2]int{a[i][1], a[i][3]}
	}
	f := func(arr [][2]int) bool {
		sort.Slice(arr, func(i, j int) bool {
			return arr[i][0] < arr[j][0] || (arr[i][0] == arr[j][0] && arr[i][1] < arr[j][1])
		})
		cnt := 0
		R := -1
		for _, pair := range arr {
			l, r := pair[0], pair[1]
			if l >= R {
				cnt++
				R = r
			} else {
				if r > R {
					R = r
				}
			}
		}
		return cnt >= 3
	}
	return f(x) || f(y)
}
