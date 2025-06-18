class Solution {
public:
    using ll = long long;
    int countDigitOne(int n) {
        ll res {0};
        for(ll i{1}; i <= n; i *= 10){
            ll l = n % i, c = (n / i) % 10, h = n / (i * 10);
            if(c == 0){
                res += h * i;
            } else if(c == 1){
                res += h * i + l + 1;
            } else {
                res += (h + 1) * i;
            }
        }
        return (int)res;
    }
};