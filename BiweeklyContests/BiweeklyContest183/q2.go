func minOperations(nums []int, k int) int {
	a, b := make([]int, k), make([]int, k)
	for i, v := range nums {
		r := (v%k + k) % k
		for j := range k {
			d := (r - j + k) % k
			if d > k-d {
				d = k - d
			}
			if i%2 == 0 {
				a[j] += d
			} else {
				b[j] += d
			}
		}
	}
	res := int(1e9)
	for i := range k {
		for j := range k {
			if i != j && a[i]+b[j] < res {
				res = a[i] + b[j]
			}
		}
	}
	return res
}