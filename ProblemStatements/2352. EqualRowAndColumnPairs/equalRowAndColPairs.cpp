class Solution {
public:
    int equalPairs(vector<vector<int>>& grid) {
        int n = grid.size();
        int res {0};        
        for(int r {0}; r < n; r++) {
            for(int c {0}; c < n; c++) {
                bool eq = true;
                for(int k = 0; k < n; k++) {
                    if(grid[r][k] != grid[k][c]) {
                        eq = false;
                        break;
                    }
                }
                if(eq){
                    res++;
                }
            }
        }        
        return res;
    }
};