class Solution {
public:
    int nthUglyNumber(int n) {
        int a {0}, b {0}, c {0};
        vector<int> d(n);
        d[0] = 1;
        for(int i {1} ; i < n; i++){
            int x2 = d[a] * 2, x3 = d[b] * 3, x5 = d[c] * 5;
            d[i] = min({x2, x3, x5});
            if(d[i] == x2){
                a++;
            }
            if(d[i] == x3){
                b++;
            }
            if(d[i] == x5){
                c++;
            }
        }
        return d[n-1];
    }
};