class Solution {
public:
    int numOfMinutes(int n, int headID, vector<int>& manager, vector<int>& informTime) {
        vector<vector<int>> sub(n);
        for(int i{0}; i < n; i++){
            if(manager[i] != -1){
                sub[manager[i]].push_back(i);
            }
        }
        return dfs(headID, sub, informTime);
    }

    int dfs(int emp, vector<vector<int>>& sub, vector<int>& time){
        if(sub[emp].empty()){
            return 0;
        }
        int res {0};
        for(int s : sub[emp]){
            res = max(res, dfs(s, sub, time));
        }
        return time[emp] + res;
    }
};