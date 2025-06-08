// func robotWithString(s string) string {
//     freq := make([]int, 26)
//     for _, ch := range s {
//         freq[ch-'a']++
//     }
//     var paper, robot []byte
//     s_arr := []byte(s)
//     for len(s_arr) > 0 || len(robot) > 0 {
//         minChar := byte('z' + 1)
//         for i := len(s_arr) - 1; i >= 0; i-- {
//             minChar = min(minChar, s_arr[i])
//         }
//         if (len(robot) > 0 && (len(s_arr) == 0 || robot[len(robot)-1] <= minChar)) {
//             paper = append(paper, robot[len(robot)-1])
//             robot = robot[:len(robot)-1]
//         } else {
//         curr := s_arr[0]
//             s_arr = s_arr[1:]
//             robot = append(robot, curr)
//         }
//     }
//     return string(paper)
// }

func robotWithString(s string) string {
	a := 'a'
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-a]++
	}
	stack := []rune{}
	result := []rune{}
	hasSmaller := func(top int) bool {
		for i := 0; i < top; i++ {
			if freq[i] > 0 {
				return true
			}
		}
		return false
	}
	for _, ch := range s {
		idx := ch - a
		freq[idx]--
		stack = append(stack, ch)
		for len(stack) > 0 && !hasSmaller(int(stack[len(stack)-1]-a)) {
			result = append(result, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}
	return string(result)
}