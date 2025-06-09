class Solution {
public:
    int findKthNumber(int n, int k) {
        long num {1};
        for(int i{1}; i < k;){
            int req = h1(num, num+1, n);
            if(i + req <= k){
                i += req;
                num++;
            } else{
                i++;
                num *= 10;
            }
        }
        return num;
    }

    int h1(long a, long b, int& n){
        int g{0};
        while(a <= n){
            g += min(long(n+1), b) - a;
            a *= 10;
            b *= 10;
        }
        return g;
    }
};