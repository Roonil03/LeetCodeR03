class Solution {
public:
    int countNumbersWithUniqueDigits(int n) {
        if(n == 0){
            return 1;
        }
        int res {10}, ud {9}, a {9};
        for(int i {2}; i <= n && a > 0; i++){
            ud *= a--;
            res += ud;
        }
        return res;
    }
};