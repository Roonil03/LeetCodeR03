func maxGCDScore(nums []int, k int) int64 {
	// wes := nums
	// mx := 0
	// for _, x := range wes {
	//     if x > mx {
	//         mx = x
	//     }
	// }
	// ds := make(map[int]bool)
	// lmt := int(math.Sqrt(float64(mx))) + 1
	// if lmt > 10000000 {
	//     lmt = 10000000
	// }
	// for _, x := range wes {
	//     for i := 1; i <= lmt && i*i <= x; i++ {
	//         if x%i == 0 {
	//             ds[i] = true
	//             if x/i <= 10000000 {
	//                 ds[x/i] = true
	//             }
	//         }
	//     }
	//     x2 := 2 * x
	//     for i := 1; i <= lmt && i*i <= x2; i++ {
	//         if x2%i == 0 {
	//             ds[i] = true
	//             if x2/i <= 10000000 {
	//                 ds[x2/i] = true
	//             }
	//         }
	//     }
	// }
	// dl := make([]int, 0, len(ds))
	// for d := range ds {
	//     dl = append(dl, d)
	// }
	// sort.Sort(sort.Reverse(sort.IntSlice(dl)))
	// r := int64(0)
	// n := len(wes)
	// for _, d := range dl {
	//     if int64(d)*int64(n) <= r {
	//         break
	//     }
	//     ml := 0
	//     l := 0
	//     ch := 0
	//     for ri := 0; ri < n; ri++ {
	//         x := wes[ri]
	//         if x%d == 0 {
	//         } else if (2*x)%d == 0 {
	//             ch++
	//         } else {
	//             l = ri + 1
	//             ch = 0
	//             continue
	//         }
	//         for ch > k {
	//             if wes[l]%d != 0 && (2*wes[l])%d == 0 {
	//                 ch--
	//             }
	//             l++
	//         }
	//         cl := ri - l + 1
	//         if cl > ml {
	//             ml = cl
	//         }
	//     }
	//     sc := int64(d) * int64(ml)
	//     if sc > r {
	//         r = sc
	//     }
	// }
	// return r
	// n := len(nums)
	// prefix := make([]int64, n+1)
	// for i := 0; i < n; i++ {
	//     prefix[i+1] = prefix[i] + int64(nums[i])
	// }

	// type pair struct {
	//     start int
	//     g     int
	// }

	// var res int64 = 0
	// f := []pair{}

	// for i := 0; i < n; i++ {
	//     newF := make([]pair, 0, len(f)+1)
	//     seen := make(map[int]int)

	//     for _, p := range f {
	//         newGcd := gcd(p.g, nums[i])
	//         if newGcd == 0 {
	//             newGcd = nums[i]
	//         }
	//         if start, exists := seen[newGcd]; exists {
	//             if p.start < start {
	//                 seen[newGcd] = p.start
	//             }
	//         } else {
	//             seen[newGcd] = p.start
	//         }
	//     }

	//     newGcd := nums[i]
	//     if start, exists := seen[newGcd]; exists {
	//         if i < start {
	//             seen[newGcd] = i
	//         }
	//     } else {
	//         seen[newGcd] = i
	//     }
	//     for g, start := range seen {
	//         newF = append(newF, pair{start, g})
	//     }
	//     f = newF[:0]
	//     seenGcds := make(map[int]int)
	//     for _, p := range newF {
	//         if earliestStart, exists := seenGcds[p.g]; exists {
	//             if p.start < earliestStart {
	//                 seenGcds[p.g] = p.start
	//             }
	//         } else {
	//             seenGcds[p.g] = p.start
	//         }
	//     }
	//     for g, start := range seenGcds {
	//         f = append(f, pair{start, g})
	//     }
	//     for _, p := range f {
	//         if i - p.start + 1 >= k {
	//             current := (prefix[i+1] - prefix[p.start]) * int64(p.g)
	//             if current > res {
	//                 res = current
	//             }
	//         }
	//     }
	// }
	// if res > 9007199254740991 {
	//     return 9007199254740991
	// }
	// if res < -9007199254740991 {
	//     return -9007199254740991
	// }
	// return res
	n := len(nums)
	e := make([]int, n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		x := nums[i]
		c := 0
		for x&1 == 0 {
			x >>= 1
			c++
		}
		e[i] = c
		a[i] = nums[i] >> e[i]
	}
	var b0, b1 int64
	for i := 0; i < n; i++ {
		g := 0
		em := 1000000000
		cm := 0
		for j := i; j < n; j++ {
			g = gcd(g, a[j])
			ej := e[j]
			if ej < em {
				em = ej
				cm = 1
			} else if ej == em {
				cm++
			}
			l := j - i + 1
			s := int64(l) * int64(g) * (1 << em)
			if s > b0 {
				b0 = s
			}
			if cm <= k && s > b1 {
				b1 = s
			}
		}
	}
	res := max(b0, 2*b1)
	return res
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}

func getDivisors(n int) map[int]bool {
	dv := make(map[int]bool)
	if n == 0 {
		return dv
	}
	na := int(math.Abs(float64(n)))
	for i := 1; i <= int(math.Sqrt(float64(na))); i++ {
		if na%i == 0 {
			dv[i] = true
			dv[na/i] = true
		}
	}
	return dv
}