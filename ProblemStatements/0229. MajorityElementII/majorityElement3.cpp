class Solution {
public:
    vector<int> majorityElement(vector<int>& nums) {
        int n = nums.size();
        int c1 {0}, c2{0};
        int can1{0}, can2{1};
        for(int n : nums){
            if(n == can1){
                c1++;
            } else if(n == can2){
                c2++;
            } else if(c1 == 0){
                can1 = n;
                c1 = 1;
            } else if (c2 == 0){
                can2 = n;
                c2 = 1;
            } else{
                c1--;
                c2--;
            }
        }
        c1 = c2 = 0;
        for(int n : nums){
            if(n == can1){
                c1++;
            } else if(n == can2){
                c2++;
            }
        }
        vector<int>res;
        if(c1 > n/3){
            res.push_back(can1);
        }
        if(c2 > n/3){
            res.push_back(can2);
        }
        return res;
    }
};