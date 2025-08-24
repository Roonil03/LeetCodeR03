class Solution {
public:
    int characterReplacement(string s, int k) {
        vector<int>count(26, 0);
        int l {0}, m {0}, res {0};
        for(int r {0}; r < s.size(); r++){
            count[s[r] - 'A']++;
            m = max(m, count[s[r] - 'A']);
            if(r - l + 1 - m > k){
                count[s[l] - 'A']--;
                l++;
            }
            res = max(res, r - l + 1);
        }
        return res;
    }
};