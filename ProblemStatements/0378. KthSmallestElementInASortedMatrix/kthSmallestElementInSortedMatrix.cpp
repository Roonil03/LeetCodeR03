class Solution {
public:
    int kthSmallest(vector<vector<int>>& matrix, int k) {
        int n = matrix.size();
        auto cmp = [](tuple<int, int, int>& a, tuple<int, int, int>& b){
            return get<0>(a) > get<0>(b);
        };
        priority_queue<tuple<int, int, int>, vector<tuple<int, int, int>>, decltype(cmp)> h(cmp);
        for(int i {0}; i < n; i++){
            h.push({matrix[i][0], i, 0});
        }
        for(int i {0}; i < k; i++){
            auto[val, r, c] = h.top();
            h.pop();
            if(i == k - 1){
                return val;
            }
            if(c + 1 < n){
                h.push({matrix[r][c+1], r, c + 1});
            }
        }
        return -1;
    }
};