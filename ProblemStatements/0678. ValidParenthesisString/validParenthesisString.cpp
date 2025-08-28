class Solution {
public:
    bool checkValidString(string s) {
        int l {0}, h {0};
        for(char c : s){
            if(c == '('){
                l++;
                h++;
            } else if(c == ')'){
                l = max(l - 1, 0);
                h--;
                if(h < 0){
                    return false;
                }
            } else{
                l = max(l - 1, 0);
                h++;
            }
        }
        return l == 0;
    }
};