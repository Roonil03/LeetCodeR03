class Solution {
public:
    int maxSumSubmatrix(vector<vector<int>>& matrix, int k) {
        int r = matrix.size(), c = matrix[0].size();
        int res = INT_MIN;
        for(int l {0}; l < c; l++){
            vector<int>temp(r, 0);
            for(int h = l; h < c; h++){
                for(int i {0}; i < r; i++){
                    temp[i] += matrix[i][h];
                }
                set<int> pre;
                pre.insert(0);
                int cur {0};
                for(int s : temp){
                    cur += s;
                    auto it = pre.lower_bound(cur - k);
                    if(it != pre.end()){
                        int can = cur - *it;
                        res = max(res, can);
                    }
                    pre.insert(cur);
                }
            }
        }
        return res;
    }
};