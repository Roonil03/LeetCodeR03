class Solution {
    final long N = 100005;
    final long hell = 1000000007;
    long[] a = new long[(int) N];
    long[] seg = new long[4 * (int) N];
    long[] segsq = new long[4 * (int) N];
    long[] lazy = new long[4 * (int) N];

    void build(int node, long start, long end) {
        if (start == end) {
            seg[node] = a[(int) start];
            segsq[node] = (a[(int) start] * a[(int) start]) % hell;
            return;
        }
        long mid = (start + end) >> 1;
        build(node << 1, start, mid);
        build(node << 1 | 1, mid + 1, end);
        seg[node] = (seg[node << 1] + seg[node << 1 | 1]) % hell;
        segsq[node] = (segsq[node << 1] + segsq[node << 1 | 1]) % hell;
    }

    void update(int node, long start, long end, int l, int r, int val) {
        if (lazy[node] != 0) {
            segsq[node] += ((end - start + 1) * (lazy[node] * lazy[node]) % hell) % hell + (2 * lazy[node] * seg[node]) % hell;
            segsq[node] %= hell;
            seg[node] += ((end - start + 1) * lazy[node]) % hell;
            seg[node] %= hell;
            if (start != end) {
                lazy[node << 1] += lazy[node];
                lazy[node << 1 | 1] += lazy[node];
            }
            lazy[node] = 0;
        }
        if (start > end || start > r || end < l)
        {
            return;
        }
        if (l <= start && end <= r) {
            segsq[node] += ((end - start + 1) * (val * val) % hell) % hell + (2 * val * seg[node]) % hell;
            segsq[node] %= hell;
            seg[node] += ((end - start + 1) * val) % hell;
            seg[node] %= hell;
            if (start != end) {
                lazy[node << 1] += val;
                lazy[node << 1 | 1] += val;
            }
            return;
        }
        long mid = (start + end) >> 1;
        update(node << 1, start, mid, l, r, val);
        update(node << 1 | 1, mid + 1, end, l, r, val);
        seg[node] = (seg[node << 1] + seg[node << 1 | 1]) % hell;
        segsq[node] = (segsq[node << 1] + segsq[node << 1 | 1]) % hell;
    }

    public int sumCounts(int[] nums) {
        int n = nums.length;
        for (int i = 1; i <= n; i++) {
            a[i] = 0;
        }
        build(1, 1, n);
        HashMap<Integer, Integer> prev_seen_at = new HashMap<>();
        long ans = 0;
        for (int i = n - 1; i >= 0; i--) {
            if (!prev_seen_at.containsKey(nums[i])) {
                update(1, 1, n, i + 1, n, 1);
            } 
            else {
                update(1, 1, n, i + 1, prev_seen_at.get(nums[i]) - 1, 1);
            }

            prev_seen_at.put(nums[i], i + 1);

            ans = (ans + segsq[1]) % hell;
        }
        return (int) ans;
    }
}

/*
This code implements a segment tree with lazy propagation to efficiently calculate the sum of
 squares of distinct counts of all subarrays for a given array nums. The segment tree is used 
 to track both the sum (seg) and the sum of squares (segsq) of elements in the range. Lazy 
 propagation (lazy array) ensures efficient updates to the segment tree when range updates are 
 performed. The build function initializes the segment tree, while the update function handles 
 range updates, modifying both seg and segsq to reflect changes in the range. The algorithm 
 processes the array from right to left, using a HashMap to track the last occurrence of each 
 element (prev_seen_at). For each element, it updates the segment tree to reflect the addition 
 of the element to the distinct count, handling overlaps or previous occurrences. The result 
 for each step is the sum of squares of distinct counts for subarrays starting at that index, 
 which is added to the answer. The use of the segment tree ensures that both updates and 
 queries are efficient, with O(logn) time complexity per operation, making the approach 
 suitable for large arrays.
*/