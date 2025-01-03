class Solution {
    public int[][] substringXorQueries(String s, int[][] queries) {
        Map<Integer, int[]>m = new HashMap();
        int len = s.length();
        for(int i = 0; i < len; i++){
            if(s.charAt(i) == '0'){
                m.putIfAbsent(0, new int[] {i, i});
                continue;
            }
            int num = 0;
            for(int j = i; j < len && j < i + 32; j++){
                num *= 2;
                num += (s.charAt(j) - '0');
                m.putIfAbsent(num, new int[] {i, j});
            }
        }
        int [][] ans = new int[queries.length][2];
        int i = 0;
        for(int [] q : queries){
            int val = q[0]^q[1];
            ans[i++] = m.getOrDefault(val, new int[] {-1, -1});
        }
        return ans;
    }
}