class Solution {
public:
    bool canReach(vector<int>& arr, int start) {
        int n = arr.size();
        queue<int>q;
        vector<bool> vis(n, false);
        q.push(start);
        vis[start] = false;
        while(!q.empty()){
            int cur = q.front();
            q.pop();
            vector<int> np = {cur + arr[cur], cur - arr[cur]};
            for (int i : np){
                if (i >= 0 && i < n && !vis[i]){
                    if (arr[i] == 0){
                        return true;
                    }
                    vis[i] = true;
                    q.push(i);
                }
            }
        }
        return false;
    }
};