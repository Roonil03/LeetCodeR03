class Solution {
public:
    string addStrings(string num1, string num2) {
        int c {0};
        string res = "";
        int i = num1.size() - 1, j = num2.size() - 1;
        while(i >= 0 || j >= 0 || c > 0){
            int d1 = (i >= 0) ? num1[i] - '0' : 0;
            int d2 = (j >= 0) ? num2[j] - '0' : 0;
            int sum = d1 + d2 + c;
            c = sum / 10;
            res += to_string(sum % 10);
            i--;
            j--;
        }
        reverse(res.begin(), res.end());
        return res;
    }
};