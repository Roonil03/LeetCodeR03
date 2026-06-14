func getLength(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}
	dict := make(map[int]int)
	id := 0
	for _, x := range nums {
		if _, ok := dict[x]; !ok {
			dict[x] = id
			id++
		}
	}
	arr := make([]int, n)
	for i, v := range nums {
		arr[i] = dict[v]
	}
	res := 1
	freq, c := make([]int, id), make([]int, n+1)
	for i := range nums {
		if n-i <= res {
			break
		}
		f, df := 0, 0
		for j := i; j < n; j++ {
			x := arr[j]
			w := freq[x]
			if w > 0 {
				c[w]--
				if c[w] == 0 {
					df--
				}
			}
			w++
			freq[x] = w
			if c[w] == 0 {
				df++
			}
			c[w]++
			f = max(f, w)
			if j-i+1 > res {
				if (df == 1 && c[f] == 1) || (df == 2 && f&1 == 0 && c[f>>1] > 0) {
					res = j - i + 1
				}
			}
		}
		for j := i; j < n; j++ {
			freq[arr[j]] = 0
		}
		clear(c[:f+1])
	}
	return res
}