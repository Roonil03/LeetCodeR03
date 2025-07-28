func minimumScore(nums []int, edges [][]int) int {
	n := len(nums)
	bidiAdj := make([][]int, n)
	for _, e := range edges {
		if e[0] > e[1] {
			e[0], e[1] = e[1], e[0]
		}
		bidiAdj[e[0]] = append(bidiAdj[e[0]], e[1])
		bidiAdj[e[1]] = append(bidiAdj[e[1]], e[0])
	}
	curr := []int{0}
	next := []int{}
	seen := make([]bool, n)
	adj := make([][]int, n)
	seen[0] = true
	for len(curr) > 0 {
		next = next[:0]
		for _, i := range curr {
			for _, j := range bidiAdj[i] {
				if seen[j] {
					continue
				}
				seen[j] = true
				next = append(next, j)
				adj[i] = append(adj[i], j)
			}
		}
		curr, next = next, curr
	}
	xors := make([]int, n)
	copy(xors, nums)
	var calcXOR func(i int) int
	calcXOR = func(i int) int {
		for _, nei := range adj[i] {
			xors[i] ^= calcXOR(nei)
		}
		return xors[i]
	}
	calcXOR(0)
	isParent := make([]bool, n)
	isParent[0] = true
	isChild := make([]bool, n)
	res := math.MaxInt64
	var visitSubChild func(subRoot, child int)
	visitSubChild = func(subRoot, child int) {
		subRootXOR := xors[subRoot] ^ xors[child]
		childXOR := xors[child]
		treeXOR := xors[0] ^ xors[subRoot]
		maxXOR := max(subRootXOR, max(childXOR, treeXOR))
		minXOR := min(subRootXOR, min(childXOR, treeXOR))
		res = min(res, maxXOR-minXOR)
		isChild[child] = true
		for _, nei := range adj[child] {
			visitSubChild(subRoot, nei)
		}
	}
	var visitSubRoot func(subRoot int)
	visitSubRoot = func(subRoot int) {
		for i := range isChild {
			isChild[i] = false
		}
		for _, child := range adj[subRoot] {
			visitSubChild(subRoot, child)
		}
		isParent[subRoot] = true
		for other := 1; other < len(xors); other++ {
			if isParent[other] || isChild[other] {
				continue
			}
			a := xors[other]
			b := xors[subRoot]
			c := xors[0] ^ a ^ b
			res = min(res, max(a, max(b, c))-min(a, min(b, c)))
		}
		for _, child := range adj[subRoot] {
			visitSubRoot(child)
		}
		isParent[subRoot] = false
	}
	for _, subRoot := range adj[0] {
		visitSubRoot(subRoot)
	}
	return res
}