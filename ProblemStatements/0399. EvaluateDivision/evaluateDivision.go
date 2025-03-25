// type key struct{
//     a, b string
// }

// func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
//     m := map[key]float64{}
//     for i, eq := range equations{
//         a, b := eq[0], eq[1]
//         m[key{a,b}] = values[i]
//         m[key{b, a}] = 1 / values[i]
//         m[key{a, a}] = 1.0
//         m[key{b, b}] = 1.0
//     }
//     for k, v := range m{
//         dfs(k.a, k.b, v, m)
//     }
//     var res []float64
//     for _, q := range queries{
//         v, ok := m[key{q[0], q[1]}]
//         if ok{
//             res = append(res, v)
//         } else{
//             res = append(res, -1.0)
//         }
//     }
//     return res
// }

// func dfs(a, b string, v float64, m map[key]float64){
//     for k, vk := range m{
//         if b == k.a && a != k.b{
//             nk := key{a, k.b}
//             _, ok := m[nk]
//             if !ok{
//                 m[nk] = v * vk
//                 dfs(a, k.b, v * vk, m)
//             }
//         }
//     }
// }

type Item struct {
	Idx int
	Val float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	store := make(map[string]int)
	graphs := make([][]Item, 0)
	n := len(equations)
	for i := 0; i < n; i++ {
		_, ok := store[equations[i][0]]
		if !ok {
			store[equations[i][0]] = len(store)
			graphs = append(graphs, make([]Item, 0))
		}
		_, ok = store[equations[i][1]]
		if !ok {
			store[equations[i][1]] = len(store)
			graphs = append(graphs, make([]Item, 0))
		}
		graphs[store[equations[i][0]]] = append(graphs[store[equations[i][0]]], Item{
			Idx: store[equations[i][1]],
			Val: values[i],
		})
		graphs[store[equations[i][1]]] = append(graphs[store[equations[i][1]]], Item{
			Idx: store[equations[i][0]],
			Val: 1.0 / values[i],
		})
	}
	var bfs func(from, to string) float64
	bfs = func(from, to string) float64 {
		if _, ok := store[from]; !ok {
			return -1.0
		}
		if _, ok := store[to]; !ok {
			return -1.0
		}
		queue := []Item{Item{
			Idx: store[from],
			Val: 1.0,
		}}
		used := make([]bool, len(store))
		for len(queue) > 0 {
			for _, q := range queue {
				queue = queue[1:]
				if used[q.Idx] {
					continue
				}
				used[q.Idx] = true
				if q.Idx == store[to] {
					return q.Val
				}
				for _, next := range graphs[q.Idx] {
					if used[next.Idx] {
						continue
					}
					queue = append(queue, Item{
						Idx: next.Idx,
						Val: q.Val * next.Val,
					})
				}
			}
		}
		return -1.0
	}
	var ans []float64
	for _, query := range queries {
		ans = append(ans, bfs(query[0], query[1]))
	}
	return ans
}