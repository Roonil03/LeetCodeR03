func fib(n int) int {
	mem := make(map[int]int)
	mem[0], mem[1], mem[2] = 0, 1, 1
	var h func(int) int
	h = func(n int) int {
		if v, ex := mem[n]; ex {
			return v
		}
		res := h(n-1) + h(n-2)
		mem[n] = res
		return res
	}
	return h(n)
}