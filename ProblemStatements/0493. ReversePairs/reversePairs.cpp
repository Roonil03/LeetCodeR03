class Solution {
public:
    int reversePairs(vector<int>& nums) {
        return h1(nums.begin(), nums.end());
    }

    int h1(vector<int>::iterator b, vector<int>::iterator e){
        if(e - b <= 1){
            return 0;
        }
        auto m = b + (e - b) / 2;
        int count = h1(b, m) + h1(m , e);
        for(auto i = b, j = m; i != m; i++){
            while(j != e && *i > 2LL * *j){
                j++;
            }
            count += j - m;
        }
        inplace_merge(b, m, e);
        return count;
    }
};