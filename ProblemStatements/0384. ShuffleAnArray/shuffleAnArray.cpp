class Solution {
public:

    vector<int> og, cur;
    random_device rd;
    mt19937 gen;

    Solution(vector<int>& nums) :
        og(nums),
        cur(nums),
        gen(rd())
    {}
    
    vector<int> reset() {
        cur = og;
        return cur;
    }
    
    vector<int> shuffle() {
        int n = cur.size();
        for(int i = n - 1; i > 0; i--){
            uniform_int_distribution<> dis(0, i);
            int j = dis(gen);
            swap(cur[i], cur[j]);
        }
        return cur;
    }
};

/**
 * Your Solution object will be instantiated and called as such:
 * Solution* obj = new Solution(nums);
 * vector<int> param_1 = obj->reset();
 * vector<int> param_2 = obj->shuffle();
 */