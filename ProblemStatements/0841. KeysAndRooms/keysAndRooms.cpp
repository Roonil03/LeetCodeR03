class Solution {
public:
    bool canVisitAllRooms(vector<vector<int>>& rooms) {
        int n = rooms.size();
        vector<bool> vis(n, false);
        queue<int> q;
        q.push(0);
        vis[0] = true;
        int temp {1};
        while (!q.empty()) {
            int c = q.front(); q.pop();
            for (int k : rooms[c]) {
                if (!vis[k]) {
                    vis[k] = true;
                    q.push(k);
                    temp++;
                }
            }
        }
        return temp == n;
    }
};