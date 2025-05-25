using ll = long long;

class Solution {
public:
    string fractionToDecimal(int numerator, int denominator) {
        if (numerator == 0){
            return "0";
        }
        string res;
        if ((numerator < 0) ^ (denominator < 0)){
            res  += '-';
        }
        ll n = llabs((ll)numerator);
        ll d = llabs((ll)denominator);
        res += to_string(n/d);
        ll rem = n%d;
        if (rem == 0){
            return res;
        }
        res += ".";
        unordered_map<ll, int> m;
        while(rem){
            if(m.count(rem)){
                res.insert(m[rem], "(");
                res += ")";
                break;
            }
            m[rem] = res.size();
            rem *= 10;
            res += to_string(rem / d);
            rem %= d;
        }
        return res;
    }
};