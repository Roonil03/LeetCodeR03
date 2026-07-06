func divisibleGame(nums []int) int {
	n := len(nums)
	mod := int64(1e9 + 7)
	pre := make([]int64, n+1)
	for i, v := range nums {
		pre[i+1] = pre[i] + int64(v)
	}
	pid := make(map[int][]int)
	mem := make(map[int][]int)
	for i, v := range nums {
		if v == 1 {
			continue
		}
		if _, ok := mem[v]; !ok {
			var fact []int
			temp := v
			for j := 2; j*j <= temp; j++ {
				if temp%j == 0 {
					fact = append(fact, j)
					for temp%j == 0 {
						temp /= j
					}
				}
			}
			if temp > 1 {
				fact = append(fact, temp)
			}
			mem[v] = fact
		}
		for _, p := range mem[v] {
			pid[p] = append(pid[p], i)
		}
	}
	if len(pid) == 0 {
		a := nums[0]
		for _, v := range nums {
			a = min(a, v)
		}
		res := (-int64(a)*2 + mod) % mod
		return int(res)
	}
	b, c := int64(-1e10), -1
	for i, v := range pid {
		cur := int64(0)
		temp := int64(-1e18)
		for j, id := range v {
			val := int64(nums[id])
			if j == 0 {
				cur = val
			} else {
				prev := v[j-1]
				gap := pre[id] - pre[prev+1]
				d := cur - gap + val
				cur = max(val, d)
			}
			temp = max(temp, cur)
		}
		if temp > b {
			b = temp
			c = i
		} else if temp == b && i < c {
			c = i
		}
	}
	res := (b % mod) * int64(c) % mod
	if res < 0 {
		res = (res + mod) % mod
	}
	return int(res)
}