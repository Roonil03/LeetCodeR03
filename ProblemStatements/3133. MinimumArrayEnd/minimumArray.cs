public class Solution {
    public long MinEnd(int n, int x) {
        long num = x;
        for (int i = 1; i < n; i++) {
            num = (num + 1) | x;
        }
        return num;
    }
}