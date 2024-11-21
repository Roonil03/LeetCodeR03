class Solution{  
    public int CountUnguarded(int m, int n, int[][] guards, int[][] walls) {
        int total = m * n, guarded = 0;
        char[,] grid = new char[m, n];
        foreach (var guard in guards)
            grid[guard[0], guard[1]] = 'G';
        foreach (var wall in walls)
            grid[wall[0], wall[1]] = 'W';
        foreach (var guard in guards) {
            int x = guard[0], y = guard[1];
            for (int d = 0; d < 4; d++) {
                int dx = d == 0 ? -1 : d == 1 ? 1 : 0;
                int dy = d < 2 ? 0 : d == 2 ? -1 : 1;
                int nx = x + dx, ny = y + dy;
                while (nx >= 0 && ny >= 0 && nx < m && ny < n) {
                    if (grid[nx, ny] == 'W' || grid[nx, ny] == 'G') break;
                    if (grid[nx, ny] == '\0') {
                        grid[nx, ny] = 'V';
                        guarded++;
                    }
                    nx += dx;
                    ny += dy;
                }
            }
        }
        return total - guarded - guards.Length - walls.Length;
    }
}