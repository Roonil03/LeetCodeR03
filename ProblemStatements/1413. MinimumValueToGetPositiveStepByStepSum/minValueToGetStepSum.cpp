class Solution {
public:
    int minStartValue(vector<int>& nums) {
        int a {0}, b {0};
        for(int i : nums){
            a += i;
            b = min(b, a);
        }
        return 1 - b;
    }
};
