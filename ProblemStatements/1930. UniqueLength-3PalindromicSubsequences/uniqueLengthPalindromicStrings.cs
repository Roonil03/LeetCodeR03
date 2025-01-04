public class Solution {
    public int CountPalindromicSubsequence(string s) {
        int[] f = new int[26];
        int[] l = new int[26];
        Array.Fill(f, -1);
        Array.Fill(l, -1);        
        for (int i = 0; i < s.Length; i++) {
            int idx = s[i] - 'a';
            if (f[idx] == -1) {
                f[idx] = i;
            }
            l[idx] = i;
        }        
        int ans = 0;
        for (int i = 0; i < 26; i++) {
            if (f[i] == -1){
                continue;            
            }
            bool[] mid = new bool[26];
            for (int j = f[i] + 1; j < l[i]; j++) {
                mid[s[j] - 'a'] = true;
            }            
            for (int j = 0; j < 26; j++) {
                if (mid[j]){
                    ans++;
                }
            }
        }        
        return ans;
    }
}