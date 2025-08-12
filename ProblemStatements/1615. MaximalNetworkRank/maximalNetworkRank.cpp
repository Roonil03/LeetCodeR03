class Solution {
public:
    int maximalNetworkRank(int n, vector<vector<int>>& roads) {
        vector<int> deg(n, 0);
        vector<vector<bool>> con(n, vector<bool>(n, false));
        for(auto r : roads){
            int a = r[0], b = r[1];
            deg[a]++;
            deg[b]++;
            con[a][b] = true;
            con[b][a] = true;
        }
        int res {0};
        for (int i{0}; i < n ; i++){
            for(int j = i + 1; j < n; j++){
                int temp = deg[i] + deg[j];
                if(con[i][j]){
                    temp--;
                }
                res = max(res, temp);
            }
        }
        return res;
    }
};