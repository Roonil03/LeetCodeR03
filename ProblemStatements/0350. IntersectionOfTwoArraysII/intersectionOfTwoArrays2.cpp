class Solution {
public:
    vector<int> intersect(vector<int>& nums1, vector<int>& nums2) {
        if(nums1.size() > nums2.size()){
            intersect(nums2, nums1);
        }
        unordered_map<int, int> count;
        for(int x : nums1){
            count[x]++;
        }
        vector<int> res;
        for(int x : nums2){
            auto it = count.find(x);
            if(it != count.end() && it->second > 0){
                res.push_back(x);
                --(it->second);
            }
        }
        return res;
    }
};