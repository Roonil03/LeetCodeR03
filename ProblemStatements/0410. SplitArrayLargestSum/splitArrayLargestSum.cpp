class Solution {
public:
    int splitArray(vector<int>& nums, int k) {
        int l {0}, r {0};
        for(int i : nums){
            l = max(l, i);
            r += i;
        }
        while(l < r){
            int m = l + (r - l) / 2;
            if(h1(nums, m, k)){
                r = m;
            } else{
                l = m + 1;
            }
        }
        return l;
    }

    bool h1(vector<int>& nums, int m, int k){
        int cur {0}, p {1};
        for(int i : nums){
            if(cur + i > m){
                p++;
                cur = i;
                if(p > k){
                    return false;
                }
            } else{
                cur += i;
            }
        }
        return true;
    }
};