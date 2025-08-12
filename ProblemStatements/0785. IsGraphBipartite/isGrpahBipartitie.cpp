class Solution {
public:
    bool isBipartite(vector<vector<int>>& graph) {
        int n = graph.size();
        vector<int> col(n, -1);
        for(int i {0}; i < n; i++){
            if(col[i] == -1){
                if(!bfs(graph, i, col)){
                    return false;
                }
            }
        }
        return true;
    }
    
    bool bfs(vector<vector<int>>& g, int s, vector<int>& col){
        queue<int>q;
        q.push(s);
        col[s] = 0;
        while(!q.empty()){
            int n = q.front();
            q.pop();
            for(int nei : g[n]){
                if(col[nei] == -1){
                    col[nei] = 1 - col[n];
                    q.push(nei);
                } else if(col[nei] == col[n]){
                    return false;
                }
            }
        }
        return true;
    }
};