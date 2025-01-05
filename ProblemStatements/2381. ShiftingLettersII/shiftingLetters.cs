public class Solution {
    public string ShiftingLetters(string s, int[][] shifts) {
        int n = s.Length;
        int[] a = new int[n + 1];

        foreach (var w in shifts)
        {
            int l = w[0];
            int r = w[1];
            int d = w[2];
            if (d == 1)
            {
                a[l] += 1;
                if (r + 1 < n){
                    a[r + 1] -= 1;
                }
            }
            else
            {
                a[l] -= 1;
                if (r + 1 < n){
                    a[r + 1] += 1;
                }
            }
        }
        for (int i = 1; i < n; i++)
        {
            a[i] += a[i - 1];
        }
        char[] chs = s.ToCharArray();
        for (int i = 0; i < n; i++)
        {
            int shift = (chs[i] - 'a' + a[i]) % 26;
            if (shift < 0) shift += 26;
            chs[i] = (char)(shift + 'a');
        }
        return new string(chs);
    }
}