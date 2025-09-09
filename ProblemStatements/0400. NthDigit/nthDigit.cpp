class Solution {
    using ll = long long;
public:
    int findNthDigit(int n) {
        ll len {1}, count {9}, st {1};
        while(n > len * count){
            n -= len * count;
            len += 1;
            count *= 10;
            st *= 10;
        }
        ll num = st + (n - 1) / len, id = (n - 1) % len;
        ll a = (len - 1)  - id;
        while(a--){
            num /= 10;
        }
        return num%10;
    }
};