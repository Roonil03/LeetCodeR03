class Solution {
public:
    bool validPalindrome(string s) {
        for(int i {0}, j = s.size() - 1; i < j; i++, j--){
            if(s[i] != s[j]){
                auto g1 = [&](int l, int r){
                    for(; l < r; l++, r--){
                        if(s[l] != s[r]){
                            return false;
                        }
                    }
                    return true;
                };
                return g1(i + 1, j) || g1(i, j - 1);
            }
        }
        return true;
    }
};