public class Solution
{
    private const int MOD = 1000000007;

    private long ModPow(long baseValue, long exp, long mod)
    {
        long result = 1;
        while (exp > 0)
        {
            if ((exp & 1) > 0)
                result = (result * baseValue) % mod;
            baseValue = (baseValue * baseValue) % mod;
            exp >>= 1;
        }
        return result;
    }

    public int CountGoodArrays(int n, int m, int k)
    {
        if (k >= n) return 0;
        long total = (m * ModPow(m - 1, n - 1 - k, MOD)) % MOD;
        for (int i = 1; i <= k; i++)
        {
            total = (total * (n - i)) % MOD;
            total = (total * ModPow(i, MOD - 2, MOD)) % MOD;
        }
        return (int)total;
    }
}