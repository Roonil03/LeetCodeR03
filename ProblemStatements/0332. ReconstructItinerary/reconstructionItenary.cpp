class Solution {
public:
    vector<string> findItinerary(vector<vector<string>>& tickets) {
        unordered_map<string, priority_queue<string, vector<string>, greater<string>>> g;
        for (auto& ticket : tickets) {
            g[ticket[0]].push(ticket[1]);
        }
        vector<string> res;
        dfs("JFK", g, res);
        reverse(res.begin(), res.end());
        return res;
    }

    void dfs(const string& a, unordered_map<string, priority_queue<string, vector<string>, greater<string>>>& g, vector<string>& res){
        while (!g[a].empty()) {
            string next = g[a].top();
            g[a].pop();
            dfs(next, g, res);
        }
        res.push_back(a);
    }
};