public class Solution {
    public int MinLength(string s, int k) {
        List<int> nums = new List<int>();
        foreach (char c in s) {
            nums.Add(c - '0');
        }        
        int n = nums.Count;
        int f1 = 0, f2 = 0;
        for (int i = 0; i < n; i++) {
            if ((nums[i] ^ (i & 1)) != 0)
            {
                f1++;  
            }
            else{
                f2++;
            }
        }

        if (f1 <= k || f2 <= k) {
            return 1;
        }

        List<int> tmp = new List<int>();
        int cur = nums[0], cnt = 0;
        foreach (int v in nums) {
            if (v == cur) {
                cnt++;
            } else {
                tmp.Add(cnt);
                cur = v;
                cnt = 1;
            }
        }
        tmp.Add(cnt);        
        int l = 2, r = n;
        while (l <= r) {
            int m = (l + r) / 2;
            int c = 0;
            foreach (int x in tmp) {
                c += x / (m + 1);
            }
            if (c > k) l = m + 1;
            else r = m - 1;
        }
        return Math.Max(l, 2);
    }
}
