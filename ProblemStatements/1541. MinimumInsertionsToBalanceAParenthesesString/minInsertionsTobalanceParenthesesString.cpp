class Solution {
public:
    int minInsertions(string s) {
        int res {0}, req{0};
        for(char ch : s){
            if(ch == '('){
                res += req & 1;
                req += 2 - (req & 1);
            } else if(--req < 0){
                res++;
                req += 2;                
            }
        }
        return res + req;
    }
};