class Solution {
public:
    int calculate(string s) {
        stack<int> nums;
        long c = 0;
        char op = '+';        
        for (int i = 0; i < s.length(); i++) {
            char ch = s[i];            
            if (isdigit(ch)) {
                c = c * 10 + (ch - '0');
            }            
            if ((!isdigit(ch) && !isspace(ch)) || i == s.length() - 1) {
                if (op == '+') {
                    nums.push(c);
                }
                else if (op == '-') {
                    nums.push(-c);
                }
                else if (op == '*') {
                    int top = nums.top();
                    nums.pop();
                    nums.push(top * c);
                }
                else if (op == '/') {
                    int top = nums.top();
                    nums.pop();
                    nums.push(top / c);
                }
                op = ch;
                c = 0;
            }
        }        
        int res = 0;
        while (!nums.empty()) {
            res += nums.top();
            nums.pop();
        }        
        return res;
    }
};