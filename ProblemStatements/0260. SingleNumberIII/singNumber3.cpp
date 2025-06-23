class Solution {
public:
    vector<int> singleNumber(vector<int>& nums) {
        long x = 0;
        for (int n : nums){
            x ^= n;
        }
        int d = x & -x;
        int a {0}, b {0};
        for(int n : nums){
            if(n & d){
                a ^= n;
            } else{
                b ^= n;
            }
        }
        return {a,b};
    }
};