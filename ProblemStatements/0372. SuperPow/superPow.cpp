class Solution {
public:
    int superPow(int a, vector<int>& b) {
        if(b.empty()){
            return 1;
        }
        int temp = b.back();
        b.pop_back();
        return h1(superPow(a, b), 10) * h1(a, temp) % 1337;
    }

    int h1(int a, int k){
        a %= 1337;
        int res {1};
        for(int i {0}; i < k; i++){
            res = (res * a) % 1337;
        }
        return res;
    }
};