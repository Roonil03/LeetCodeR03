class Solution {
    public int findChampion(int n, int[][] edges) {
        int[] indegree = new int[n];
        for (int[] edge: edges) indegree[edge[1]]++;        
        int champ = -1;
        for (int i = 0; i < n; i++) {
            if (indegree[i] == 0) {
                if (champ != -1) return -1; 
                champ = i;
            }
        }
        return champ;
    }
}