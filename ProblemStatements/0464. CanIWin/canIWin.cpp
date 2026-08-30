class Solution {
public:
    bool canIWin(int maxChoosableInteger, int desiredTotal) {
        int sum = (maxChoosableInteger * (maxChoosableInteger + 1)) / 2;
        if(sum < desiredTotal){
            return false;
        }
        vector<int>mem(1<<maxChoosableInteger, -1);
        auto dfs = [&](this auto& self, int mask, int cur) -> bool{
            if(mem[mask] != -1){
                return mem[mask];
            }
            for(int i {1}; i <= maxChoosableInteger; i++){
                int b = 1 << (i - 1);
                if((mask & b) == 0){
                    if(cur + i >= desiredTotal || !self(mask | b, cur + i)){
                        return mem[mask] = 1;
                    }
                }
            }
            return mem[mask] = 0;
        };
        return dfs(0, 0);
    }
};