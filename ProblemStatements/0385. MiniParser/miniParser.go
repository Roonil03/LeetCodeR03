/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func deserialize(s string) *NestedInteger {
	// if s[0] !='['{
	//     a := 1
	//     i := 0
	//     if s[0] == '-'{
	//         a = -1
	//         i = 1
	//     }
	//     num := 0
	//     for ; i < len(s); i++{
	//         num = num * 10 + int(s[i] - '0')
	//     }
	//     ni := &NestedInteger{}
	//     ni.SetInteger(a * num)
	//     return ni
	// }
	// st := []*NestedInteger{}
	// num := 0
	// neg, h2 := false, false
	// for i := 0; i < len(s); i++{

	// }
	stack := []NestedInteger{{}}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-':
			j := i + 1
			for ; j < len(s) && unicode.IsDigit(rune(s[j])); j++ {
			}
			integer, _ := strconv.Atoi(s[i:j])
			nestedInteger := NestedInteger{}
			nestedInteger.SetInteger(integer)
			stack[len(stack)-1].Add(nestedInteger)
			i = j - 1
		case '[':
			stack = append(stack, NestedInteger{})
		case ']':
			nestedInteger := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1].Add(nestedInteger)
		}
	}
	return stack[len(stack)-1].GetList()[0]
}