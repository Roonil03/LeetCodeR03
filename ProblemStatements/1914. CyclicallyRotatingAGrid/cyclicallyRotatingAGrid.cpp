class Solution {
public:
    vector<vector<int>> rotateGrid(vector<vector<int>>& grid, int k) {
        int m = grid.size(), n = grid[0].size();
        int l = min(m, n) / 2;
        vector<int>temp;
        for(int i {0}; i < l; i++){
            temp.clear();
            int top = i, bot = m - 1 - i, lef = i, rig = n - 1 - i;
            for(int j = lef; j <= rig; j++){
                temp.push_back(grid[top][j]);
            }
            for(int j = top + 1; j <= bot; j++){
                temp.push_back(grid[j][rig]);
            }
            for(int j = rig - 1; j >= lef; j--){
                temp.push_back(grid[bot][j]);
            }
            for(int j = bot - 1; j > top; j--){
                temp.push_back(grid[j][lef]);
            }
            int a = temp.size();
            int k1 = k % a;
            if(k1 > 0){
                ranges::rotate(temp, temp.begin() + k1);
            }
            int id {0};
            for(int j = lef; j <= rig; j++){
                grid[top][j] = temp[id++];
            }
            for(int j = top + 1; j <= bot; j++){
                grid[j][rig] = temp[id++];
            }
            for(int j = rig - 1; j >= lef; j--){
                grid[bot][j] = temp[id++];
            }
            for(int j = bot - 1; j > top; j--){
                grid[j][lef] = temp[id++];
            }
        }
        return grid;
    }
};