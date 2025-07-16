class Solution {
public:
    int maxProduct(vector<string>& words) {
        int n = words.size();
        vector<int> a(n, 0), b(n);
        for (int i {0}; i <n; i++){
            int m {0};
            for (char ch : words[i]){
                m |= 1 << (ch - 'a');
            }
            a[i] = m;
            b[i] = words[i].size();
        }
        int res {0};
        for (int i{0}; i < n; i++){
            for(int j = i + 1; j < n; j++){
                if ((a[i] & a[j]) == 0){
                    int prod = b[i] * b [j];
                    if(prod > res){
                        res = prod;
                    }
                }
            }
        }
        return res;
    }
};