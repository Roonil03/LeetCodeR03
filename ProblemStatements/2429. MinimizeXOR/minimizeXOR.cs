public class Solution {
    public int MinimizeXor(int num1, int num2) {
        int s = bits(num2);
        int res = 0;
        for (int i = 30; i >= 0 && s > 0; i--) {
            if ((num1 & (1 << i)) != 0) {
                res |= (1 << i);
                s--;
            }
        }
        for (int i = 0; i <= 30 && s > 0; i++) {
            if ((res & (1 << i)) == 0) {
                res |= (1 << i);
                s--;
            }
        }

        return res;
    }

    int bits(int n) {
        int c = 0;
        while (n > 0) {
            c += n & 1;
            n >>= 1;
        }
        return c;
    }
}