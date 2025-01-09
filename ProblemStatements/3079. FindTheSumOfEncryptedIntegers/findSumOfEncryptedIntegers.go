func sumOfEncryptedInt(nums []int) int {
	sum := 0
	for _, n := range nums {
		maxD := 0
		n1 := n
		for n1 > 0 {
			if d := n1 % 10; d > maxD {
				maxD = d
			}
			n1 /= 10
		}
		enc := 0
		n1 = n
		pow := 1
		for n1 > 0 {
			enc += maxD * pow
			pow *= 10
			n1 /= 10
		}
		sum += enc
	}
	return sum
}