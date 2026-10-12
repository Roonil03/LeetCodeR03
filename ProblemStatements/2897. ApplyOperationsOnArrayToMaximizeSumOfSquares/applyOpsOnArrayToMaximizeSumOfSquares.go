func maxSum(nums []int, k int) int {
    count := make([]int, 31)
    for _, v := range nums{
        for i := 0; v > 0; i++{
            count[i] += v & 1
            v >>= 1
        }
    }
    mod := int(1e9+7)
    res :=0 
    for range k{
        cur := 0
        for i := range 31{
            if count[i] > 0{
                cur |= 1<< i
                count[i]--
            }
        }
        res = AddMod(res, MulMod(cur, cur, mod), mod)
    }
    return res
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func AddMod[T Integer](a, b, mod T) T {
	a %= mod
	b %= mod
	res := a + b
	if res >= mod {
		res -= mod
	}
	return res
}

func SubMod[T Integer](a, b, mod T) T {
	a %= mod
	b %= mod
	res := a - b
	if res < 0 {
		res += mod
	}
	return res
}

func MulMod[T Integer](a, b, mod T) T {
	a %= mod
	b %= mod
	return (a * b) % mod
}

func PowerMod[T Integer](base, exp, mod T) T {
	var res T = 1
	base %= mod
	for exp > 0 {
		if exp&1 == 1 {
			res = (res * base) % mod
		}
		base = (base * base) % mod
		exp >>= 1
	}
	return res
}