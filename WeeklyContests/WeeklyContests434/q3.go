// Did not solve during contest time, has been solved post the contest time
//func maxFrequency(nums []int, k int) int {
// n := len(nums)
// s := make([]int, n)
// copy(s, nums)
// sort.Ints(s)

// l := 0
// t := 0
// m := 0
// for r := 0; r < n; r++ {
//     t += s[r]
//     for l <= r && s[r]*(r-l+1) - t > k {
//         t -= s[l]
//         l++
//     }
//     if r-l+1 > m {
//         m = r-l+1
//     }
// }
// return m

/*
   int res = 0;
       for (int i = 1; i <= 50; i++) {
           int cur = 0;
           for (const auto& x : nums) {
               cur = max(0, cur + (x == i) - (x == k));
               res = max(res, cur);
           }
       }
       return res + count(nums.begin(), nums.end(), k);
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxFrequency(nums []int, k int) int {
	n := len(nums)
	kCount := 0
	for _, x := range nums {
		if x == k {
			kCount++
		}
	}
	answer := kCount
	for x := 1; x <= 50; x++ {
		currK := 0
		myMin := 0
		xCount := 0
		for i := 0; i < n; i++ {
			if nums[i] == x {
				xCount++
			}
			if nums[i] == k {
				currK++
			}
			myMin = min(myMin, xCount-currK)
			answer = max(answer, xCount-myMin+kCount-currK)
		}
	}
	return answer
}

/*
The provided program leverages basic data structures and algorithms to determine the
maximum frequency of any integer in the nums array, considering the effect of a specific
integer k. The max function is a utility to find the maximum of two integers, while the count
function counts occurrences of integer k in the array. The core function, maxFrequency,
iterates through potential integers (from 1 to 50), and for each integer, it traverses the
nums array to calculate a frequency-like score that considers occurrences of the current
integer and adjusts based on the presence of k. The algorithm uses nested loops to scan and
update a running count (cur) which is then compared to find the maximum frequency (res). The
final result is derived by adding the occurrences of k to the maximal frequency count. This
approach utilizes simple iteration and comparison operations to solve the problem efficiently
within the given constraints.
*/