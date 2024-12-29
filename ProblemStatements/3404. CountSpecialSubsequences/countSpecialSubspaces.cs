public class Solution
{
    private int[,] cnt = new int[1001, 1001];

    private int Gcd(int a, int b)
    {
        return b == 0 ? a : Gcd(b, a % b);
    }

    public long NumberOfSubsequences(int[] n)
    {
        long res = 0;
        Array.Clear(cnt, 0, cnt.Length);
        for (int q = n.Length - 5; q > 1; --q)
        {
            for (int s = q + 4, r = q + 2; s < n.Length; ++s)
            {
                int f = Gcd(n[r], n[s]);
                ++cnt[n[r] / f, n[s] / f];
            }
            for (int p = 0; p + 1 < q; ++p)
            {
                int f = Gcd(n[p], n[q]);
                res += cnt[n[q] / f, n[p] / f];
            }
        }
        return res;
    }
}