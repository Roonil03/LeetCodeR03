// Could not solve during the time period of the contest
// Has been solved later

public class Solution {
    private long[] a = new long[400005];
    private long[] b = new long[400005];
    private long[] s = new long[400005];
    private long[] lz = new long[400005];
    private int[] v = new int[100005];
    private int n;

    private struct Val {
        public int x, y;
    }
    private Val[] t = new Val[100005];

    private void Down(int x) {
        if (lz[x] == 0) return;
        lz[x * 2] += lz[x];
        lz[x * 2 + 1] += lz[x];
        a[x * 2] += lz[x];
        a[x * 2 + 1] += lz[x];
        b[x * 2] += lz[x];
        b[x * 2 + 1] += lz[x];
        lz[x] = 0;
    }

    private void Up(int x) {
        a[x] = Math.Min(a[x * 2], a[x * 2 + 1]);
        b[x] = Math.Max(b[x * 2], b[x * 2 + 1]);
        s[x] = Math.Max(s[x * 2], s[x * 2 + 1]);
        s[x] = Math.Max(s[x], b[x * 2 + 1] - a[x * 2]);
    }

    private void Upd(int x, int l, int r, int L, int R, int y) {
        if (L == l && r == R) {
            a[x] += y;
            b[x] += y;
            lz[x] += y;
            return;
        }
        Down(x);
        int mid = (l + r) >> 1;
        if (R <= mid) Upd(x * 2, l, mid, L, R, y);
        else if (L >= mid + 1) Upd(x * 2 + 1, mid + 1, r, L, R, y);
        else {
            Upd(x * 2, l, mid, L, mid, y);
            Upd(x * 2 + 1, mid + 1, r, mid + 1, R, y);
        }
        Up(x);
    }

    private long Ask() {
        return s[1];
    }

    public long MaxSubarraySum(int[] nums) {
        n = nums.Length;
        for (int i = 0; i < n; ++i) {
            v[i + 1] = nums[i];
            t[i + 1].x = i + 1;
            t[i + 1].y = nums[i];
        }
        bool flag = true;
        for (int i = 0; i < n; ++i) {
            if (nums[i] > 0) flag = false;
        }
        if (flag) {
            long mx = nums[0];
            for (int i = 1; i < n; ++i) mx = Math.Max(mx, (long)nums[i]);
            return mx;
        }
        for (int i = 1; i <= n; ++i) Upd(1, 0, n, i, n, v[i]);
        long ans = Ask();
        Array.Sort(t, 1, n + 1, Comparer<Val>.Create((a, b) => a.y.CompareTo(b.y)));
        for (int i = 1, j; i <= n; i = j + 1) {
            j = i;
            while (j < n && t[j + 1].y == t[j].y) ++j;
            for (int k = i; k <= j; ++k) Upd(1, 0, n, t[k].x, n, -t[k].y);
            ans = Math.Max(ans, Ask());
            for (int k = i; k <= j; ++k) Upd(1, 0, n, t[k].x, n, t[k].y);
        }
        return ans;
    }
}

// The provided code defines a Solution class with a method MaxSubarraySum that calculates the maximum subarray 
// sum for an array nums using a segment tree with lazy propagation. The segment tree is used to efficiently 
// manage range updates and queries. The Down method propagates the lazy updates to child nodes, while the Up 
// method updates the segment tree nodes to maintain minimum, maximum, and maximum difference values. The Upd 
// method applies range updates, and Ask retrieves the maximum subarray sum. The MaxSubarraySum method initializes 
// the segment tree, handles the case where all elements are non-positive, updates the segment tree with the 
// array values, and iterates through the sorted values to adjust the segment tree and find the maximum subarray 
// sum. The time complexity is (O(nlog n)) due to sorting and segment tree operations, and the space complexity 
// is (O(n)).