func maxValue(nums1 []int, nums0 []int) int {
	mod := int(1e9 + 7)
	n := len(nums1)
	s := make([]seg, 0, n)
	for i := range n {
		a := int64(nums1[i])
		b := int64(nums0[i])
		if a == 0 && b == 0 {
			continue
		}
		s = append(s, seg{o: a, z: b})
	}
	sort.Slice(s, func(i, j int) bool {
		x, y := s[i], s[j]
		a, b := cat(x), cat(y)
		if a != b {
			return a < b
		}
		if a == 1 {
			if x.o != y.o {
				return x.o > y.o
			}
			if x.z != y.z {
				return x.z < y.z
			}
		}
		return false
	})
	mem := map[int64]int64{0: 1}
	res := int64(0)
	for _, v := range s {
		l := pow2(v.o+v.z, mem, mod)
		m := pow2(v.z, mem, mod)
		c := (l - m + int64(mod)) % int64(mod)
		res = (res*l + c) % int64(mod)
	}
	return int(res)
}

type seg struct {
	o, z int64
}

func cat(s seg) int {
	if s.z == 0 {
		return 0
	}
	if s.o == 0 {
		return 2
	}
	return 1
}

func pow2(exp int64, mem map[int64]int64, mod int) int64 {
	if v, ok := mem[exp]; ok {
		return v
	}
	v := modPow(2, exp, mod)
	mem[exp] = v
	return v
}

func modPow(base, exp int64, mod int) int64 {
	base %= int64(mod)
	res := int64(1)
	for exp > 0 {
		if exp&1 == 1 {
			res = (res * base) % int64(mod)
		}
		base = (base * base) % int64(mod)
		exp >>= 1
	}
	return res
}