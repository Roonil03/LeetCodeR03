class Solution {
    public long numberOfWays(String s) {
        int len = s.length();
        long a = 0, b = 0;
        long ans = 0;        
        for (int i = 0; i < len; i++) {
            if (s.charAt(i) == '0') {
                a++;
            } else {
                b++;
            }
        }        
        long c = 0, d = 0;        
        for (int i = 0; i < len; i++) {
            if (s.charAt(i) == '0') {
                ans += d * (b - d);
                c++;
            } 
            else {
                ans += c * (a - c);
                d++;
            }
        }        
        return ans;
    }
}