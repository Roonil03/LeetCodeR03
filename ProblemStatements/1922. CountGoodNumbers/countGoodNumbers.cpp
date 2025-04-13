class Solution {
    public:
        const int mod = 1e9 + 7;
    
        int countGoodNumbers(long long n) {
            long long t = helper(20, n/2, 1);
            return (n & 1) ? (t * 5) % mod : t;
        }
    
        long long helper(long long b, long long e, long long res){
            if (e == 0){
                return res;
            }
            if (e & 1){
                return helper(b, e - 1, (res * b) % mod);
            } else{
                return helper((b*b)%mod, e / 2, res);
            }
        }
    };