class Solution {
public:
    vector<int> selfDividingNumbers(int left, int right) {
        vector<int> res;
        for(int i = left; i <= right; i++){
            if(h1(i)){
                res.push_back(i);
            }
        }
        return res;
    }

    bool h1(int i){
        int o = i;
        while(i > 0){
            int d = i % 10;
            if(d == 0 || (o % d) != 0){
                return false;
            }
            i /= 10;
        }
        return true;
    }
};