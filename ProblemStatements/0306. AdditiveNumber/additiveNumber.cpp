class Solution {
public:
    bool isAdditiveNumber(string num) {
        int n = num.size();
        for (int i {1}; i <= n/2; ++i) {
            for (int j = i+1; j < n; ++j) {
                if (isValid(i, j, num) && j < n){
                    return true;
                }
            }
        }
        return false;
    }

    bool isValid(int i, int j, string& num) {
        int n = num.size();
        string num1 = num.substr(0, i), num2 = num.substr(i, j - i);
        if ((num1.size() > 1 && num1[0] == '0') || (num2.size() > 1 && num2[0] == '0')){
            return false;
        }
        int k = j;
        while (k < n) {
            string sum = addStrings(num1, num2);
            if (num.substr(k, sum.size()) != sum){
                return false;
            }
            k += sum.size();
            num1 = num2;
            num2 = sum;
        }
        return true;
    }

    string addStrings(string& num1, string& num2) {
        string res;
        int c {0}, i = num1.size()-1, j = num2.size()-1;
        while (i >= 0 || j >= 0 || c) {
            int n1 = i >= 0 ? num1[i--] - '0' : 0;
            int n2 = j >= 0 ? num2[j--] - '0' : 0;
            int sum = n1 + n2 + c;
            res.insert(res.begin(), sum % 10 + '0');
            c = sum / 10;
        }
        return res;
    }
};