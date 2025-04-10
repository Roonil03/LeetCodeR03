class Solution {
    public:
        long long dp[2][16];
        long long numberOfPowerfulInt(long long start, long long finish, int limit, string s) {
            long long res {0};
            string L = to_string(start - 1), R = to_string(finish);
            if(L.size() >= s.size()){
                memset(dp, -1, sizeof(dp));
                res = find(0, L, true, limit, s);
            }
            memset(dp, -1, sizeof(dp));
            long long sc = find(0, R, true, limit, s);
            return sc - res;
        }
        long long find(int id, string& num, bool t, int& limit, string& s){
            if(dp[t][id] != -1){
                return dp[t][id];
            }
            long long ub = t ? num[id] - '0' : 9;
            long long res {0};
            if(id == num.size() - s.size()){
                if(!t){
                    return dp[t][id] = 1;
                }
                return dp[t][id] = (stoll(num.substr(id, (int)s.size())) >= stoll(s));
            }
            for (int dig {0}; dig <= min(limit, (int)ub); dig++){
                res += find(id + 1, num, t & (dig == ub), limit, s);
            }
            return dp[t][id] = res;
        }
    };