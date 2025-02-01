// func wiggleSort(nums []int)  {
//     if nums == nil{
//         return
//     }
//     n := len(nums)
//     temp := make([]int, n)
//     copy(temp, nums)
//     mid := (n - 1) / 2
//     median := quickSelect(temp, 0, n-1, mid)
//     m := func(idx int) int {
//         return (2*idx + 1) % (n | 1)
//     }
//     first, mid, last := 0, 0, n-1
//     for mid <= last {
//         if nums[m(mid)] > median {
//             nums[m(first)], nums[m(mid)] = nums[m(mid)], nums[m(first)]
//             first++
//             mid++
//         } else if nums[m(mid)] < median {
//             nums[m(mid)], nums[m(last)] = nums[m(last)], nums[m(mid)]
//             last--
//         } else {
//             mid++
//         }
//     }
// }

// func quickSelect(nums []int, left, right, k int) int {
//     if left == right {
//         return nums[left]
//     }
//     pivotIndex := partition(nums, left, right)
//     if k == pivotIndex {
//         return nums[k]
//     } else if k < pivotIndex {
//         return quickSelect(nums, left, pivotIndex-1, k)
//     }
//     return quickSelect(nums, pivotIndex+1, right, k)
// }

// func partition(nums []int, left, right int) int {
//     pivot := nums[right]
//     i := left - 1
//     for j := left; j < right; j++ {
//         if nums[j] <= pivot {
//             i++
//             nums[i], nums[j] = nums[j], nums[i]
//         }
//     }
//     nums[i+1], nums[right] = nums[right], nums[i+1]
//     return i + 1
// }

// func wiggleSort(ns []int) {
//     if len(ns) <= 1 {
//         return
//     }
//     tm := make([]int, len(ns))
//     copy(tm, ns)
//     ms(tm, 0, len(ns)-1)
//     md := (len(ns) + 1) / 2
//     lf := tm[:md]
//     rt := tm[md:]
//     j, k := len(lf)-1, len(rt)-1
//     for i := 0; i < len(ns); i++ {
//         if i%2 == 0 {
//             ns[i] = lf[j]
//             j--
//         } else {
//             ns[i] = rt[k]
//             k--
//         }
//     }
// }

// func ms(ns []int, lf, rt int) {
//     if lf < rt {
//         md := lf + (rt-lf)/2
//         ms(ns, lf, md)
//         ms(ns, md+1, rt)
//         mg(ns, lf, md, rt)
//     }
// }

// func mg(ns []int, lf, md, rt int) {
//     tm := make([]int, rt-lf+1)
//     i, j, k := lf, md+1, 0
//     for i <= md && j <= rt {
//         if ns[i] <= ns[j] {
//             tm[k] = ns[i]
//             i++
//         } else {
//             tm[k] = ns[j]
//             j++
//         }
//         k++
//     }
//     for i <= md {
//         tm[k] = ns[i]
//         i++
//         k++
//     }
//     for j <= rt {
//         tm[k] = ns[j]
//         j++
//         k++
//     }
//     for i := 0; i < len(tm); i++ {
//         ns[lf+i] = tm[i]
//     }
// }

func wiggleSort(nums []int) {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	mid := (len(nums) - 1) / 2
	right := len(nums) - 1

	for i := range nums {
		if i%2 == 0 {
			nums[i] = sorted[mid]
			mid--
		} else {
			nums[i] = sorted[right]
			right--
		}
	}
}