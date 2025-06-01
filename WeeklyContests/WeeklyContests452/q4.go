// Couldn't solve during the contest, has been solved post the contest period.

const MAXV = 100001

var isPrime [MAXV]bool

func init() {
	for i := 0; i < MAXV; i++ {
		isPrime[i] = true
	}
	isPrime[0] = false
	isPrime[1] = false
	for i := 2; i*i < MAXV; i++ {
		if isPrime[i] {
			for j := i * i; j < MAXV; j += i {
				isPrime[j] = false
			}
		}
	}
}

type SegmentTree struct {
	sum  []int
	max  []int
	size int
}

func NewSegmentTree(n int) *SegmentTree {
	size := 1
	for size < n+2 {
		size <<= 1
	}
	return &SegmentTree{
		sum:  make([]int, 2*size),
		max:  make([]int, 2*size),
		size: size,
	}
}

func (st *SegmentTree) pointAdd(pos, delta int) {
	i := pos + st.size
	st.sum[i] += delta
	st.max[i] = max(st.sum[i], 0)
	for i >>= 1; i > 0; i >>= 1 {
		st.sum[i] = st.sum[2*i] + st.sum[2*i+1]
		st.max[i] = max(st.max[2*i], st.sum[2*i]+st.max[2*i+1])
	}
}

func (st *SegmentTree) applyRange(from, to, delta int) {
	if from+1 <= to {
		st.pointAdd(from+1, delta)
		st.pointAdd(to+1, -delta)
	}
}

func (st *SegmentTree) getMax() int {
	return st.max[1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumCount(nums []int, queries [][]int) []int {
	n := len(nums)
	segTree := NewSegmentTree(n)
	primeIndices := make(map[int][]int)
	distinct := 0

	for i, v := range nums {
		if v < MAXV && isPrime[v] {
			primeIndices[v] = append(primeIndices[v], i)
		}
	}

	for p := range primeIndices {
		indices := primeIndices[p]
		sort.Ints(indices)
		first, last := indices[0], indices[len(indices)-1]
		if first < last {
			segTree.applyRange(first, last, 1)
		}
		distinct++
	}

	result := make([]int, len(queries))
	for q, query := range queries {
		idx, val := query[0], query[1]
		old := nums[idx]

		if old == val {
			result[q] = distinct + segTree.getMax()
			continue
		}

		if old < MAXV && isPrime[old] {
			indices := primeIndices[old]
			pos := sort.SearchInts(indices, idx)
			if pos < len(indices) && indices[pos] == idx {
				first, last := indices[0], indices[len(indices)-1]
				if first < last {
					segTree.applyRange(first, last, -1)
				}
				indices = append(indices[:pos], indices[pos+1:]...)
				if len(indices) == 0 {
					delete(primeIndices, old)
					distinct--
				} else {
					primeIndices[old] = indices
					first, last = indices[0], indices[len(indices)-1]
					if first < last {
						segTree.applyRange(first, last, 1)
					}
				}
			}
		}

		if val < MAXV && isPrime[val] {
			indices, ok := primeIndices[val]
			if !ok {
				primeIndices[val] = []int{idx}
				distinct++
			} else {
				first, last := indices[0], indices[len(indices)-1]
				if first < last {
					segTree.applyRange(first, last, -1)
				}
				pos := sort.SearchInts(indices, idx)
				indices = append(indices, 0)
				copy(indices[pos+1:], indices[pos:])
				indices[pos] = idx
				primeIndices[val] = indices
				first, last = indices[0], indices[len(indices)-1]
				if first < last {
					segTree.applyRange(first, last, 1)
				}
			}
		}

		nums[idx] = val
		result[q] = distinct + segTree.getMax()
	}

	return result
}

// For any split point k, we want to maximize: distinct_primes_in_prefix + distinct_primes_in_suffix

// Key Insight:
// If a prime appears at positions [a, b, c, d], then:
// For splits k where a < k ≤ b, the prime contributes to both prefix and suffix (overlap = +1)
// The maximum overlap across all split points gives us the optimal solution

// Algorithm:
// Sieve: Precompute all primes up to 100,000
// Segment Tree: Uses difference array technique for range updates
// applyRange(from, to, delta) adds delta to range [from+1, to]
// Tracks maximum prefix sum (representing maximum overlap)
// Prime Tracking: Map each prime to its sorted indices
// Updates: For each query:
// Remove old value's contribution to segment tree
// Add new value's contribution
// Answer = total_distinct_primes + max_overlap

// Time Complexity: O(Q log N log V) where Q = queries, N = array length, V = max value

// The segment tree efficiently tracks the optimal split point as the array changes,
// making this solution much faster than the naive O(QN) approach.