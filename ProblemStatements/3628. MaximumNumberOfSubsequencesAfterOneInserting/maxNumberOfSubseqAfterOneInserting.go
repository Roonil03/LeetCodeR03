// func numOfSubsequences(s string) int64 {
//     if s == "LT"{
//         return 1
//     }
//     n := len(s)
//     cL := make([]int, n + 1)
//     cT := make([]int, n + 1)
//     ct := 0
//     for i := n - 1; i >= 0; i--{
//         cT[i] = ct
//         if s[i] == 'T'{
//             ct++
//         }
//     }
//     cl := 0
//     for i := 0; i< n; i++{
//         cL[i] = cl
//         if s[i] == 'L'{
//             cl++
//         }
//     }
//     cL[n] = cl
//     init := 0
//     for i := 0; i < n; i++{
//         if s[i] == 'C'{
//             init += cL[i] * cT[i]
//         }
//     }
//     ctSubseq := make([]int, n+1)
//     for i := n - 1; i >= 0; i-- {
//         ctSubseq[i] = ctSubseq[i+1]
//         if s[i] == 'C' {
//             ctSubseq[i] += cT[i]
//         }
//     }
//     lcSubseq := make([]int, n+1)
//     for i := 1; i <= n; i++ {
//         lcSubseq[i] = lcSubseq[i-1]
//         if i > 0 && s[i-1] == 'C' {
//             lcSubseq[i] += cL[i-1]
//         }
//     }
//     res := int64(init)
//     for i := 0; i <= n; i++{
//         // a := int64(init) + int64(cT[i])
//         // if a > res{
//         //     res = a
//         // }
//         // a = int64(init) + int64(cL[i]) * int64(cT[i])
//         // if a > res{
//         //     res = a
//         // }
//         // a = int64(init) + int64(cL[i])
//         // if a > res{
//         //     res = a
//         // }
//         a := int64(init) + int64(ctSubseq[i])
//         if a > res {
//             res = a
//         }
//         a = int64(init) + int64(cL[i]) * int64(cT[i])
//         if a > res {
//             res = a
//         }
//         a = int64(init) + int64(lcSubseq[i])
//         if a > res {
//             res = a
//         }
//     }
//     return res
// }

func numOfSubsequences(s string) int64 {
	LPrefixCount := make([]int, len(s))
	if s[0] == 'L' {
		LPrefixCount[0] = 1
	}
	for i := 1; i < len(s); i++ {
		LPrefixCount[i] = LPrefixCount[i-1]
		if s[i] == 'L' {
			LPrefixCount[i]++
		}
	}
	TSuffixCount := make([]int, len(s))
	if s[len(s)-1] == 'T' {
		TSuffixCount[len(s)-1] = 1
	}
	for i := len(s) - 2; i >= 0; i-- {
		TSuffixCount[i] = TSuffixCount[i+1]
		if s[i] == 'T' {
			TSuffixCount[i]++
		}
	}
	var Cres int64
	var maxC int64
	var Lres int64
	var Tres int64
	for i := range s {
		if s[i] == 'C' {
			Cres += int64(LPrefixCount[i] * TSuffixCount[i])
			Lres += int64((LPrefixCount[i] + 1) * TSuffixCount[i])
			Tres += int64(LPrefixCount[i] * (TSuffixCount[i] + 1))
		}
		maxC = max(maxC, int64(LPrefixCount[i]*TSuffixCount[i]))
	}
	return max(Cres+maxC, max(Lres, Tres))
}