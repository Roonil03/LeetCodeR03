class Solution {
public:
    int longestPalindrome(string s) {
        int n = s.length();
        unordered_map<char, int>mp;
        for(char ch : s){
            mp[ch]++;
        }
        int res {0};
        bool o =false;
        for(auto& a : mp){
            if (a.second & 1){
                o = true;
                res += a.second - 1;
            } else{
                res += a.second;
            }
            // cout << a.first << ": "<< a.second<< " " /*<< mOdd*/ << "/" << res << endl;
        }
        return res + (o?1 : 0);
    }
};
