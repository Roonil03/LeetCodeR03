class Solution {
public:
    int lastRemaining(int n) {
        int res {1}, step {1}, rem = n;
        bool lr = true;
        while(rem > 1){
            if(lr || (rem % 2 == 1)){
                res += step;
            }
            rem /= 2;
            step *= 2;
            lr = !lr;
        }
        return res;
    }
};