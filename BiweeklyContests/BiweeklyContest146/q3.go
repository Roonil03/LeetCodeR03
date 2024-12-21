/* Could not solve it entirely during the time of the contest,
 * it has been solved post the time limit of the contest.
 */

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

/* 	The checkValidCuts function checks whether there are at least three non-overlapping intervals in a set of 2D
	intervals represented by the array a. The function first separates the intervals into two sets: one for the x-coordinates
	(x[i][0] to x[i][2]) and one for the y-coordinates (y[i][1] to y[i][3]). It then defines a helper function f that
	takes an array of intervals, sorts them by their starting points, and iteratively checks for non-overlapping intervals
	by tracking the rightmost boundary of the previous selected interval (R). If the current interval's left boundary is
   	greater than or equal to R, it is counted as a valid non-overlapping interval, and the right boundary (R) is updated.
   	The function returns true if there are at least 3 such valid intervals, and the main function returns true if either
   	the x-coordinates or y-coordinates satisfy this condition. The algorithm primarily uses arrays for interval storage,
   	sorting, and a greedy approach to check for valid non-overlapping intervals.
*/
