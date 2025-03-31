const mod = 1e9 + 7
const mx = 100_000
const inf = int(1e18)

func maximumScore(nums []int, k int) int {
	// n := len(nums)
	// maxNum := 0
	// for _, num := range nums {
	// 	if num > maxNum {
	// 		maxNum = num
	// 	}
	// }
	// primeScores := make([]int, maxNum+1)
	// for i := 2; i <= maxNum; i++ {
	// 	if primeScores[i] == 0 {
	// 		for j := i; j <= maxNum; j += i {
	// 			primeScores[j]++
	// 		}
	// 	}
	// }
	// nextGreater := make([]int, n)
	// prevGreater := make([]int, n)
	// stack := []int{}
	// for i := 0; i < n; i++ {
	// 	for len(stack) > 0 && primeScores[nums[stack[len(stack)-1]]] < primeScores[nums[i]] {
	// 		nextGreater[stack[len(stack)-1]] = i
	// 		stack = stack[:len(stack)-1]
	// 	}
	// 	stack = append(stack, i)
	// }
	// for len(stack) > 0 {
	// 	nextGreater[stack[len(stack)-1]] = n
	// 	stack = stack[:len(stack)-1]
	// }
	// for i := n - 1; i >= 0; i-- {
	// 	for len(stack) > 0 && primeScores[nums[stack[len(stack)-1]]] <= primeScores[nums[i]] {
	// 		prevGreater[stack[len(stack)-1]] = i
	// 		stack = stack[:len(stack)-1]
	// 	}
	// 	stack = append(stack, i)
	// }
	// for len(stack) > 0 {
	// 	prevGreater[stack[len(stack)-1]] = -1
	// 	stack = stack[:len(stack)-1]
	// }
	// type Element struct {
	// 	val    int
	// 	count  int
	// }
	// elements := make([]Element, n)
	// for i := 0; i < n; i++ {
	// 	left := prevGreater[i]
	// 	right := nextGreater[i]
	// 	count := (i - left) * (right - i)
	// 	elements[i] = Element{nums[i], count}
	// }
	// sort.Slice(elements, func(i, j int) bool {
	// 	return elements[i].val > elements[j].val
	// })
	// score := 1
	// remaining := k
	// for _, elem := range elements {
	// 	if remaining <= 0 {
	// 		break
	// 	}
	// 	take := elem.count
	// 	if take > remaining {
	// 		take = remaining
	// 	}
	// 	score = (score * pow(elem.val, take)) % mod
	// 	remaining -= take
	// }
	// return score
	nums = append(nums, 0)
	stack, count := []int{-1}, make(map[int]int)
	for i, v := range nums {
		for len(stack) > 1 && primes[v] > primes[nums[stack[len(stack)-1]]] {
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			count[nums[p]] += (p - stack[len(stack)-1]) * (i - p)
		}
		stack = append(stack, i)
	}
	keys := []int{}
	for v := range count {
		keys = append(keys, v)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	res := 1
	for i := 0; k > 0; i++ {
		t := min(count[keys[i]], k)
		k -= t
		res = res * pow(keys[i], t) % mod
	}
	return res
}

func pow(x, n int) int {
	res := 1
	for n > 0 {
		if n%2 == 1 {
			res = (res * x) % mod
		}
		x = (x * x) % mod
		n /= 2
	}
	return res
}

var primes [mx + 1]int

func init() {
	primes[0] = inf
	for i := 2; i <= mx; i++ {
		if primes[i] == 0 {
			for j := i; j <= mx; j += i {
				primes[j]++
			}
		}
	}
}