class Solution {
public:
    int countBinarySubstrings(string s) {
        int a {1}, res {0}, prev {0};
        for(int i {1}; i < s.size(); i++){
            if(s[i] == s[i-1]){
                a++;
            } else{
                prev = a;
                a = 1;
            }
            if(a <= prev){
                res++;
            }
        }
        return res;
    }
};