func findEvenNumbers(digits []int) []int {
	res := []int{}
	count := make([]int, 10)
	for _, i := range digits {
		count[i]++
	}
	for i := 1; i < 10; i++ {
		for j := 0; count[i] > 0 && j < 10; j++ {
			if count[i] > 0 && count[j] > 0 {
				for k := 0; k < 10; k += 2 {
					if count[k] > 0 {
						a := 0
						if i == j {
							a = 1
						}
						b := 0
						if k == i {
							b = 1
						}
						c := 0
						if k == j {
							c = 1
						}
						if count[j] > a && count[k] > b+c {
							n := i*100 + j*10 + k
							res = append(res, n)
						}
					}
				}
			}
		}
	}
	return res
}