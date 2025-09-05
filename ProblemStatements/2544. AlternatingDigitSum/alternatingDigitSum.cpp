class Solution {
public:
    int alternateDigitSum(int n) {
        int sum {0};
        string s = to_string(n);
        int sign = 1;
        for(char c : s){
            int dig = c - '0';
            sum += sign * dig;
            sign = -sign;
        }
        return sum;
    }
};