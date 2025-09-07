class Solution {
public:
    int integerReplacement(int num) {
        int res {0};
        long long n = num;
        while(n != 1){
            if((n & 1) == 0){
                n >>= 1;
            } else{
                if(n == 3 || ((n >> 1) & 1) == 0){
                    n--;
                } else {
                    n++;
                }
            }
            res++;
        }
        return res;
    }
};