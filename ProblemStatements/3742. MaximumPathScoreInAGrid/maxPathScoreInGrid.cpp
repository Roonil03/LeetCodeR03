class Solution {
public:
    int maxPathScore(vector<vector<int>>& grid, int k) {
        int m = grid.size();
        int n = grid[0].size();
        int eff_k = min(k, m + n);
        int stride = eff_k + 1;
        vector<int> dp_prev(n * stride, -1);
        vector<int> dp_curr(n * stride, -1);
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                int val = grid[i][j];
                int ci = (val > 0 ? 1 : 0);
                int* cur = &dp_curr[j * stride];
                if (i == 0 && j == 0) {
                    if (ci <= eff_k) {
                        for (int c = ci; c <= eff_k; ++c) cur[c] = val;
                    }
                    if (ci == 1 && eff_k >= 0) cur[0] = -1;
                    continue;
                }
                int* p_row = (i > 0) ? &dp_prev[j * stride] : nullptr;
                int* p_col = (j > 0) ? &dp_curr[(j - 1) * stride] : nullptr;
                if (ci == 1 && eff_k >= 0){
                    cur[0] = -1;
                }
                if (p_row && p_col) {
                    for (int c = ci; c <= eff_k; ++c) {
                        int bp = max(p_row[c - ci], p_col[c - ci]);
                        cur[c] = (bp == -1) ? -1 : bp + val;
                    }
                } else if (p_row) {
                    for (int c = ci; c <= eff_k; ++c) {
                        int bp = p_row[c - ci];
                        cur[c] = (bp == -1) ? -1 : bp + val;
                    }
                } else if (p_col) {
                    for (int c = ci; c <= eff_k; ++c) {
                        int bp = p_col[c - ci];
                        cur[c] = (bp == -1) ? -1 : bp + val;
                    }
                }
            }
            dp_prev.swap(dp_curr);
        }
        return dp_prev[(n - 1) * stride + eff_k];
    }
};