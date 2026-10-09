class Solution {
public:
    string removeOuterParentheses(string s) {
        int count = 0, k = 0;
        for(char ch : s){
            if(ch == '(' ? count++ > 0 : --count > 0){
                s[k++] = ch;
            }
        }
        s.resize(k);
        return s;
    }
};