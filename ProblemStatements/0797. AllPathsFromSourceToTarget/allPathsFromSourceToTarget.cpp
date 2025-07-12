class Solution {
public:
    vector<vector<int>> allPathsSourceTarget(vector<vector<int>>& graph) {
        vector<vector<int>> res;
        vector<int>path;
        dfs(graph, 0, path, res);
        return res;
    }

    void dfs(vector<vector<int>>& g, int n, vector<int>& p, vector<vector<int>>& res){
        p.push_back(n);
        if(n == g.size() - 1){
            res.push_back(p);
        } else{
            for(int a : g[n]){
                dfs(g, a, p, res);
            }
        }
        p.pop_back();
    }
};