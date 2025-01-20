public class Solution {
    public int FirstCompleteIndex(int[] arr, int[][] mat) {
        int m = mat.Length, n = mat[0].Length;
        int[] map = new int[m * n + 1];        
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                map[mat[i][j]] = i * n + j;
            }
        }
        int[] r = new int[m];
        int[] c = new int[n];
        for (int i = 0; i < arr.Length; i++) {
            int x = map[arr[i]];
            int row = x / n;
            int col = x % n;            
            r[row]++;
            c[col]++;            
            if (r[row] == n || c[col] == m) {
                return i;
            }
        }
        return -1;
    }
}