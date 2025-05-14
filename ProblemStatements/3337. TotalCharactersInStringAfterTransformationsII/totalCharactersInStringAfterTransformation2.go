type vec26 [26]int
type matrix [26][26]int

const mod int = 1e9 + 7

func identityMatrix() matrix {
	var I matrix
	for i := 0; i < 26; i++ {
		I[i][i] = 1
	}
	return I
}

func multiplyMatrix(A, B matrix) matrix {
	var C matrix
	for i := 0; i < 26; i++ {
		for k := 0; k < 26; k++ {
			if A[i][k] == 0 {
				continue
			}
			for j := 0; j < 26; j++ {
				C[i][j] = (C[i][j] + A[i][k]*B[k][j]) % mod
			}
		}
	}
	return C
}

func multiplyMatrixVector(A matrix, X vec26) vec26 {
	var ans vec26
	for i := 0; i < 26; i++ {
		sum := 0
		for j := 0; j < 26; j++ {
			sum += A[i][j] * X[j]
		}
		ans[i] = sum % mod
	}
	return ans
}

func powLSBF(M matrix, n int) matrix {
	ans := identityMatrix()
	for n > 0 {
		if n&1 == 1 {
			ans = multiplyMatrix(ans, M)
		}
		M = multiplyMatrix(M, M)
		n >>= 1
	}
	return ans
}

func buildTransitionMatrix(nums []int) matrix {
	var M matrix
	for i := 0; i < 26; i++ {
		jN := i + nums[i]
		for j := i + 1; j <= jN; j++ {
			M[j%26][i] = 1
		}
	}
	return M
}

func lengthAfterTransformations(s string, t int, nums []int) int {
	var freq vec26
	for _, c := range s {
		freq[c-'a']++
	}
	M := buildTransitionMatrix(nums)
	B := powLSBF(M, t)
	C := multiplyMatrixVector(B, freq)
	res := 0
	for _, val := range C {
		res = (res + val) % mod
	}
	return res
}