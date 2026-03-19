class Solution {
public:
    string minRemoveToMakeValid(string s) {
        string t;
        int open {0};
        for(char c : s){
            if(c == '('){
                open++;
                t.push_back(c);
            } else if(c == ')'){
                if(open > 0){
                    open--;
                    t.push_back(c);
                }
            } else{
                t.push_back(c);
            }
        }
        string res;
        for(int i = t.size() - 1;i >= 0; i--){
            if(t[i] == '(' && open > 0){
                open--;
            } else{
                res.push_back(t[i]);
            }
        }
        reverse(res.begin(), res.end());
        return res;
    }
};