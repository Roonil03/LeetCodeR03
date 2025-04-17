// func countPairs(nums []int, k int) int {
// 	pairs := 0
// 	mpp := make(map[int][]int)
// 	for i, num := range nums {
// 		mpp[num] = append(mpp[num], i)
// 	}
// 	divisors := []int{}
// 	for d := 1; d*d <= k; d++ {
// 		if k%d == 0 {
// 			divisors = append(divisors, d)
// 			if d != k/d {
// 				divisors = append(divisors, k/d)
// 			}
// 		}
// 	}
// 	for _, vec := range mpp {
// 		mpp := make(map[int]int)
// 		for _, i := range vec {
// 			gcdd := gcd(i, k)
// 			add := k / gcdd
// 			pairs += mpp[add]
// 			for _, d := range divisors {
// 				if i%d == 0 {
// 					mpp[d]++
// 				}
// 			}
// 		}
// 	}
// 	return pairs
// }

// func gcd(a, b int) int {
// 	return int(big.NewInt(0).GCD(nil, nil, big.NewInt(int64(a)), big.NewInt(int64(b))).Int64())
// }

func countPairs(nums []int, k int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && (i*j)%k == 0 {
				res++
			}
		}
	}
	return res
}