class Solution {
public:
    int countSubstrings(string s) {
        int n = s.length();
        int count {0};
        for(int cen {0}; cen < 2 * n - 1; cen++){
            int l = cen / 2;
            int r = l + cen % 2;
            while(l >= 0 && r < n && s[l] == s[r]){
                count++;
                l--;
                r++;
            }
        }
        return count;
    }
};