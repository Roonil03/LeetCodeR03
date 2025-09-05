class Solution {
public:
    int longestSubstring(string s, int k) {
        int freq[26] = {0};
        for(char c : s){
            freq[c - 'a']++;
        }
        for(int i {0}; i< s.length(); i++){
            if(freq[s[i] - 'a'] < k){
                int l = longestSubstring(s.substr(0, i), k);
                int r = longestSubstring(s.substr(i+1), k);
                return max(l, r);
            }
        }
        return s.size();
    }
};