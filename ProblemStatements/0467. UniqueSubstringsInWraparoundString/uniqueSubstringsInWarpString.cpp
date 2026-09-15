class Solution {
public:
    int findSubstringInWraproundString(string s) {
        int dp[26] = {};
        int l {0};
        for(int i {0}; i < s.size(); i++){
            l = (i && (s[i] - s[i - 1] == 1 || s[i - 1] - s[i] == 25)) ? l + 1 : 1;
            dp[s[i] - 'a'] = max(dp[s[i] - 'a'], l);
        }
        return reduce(dp, dp + 26);
    }
};