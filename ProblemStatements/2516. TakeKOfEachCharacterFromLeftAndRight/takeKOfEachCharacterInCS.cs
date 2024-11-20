public class Solution {
    public int TakeCharacters(string s, int k) {
        int[] count = new int[3];
        foreach (char c in s) {
            count[c - 'a']++;
        }        
        if (Math.Min(count[0], Math.Min(count[1], count[2])) < k) {
            return -1;
        }        
        int n = s.Length;
        int res = n + 1, l = 0;        
        for (int r = 0; r < n; r++) {
            count[s[r] - 'a']--;            
            while (Math.Min(count[0], Math.Min(count[1], count[2])) < k) {
                count[s[l] - 'a']++;
                l++;
            }            
            res = Math.Min(res, n - (r - l + 1));
        }        
        return res;
    }
}
