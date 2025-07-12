class Solution {
public:
    int nthSuperUglyNumber(int n, vector<int>& primes) {
        vector<int> id(primes.size(), 0);
        vector<int> ugly(n);
        ugly[0] = 1;
        for(int i{1}; i < n; i++){
            long long  nv = numeric_limits<long long>::max();
            for(int j{0}; j < primes.size(); j++){
                nv = min(nv, (long long)primes[j] * ugly[id[j]]);
            }
            ugly[i] = (int)nv;
            for(int j{0}; j < primes.size(); j++){
                if((long long)primes[j] * ugly[id[j]] == nv){
                    id[j]++;
                }
            }
        }
        return ugly[n - 1];
    }
};