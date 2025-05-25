class Solution {
public:
    vector<vector<int>> kSmallestPairs(vector<int>& nums1, vector<int>& nums2, int k) {
        vector<vector<int>> res;
        if(nums1.empty() || nums2.empty() || k == 0){
            return res;
        }
        auto c = [&](const vector<int>& a, const vector<int>& b){
            return a[0] > b[0];
        };
        priority_queue<vector<int>, vector<vector<int>>, decltype(c)> h(c);
        for (int i {0}; i < nums1.size() && i < k; i++){
            h.push({nums1[i] + nums2[0], i, 0});
        }
        while(k -- && !h.empty()){
            auto temp = h.top();
            h.pop();
            int i {temp[1]}, j {temp[2]};
            res.push_back({nums1[i], nums2[j]});
            if (j + 1 < nums2.size()){
                h.push({nums1[i] + nums2[j + 1], i, j+1});
            }
        }
        return res;
    }
};