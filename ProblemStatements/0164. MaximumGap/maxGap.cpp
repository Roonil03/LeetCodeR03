class Solution {
public:
    int maximumGap(vector<int>& nums) {
        int n = nums.size();
        if (n < 2){
            return 0;
        }
        int m1 = *min_element(nums.begin(), nums.end());
        int m2 = *max_element(nums.begin(), nums.end());
        if(m1 == m2){
            return 0;
        }
        int buck = max(1, (m2 - m1)/(n-1));
        int count = (m2 - m1) / buck + 1;
        vector<int> buckMin(count, INT_MAX);
        vector<int> buckMax(count, INT_MIN);
        vector<bool> vis(count, false);
        for(int n : nums){
            int id = (n - m1) / buck;
            buckMin[id] = min(buckMin[id], n);
            buckMax[id] = max(buckMax[id], n);
            vis[id] = true;
        }
        int res {0}, prev {m1};
        for (int i {0}; i < count; i++){
            if(!vis[i]){
                continue;
            }
            res = max(res, buckMin[i] - prev);
            prev = buckMax[i];
        }
        return res;
    }
};