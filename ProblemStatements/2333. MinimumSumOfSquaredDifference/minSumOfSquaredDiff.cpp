class Solution {
public:
    long long minSumSquareDiff(vector<int>& nums1, vector<int>& nums2, int k1, int k2) {
        int m {0}, n = nums1.size();
        for(int i {0}; i <n ;i++){
            m = max(m, abs(nums1[i] - nums2[i]));
        }
        if(!m){
            return 0;
        }
        vector<int> count(m + 1);
        for(int i {0}; i < n; i++){
            count[abs(nums1[i] - nums2[i])]++;
        }
        long long k = 1LL * k1 + k2;
        for(int i = m ; i> 0 && k > 0; i--){
            if(!count[i]){
                continue;
            }
            long long take = min<long long>(k, count[i]);
            count[i] -= take;
            count[i - 1] += take;
            k -= take;
        }
        long long res {0};
        for(long long i {1}; i <= m; i++){
            if(count[i]){
                res += count[i] * i * i;
            }
        }
        return res;
    }
};