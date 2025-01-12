// func canBeValid(s string, locked string) bool {
// n := len(s)
// if n%2==1{
//     return false
// }
// b := 0
// for i:= 0; i< n; i++{
//     if s[i] == '(' || locked[i] == '0'{
//         b++
//     } else {
//         b--
//     }
//     if b < 0{
//         return false
//     }
// }
// b = 0
// for i := n - 1; i >= 0; i-- {
//     if locked[i] == '0' || s[i] == ')' {
//         b++
//     } else {
//         b--
//     }
//     if b < 0 {
//         return false
//     }
// }
// return true
// }
func validate(s, locked string, op byte) bool {
	bal, wild := 0, 0
	sz := len(s)
	start, dir := 0, 1
	if op == '(' {
		start, dir = 0, 1
	} else {
		start, dir = sz-1, -1
	}
	for i := start; i >= 0 && i < sz && wild+bal >= 0; i += dir {
		if locked[i] == '1' {
			if s[i] == op {
				bal++
			} else {
				bal--
			}
		} else {
			wild++
		}
	}
	return abs(bal) <= wild
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func canBeValid(s, locked string) bool {
	return len(s)%2 == 0 && validate(s, locked, '(') && validate(s, locked, ')')
}