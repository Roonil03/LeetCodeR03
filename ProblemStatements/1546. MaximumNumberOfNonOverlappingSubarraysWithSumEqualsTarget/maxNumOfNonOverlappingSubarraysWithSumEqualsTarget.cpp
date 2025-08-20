class Solution {
public:
    int maxNonOverlapping(vector<int>& nums, int target) {
        unordered_set<int> ss;
        int pre {0}, count {0};
        ss.insert(0);
        for(int i : nums){
            pre += i;
            if(ss.count(pre - target)){
                count++;
                ss.clear();
                ss.insert(0);
                pre = 0;
            } else{
                ss.insert(pre);
            }
        }
        return count;
    }
};