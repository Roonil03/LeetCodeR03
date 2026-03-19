class Solution {
public:
    int minOperations(int n) {
        int res {0};
        while(n > 0){
            if(n & 1){
                if(n == 1){
                    res++;
                    break;
                }
                if(n % 4 == 1){
                    n -= 1;
                } else{
                    n += 1;
                }
                res++;
            }
            n >>= 1;
        }
        return res;
    }
};