class Solution {
public:
    int numberOfArithmeticSlices(vector<int>& nums) {
        using ll = long long;
        int n = nums.size();
        vector<unordered_map<ll, ll>> dp(n);
        ll res {0};
        for(int i {0}; i < n; i++){
            for(int j {0}; j < i; j++){
                ll dif = (ll)nums[i] - (ll)nums[j];
                ll a = dp[j].count(dif) ? dp[j][dif] : 0;
                dp[i][dif] += a + 1;
                res += a;
            }
        }
        return (int)res;
    }
};